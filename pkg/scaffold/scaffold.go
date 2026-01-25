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
	github.com/conductorone/baton-sdk v0.7.1
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
//
// Required permissions for sync operations:
// - Read access to {{.NameTitle}} users, groups, and roles
//
// Required permissions for provisioning operations:
// - Write access to create/modify/delete users and group memberships
type Config struct {
	// BaseURL is the API base URL. Override for testing against mocks.
	BaseURL string ` + "`" + `mapstructure:"base-url"` + "`" + `
	// Insecure skips TLS certificate verification. ONLY use for testing.
	Insecure bool ` + "`" + `mapstructure:"insecure"` + "`" + `
	// Add connector-specific fields here as needed.
	// Example: APIKey string ` + "`" + `mapstructure:"api-key"` + "`" + `
}

// Implement field.Configurable interface.
// These methods allow the SDK to read configuration values.
// Each method should return the appropriate value for the given key.
func (c *Config) GetString(key string) string {
	switch key {
	case "base-url":
		return c.BaseURL
	default:
		return ""
	}
}

func (c *Config) GetBool(key string) bool {
	switch key {
	case "insecure":
		return c.Insecure
	default:
		return false
	}
}

func (c *Config) GetInt(key string) int               { return 0 }
func (c *Config) GetStringSlice(key string) []string  { return nil }
func (c *Config) GetStringMap(key string) map[string]any { return nil }

// Configuration fields for the connector.
// These define CLI flags and environment variables.
var configFields = []field.SchemaField{
	// Testability: Allow overriding base URL for mock servers
	field.StringField(
		"base-url",
		field.WithDescription("Base URL for the {{.NameTitle}} API (override for testing)"),
		field.WithDefaultValue("https://api.example.com"), // TODO: Set your API's default URL
	),
	// Testability: Allow skipping TLS verification for self-signed certs
	field.BoolField(
		"insecure",
		field.WithDescription("Skip TLS certificate verification (for testing only - DO NOT USE IN PRODUCTION)"),
		field.WithDefaultValue(false),
	),
	// TODO: Add your connector-specific fields here, e.g.:
	// field.StringField("api-key", field.WithRequired(true), field.WithDescription("API key for authentication")),
}

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

	// Log warning if insecure mode is enabled
	if cfg.Insecure {
		l.Warn("baton-{{.Name}}: TLS certificate verification disabled - DO NOT USE IN PRODUCTION")
	}

	cb, err := connector.New(ctx, cfg.BaseURL, cfg.Insecure)
	if err != nil {
		l.Error("baton-{{.Name}}: error creating connector", zap.Error(err))
		return nil, fmt.Errorf("baton-{{.Name}}: failed to create connector: %w", err)
	}

	// IMPORTANT: connectorbuilder.NewConnector wraps your connector with SDK infrastructure.
	// This is REQUIRED - without it, the connector won't function.
	// The wrapper provides: gRPC server, sync orchestration, pagination handling.
	c, err := connectorbuilder.NewConnector(ctx, cb)
	if err != nil {
		l.Error("baton-{{.Name}}: error wrapping connector", zap.Error(err))
		return nil, fmt.Errorf("baton-{{.Name}}: failed to initialize connector (connectorbuilder.NewConnector failed): %w", err)
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
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

// Connector implements the {{.Name}} connector.
type Connector struct {
	baseURL    string
	httpClient *http.Client
	// TODO: Add API client or other state here.
	// Example: client *{{.NamePascal}}Client
}

// ResourceSyncers returns a ResourceSyncer for each resource type that should be synced.
//
// Resource types:
// - user: Users from {{.NameTitle}} (principals that receive grants)
// - group: Groups with "member" entitlement
// - role: Roles with "assigned" entitlement
//
// The three fundamental resource types are:
// 1. Users - principals that can be granted access
// 2. Groups - collections of users with membership entitlement
// 3. Roles - permissions that can be assigned to users
func (c *Connector) ResourceSyncers(ctx context.Context) []connectorbuilder.ResourceSyncer {
	return []connectorbuilder.ResourceSyncer{
		newUserBuilder(c),
		newGroupBuilder(c),
		newRoleBuilder(c),
	}
}

// Asset takes an input AssetRef and attempts to fetch it.
// Most connectors don't need to implement this.
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
// This runs before every sync to fail fast on bad credentials.
//
// Required permissions:
// - Read access to {{.NameTitle}} API (basic endpoint like /users or /me)
func (c *Connector) Validate(ctx context.Context) (annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)
	l.Debug("baton-{{.Name}}: validating connection")

	// TODO: Implement validation - test API connection
	// Example:
	//   _, err := c.client.GetCurrentUser(ctx)
	//   if err != nil {
	//       return nil, fmt.Errorf("baton-{{.Name}}: validation failed: %w", err)
	//   }

	l.Info("baton-{{.Name}}: connection validated")
	return nil, nil
}

// New returns a new instance of the connector.
//
// Parameters:
// - baseURL: API base URL (can be overridden for testing)
// - insecure: Skip TLS verification (for testing with self-signed certs)
func New(ctx context.Context, baseURL string, insecure bool) (*Connector, error) {
	l := ctxzap.Extract(ctx)

	// Configure HTTP client with optional insecure TLS
	httpClient := &http.Client{}
	if insecure {
		l.Warn("baton-{{.Name}}: TLS certificate verification disabled")
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, //nolint:gosec // Intentional for testing
			},
		}
	}

	l.Info("baton-{{.Name}}: creating connector",
		zap.String("base_url", baseURL),
		zap.Bool("insecure", insecure),
	)

	// TODO: Create API client
	// Example:
	//   client, err := NewClient(baseURL, httpClient)
	//   if err != nil {
	//       return nil, fmt.Errorf("baton-{{.Name}}: failed to create client: %w", err)
	//   }

	return &Connector{
		baseURL:    baseURL,
		httpClient: httpClient,
	}, nil
}
`,
	},
	{
		Path: "pkg/connector/resource_types.go",
		Template: `package connector

import (
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
)

// Resource type definitions for {{.NameTitle}}.
// Each resource type maps to an entity in the target system.

// userResourceType defines the user resource type.
// Users are principals that can receive grants to entitlements.
var userResourceType = &v2.ResourceType{
	Id:          "user",
	DisplayName: "User",
	Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_USER},
}

// groupResourceType defines the group resource type.
// Groups have a "member" entitlement that users can be granted.
var groupResourceType = &v2.ResourceType{
	Id:          "group",
	DisplayName: "Group",
	Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_GROUP},
}

// roleResourceType defines the role resource type.
// Roles have an "assigned" entitlement that users can be granted.
var roleResourceType = &v2.ResourceType{
	Id:          "role",
	DisplayName: "Role",
	Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_ROLE},
}
`,
	},
	{
		Path: "pkg/connector/users.go",
		Template: `package connector

import (
	"context"
	"fmt"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	rs "github.com/conductorone/baton-sdk/pkg/types/resource"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

type userBuilder struct {
	conn *Connector
}

func (u *userBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return userResourceType
}

// List returns all the users from the upstream service as resource objects.
//
// Required permissions:
// - Read access to users in {{.NameTitle}}
func (u *userBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)
	l.Debug("baton-{{.Name}}: listing users")

	// TODO: Implement user listing with pagination
	// Example:
	//
	// page := ""
	// if pToken != nil && pToken.Token != "" {
	//     page = pToken.Token
	// }
	//
	// users, nextPage, err := u.conn.client.ListUsers(ctx, page, 100)
	// if err != nil {
	//     return nil, "", nil, fmt.Errorf("baton-{{.Name}}: failed to list users: %w", err)
	// }
	//
	// var rv []*v2.Resource
	// for _, user := range users {
	//     // Check context for cancellation in loops
	//     select {
	//     case <-ctx.Done():
	//         return nil, "", nil, fmt.Errorf("baton-{{.Name}}: context cancelled: %w", ctx.Err())
	//     default:
	//     }
	//
	//     displayName := user.Name
	//     if displayName == "" {
	//         displayName = user.Email // Fall back to email if no name
	//     }
	//
	//     resource, err := rs.NewUserResource(
	//         displayName,
	//         userResourceType,
	//         user.ID,
	//         []rs.UserTraitOption{
	//             rs.WithEmail(user.Email, true),
	//             rs.WithUserLogin(user.Username),
	//             rs.WithStatus(v2.UserTrait_Status_STATUS_ENABLED),
	//         },
	//         // ExternalId is CRITICAL for provisioning - stores native identifier
	//         rs.WithExternalID(&v2.ExternalId{
	//             Id:   user.ID,
	//             Link: fmt.Sprintf("%s/users/%s", u.conn.baseURL, user.ID),
	//         }),
	//     )
	//     if err != nil {
	//         return nil, "", nil, fmt.Errorf("baton-{{.Name}}: failed to create user resource: %w", err)
	//     }
	//     rv = append(rv, resource)
	// }
	//
	// l.Info("baton-{{.Name}}: listed users", zap.Int("count", len(rv)))
	// return rv, nextPage, nil, nil

	// Placeholder - remove after implementing
	_ = rs.NewUserResource
	_ = fmt.Sprintf
	l.Info("baton-{{.Name}}: user listing not implemented yet")
	return nil, "", nil, nil
}

// Entitlements always returns an empty slice for users.
// Users are principals that receive grants, not resources with grantable permissions.
func (u *userBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

// Grants always returns an empty slice for users.
// Grants flow from entitlements to users, not from users.
func (u *userBuilder) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

// =============================================================================
// PROVISIONING: Create/Delete Users (ResourceManager interface)
// =============================================================================
// Uncomment and implement these methods to support user lifecycle management.
//
// func (u *userBuilder) Create(ctx context.Context, resource *v2.Resource) (*v2.Resource, annotations.Annotations, error) {
//     l := ctxzap.Extract(ctx)
//     l.Info("baton-{{.Name}}: creating user")
//
//     // Get user traits from the resource
//     userTrait, err := rs.GetUserTrait(resource)
//     if err != nil {
//         return nil, nil, fmt.Errorf("baton-{{.Name}}: failed to get user trait: %w", err)
//     }
//
//     // TODO: Create user in upstream system
//     // newUser, err := u.conn.client.CreateUser(ctx, &CreateUserRequest{
//     //     Email:    userTrait.GetEmail().GetAddress(),
//     //     Username: userTrait.GetLogin(),
//     // })
//     // if err != nil {
//     //     return nil, nil, fmt.Errorf("baton-{{.Name}}: failed to create user: %w", err)
//     // }
//
//     // Return the created resource with its new ID
//     // return rs.NewUserResource(newUser.Name, userResourceType, newUser.ID, ...)
//     return nil, nil, fmt.Errorf("baton-{{.Name}}: user creation not implemented")
// }
//
// func (u *userBuilder) Delete(ctx context.Context, resourceId *v2.ResourceId) (annotations.Annotations, error) {
//     l := ctxzap.Extract(ctx)
//     l.Info("baton-{{.Name}}: deleting user", zap.String("id", resourceId.Resource))
//
//     // TODO: Delete user from upstream system
//     // err := u.conn.client.DeleteUser(ctx, resourceId.Resource)
//     // if err != nil {
//     //     return nil, fmt.Errorf("baton-{{.Name}}: failed to delete user: %w", err)
//     // }
//     // return nil, nil
//     return nil, fmt.Errorf("baton-{{.Name}}: user deletion not implemented")
// }

func newUserBuilder(conn *Connector) *userBuilder {
	return &userBuilder{conn: conn}
}
`,
	},
	{
		Path: "pkg/connector/groups.go",
		Template: `package connector

import (
	"context"
	"fmt"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	ent "github.com/conductorone/baton-sdk/pkg/types/entitlement"
	"github.com/conductorone/baton-sdk/pkg/types/grant"
	rs "github.com/conductorone/baton-sdk/pkg/types/resource"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

const memberEntitlement = "member"

type groupBuilder struct {
	conn *Connector
}

func (g *groupBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return groupResourceType
}

// List returns all groups from the upstream service.
//
// Required permissions:
// - Read access to groups in {{.NameTitle}}
func (g *groupBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)
	l.Debug("baton-{{.Name}}: listing groups")

	// TODO: Implement group listing with pagination
	// Example:
	//
	// page := ""
	// if pToken != nil && pToken.Token != "" {
	//     page = pToken.Token
	// }
	//
	// groups, nextPage, err := g.conn.client.ListGroups(ctx, page, 100)
	// if err != nil {
	//     return nil, "", nil, fmt.Errorf("baton-{{.Name}}: failed to list groups: %w", err)
	// }
	//
	// var rv []*v2.Resource
	// for _, group := range groups {
	//     resource, err := rs.NewGroupResource(
	//         group.Name,
	//         groupResourceType,
	//         group.ID,
	//         []rs.GroupTraitOption{
	//             rs.WithGroupProfile(map[string]interface{}{
	//                 "description": group.Description,
	//             }),
	//         },
	//         rs.WithExternalID(&v2.ExternalId{Id: group.ID}),
	//     )
	//     if err != nil {
	//         return nil, "", nil, fmt.Errorf("baton-{{.Name}}: failed to create group resource: %w", err)
	//     }
	//     rv = append(rv, resource)
	// }
	//
	// l.Info("baton-{{.Name}}: listed groups", zap.Int("count", len(rv)))
	// return rv, nextPage, nil, nil

	_ = rs.NewGroupResource
	_ = fmt.Sprintf
	l.Info("baton-{{.Name}}: group listing not implemented yet")
	return nil, "", nil, nil
}

// Entitlements returns the "member" entitlement for the group.
// This entitlement can be granted to users to make them members of the group.
func (g *groupBuilder) Entitlements(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	// Create the "member" entitlement for this group
	entitlement := ent.NewAssignmentEntitlement(
		resource,
		memberEntitlement,
		ent.WithGrantableTo(userResourceType),
		ent.WithDisplayName(fmt.Sprintf("%s Group Member", resource.DisplayName)),
		ent.WithDescription(fmt.Sprintf("Member of the %s group", resource.DisplayName)),
	)

	return []*v2.Entitlement{entitlement}, "", nil, nil
}

// Grants returns all users who are members of this group.
//
// Required permissions:
// - Read access to group memberships in {{.NameTitle}}
func (g *groupBuilder) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)
	l.Debug("baton-{{.Name}}: listing group members", zap.String("group_id", resource.Id.Resource))

	// TODO: Implement group membership listing
	// Example:
	//
	// groupID := resource.Id.Resource
	//
	// page := ""
	// if pToken != nil && pToken.Token != "" {
	//     page = pToken.Token
	// }
	//
	// members, nextPage, err := g.conn.client.ListGroupMembers(ctx, groupID, page, 100)
	// if err != nil {
	//     return nil, "", nil, fmt.Errorf("baton-{{.Name}}: failed to list group members: %w", err)
	// }
	//
	// var rv []*v2.Grant
	// for _, member := range members {
	//     grant := grant.NewGrant(
	//         resource,
	//         memberEntitlement,
	//         &v2.ResourceId{
	//             ResourceType: userResourceType.Id,
	//             Resource:     member.UserID,
	//         },
	//     )
	//     rv = append(rv, grant)
	// }
	//
	// l.Info("baton-{{.Name}}: listed group members", zap.Int("count", len(rv)))
	// return rv, nextPage, nil, nil

	_ = grant.NewGrant
	l.Info("baton-{{.Name}}: group grants not implemented yet")
	return nil, "", nil, nil
}

// =============================================================================
// PROVISIONING: Grant/Revoke group membership (ResourceProvisioner interface)
// =============================================================================
// Uncomment to support group membership provisioning.
//
// func (g *groupBuilder) Grant(ctx context.Context, principal *v2.Resource, entitlement *v2.Entitlement) (annotations.Annotations, error) {
//     l := ctxzap.Extract(ctx)
//     groupID := entitlement.Resource.Id.Resource
//     userID := principal.Id.Resource
//     l.Info("baton-{{.Name}}: granting group membership", zap.String("group", groupID), zap.String("user", userID))
//     // TODO: Add user to group in upstream system
//     // err := g.conn.client.AddGroupMember(ctx, groupID, userID)
//     // if err != nil {
//     //     return nil, fmt.Errorf("baton-{{.Name}}: failed to grant membership: %w", err)
//     // }
//     // return nil, nil
//     return nil, fmt.Errorf("baton-{{.Name}}: grant not implemented")
// }
//
// func (g *groupBuilder) Revoke(ctx context.Context, grantToRevoke *v2.Grant) (annotations.Annotations, error) {
//     l := ctxzap.Extract(ctx)
//     groupID := grantToRevoke.Entitlement.Resource.Id.Resource
//     userID := grantToRevoke.Principal.Id.Resource
//     l.Info("baton-{{.Name}}: revoking group membership", zap.String("group", groupID), zap.String("user", userID))
//     // TODO: Remove user from group in upstream system
//     // err := g.conn.client.RemoveGroupMember(ctx, groupID, userID)
//     // if err != nil {
//     //     return nil, fmt.Errorf("baton-{{.Name}}: failed to revoke membership: %w", err)
//     // }
//     // return nil, nil
//     return nil, fmt.Errorf("baton-{{.Name}}: revoke not implemented")
// }

func newGroupBuilder(conn *Connector) *groupBuilder {
	return &groupBuilder{conn: conn}
}
`,
	},
	{
		Path: "pkg/connector/roles.go",
		Template: `package connector

import (
	"context"
	"fmt"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	ent "github.com/conductorone/baton-sdk/pkg/types/entitlement"
	"github.com/conductorone/baton-sdk/pkg/types/grant"
	rs "github.com/conductorone/baton-sdk/pkg/types/resource"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

const assignedEntitlement = "assigned"

type roleBuilder struct {
	conn *Connector
}

func (r *roleBuilder) ResourceType(ctx context.Context) *v2.ResourceType {
	return roleResourceType
}

// List returns all roles from the upstream service.
//
// Required permissions:
// - Read access to roles in {{.NameTitle}}
func (r *roleBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)
	l.Debug("baton-{{.Name}}: listing roles")

	// TODO: Implement role listing with pagination
	// Example:
	//
	// page := ""
	// if pToken != nil && pToken.Token != "" {
	//     page = pToken.Token
	// }
	//
	// roles, nextPage, err := r.conn.client.ListRoles(ctx, page, 100)
	// if err != nil {
	//     return nil, "", nil, fmt.Errorf("baton-{{.Name}}: failed to list roles: %w", err)
	// }
	//
	// var rv []*v2.Resource
	// for _, role := range roles {
	//     resource, err := rs.NewRoleResource(
	//         role.Name,
	//         roleResourceType,
	//         role.ID,
	//         []rs.RoleTraitOption{
	//             rs.WithRoleProfile(map[string]interface{}{
	//                 "description": role.Description,
	//                 "permissions": role.Permissions,
	//             }),
	//         },
	//         rs.WithExternalID(&v2.ExternalId{Id: role.ID}),
	//     )
	//     if err != nil {
	//         return nil, "", nil, fmt.Errorf("baton-{{.Name}}: failed to create role resource: %w", err)
	//     }
	//     rv = append(rv, resource)
	// }
	//
	// l.Info("baton-{{.Name}}: listed roles", zap.Int("count", len(rv)))
	// return rv, nextPage, nil, nil

	_ = rs.NewRoleResource
	_ = fmt.Sprintf
	l.Info("baton-{{.Name}}: role listing not implemented yet")
	return nil, "", nil, nil
}

// Entitlements returns the "assigned" entitlement for the role.
// This entitlement can be granted to users to assign them the role.
func (r *roleBuilder) Entitlements(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	// Create the "assigned" entitlement for this role
	entitlement := ent.NewAssignmentEntitlement(
		resource,
		assignedEntitlement,
		ent.WithGrantableTo(userResourceType),
		ent.WithDisplayName(fmt.Sprintf("%s Role", resource.DisplayName)),
		ent.WithDescription(fmt.Sprintf("Assigned the %s role", resource.DisplayName)),
	)

	return []*v2.Entitlement{entitlement}, "", nil, nil
}

// Grants returns all users who are assigned this role.
//
// Required permissions:
// - Read access to role assignments in {{.NameTitle}}
func (r *roleBuilder) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)
	l.Debug("baton-{{.Name}}: listing role assignments", zap.String("role_id", resource.Id.Resource))

	// TODO: Implement role assignment listing
	// Example:
	//
	// roleID := resource.Id.Resource
	//
	// page := ""
	// if pToken != nil && pToken.Token != "" {
	//     page = pToken.Token
	// }
	//
	// assignments, nextPage, err := r.conn.client.ListRoleAssignments(ctx, roleID, page, 100)
	// if err != nil {
	//     return nil, "", nil, fmt.Errorf("baton-{{.Name}}: failed to list role assignments: %w", err)
	// }
	//
	// var rv []*v2.Grant
	// for _, assignment := range assignments {
	//     grant := grant.NewGrant(
	//         resource,
	//         assignedEntitlement,
	//         &v2.ResourceId{
	//             ResourceType: userResourceType.Id,
	//             Resource:     assignment.UserID,
	//         },
	//     )
	//     rv = append(rv, grant)
	// }
	//
	// l.Info("baton-{{.Name}}: listed role assignments", zap.Int("count", len(rv)))
	// return rv, nextPage, nil, nil

	_ = grant.NewGrant
	l.Info("baton-{{.Name}}: role grants not implemented yet")
	return nil, "", nil, nil
}

// =============================================================================
// PROVISIONING: Grant/Revoke role assignment (ResourceProvisioner interface)
// =============================================================================
// Uncomment to support role assignment provisioning.
//
// func (r *roleBuilder) Grant(ctx context.Context, principal *v2.Resource, entitlement *v2.Entitlement) (annotations.Annotations, error) {
//     l := ctxzap.Extract(ctx)
//     roleID := entitlement.Resource.Id.Resource
//     userID := principal.Id.Resource
//     l.Info("baton-{{.Name}}: granting role", zap.String("role", roleID), zap.String("user", userID))
//     // TODO: Assign role to user in upstream system
//     return nil, fmt.Errorf("baton-{{.Name}}: grant not implemented")
// }
//
// func (r *roleBuilder) Revoke(ctx context.Context, grantToRevoke *v2.Grant) (annotations.Annotations, error) {
//     l := ctxzap.Extract(ctx)
//     roleID := grantToRevoke.Entitlement.Resource.Id.Resource
//     userID := grantToRevoke.Principal.Id.Resource
//     l.Info("baton-{{.Name}}: revoking role", zap.String("role", roleID), zap.String("user", userID))
//     // TODO: Unassign role from user in upstream system
//     return nil, fmt.Errorf("baton-{{.Name}}: revoke not implemented")
// }

func newRoleBuilder(conn *Connector) *roleBuilder {
	return &roleBuilder{conn: conn}
}
`,
	},
	{
		Path: "pkg/connector/actions.go",
		Template: `package connector

import (
	"context"
	"fmt"

	config "github.com/conductorone/baton-sdk/pb/c1/config/v1"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/actions"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"
)

// =============================================================================
// BATON ACTIONS: Custom operations exposed to ConductorOne
// =============================================================================
//
// Actions are arbitrary operations your connector can perform. Unlike Grant/Revoke
// (which modify access) or Create/Delete (which manage resources), Actions are
// general-purpose operations that ConductorOne can trigger.
//
// Common action types:
// - ACTION_TYPE_ACCOUNT_ENABLE / ACTION_TYPE_ACCOUNT_DISABLE
// - ACTION_TYPE_ACCOUNT_UPDATE_PROFILE
// - Custom operations specific to your system
//
// To enable actions, uncomment GlobalActions and the action handlers below.

// Example: Disable account action schema
var disableAccountAction = &v2.BatonActionSchema{
	Name: "disableAccount",
	Arguments: []*config.Field{
		{
			Name:        "accountId",
			DisplayName: "Account ID",
			Description: "The ID of the account to disable",
			Field:       &config.Field_StringField{},
			IsRequired:  true,
		},
	},
	ReturnTypes: []*config.Field{
		{
			Name:        "success",
			DisplayName: "Success",
			Field:       &config.Field_BoolField{},
		},
	},
	ActionType: []v2.ActionType{
		v2.ActionType_ACTION_TYPE_ACCOUNT,
		v2.ActionType_ACTION_TYPE_ACCOUNT_DISABLE,
	},
}

// Example: Enable account action schema
var enableAccountAction = &v2.BatonActionSchema{
	Name: "enableAccount",
	Arguments: []*config.Field{
		{
			Name:        "accountId",
			DisplayName: "Account ID",
			Description: "The ID of the account to enable",
			Field:       &config.Field_StringField{},
			IsRequired:  true,
		},
	},
	ReturnTypes: []*config.Field{
		{
			Name:        "success",
			DisplayName: "Success",
			Field:       &config.Field_BoolField{},
		},
	},
	ActionType: []v2.ActionType{
		v2.ActionType_ACTION_TYPE_ACCOUNT,
		v2.ActionType_ACTION_TYPE_ACCOUNT_ENABLE,
	},
}

// GlobalActions registers custom actions with the SDK.
// Uncomment to enable actions.
//
// func (c *Connector) GlobalActions(ctx context.Context, registry actions.ActionRegistry) error {
//     if err := registry.Register(ctx, disableAccountAction, c.disableAccount); err != nil {
//         return err
//     }
//     if err := registry.Register(ctx, enableAccountAction, c.enableAccount); err != nil {
//         return err
//     }
//     return nil
// }

func (c *Connector) disableAccount(ctx context.Context, args *structpb.Struct) (*structpb.Struct, annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)

	accountId, ok := args.Fields["accountId"]
	if !ok {
		return nil, nil, fmt.Errorf("missing required argument accountId")
	}

	l.Info("baton-{{.Name}}: disabling account", zap.String("accountId", accountId.GetStringValue()))

	// TODO: Implement account disabling in upstream system
	// err := c.client.DisableUser(ctx, accountId.GetStringValue())
	// if err != nil {
	//     return nil, nil, fmt.Errorf("baton-{{.Name}}: failed to disable account: %w", err)
	// }

	response := &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"success": structpb.NewBoolValue(true),
		},
	}
	return response, nil, nil
}

func (c *Connector) enableAccount(ctx context.Context, args *structpb.Struct) (*structpb.Struct, annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)

	accountId, ok := args.Fields["accountId"]
	if !ok {
		return nil, nil, fmt.Errorf("missing required argument accountId")
	}

	l.Info("baton-{{.Name}}: enabling account", zap.String("accountId", accountId.GetStringValue()))

	// TODO: Implement account enabling in upstream system
	// err := c.client.EnableUser(ctx, accountId.GetStringValue())
	// if err != nil {
	//     return nil, nil, fmt.Errorf("baton-{{.Name}}: failed to enable account: %w", err)
	// }

	response := &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"success": structpb.NewBoolValue(true),
		},
	}
	return response, nil, nil
}

// Ensure imports are used (remove after implementing)
var _ = actions.ActionRegistry(nil)
var _ = disableAccountAction
var _ = enableAccountAction
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
		Path: ".env.example",
		Template: `# {{.NameTitle}} Connector Configuration
# Copy this file to .env and fill in your values
# All variables use the BATON_ prefix

# =============================================================================
# Target System Authentication
# =============================================================================

# API key for {{.NameTitle}} (if using API key auth)
# BATON_API_KEY=your-api-key-here

# Bearer token for {{.NameTitle}} (if using bearer auth)
# BATON_BEARER_TOKEN=your-bearer-token-here

# =============================================================================
# ConductorOne Authentication (for daemon mode)
# =============================================================================

# Client credentials for ConductorOne integration
# Get these from ConductorOne admin console
# BATON_CLIENT_ID=your-client-id
# BATON_CLIENT_SECRET=your-client-secret

# =============================================================================
# Testing and Development
# =============================================================================

# Override base URL (for testing against mocks)
# BATON_BASE_URL=http://localhost:8089

# Skip TLS verification (ONLY for local testing with self-signed certs)
# BATON_INSECURE=true

# =============================================================================
# Logging and Observability
# =============================================================================

# Log level: debug, info, warn, error
BATON_LOG_LEVEL=info

# Log format: json, console
BATON_LOG_FORMAT=json

# OpenTelemetry collector endpoint (optional)
# BATON_OTEL_COLLECTOR_ENDPOINT=http://localhost:4317
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

## Configuration

Copy ` + "`.env.example`" + ` to ` + "`.env`" + ` and fill in your values:

` + "```" + `bash
cp .env.example .env
# Edit .env with your credentials
` + "```" + `

Or pass configuration via CLI flags:

` + "```" + `bash
baton-{{.Name}} --api-key=your-key
` + "```" + `

See all options with ` + "`baton-{{.Name}} --help`" + `.

## Usage

` + "```" + `bash
# Run sync (outputs to sync.c1z)
baton-{{.Name}}

# Run with specific output file
baton-{{.Name}} -f output.c1z

# See all options
baton-{{.Name}} --help
` + "```" + `

## Development

` + "```" + `bash
# Build
make build

# Format, vet, and build
make check

# Run all validations (tidy, fmt, vet, lint, test, build)
make all

# Run tests
make test

# Run tests with coverage
make test-cover

# Format code
make fmt

# Run linter
make lint
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
		Template: `.PHONY: build test test-mock clean lint fmt vet tidy check all

BINARY_NAME=baton-{{.Name}}
MOCK_PORT?=8089

# Build the connector binary
build:
	go build -o $(BINARY_NAME) .

# Run unit tests
test:
	go test -v ./...

# Run tests with coverage
test-cover:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run integration tests against a mock server
# Assumes mock server is running on localhost:$(MOCK_PORT)
test-mock: build
	./$(BINARY_NAME) --base-url=http://localhost:$(MOCK_PORT) --insecure

# Remove build artifacts
clean:
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html
	rm -rf dist/

# Run golangci-lint
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Run go vet
vet:
	go vet ./...

# Tidy and verify dependencies
tidy:
	go mod tidy
	go mod verify

# Quick check: fmt, vet, build
check: fmt vet build

# Full validation: tidy, fmt, vet, lint, test, build
all: tidy fmt vet lint test build

.DEFAULT_GOAL := build
`,
	},
	{
		Path: "CLAUDE.md",
		Template: `# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with this connector.

## Project Overview

This is a ConductorOne Baton connector for {{.NameTitle}}. It syncs identity and access data from {{.NameTitle}} into ConductorOne.

## Build and Run

` + "```" + `bash
# Build
go build -o baton-{{.Name}} .

# Run sync
./baton-{{.Name}}

# Run with verbose logging
./baton-{{.Name}} --log-level debug
` + "```" + `

## Standard Connector Structure

` + "```" + `
baton-{{.Name}}/
    main.go                     # Entry point: config, connector init
    pkg/connector/
        connector.go            # Metadata, ResourceSyncers(), Validate()
        resource_types.go       # Resource type definitions
        users.go                # User resource syncer
        groups.go               # Group resource syncer (add as needed)
    pkg/client/                 # Optional: API wrapper
    CLAUDE.md                   # This file
` + "```" + `

## Key Patterns

### 1. ResourceSyncer Interface

Every resource type implements:
- ` + "`" + `List()` + "`" + ` - Return all resources (with pagination)
- ` + "`" + `Entitlements()` + "`" + ` - Return available permissions for a resource
- ` + "`" + `Grants()` + "`" + ` - Return who has what permissions

### 2. Error Wrapping

Always prefix errors with connector name:
` + "```" + `go
return nil, fmt.Errorf("baton-{{.Name}}: failed to list users: %w", err)
` + "```" + `

### 3. Pagination

Always paginate API calls:
` + "```" + `go
func (u *userBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
    // Get page from token
    page, _ := parsePageToken(pToken.Token)

    // Fetch one page
    users, nextCursor, err := u.client.ListUsers(ctx, page, pageSize)

    // Return next token
    return resources, nextCursor, nil, nil
}
` + "```" + `

### 4. User Resource Creation

` + "```" + `go
import rs "github.com/conductorone/baton-sdk/pkg/types/resource"

resource, err := rs.NewUserResource(
    user.DisplayName,
    userResourceType,
    user.ID,
    []rs.UserTraitOption{
        rs.WithEmail(user.Email, true),
        rs.WithUserLogin(user.Username),
        rs.WithStatus(v2.UserTrait_Status_STATUS_ENABLED),
    },
)
` + "```" + `

## Testing Requirements

### Configurable Base URL

The connector MUST support ` + "`" + `--base-url` + "`" + ` flag for testing against mocks:
` + "```" + `go
field.StringField("base-url",
    field.WithDescription("Base URL for API (for testing)"),
    field.WithDefaultValue("https://api.example.com"),
),
` + "```" + `

### Insecure TLS Option

Support ` + "`" + `--insecure` + "`" + ` for self-signed certs in testing:
` + "```" + `go
field.BoolField("insecure",
    field.WithDescription("Skip TLS verification (testing only)"),
),
` + "```" + `

## Common Pitfalls

1. **Don't swallow errors** - Always return errors, don't log and continue
2. **Don't buffer entire datasets** - Always paginate
3. **Don't ignore context** - Pass ctx to all API calls
4. **Don't log credentials** - Never log tokens or API keys
5. **Don't use empty display names** - Fall back to ID if name is empty

## Work Tracking

Track work in TODO.md:
- Create TODO.md for pending tasks
- Move completed items to COMPLETED section
- Add QUESTIONS section for clarifications needed

## Reference

For comprehensive patterns and best practices, see:
- baton-demo: https://github.com/conductorone/baton-demo
- baton-github: https://github.com/conductorone/baton-github
- baton-sdk docs: https://github.com/conductorone/baton-sdk
`,
	},
	{
		Path: "docs/README.md",
		Template: `# Documentation

Place documentation about the downstream {{.NameTitle}} API here.

This helps Claude Code understand the API and build appropriate mocks for testing.

## Suggested Content

1. **API Authentication** - How to authenticate (API key, OAuth, etc.)
2. **Endpoints** - Key endpoints the connector needs
3. **Data Models** - User, group, role structures
4. **Rate Limits** - API rate limiting behavior
5. **Pagination** - How the API paginates results

## Example

` + "```" + `
GET /api/v1/users
Authorization: Bearer {token}

Response:
{
  "users": [...],
  "next_cursor": "abc123"
}
` + "```" + `
`,
	},
	{
		Path: "docs/.gitignore",
		Template: `# Ignore everything in docs except README and API_NOTES
*
!.gitignore
!README.md
!API_NOTES.md
`,
	},
	{
		Path: "docs/API_NOTES.md",
		Template: `# {{.NameTitle}} API Notes

This document captures API behavior, quirks, and implementation notes discovered during connector development.

## Authentication

` + "```" + `
# TODO: Document authentication method
# Example:
# Authorization: Bearer {api_key}
# X-API-Key: {api_key}
` + "```" + `

## Pagination

` + "```" + `
# TODO: Document pagination pattern
# Example cursor-based:
# GET /users?cursor=abc123&limit=100
# Response: { "users": [...], "next_cursor": "def456" }
#
# Example offset-based:
# GET /users?page=2&per_page=100
# Response: { "users": [...], "total": 1234 }
` + "```" + `

## Rate Limits

` + "```" + `
# TODO: Document rate limits
# Example:
# X-RateLimit-Limit: 1000
# X-RateLimit-Remaining: 999
# X-RateLimit-Reset: 1234567890
` + "```" + `

## Key Endpoints

### Users

` + "```" + `
# List users
GET /api/v1/users

# Get single user
GET /api/v1/users/{id}

# Response shape
{
  "id": "user-123",
  "email": "user@example.com",
  "name": "Display Name",
  "status": "active"
}
` + "```" + `

### Groups (if applicable)

` + "```" + `
# TODO: Document group endpoints
` + "```" + `

### Roles/Permissions (if applicable)

` + "```" + `
# TODO: Document role/permission endpoints
` + "```" + `

## Quirks and Gotchas

- TODO: Document any API quirks discovered during implementation
- Example: "User IDs are case-sensitive"
- Example: "Empty arrays are returned as null, not []"
- Example: "Deleted users still appear in list with status=deleted"

## Mock Server Notes

When building a mock server for testing, ensure it:
1. Returns proper pagination tokens
2. Handles the authentication header
3. Returns realistic response shapes

See ` + "`" + `mocks/` + "`" + ` directory for mock server implementation.
`,
	},
	{
		Path: ".github/workflows/ci.yaml",
		Template: `name: ci

on:
  pull_request:
    types: [opened, reopened, synchronize]
  push:
    branches:
      - main

jobs:
  go-lint:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
      - name: Install Go
        uses: actions/setup-go@v5
        with:
          go-version-file: 'go.mod'
      - name: Run linters
        uses: golangci/golangci-lint-action@v8
        with:
          version: latest
          args: --timeout=3m

  go-test:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
      - name: Install Go
        uses: actions/setup-go@v5
        with:
          go-version-file: 'go.mod'
      - name: Run tests
        run: go test -v -covermode=count ./...

  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
      - name: Install Go
        uses: actions/setup-go@v5
        with:
          go-version-file: 'go.mod'
      - name: Build connector
        run: go build -o baton-{{.Name}} .
      - name: Verify binary runs
        run: ./baton-{{.Name}} --help
`,
	},
	{
		Path: ".github/workflows/release.yaml",
		Template: `name: Release

on:
  push:
    tags:
      - "v*"

jobs:
  release:
    # Uses ConductorOne's shared release workflow
    # Documentation: https://github.com/ConductorOne/github-workflows
    uses: ConductorOne/github-workflows/.github/workflows/release.yaml@v2
    with:
      tag: ${{"{{"}} github.ref_name {{"}}"}}
      lambda: false  # Set to true if you need Lambda deployment
    secrets:
      RELENG_GITHUB_TOKEN: ${{"{{"}} secrets.RELENG_GITHUB_TOKEN {{"}}"}}
      APPLE_SIGNING_KEY_P12: ${{"{{"}} secrets.APPLE_SIGNING_KEY_P12 {{"}}"}}
      APPLE_SIGNING_KEY_P12_PASSWORD: ${{"{{"}} secrets.APPLE_SIGNING_KEY_P12_PASSWORD {{"}}"}}
      AC_PASSWORD: ${{"{{"}} secrets.AC_PASSWORD {{"}}"}}
      AC_PROVIDER: ${{"{{"}} secrets.AC_PROVIDER {{"}}"}}
      DATADOG_API_KEY: ${{"{{"}} secrets.DATADOG_API_KEY {{"}}"}}
`,
	},
	{
		Path: ".golangci.yml",
		Template: `version: "2"
linters:
  default: none
  enable:
    - asasalint
    - asciicheck
    - bidichk
    - bodyclose
    - durationcheck
    - errcheck
    - errorlint
    - exhaustive
    - goconst
    - gocritic
    - godot
    - gosec
    - govet
    - ineffassign
    - nakedret
    - nilerr
    - noctx
    - revive
    - staticcheck
    - unconvert
    - unused
    - whitespace
  settings:
    exhaustive:
      default-signifies-exhaustive: true
    govet:
      enable-all: true
      disable:
        - fieldalignment
        - shadow
    nakedret:
      max-func-lines: 0
`,
	},
	{
		Path: "Dockerfile",
		Template: `# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/baton-{{.Name}} .

# Runtime stage - distroless for security
# - No shell, no package manager (minimal attack surface)
# - Runs as nonroot user (uid 65534)
# - Only includes: CA certificates, timezone data
FROM gcr.io/distroless/static-debian11:nonroot

# Copy binary from build stage
COPY --from=builder /build/baton-{{.Name}} /

# Run as nonroot user (distroless default)
USER 65534

# Set entrypoint
ENTRYPOINT ["/baton-{{.Name}}"]

# OCI metadata labels
LABEL org.opencontainers.image.title="baton-{{.Name}}"
LABEL org.opencontainers.image.description="{{.Description}}"
LABEL org.opencontainers.image.source="{{.ModulePath}}"
`,
	},
	{
		Path: "Dockerfile.lambda",
		Template: `# Lambda deployment variant
# Use this for AWS Lambda deployments

# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build with Lambda support tag
RUN CGO_ENABLED=0 GOOS=linux go build -tags baton_lambda_support -o /build/baton-{{.Name}} .

# Runtime stage - AWS Lambda provided runtime
FROM public.ecr.aws/lambda/provided:al2023

# Copy binary
COPY --from=builder /build/baton-{{.Name}} /var/task/

# Lambda entrypoint (note: "lambda" argument triggers Lambda mode)
ENTRYPOINT ["/var/task/baton-{{.Name}}", "lambda"]
`,
	},
	{
		Path: "docker-compose.yml",
		Template: `# Docker Compose for local development and testing
#
# Usage:
#   docker compose up          # Run connector in daemon mode
#   docker compose run baton   # Run one-shot sync
#
version: '3.9'

services:
  baton:
    build: .
    environment:
      # ConductorOne daemon mode credentials
      # Required for long-running daemon mode
      - BATON_CLIENT_ID=${BATON_CLIENT_ID:-}
      - BATON_CLIENT_SECRET=${BATON_CLIENT_SECRET:-}
      # Connector-specific credentials
      # TODO: Add your API credentials here
      # - BATON_API_KEY=${BATON_API_KEY:-}
    # Uncomment for one-shot mode (sync to file):
    # volumes:
    #   - ./output:/work
    # command: ["-f", "/work/sync.c1z"]

  # Optional: Mock server for testing
  # mock:
  #   build:
  #     context: ./mocks
  #   ports:
  #     - "8089:8089"
`,
	},
	{
		Path: "pkg/client/client.go",
		Template: `// Package client provides an HTTP client for the {{.NameTitle}} API.
//
// This package wraps the {{.NameTitle}} REST API with Go types and handles:
// - Authentication
// - Pagination
// - Error handling
// - Rate limiting (optional)
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client wraps the {{.NameTitle}} API.
type Client struct {
	baseURL    string
	httpClient *http.Client
	// TODO: Add authentication fields
	// Example: apiKey string
}

// New creates a new {{.NameTitle}} API client.
//
// Parameters:
// - baseURL: API base URL (e.g., "https://api.example.com")
// - httpClient: HTTP client (can be configured for insecure TLS in tests)
func New(baseURL string, httpClient *http.Client) (*Client, error) {
	if baseURL == "" {
		return nil, fmt.Errorf("baseURL is required")
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Parse and validate base URL
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid baseURL: %w", err)
	}

	return &Client{
		baseURL:    u.String(),
		httpClient: httpClient,
	}, nil
}

// User represents a user from the {{.NameTitle}} API.
// TODO: Update fields to match actual API response.
type User struct {
	ID       string ` + "`" + `json:"id"` + "`" + `
	Email    string ` + "`" + `json:"email"` + "`" + `
	Name     string ` + "`" + `json:"name"` + "`" + `
	Username string ` + "`" + `json:"username,omitempty"` + "`" + `
	Status   string ` + "`" + `json:"status"` + "`" + `
}

// ListUsersResponse is the API response for listing users.
// TODO: Update to match actual API response shape.
type ListUsersResponse struct {
	Users      []User ` + "`" + `json:"users"` + "`" + `
	NextCursor string ` + "`" + `json:"next_cursor,omitempty"` + "`" + `
}

// ListUsers returns a page of users from the API.
//
// Parameters:
// - cursor: Pagination cursor (empty for first page)
// - limit: Maximum users to return per page
//
// Returns:
// - users: List of users
// - nextCursor: Cursor for next page (empty if no more pages)
func (c *Client) ListUsers(ctx context.Context, cursor string, limit int) ([]User, string, error) {
	// Build request URL
	reqURL := fmt.Sprintf("%s/api/v1/users?limit=%d", c.baseURL, limit)
	if cursor != "" {
		reqURL += "&cursor=" + url.QueryEscape(cursor)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create request: %w", err)
	}

	// TODO: Add authentication
	// req.Header.Set("Authorization", "Bearer "+c.apiKey)

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check status
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, "", fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Parse response
	var result ListUsersResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Users, result.NextCursor, nil
}

// GetUser returns a single user by ID.
func (c *Client) GetUser(ctx context.Context, userID string) (*User, error) {
	reqURL := fmt.Sprintf("%s/api/v1/users/%s", c.baseURL, url.PathEscape(userID))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// TODO: Add authentication
	// req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("user not found: %s", userID)
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &user, nil
}
`,
	},
}
