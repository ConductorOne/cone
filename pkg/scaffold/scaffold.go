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
		"NameTitle":   strings.Title(strings.ReplaceAll(cfg.Name, "-", " ")),
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

// templateFiles contains all the files to generate for a new connector.
var templateFiles = []templateFile{
	{
		Path: "go.mod",
		Template: `module {{.ModulePath}}

go 1.21

require (
	github.com/conductorone/baton-sdk v0.2.0
	github.com/conductorone/conductorone-sdk-go v1.0.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	go.uber.org/zap v1.26.0
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

	"{{.ModulePath}}/cmd/baton-{{.Name}}/config"
	"{{.ModulePath}}/pkg/connector"
	"github.com/conductorone/baton-sdk/pkg/cli"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
	"github.com/conductorone/baton-sdk/pkg/types"
)

var version = "dev"

func main() {
	ctx := context.Background()

	cfg := &config.Config{}
	app, err := cli.NewApp(
		"baton-{{.Name}}",
		version,
		cfg,
		cli.WithConnector(func(ctx context.Context, cfg *config.Config) (types.ConnectorServer, error) {
			return connector.New(ctx, cfg)
		}),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	err = app.Run(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
`,
	},
	{
		Path: "cmd/baton-{{.Name}}/config/config.go",
		Template: `package config

import "github.com/conductorone/baton-sdk/pkg/field"

// Config holds the configuration for the {{.Name}} connector.
type Config struct {
	// APIKey is the API key for authenticating with {{.NameTitle}}.
	APIKey string ` + "`" + `mapstructure:"api-key"` + "`" + `
	// BaseURL is the base URL for the {{.NameTitle}} API (optional).
	BaseURL string ` + "`" + `mapstructure:"base-url"` + "`" + `
}

// Fields returns the configuration fields for the connector.
func (c *Config) Fields() []field.SchemaField {
	return []field.SchemaField{
		field.StringField(
			"api-key",
			field.WithRequired(true),
			field.WithDescription("API key for {{.NameTitle}}"),
		),
		field.StringField(
			"base-url",
			field.WithDescription("Base URL for the {{.NameTitle}} API"),
		),
	}
}
`,
	},
	{
		Path: "pkg/connector/connector.go",
		Template: `package connector

import (
	"context"
	"fmt"

	"{{.ModulePath}}/cmd/baton-{{.Name}}/config"
	"{{.ModulePath}}/pkg/client"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
)

// Connector implements the {{.Name}} connector.
type Connector struct {
	client *client.Client
}

// New creates a new {{.Name}} connector.
func New(ctx context.Context, cfg *config.Config) (*Connector, error) {
	c, err := client.New(ctx, cfg.APIKey, cfg.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("{{.Name}}: failed to create client: %w", err)
	}
	return &Connector{client: c}, nil
}

// ResourceSyncers returns the resource syncers for this connector.
func (c *Connector) ResourceSyncers(ctx context.Context) []connectorbuilder.ResourceSyncer {
	return []connectorbuilder.ResourceSyncer{
		newUserSyncer(c.client),
		newGroupSyncer(c.client),
		newRoleSyncer(c.client),
	}
}

// Metadata returns connector metadata.
func (c *Connector) Metadata(ctx context.Context) (*v2.ConnectorMetadata, error) {
	return &v2.ConnectorMetadata{
		DisplayName: "{{.NameTitle}}",
		Description: "{{.Description}}",
	}, nil
}

// Validate validates the connector configuration.
func (c *Connector) Validate(ctx context.Context) (annotations.Annotations, error) {
	// TODO: Implement validation (e.g., test API connection)
	return nil, nil
}
`,
	},
	{
		Path: "pkg/connector/users.go",
		Template: `package connector

import (
	"context"

	"{{.ModulePath}}/pkg/client"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	"github.com/conductorone/baton-sdk/pkg/types/resource"
)

const userResourceTypeID = "user"

var userResourceType = &v2.ResourceType{
	Id:          userResourceTypeID,
	DisplayName: "User",
	Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_USER},
}

type userSyncer struct {
	client *client.Client
}

func newUserSyncer(c *client.Client) *userSyncer {
	return &userSyncer{client: c}
}

func (s *userSyncer) ResourceType(ctx context.Context) *v2.ResourceType {
	return userResourceType
}

func (s *userSyncer) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	// TODO: Implement user listing
	// users, nextToken, err := s.client.ListUsers(ctx, pToken.Token)
	// if err != nil {
	//     return nil, "", nil, err
	// }

	var resources []*v2.Resource
	// for _, user := range users {
	//     r, err := resource.NewUserResource(
	//         user.Name,
	//         userResourceType,
	//         user.ID,
	//         []resource.UserTraitOption{
	//             resource.WithEmail(user.Email, true),
	//         },
	//     )
	//     if err != nil {
	//         return nil, "", nil, err
	//     }
	//     resources = append(resources, r)
	// }

	return resources, "", nil, nil
}

func (s *userSyncer) Entitlements(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func (s *userSyncer) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}
`,
	},
	{
		Path: "pkg/connector/groups.go",
		Template: `package connector

import (
	"context"

	"{{.ModulePath}}/pkg/client"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
)

const groupResourceTypeID = "group"

var groupResourceType = &v2.ResourceType{
	Id:          groupResourceTypeID,
	DisplayName: "Group",
	Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_GROUP},
}

type groupSyncer struct {
	client *client.Client
}

func newGroupSyncer(c *client.Client) *groupSyncer {
	return &groupSyncer{client: c}
}

func (s *groupSyncer) ResourceType(ctx context.Context) *v2.ResourceType {
	return groupResourceType
}

func (s *groupSyncer) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	// TODO: Implement group listing
	return nil, "", nil, nil
}

func (s *groupSyncer) Entitlements(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	// TODO: Implement group entitlements (membership)
	return nil, "", nil, nil
}

func (s *groupSyncer) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	// TODO: Implement group grants (members)
	return nil, "", nil, nil
}
`,
	},
	{
		Path: "pkg/connector/roles.go",
		Template: `package connector

import (
	"context"

	"{{.ModulePath}}/pkg/client"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
)

const roleResourceTypeID = "role"

var roleResourceType = &v2.ResourceType{
	Id:          roleResourceTypeID,
	DisplayName: "Role",
	Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_ROLE},
}

type roleSyncer struct {
	client *client.Client
}

func newRoleSyncer(c *client.Client) *roleSyncer {
	return &roleSyncer{client: c}
}

func (s *roleSyncer) ResourceType(ctx context.Context) *v2.ResourceType {
	return roleResourceType
}

func (s *roleSyncer) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	// TODO: Implement role listing
	return nil, "", nil, nil
}

func (s *roleSyncer) Entitlements(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	// TODO: Implement role entitlements (assignment)
	return nil, "", nil, nil
}

func (s *roleSyncer) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	// TODO: Implement role grants (assignees)
	return nil, "", nil, nil
}
`,
	},
	{
		Path: "pkg/client/client.go",
		Template: `package client

import (
	"context"
	"fmt"
	"net/http"
)

const defaultBaseURL = "https://api.example.com/v1"

// Client is an API client for {{.NameTitle}}.
type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

// New creates a new {{.NameTitle}} API client.
func New(ctx context.Context, apiKey, baseURL string) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("client: API key is required")
	}
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	return &Client{
		httpClient: &http.Client{},
		baseURL:    baseURL,
		apiKey:     apiKey,
	}, nil
}

// TODO: Implement API methods
// func (c *Client) ListUsers(ctx context.Context, pageToken string) ([]*User, string, error) { ... }
// func (c *Client) ListGroups(ctx context.Context, pageToken string) ([]*Group, string, error) { ... }
// func (c *Client) ListRoles(ctx context.Context, pageToken string) ([]*Role, string, error) { ... }
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

- Go 1.21+
- {{.NameTitle}} API key

## Installation

` + "```" + `bash
go install {{.ModulePath}}@latest
` + "```" + `

## Usage

` + "```" + `bash
# Set credentials
export BATON_API_KEY="your-api-key"

# Run sync
baton-{{.Name}}

# Or use flags
baton-{{.Name}} --api-key "your-api-key"
` + "```" + `

## Development

` + "```" + `bash
# Build
go build -o baton-{{.Name}} .

# Run locally
./baton-{{.Name}} --api-key "your-api-key"

# Run with hot reload (using cone)
cone connector dev
` + "```" + `

## Resources

This connector syncs the following resources:

| Resource Type | Description |
|---------------|-------------|
| User | {{.NameTitle}} users |
| Group | {{.NameTitle}} groups |
| Role | {{.NameTitle}} roles |

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
