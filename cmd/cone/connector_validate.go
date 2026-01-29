package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func connectorValidateConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate-config <file>",
		Short: "Validate a meta-connector mapping configuration",
		Long: `Validate a mapping configuration file for meta-connectors like baton-openapi.

This checks:
  - Required fields are present
  - Field values are valid
  - At least one TRAIT_USER resource exists
  - Entitlements have grantable_to defined
  - No duplicate resource types`,
		Example: `  cone connector validate-config mapping.yaml
  cone connector validate-config examples/github/mapping.yaml`,
		Args: cobra.ExactArgs(1),
		RunE: runValidateConfig,
	}
	return cmd
}

func runValidateConfig(cmd *cobra.Command, args []string) error {
	configPath := args[0]

	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var config MappingConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	if err := config.Validate(); err != nil {
		return err
	}

	fmt.Printf("Valid: %s\n", configPath)
	fmt.Printf("  Name: %s\n", config.Name)
	fmt.Printf("  Resources: %d\n", len(config.Resources))
	for _, r := range config.Resources {
		entCount := len(r.Entitlements)
		fmt.Printf("    - %s (%s) [%d entitlements]\n", r.Type, r.Trait, entCount)
	}

	return nil
}

// MappingConfig is the root configuration for meta-connectors.
type MappingConfig struct {
	Name        string           `yaml:"name"`
	Description string           `yaml:"description"`
	Resources   []ResourceConfig `yaml:"resources"`
}

// ResourceConfig defines how to sync a resource type.
type ResourceConfig struct {
	Type         string              `yaml:"type"`
	DisplayName  string              `yaml:"display_name"`
	Trait        string              `yaml:"trait"`
	List         ListConfig          `yaml:"list"`
	Fields       FieldMapping        `yaml:"fields"`
	Entitlements []EntitlementConfig `yaml:"entitlements,omitempty"`
}

// ListConfig defines how to list resources.
type ListConfig struct {
	Endpoint     string            `yaml:"endpoint"`
	Method       string            `yaml:"method,omitempty"`
	ResponsePath string            `yaml:"response_path,omitempty"`
	Pagination   *PaginationConfig `yaml:"pagination,omitempty"`
}

// PaginationConfig defines pagination behavior.
type PaginationConfig struct {
	Type        string `yaml:"type"`
	CursorParam string `yaml:"cursor_param,omitempty"`
	CursorPath  string `yaml:"cursor_path,omitempty"`
	OffsetParam string `yaml:"offset_param,omitempty"`
	LimitParam  string `yaml:"limit_param,omitempty"`
	PageSize    int    `yaml:"page_size,omitempty"`
}

// FieldMapping maps API fields to baton fields.
type FieldMapping struct {
	ID          string            `yaml:"id"`
	DisplayName string            `yaml:"display_name"`
	Email       string            `yaml:"email,omitempty"`
	Description string            `yaml:"description,omitempty"`
	Status      string            `yaml:"status,omitempty"`
	Profile     map[string]string `yaml:"profile,omitempty"`
}

// EntitlementConfig defines an entitlement.
type EntitlementConfig struct {
	ID           string           `yaml:"id"`
	DisplayName  string           `yaml:"display_name"`
	Description  string           `yaml:"description,omitempty"`
	GrantableTo  []string         `yaml:"grantable_to"`
	Grants       *GrantsConfig    `yaml:"grants,omitempty"`
	Provisioning *ProvisionConfig `yaml:"provisioning,omitempty"`
}

// GrantsConfig defines how to fetch grants.
type GrantsConfig struct {
	Endpoint        string `yaml:"endpoint"`
	Method          string `yaml:"method,omitempty"`
	ResponsePath    string `yaml:"response_path,omitempty"`
	PrincipalIDPath string `yaml:"principal_id_path"`
	PrincipalType   string `yaml:"principal_type"`
}

// ProvisionConfig defines provisioning actions.
type ProvisionConfig struct {
	Grant  *ProvisionAction `yaml:"grant,omitempty"`
	Revoke *ProvisionAction `yaml:"revoke,omitempty"`
}

// ProvisionAction defines a single provisioning call.
type ProvisionAction struct {
	Endpoint string         `yaml:"endpoint"`
	Method   string         `yaml:"method"`
	Body     map[string]any `yaml:"body,omitempty"`
}

// Validate checks the configuration for errors.
func (c *MappingConfig) Validate() error {
	var errs []string

	if c.Name == "" {
		errs = append(errs, "name is required")
	}

	if len(c.Resources) == 0 {
		errs = append(errs, "at least one resource is required")
	}

	hasUser := false
	resourceTypes := make(map[string]bool)
	validTraits := map[string]bool{
		"":            true,
		"TRAIT_USER":  true,
		"TRAIT_GROUP": true,
		"TRAIT_ROLE":  true,
		"TRAIT_APP":   true,
	}

	for i, r := range c.Resources {
		prefix := fmt.Sprintf("resources[%d]", i)
		if r.Type != "" {
			prefix = fmt.Sprintf("resources[%d] (%s)", i, r.Type)
		}

		switch {
		case r.Type == "":
			errs = append(errs, fmt.Sprintf("%s: type is required", prefix))
		case resourceTypes[r.Type]:
			errs = append(errs, fmt.Sprintf("%s: duplicate resource type", prefix))
		default:
			resourceTypes[r.Type] = true
		}

		if !validTraits[r.Trait] {
			errs = append(errs, fmt.Sprintf("%s: invalid trait %q", prefix, r.Trait))
		}

		if r.Trait == "TRAIT_USER" {
			hasUser = true
		}

		if r.List.Endpoint == "" {
			errs = append(errs, fmt.Sprintf("%s: list.endpoint is required", prefix))
		}

		if r.Fields.ID == "" {
			errs = append(errs, fmt.Sprintf("%s: fields.id is required", prefix))
		}

		if r.List.Pagination != nil {
			pag := r.List.Pagination
			validPagTypes := map[string]bool{"cursor": true, "offset": true, "page": true, "link": true}
			if !validPagTypes[pag.Type] {
				errs = append(errs, fmt.Sprintf("%s: invalid pagination type %q", prefix, pag.Type))
			}
			if pag.Type == "cursor" && (pag.CursorParam == "" || pag.CursorPath == "") {
				errs = append(errs, fmt.Sprintf("%s: cursor pagination requires cursor_param and cursor_path", prefix))
			}
			if pag.Type == "offset" && (pag.OffsetParam == "" || pag.LimitParam == "") {
				errs = append(errs, fmt.Sprintf("%s: offset pagination requires offset_param and limit_param", prefix))
			}
		}

		for j, e := range r.Entitlements {
			eprefix := fmt.Sprintf("%s.entitlements[%d]", prefix, j)
			if e.ID != "" {
				eprefix = fmt.Sprintf("%s.entitlements[%d] (%s)", prefix, j, e.ID)
			}

			if e.ID == "" {
				errs = append(errs, fmt.Sprintf("%s: id is required", eprefix))
			}

			if len(e.GrantableTo) == 0 {
				errs = append(errs, fmt.Sprintf("%s: grantable_to is required", eprefix))
			}

			if e.Grants != nil {
				if e.Grants.Endpoint == "" {
					errs = append(errs, fmt.Sprintf("%s.grants: endpoint is required", eprefix))
				}
				if e.Grants.PrincipalIDPath == "" {
					errs = append(errs, fmt.Sprintf("%s.grants: principal_id_path is required", eprefix))
				}
				if e.Grants.PrincipalType == "" {
					errs = append(errs, fmt.Sprintf("%s.grants: principal_type is required", eprefix))
				}
			}

			if e.Provisioning != nil {
				if e.Provisioning.Grant != nil {
					if e.Provisioning.Grant.Endpoint == "" {
						errs = append(errs, fmt.Sprintf("%s.provisioning.grant: endpoint is required", eprefix))
					}
					if e.Provisioning.Grant.Method == "" {
						errs = append(errs, fmt.Sprintf("%s.provisioning.grant: method is required", eprefix))
					}
				}
				if e.Provisioning.Revoke != nil {
					if e.Provisioning.Revoke.Endpoint == "" {
						errs = append(errs, fmt.Sprintf("%s.provisioning.revoke: endpoint is required", eprefix))
					}
					if e.Provisioning.Revoke.Method == "" {
						errs = append(errs, fmt.Sprintf("%s.provisioning.revoke: method is required", eprefix))
					}
				}
			}
		}
	}

	if !hasUser && len(c.Resources) > 0 {
		errs = append(errs, "at least one resource with TRAIT_USER is required")
	}

	if len(errs) > 0 {
		msg := fmt.Sprintf("validation failed (%d errors):", len(errs))
		for _, e := range errs {
			msg += fmt.Sprintf("\n  - %s", e)
		}
		return fmt.Errorf("%s", msg)
	}

	return nil
}
