package main

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// aliasStats tracks the progress of alias generation.
type aliasStats struct {
	Total     int
	Processed int
	Skipped   int
	Updated   int
	Failed    int
	Errors    []string
}

// Generic entitlement words that should be replaced with resource names.
var genericEntitlementWords = []string{
	"member",
	"assignment",
	"access",
	"role",
	"group",
}

// Words to remove from display names.
var wordsToRemove = []string{
	"Role member", "role member", "RoleMember", "roleMember",
	"Group member", "group member", "GroupMember", "groupMember",
	"Member", "member",
	"Role", "role",
	"Group", "group",
}

// Pre-compiled regexps for cleanText.
var (
	reParens     = regexp.MustCompile(`\s*\([^)]*\)`)
	reWhitespace = regexp.MustCompile(`\s+`)
	reInvalid    = regexp.MustCompile(`[^a-z0-9\-_\.]`)
)

var validSchemas = []string{
	"resource-entitlement",
	"app-entitlement",
	"app-resource-entitlement",
	"resource-type-entitlement",
	"custom",
}

// cleanText processes a string to create a valid alias.
func cleanText(text string) string {
	text = reParens.ReplaceAllString(text, "")

	for _, word := range wordsToRemove {
		text = strings.ReplaceAll(text, word, "")
	}

	text = strings.TrimSpace(text)
	text = reWhitespace.ReplaceAllString(text, " ")
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "-")
	text = reInvalid.ReplaceAllString(text, "")
	text = strings.ReplaceAll(text, "--", "-")
	text = strings.Trim(text, "-_.")

	if text == "" {
		text = "entitlement"
	}

	if len(text) > 63 {
		text = text[:63]
		text = strings.TrimRight(text, "-")
	}

	return text
}

// isGenericEntitlement checks if an entitlement name is too generic to be useful.
func isGenericEntitlement(name string) bool {
	return slices.Contains(genericEntitlementWords, strings.ToLower(strings.TrimSpace(name)))
}

// getResourceName extracts and cleans the resource name from an entitlement.
func getResourceName(e *client.EntitlementWithBindings) string {
	if appResource := client.GetExpanded[shared.AppResource](e, client.ExpandedAppResource); appResource != nil && appResource.DisplayName != nil {
		return cleanText(*appResource.DisplayName)
	}
	if app := client.GetExpanded[shared.App](e, client.ExpandedApp); app != nil && app.DisplayName != nil {
		return cleanText(*app.DisplayName)
	}
	return "resource"
}

// generateAliasRun is the main function that handles alias generation.
func generateAliasRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := verifyAdminPermissions(ctx, c); err != nil {
		return err
	}

	flags := getCommandFlags(v)

	if !slices.Contains(validSchemas, flags.Schema) {
		return fmt.Errorf("unknown schema %q; valid values: %s", flags.Schema, strings.Join(validSchemas, ", "))
	}

	entitlements, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
		GrantedStatus:            shared.GrantedStatusAll,
		AppEntitlementExpandMask: shared.AppEntitlementExpandMask{Paths: []string{"app_id", "app_resource_type_id", "app_resource_id"}},
	})
	if err != nil {
		return fmt.Errorf("failed to search entitlements: %w", err)
	}
	if len(entitlements) == 0 {
		pterm.Warning.Println("No requestable entitlements found.")
		return nil
	}

	stats := &aliasStats{
		Total:  len(entitlements),
		Errors: make([]string, 0),
	}

	pterm.Info.Printf("Processing %d entitlements...\n", stats.Total)

	processedEntitlements := make(map[string]bool)
	for i, e := range entitlements {
		if err := processEntitlement(ctx, c, e, flags, stats, processedEntitlements); err != nil {
			stats.Failed++
			stats.Errors = append(stats.Errors, err.Error())
		}

		if (i+1)%10 == 0 {
			pterm.Info.Printf("Processed %d/%d entitlements...\n", i+1, stats.Total)
		}
	}

	printSummary(stats)

	return nil
}

// commandFlags holds all the command line flags.
type commandFlags struct {
	ResourceTypes  []string
	EntitlementIDs []string
	Schema         string
	Format         string
	Separator      string
	Force          bool
	ForceNonAWS    bool
	SkipAWS        bool
	DryRun         bool
}

func getCommandFlags(v *viper.Viper) commandFlags {
	return commandFlags{
		ResourceTypes:  v.GetStringSlice("resource-type"),
		EntitlementIDs: v.GetStringSlice("entitlement-id"),
		Schema:         v.GetString("schema"),
		Format:         v.GetString("format"),
		Separator:      v.GetString("separator"),
		Force:          v.GetBool("force"),
		ForceNonAWS:    v.GetBool("force-non-aws"),
		SkipAWS:        v.GetBool("skip-aws"),
		DryRun:         v.GetBool("dry-run"),
	}
}

func processEntitlement(ctx context.Context, c client.C1Client, e *client.EntitlementWithBindings, flags commandFlags, stats *aliasStats, processedEntitlements map[string]bool) error {
	ent := e.Entitlement
	if ent.DisplayName == nil || ent.AppID == nil || ent.ID == nil {
		stats.Skipped++
		return nil
	}

	if len(flags.EntitlementIDs) > 0 && !slices.Contains(flags.EntitlementIDs, *ent.ID) {
		stats.Skipped++
		return nil
	}

	app := client.GetExpanded[shared.App](e, client.ExpandedApp)
	appResourceType := client.GetExpanded[shared.AppResourceType](e, client.ExpandedAppResourceType)
	if app == nil || app.DisplayName == nil || appResourceType == nil || appResourceType.DisplayName == nil {
		stats.Skipped++
		return nil
	}

	if len(flags.ResourceTypes) > 0 && !slices.Contains(flags.ResourceTypes, *appResourceType.DisplayName) {
		stats.Skipped++
		return nil
	}

	// Resolve display name: use resource name for generic entitlements.
	displayName := cleanText(*ent.DisplayName)
	displayName = strings.TrimSuffix(displayName, "-access")
	displayName = strings.TrimSuffix(displayName, "-permissionset")

	resourceName := getResourceName(e)
	if isGenericEntitlement(*ent.DisplayName) {
		displayName = resourceName
		pterm.Info.Printf("Using resource name '%s' instead of generic entitlement name '%s'\n",
			resourceName, *ent.DisplayName)
	}

	// Build alias from schema.
	appName := cleanText(*app.DisplayName)
	values := map[string]string{
		"a": appName,
		"r": resourceName,
		"t": cleanText(*appResourceType.DisplayName),
		"e": displayName,
	}

	var alias string
	switch flags.Schema {
	case "app-entitlement":
		alias = generateAlias("%a-%e", flags.Separator, values)
	case "resource-entitlement":
		alias = generateAlias("%r-%e", flags.Separator, values)
	case "app-resource-entitlement":
		alias = generateAlias("%a-%r-%e", flags.Separator, values)
	case "resource-type-entitlement":
		alias = generateAlias("%t-%e", flags.Separator, values)
	case "custom":
		alias = generateAlias(flags.Format, flags.Separator, values)
	}

	isAWSPermissionSet := strings.ToLower(*appResourceType.DisplayName) == "account"

	if flags.SkipAWS && isAWSPermissionSet {
		stats.Skipped++
		return nil
	}

	if ent.Alias != nil && *ent.Alias != "" {
		if !flags.Force && !flags.ForceNonAWS && !isAWSPermissionSet {
			stats.Skipped++
			return nil
		}
	}

	if processedEntitlements[alias] {
		stats.Skipped++
		return nil
	}
	processedEntitlements[alias] = true

	if !flags.DryRun {
		req := &shared.UpdateAppEntitlementRequest{
			AppEntitlement: &shared.AppEntitlementInput{
				Alias: &alias,
			},
			UpdateMask: stringPtr("alias"),
		}
		if err := c.UpdateEntitlement(ctx, *ent.AppID, *ent.ID, req); err != nil {
			return fmt.Errorf("failed to update %s: %w", *ent.DisplayName, err)
		}
	}
	stats.Updated++
	stats.Processed++

	return nil
}

// Admin role prefixes that grant permission to update entitlements.
var adminRolePrefixes = []string{
	"role/c1.api.tenant.v1.Tenant:owner",
	"role/c1.api.auth.v1.Auth:owner",
	"role/c1.api.app.v1.Apps:owner",
	"role/c1.api.app.v1.AppEntitlements:owner",
}

// verifyAdminPermissions checks the caller's roles from AuthIntrospect.
func verifyAdminPermissions(ctx context.Context, c client.C1Client) error {
	userInfo, err := c.AuthIntrospect(ctx)
	if err != nil {
		return fmt.Errorf("failed to check permissions: %w", err)
	}

	for _, role := range userInfo.Roles {
		for _, prefix := range adminRolePrefixes {
			if strings.HasPrefix(role, prefix) {
				return nil
			}
		}
	}

	return fmt.Errorf("this command requires super admin or app admin permissions")
}

func printSummary(stats *aliasStats) {
	pterm.Info.Printf("\nSummary:\n")
	pterm.Info.Printf("Total entitlements: %d\n", stats.Total)
	pterm.Info.Printf("Processed: %d\n", stats.Processed)
	pterm.Info.Printf("Skipped: %d\n", stats.Skipped)
	pterm.Info.Printf("Updated: %d\n", stats.Updated)
	pterm.Info.Printf("Failed: %d\n", stats.Failed)

	if len(stats.Errors) > 0 {
		pterm.Error.Printf("\nErrors:\n")
		for _, err := range stats.Errors {
			pterm.Error.Printf("- %s\n", err)
		}
	}
}

func stringPtr(s string) *string { return &s }

func generateAliasCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-alias",
		Short: "Generate aliases for entitlements in ConductorOne",
		Long: `Generate aliases for entitlements in ConductorOne.
This command will:
- Generate aliases based on the selected schema
- By default, skip entitlements that already have aliases (except AWS permission sets)
- Show progress and summary statistics

Available alias schemas (--schema):
- resource-entitlement: resource name + entitlement name (default)
- app-entitlement: app name + entitlement name
- app-resource-entitlement: app name + resource name + entitlement name
- resource-type-entitlement: resource type + entitlement name
- custom: Use a custom format string with --format

Format placeholders:
- %a: App name
- %r: Resource name
- %t: Resource type
- %e: Entitlement name
- %s: Any separator (default: "-")

Example formats:
- "%r-%e" (default): resource-entitlement
- "%a-%e": app-entitlement
- "%a-%r-%e": app-resource-entitlement
- "%t-%e": resource-type-entitlement

The alias format can be customized using flags:
- --schema: Choose a predefined schema or "custom"
- --format: Custom format string (only used with schema=custom)
- --separator: Custom separator (default: "-")
- --force: Override ALL existing aliases (including AWS permission sets)
- --force-non-aws: Override existing aliases for non-AWS entitlements
- --skip-aws: Skip AWS permission sets entirely
- --dry-run: Preview changes without making them

Filtering options:
- --resource-type: Only process entitlements with specific resource types
- --entitlement-id: Process only specific entitlements`,
		RunE: generateAliasRun,
	}

	cmd.Flags().String("schema", "resource-entitlement", "Alias schema to use (resource-entitlement, app-entitlement, app-resource-entitlement, resource-type-entitlement, custom)")
	cmd.Flags().String("format", "%r-%e", "Custom format string for alias generation (only used with schema=custom)")
	cmd.Flags().String("separator", "-", "Separator to use between components")
	cmd.Flags().Bool("force", false, "Override ALL existing aliases (including AWS permission sets)")
	cmd.Flags().Bool("force-non-aws", false, "Override existing aliases for non-AWS entitlements")
	cmd.Flags().Bool("skip-aws", false, "Skip AWS permission sets entirely")
	cmd.Flags().Bool("dry-run", false, "Preview changes without making them")
	cmd.Flags().StringSlice("resource-type", []string{}, "Only process entitlements with these resource types")
	cmd.Flags().StringSlice("entitlement-id", []string{}, "Process only these entitlements")

	cmd.MarkFlagsMutuallyExclusive("force", "force-non-aws")

	return cmd
}

// generateAlias generates an alias using the given format and values.
func generateAlias(format, separator string, values map[string]string) string {
	result := format
	for key, value := range values {
		result = strings.ReplaceAll(result, "%"+key, value)
	}
	result = strings.ReplaceAll(result, "%s", separator)
	return result
}
