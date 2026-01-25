// Package scaffold provides templates for generating new connector projects.
package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Config holds configuration for generating a new connector project.
type Config struct {
	// Name is the connector name (e.g., "my-app")
	Name string
	// ModulePath is the Go module path (e.g., "github.com/myorg/baton-my-app")
	ModulePath string
	// OutputDir is where the project will be created
	OutputDir string
	// Description is a brief description of the connector
	Description string
}

// Generate creates a new connector project from the standard template.
func Generate(cfg *Config) error {
	if cfg.Name == "" {
		return fmt.Errorf("scaffold: connector name is required")
	}
	if cfg.ModulePath == "" {
		cfg.ModulePath = fmt.Sprintf("github.com/conductorone/baton-%s", cfg.Name)
	}
	if cfg.OutputDir == "" {
		cfg.OutputDir = fmt.Sprintf("baton-%s", cfg.Name)
	}
	if cfg.Description == "" {
		cfg.Description = fmt.Sprintf("ConductorOne connector for %s", cfg.Name)
	}

	// Create output directory
	if err := os.MkdirAll(cfg.OutputDir, 0755); err != nil {
		return fmt.Errorf("scaffold: failed to create output directory: %w", err)
	}

	// Generate all template files
	for _, tf := range templateFiles {
		if err := generateFile(cfg, tf); err != nil {
			return fmt.Errorf("scaffold: failed to generate %s: %w", tf.Path, err)
		}
	}

	return nil
}

// templateFile represents a file to be generated.
type templateFile struct {
	Path     string
	Template string
	Mode     os.FileMode
}

// generateFile generates a single file from a template.
func generateFile(cfg *Config, tf templateFile) error {
	// Parse template
	tmpl, err := template.New(tf.Path).Parse(tf.Template)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Expand path template variables
	expandedPath := strings.ReplaceAll(tf.Path, "{{.Name}}", cfg.Name)

	// Create directory structure
	fullPath := filepath.Join(cfg.OutputDir, expandedPath)
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create file
	mode := tf.Mode
	if mode == 0 {
		mode = 0644
	}
	f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	// Execute template
	data := map[string]string{
		"Name":        cfg.Name,
		"NameTitle":   toTitleCase(cfg.Name),
		"NamePascal":  toPascalCase(cfg.Name),
		"ModulePath":  cfg.ModulePath,
		"Description": cfg.Description,
	}
	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

// toPascalCase converts a kebab-case string to PascalCase.
func toPascalCase(s string) string {
	parts := strings.Split(s, "-")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, "")
}

// toTitleCase converts a kebab-case string to Title Case.
func toTitleCase(s string) string {
	parts := strings.Split(s, "-")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}

// templateFiles contains all the files to generate for a new connector.
// These templates use baton-sdk v0.4.7+ patterns with config.DefineConfiguration.
var templateFiles = []templateFile{
	{
		Path: "go.mod",
		Template: `module {{.ModulePath}}

go 1.23

require (
	github.com/conductorone/baton-sdk v0.4.7
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	go.uber.org/zap v1.27.0
)
`,
	},
	{
		Path: "main.go",
		Template: `package main

import (
	"context"
	"fmt"
	"os"

	"{{.ModulePath}}/pkg/connector"
	configSdk "github.com/conductorone/baton-sdk/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
	"github.com/conductorone/baton-sdk/pkg/field"
	"github.com/conductorone/baton-sdk/pkg/types"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

var version = "dev"

// Config holds the connector configuration.
// It implements field.Configurable to work with the SDK's configuration system.
type Config struct {
	// Add connector-specific fields here as needed.
	// Example: APIKey string ` + "`" + `mapstructure:"api-key"` + "`" + `
}

// Implement field.Configurable interface.
// These methods allow the SDK to read configuration values.
func (c *Config) GetString(key string) string         { return "" }
func (c *Config) GetBool(key string) bool             { return false }
func (c *Config) GetInt(key string) int               { return 0 }
func (c *Config) GetStringSlice(key string) []string  { return nil }
func (c *Config) GetStringMap(key string) map[string]any { return nil }

// Configuration fields for the connector.
// Add required fields here, e.g.:
//   field.StringField("api-key", field.WithRequired(true), field.WithDescription("API key")),
var configFields = []field.SchemaField{}

// ConfigSchema is the configuration schema for the connector.
var ConfigSchema = field.NewConfiguration(
	configFields,
	field.WithConnectorDisplayName("{{.NameTitle}}"),
)

func main() {
	ctx := context.Background()

	_, cmd, err := configSdk.DefineConfiguration(
		ctx,
		"baton-{{.Name}}",
		getConnector,
		ConfigSchema,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	cmd.Version = version

	err = cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func getConnector(ctx context.Context, cfg *Config) (types.ConnectorServer, error) {
	l := ctxzap.Extract(ctx)

	cb, err := connector.New(ctx)
	if err != nil {
		l.Error("error creating connector", zap.Error(err))
		return nil, err
	}

	c, err := connectorbuilder.NewConnector(ctx, cb)
	if err != nil {
		l.Error("error creating connector", zap.Error(err))
		return nil, err
	}

	return c, nil
}
`,
	},
	{
		Path: "pkg/connector/connector.go",
		Template: `package connector

import (
	"context"
	"io"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
)

// Connector implements the {{.Name}} connector.
type Connector struct {
	// Add API client or other state here.
}

// ResourceSyncers returns a ResourceSyncer for each resource type that should be synced.
func (c *Connector) ResourceSyncers(ctx context.Context) []connectorbuilder.ResourceSyncer {
	return []connectorbuilder.ResourceSyncer{
		newUserBuilder(),
	}
}

// Asset takes an input AssetRef and attempts to fetch it.
func (c *Connector) Asset(ctx context.Context, asset *v2.AssetRef) (string, io.ReadCloser, error) {
	return "", nil, nil
}

// Metadata returns metadata about the connector.
func (c *Connector) Metadata(ctx context.Context) (*v2.ConnectorMetadata, error) {
	return &v2.ConnectorMetadata{
		DisplayName: "{{.NameTitle}}",
		Description: "{{.Description}}",
	}, nil
}

// Validate is called to ensure that the connector is properly configured.
func (c *Connector) Validate(ctx context.Context) (annotations.Annotations, error) {
	// TODO: Implement validation (e.g., test API connection)
	return nil, nil
}

// New returns a new instance of the connector.
func New(ctx context.Context) (*Connector, error) {
	return &Connector{}, nil
}
`,
	},
	{
		Path: "pkg/connector/resource_types.go",
		Template: `package connector

import (
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
)

// userResourceType defines the user resource type.
var userResourceType = &v2.ResourceType{
	Id:          "user",
	DisplayName: "User",
	Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_USER},
}
`,
	},
	{
		Path: "pkg/connector/users.go",
		Template: `package connector

import (
	"context"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
)

type userBuilder struct{}

func (o *userBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return userResourceType
}

// List returns all the users from the upstream service as resource objects.
func (o *userBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	// TODO: Implement user listing
	// Example:
	//   users, nextToken, err := client.ListUsers(ctx, pToken.Token)
	//   for _, user := range users {
	//       r, _ := resource.NewUserResource(user.Name, userResourceType, user.ID,
	//           []resource.UserTraitOption{resource.WithEmail(user.Email, true)})
	//       resources = append(resources, r)
	//   }
	return nil, "", nil, nil
}

// Entitlements always returns an empty slice for users.
func (o *userBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

// Grants always returns an empty slice for users since they don't have any entitlements.
func (o *userBuilder) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func newUserBuilder() *userBuilder {
	return &userBuilder{}
}
`,
	},
	{
		Path: ".gitignore",
		Template: `# Binaries
baton-{{.Name}}
*.exe
*.dll
*.so
*.dylib

# Test coverage
*.out
coverage.html

# IDE
.idea/
.vscode/
*.swp
*.swo

# OS
.DS_Store
Thumbs.db

# Build artifacts
dist/
c1z/
*.c1z

# Environment
.env
.env.local
`,
	},
	{
		Path: "README.md",
		Template: `# baton-{{.Name}}

{{.Description}}

## Prerequisites

- Go 1.23+

## Installation

` + "```" + `bash
go install {{.ModulePath}}@latest
` + "```" + `

## Usage

` + "```" + `bash
# Run sync
baton-{{.Name}}

# See all options
baton-{{.Name}} --help
` + "```" + `

## Development

` + "```" + `bash
# Build
go build -o baton-{{.Name}} .

# Run locally
./baton-{{.Name}}

# Run with hot reload (using cone)
cone connector dev
` + "```" + `

## Resources

This connector syncs the following resources:

| Resource Type | Description |
|---------------|-------------|
| User | {{.NameTitle}} users |

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: ` + "`go test ./...`" + `
5. Submit a pull request
`,
	},
	{
		Path: "Makefile",
		Template: `.PHONY: build test clean

BINARY_NAME=baton-{{.Name}}

build:
	go build -o $(BINARY_NAME) .

test:
	go test -v ./...

clean:
	rm -f $(BINARY_NAME)
	rm -rf dist/

lint:
	golangci-lint run

.DEFAULT_GOAL := build
`,
	},
}
