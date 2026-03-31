package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
)

type virtualEntitlementYAML struct {
	Resources []virtualResourceYAML `yaml:"resources"`
}

type virtualResourceYAML struct {
	Name         string   `yaml:"name"`
	Type         string   `yaml:"type"`
	Description  string   `yaml:"description"`
	Entitlements []string `yaml:"entitlements"`
}

func virtualEntitlementsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "virtual-entitlements",
		Short: "Manage virtual (manually-managed) entitlements.",
	}
	cmd.AddCommand(virtualEntitlementsCreateCmd())
	return cmd
}

func virtualEntitlementsCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create virtual resource types, resources, and entitlements on an app.",
		Long: `Create virtual (manually-managed) entitlements on a ConductorOne app.

You can specify resources and entitlements via CLI flags or a YAML file.

CLI example:
  cone virtual-entitlements create --app "My App" --resource "My Group" --type GROUP --entitlements "Member" --entitlements "Admin"

YAML file example:
  cone virtual-entitlements create --app "My App" --from-file entitlements.yaml

YAML format:
  resources:
    - name: "My Group"
      type: GROUP
      description: "optional"
      entitlements:
        - "Member"
        - "Admin"`,
		RunE: virtualEntitlementsCreateRun,
	}
	cmd.Flags().String("app", "", "App ID or display name (required).")
	cmd.Flags().String("resource", "", "Resource display name.")
	cmd.Flags().StringP("type", "t", "CUSTOM", "Resource type: ROLE, GROUP, LICENSE, PROJECT, CATALOG, CUSTOM, VAULT, PROFILE_TYPE.")
	cmd.Flags().StringSlice("entitlements", nil, "Entitlement names (repeatable).")
	cmd.Flags().StringP("from-file", "f", "", "YAML file with resource/entitlement definitions.")
	if err := cmd.MarkFlagRequired("app"); err != nil {
		panic(err)
	}
	return cmd
}

var slugRegexp = regexp.MustCompile(`[^a-z0-9\-_.]`)

// makeSlug produces a URL-safe slug from a display name. Non-alphanumeric characters
// are replaced with hyphens, so distinct names like "foo@bar" and "foo#bar" will
// collide to the same slug. This is acceptable because the C1 API upserts entitlements
// on (app, resource_type, resource) — slug is not part of the uniqueness key.
func makeSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = slugRegexp.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	return slug
}

func virtualEntitlementsCreateRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	appIDOrName, err := cmd.Flags().GetString("app")
	if err != nil {
		return err
	}

	app, err := c.ResolveAppByNameOrID(ctx, appIDOrName)
	if err != nil {
		return err
	}

	appID := client.StringFromPtr(app.ID)
	appName := client.StringFromPtr(app.DisplayName)
	fmt.Fprintf(os.Stderr, "Using app: %s (%s)\n", appName, appID)

	fromFile, err := cmd.Flags().GetString("from-file")
	if err != nil {
		return err
	}

	var resources []virtualResourceYAML
	if fromFile != "" {
		data, err := os.ReadFile(fromFile) //nolint:gosec // path from CLI flag
		if err != nil {
			return fmt.Errorf("reading file %s: %w", fromFile, err)
		}
		var spec virtualEntitlementYAML
		if err := yaml.Unmarshal(data, &spec); err != nil {
			return fmt.Errorf("parsing YAML file %s: %w", fromFile, err)
		}
		if len(spec.Resources) == 0 {
			return fmt.Errorf("no resources found in %s", fromFile)
		}
		resources = spec.Resources
	} else {
		resourceName, err := cmd.Flags().GetString("resource")
		if err != nil {
			return err
		}
		if resourceName == "" {
			return fmt.Errorf("--resource is required when not using --from-file")
		}

		entitlementNames, err := cmd.Flags().GetStringSlice("entitlements")
		if err != nil {
			return err
		}
		if len(entitlementNames) == 0 {
			return fmt.Errorf("--entitlements is required when not using --from-file")
		}

		typeName, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}

		resources = []virtualResourceYAML{
			{
				Name:         resourceName,
				Type:         typeName,
				Entitlements: entitlementNames,
			},
		}
	}

	resp := &VirtualEntitlementsResponse{}

	for _, res := range resources {
		if res.Type == "" {
			res.Type = "CUSTOM"
		}
		resourceType, typeDisplayName := client.ResolveResourceType(res.Type)

		fmt.Fprintf(os.Stderr, "\nProcessing resource: %s (type: %s)\n", res.Name, typeDisplayName)

		rt, err := c.CreateManuallyManagedResourceType(ctx, appID, resourceType, typeDisplayName)
		if err != nil {
			return fmt.Errorf("creating resource type %q: %w", typeDisplayName, err)
		}
		rtID := client.StringFromPtr(rt.ID)
		fmt.Fprintf(os.Stderr, "  Created resource type: %s (%s)\n", typeDisplayName, rtID)

		resource, err := c.CreateManuallyManagedResource(ctx, appID, rtID, res.Name, res.Description)
		if err != nil {
			return fmt.Errorf("creating resource %q: %w", res.Name, err)
		}
		resourceID := client.StringFromPtr(resource.ID)
		fmt.Fprintf(os.Stderr, "  Created resource: %s (%s)\n", res.Name, resourceID)

		for _, entName := range res.Entitlements {
			ent, err := c.CreateAppEntitlement(ctx, appID, rtID, resourceID, entName, makeSlug(entName))
			if err != nil {
				return fmt.Errorf("creating entitlement %q: %w", entName, err)
			}
			entID := client.StringFromPtr(ent.ID)
			fmt.Fprintf(os.Stderr, "    Created entitlement: %s (%s)\n", entName, entID)

			resp.Entitlements = append(resp.Entitlements, virtualEntitlementRow{
				AppName:          appName,
				ResourceTypeName: typeDisplayName,
				ResourceName:     res.Name,
				EntitlementName:  entName,
				EntitlementID:    entID,
			})
		}
	}

	outputManager := output.NewManager(ctx, v)
	return outputManager.Output(ctx, resp)
}

type virtualEntitlementRow struct {
	AppName          string `json:"appName"`
	ResourceTypeName string `json:"resourceTypeName"`
	ResourceName     string `json:"resourceName"`
	EntitlementName  string `json:"entitlementName"`
	EntitlementID    string `json:"entitlementId"`
}

// VirtualEntitlementsResponse implements the output interfaces for table/JSON rendering.
type VirtualEntitlementsResponse struct {
	Entitlements []virtualEntitlementRow `json:"entitlements"`
}

func (r *VirtualEntitlementsResponse) Header() []string {
	return []string{"App", "Resource Type", "Resource", "Entitlement", "Entitlement ID"}
}

func (r *VirtualEntitlementsResponse) Rows() [][]string {
	rows := make([][]string, 0, len(r.Entitlements))
	for _, e := range r.Entitlements {
		rows = append(rows, []string{
			e.AppName,
			e.ResourceTypeName,
			e.ResourceName,
			e.EntitlementName,
			e.EntitlementID,
		})
	}
	return rows
}

func (r *VirtualEntitlementsResponse) WideHeader() []string {
	return r.Header()
}

func (r *VirtualEntitlementsResponse) WideRows() [][]string {
	return r.Rows()
}
