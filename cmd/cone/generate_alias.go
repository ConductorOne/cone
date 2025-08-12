package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Stats tracks the progress of alias generation.
type Stats struct {
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

// cleanText processes a string to create a valid alias by:
// 1. Removing text in parentheses
// 2. Removing unnecessary words
// 3. Converting to lowercase
// 4. Replacing spaces with hyphens
// 5. Removing invalid characters
// 6. Ensuring proper length and format.
func cleanText(text string) string {
	// Remove anything in parentheses
	text = regexp.MustCompile(`\s*\([^)]*\)`).ReplaceAllString(text, "")

	// Remove unnecessary words
	for _, word := range wordsToRemove {
		text = strings.ReplaceAll(text, word, "")
	}

	// Clean up text
	text = strings.TrimSpace(text)
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")

	// Convert to lowercase and replace spaces with hyphens
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "-")

	// Remove any characters that aren't allowed
	text = regexp.MustCompile(`[^a-z0-9\-_\.]`).ReplaceAllString(text, "")

	// Remove any double hyphens
	text = strings.ReplaceAll(text, "--", "-")

	// Ensure it starts and ends with a letter or digit
	text = strings.Trim(text, "-_.")
	if text == "" {
		text = "entitlement" // fallback if we end up with empty string
	}

	// Ensure it's not too long
	if len(text) > 63 {
		text = text[:63]
		// Make sure we don't end with a hyphen
		text = strings.TrimRight(text, "-")
	}

	return text
}

// isGenericEntitlement checks if an entitlement name is too generic to be useful.
func isGenericEntitlement(name string) bool {
	name = strings.ToLower(strings.TrimSpace(name))
	for _, word := range genericEntitlementWords {
		if name == word {
			return true
		}
	}
	return false
}

// getResourceName extracts and cleans the resource name from an entitlement.
func getResourceName(e *client.EntitlementWithBindings) string {
	if appResource := client.GetExpanded[shared.AppResource](e, client.ExpandedAppResource); appResource != nil && appResource.DisplayName != nil {
		return cleanText(*appResource.DisplayName)
	}
	// Fallback to app name if no resource name is available
	if app := client.GetExpanded[shared.App](e, client.ExpandedApp); app != nil && app.DisplayName != nil {
		return cleanText(*app.DisplayName)
	}
	return "resource" // Final fallback
}

// generateAliasRun is the main function that handles alias generation.
func generateAliasRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	// Verify admin permissions
	if err := verifyAdminPermissions(ctx, c, v); err != nil {
		return err
	}

	// Get all requestable entitlements
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

	// Initialize stats and get command flags
	stats := &Stats{
		Total:  len(entitlements),
		Errors: make([]string, 0),
	}
	flags := getCommandFlags(v)

	pterm.Info.Printf("Processing %d entitlements...\n", stats.Total)

	processedEntitlements := make(map[string]bool)
	for i, e := range entitlements {
		// Process each entitlement
		if err := processEntitlement(ctx, c, e, flags, stats, processedEntitlements); err != nil {
			stats.Failed++
			stats.Errors = append(stats.Errors, err.Error())
		}

		// Show progress every 10 items
		if (i+1)%10 == 0 {
			pterm.Info.Printf("Processed %d/%d entitlements...\n", i+1, stats.Total)
		}
	}

	// Print summary
	printSummary(stats)

	return nil
}

// CommandFlags holds all the command line flags.
type CommandFlags struct {
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

// getCommandFlags extracts all command line flags.
func getCommandFlags(v *viper.Viper) CommandFlags {
	return CommandFlags{
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

// processEntitlement handles the processing of a single entitlement.
func processEntitlement(ctx context.Context, c client.C1Client, e *client.EntitlementWithBindings, flags CommandFlags, stats *Stats, processedEntitlements map[string]bool) error {
	ent := e.Entitlement
	if ent.DisplayName == nil || ent.AppID == nil || ent.ID == nil {
		stats.Skipped++
		return nil
	}

	// Apply filters
	if len(flags.EntitlementIDs) > 0 && !contains(flags.EntitlementIDs, *ent.ID) {
		stats.Skipped++
		return nil
	}

	// Get app and resource type info
	app := client.GetExpanded[shared.App](e, client.ExpandedApp)
	appResourceType := client.GetExpanded[shared.AppResourceType](e, client.ExpandedAppResourceType)
	if app == nil || app.DisplayName == nil || appResourceType == nil || appResourceType.DisplayName == nil {
		stats.Skipped++
		return nil
	}

	// Filter by resource type
	if len(flags.ResourceTypes) > 0 && !contains(flags.ResourceTypes, *appResourceType.DisplayName) {
		stats.Skipped++
		return nil
	}

	// Clean up the display name
	displayName := cleanText(*ent.DisplayName)
	displayName = strings.TrimSuffix(displayName, "-access")
	displayName = strings.TrimSuffix(displayName, "-permissionset")

	// Get resource name and check for generic entitlements
	resourceName := getResourceName(e)
	if isGenericEntitlement(*ent.DisplayName) {
		displayName = resourceName
		pterm.Info.Printf("Using resource name '%s' instead of generic entitlement name '%s'\n",
			resourceName, *ent.DisplayName)
	}

	// Generate alias
	var alias string
	var resourceNameGenerated string

	// Get resource name if available
	if appResource := client.GetExpanded[shared.AppResource](e, client.ExpandedAppResource); appResource != nil && appResource.DisplayName != nil {
		resourceNameGenerated = cleanText(*appResource.DisplayName)
	} else {
		resourceNameGenerated = cleanText(*app.DisplayName)
	}

	// Check if we need to use resource name instead of entitlement name
	if isGenericEntitlement(*ent.DisplayName) {
		displayName = resourceNameGenerated
		pterm.Info.Printf("Using resource name '%s' instead of generic entitlement name '%s'\n",
			resourceNameGenerated, *ent.DisplayName)
	}

	// Generate the alias based on schema
	appName := cleanText(*app.DisplayName)
	switch flags.Schema {
	case "app-entitlement":
		alias = generateAlias("%a-%e", flags.Separator, map[string]string{
			"a": appName,
			"e": displayName,
		})
	case "resource-entitlement":
		alias = generateAlias("%r-%e", flags.Separator, map[string]string{
			"r": resourceNameGenerated,
			"e": displayName,
		})
	case "app-resource-entitlement":
		alias = generateAlias("%a-%r-%e", flags.Separator, map[string]string{
			"a": appName,
			"r": resourceNameGenerated,
			"e": displayName,
		})
	case "resource-type-entitlement":
		alias = generateAlias("%t-%e", flags.Separator, map[string]string{
			"t": cleanText(*appResourceType.DisplayName),
			"e": displayName,
		})
	default:
		alias = generateAlias(flags.Format, flags.Separator, map[string]string{
			"a": appName,
			"r": resourceNameGenerated,
			"t": cleanText(*appResourceType.DisplayName),
			"e": displayName,
		})
	}

	// Check if this is an AWS permission set
	isAWSPermissionSet := strings.ToLower(*appResourceType.DisplayName) == "account"

	// Skip AWS permission sets if requested
	if flags.SkipAWS && isAWSPermissionSet {
		stats.Skipped++
		return nil
	}

	// Skip if alias is already set and not forcing
	if ent.Alias != nil && *ent.Alias != "" {
		if !flags.Force && !flags.ForceNonAWS && !isAWSPermissionSet {
			stats.Skipped++
			return nil
		}
	}

	// Skip if we've already processed this alias
	if processedEntitlements[alias] {
		stats.Skipped++
		return nil
	}
	processedEntitlements[alias] = true

	// Update the alias
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

// verifyAdminPermissions checks if the user has admin permissions.
func verifyAdminPermissions(ctx context.Context, c client.C1Client, v *viper.Viper) error {
	// In non-interactive mode, skip user prompt and check permissions directly
	if v.GetBool(nonInteractiveFlag) {
		// Just check actual permissions without prompting
		isAdmin, err := checkAdminPermissions(ctx, c)
		if err != nil {
			return fmt.Errorf("failed to check admin permissions: %w", err)
		}
		if !isAdmin {
			return fmt.Errorf("you do not have super admin or app admin permissions. Use --help for more information")
		}
		return nil
	}

	// Interactive mode: prompt user first, then check permissions
	// Prompt the user
	pterm.Info.Print("Are you a super admin or app admin? (yes/no): ")
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))
	if answer != "yes" {
		return fmt.Errorf("you must be a super admin or app admin to run this command")
	}

	// Check actual permissions
	isAdmin, err := checkAdminPermissions(ctx, c)
	if err != nil {
		return fmt.Errorf("failed to check admin permissions: %w", err)
	}
	if !isAdmin {
		return fmt.Errorf("you do not have super admin or app admin permissions. Please contact your administrator")
	}

	return nil
}

// printSummary prints the final summary of the alias generation process.
func printSummary(stats *Stats) {
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

// Helper functions.
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func stringPtr(s string) *string { return &s }

// generateAliasCmd creates the cobra command for alias generation.
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

	// Add flags
	cmd.Flags().String("schema", "resource-entitlement", "Alias schema to use (resource-entitlement, app-entitlement, app-resource-entitlement, resource-type-entitlement, custom)")
	cmd.Flags().String("format", "%r-%e", "Custom format string for alias generation (only used with schema=custom)")
	cmd.Flags().String("separator", "-", "Separator to use between components")
	cmd.Flags().Bool("force", false, "Override ALL existing aliases (including AWS permission sets)")
	cmd.Flags().Bool("force-non-aws", false, "Override existing aliases for non-AWS entitlements")
	cmd.Flags().Bool("skip-aws", false, "Skip AWS permission sets entirely")
	cmd.Flags().Bool("dry-run", false, "Preview changes without making them")
	cmd.Flags().StringSlice("resource-type", []string{}, "Only process entitlements with these resource types")
	cmd.Flags().StringSlice("entitlement-id", []string{}, "Process only these entitlements")

	// Mark flags as mutually exclusive
	cmd.MarkFlagsMutuallyExclusive("force", "force-non-aws")

	return cmd
}

// generateAlias generates an alias using the given format and values.
func generateAlias(format, separator string, values map[string]string) string {
	// Replace placeholders with values
	result := format
	for key, value := range values {
		result = strings.ReplaceAll(result, "%"+key, value)
	}
	// Replace separator placeholder
	result = strings.ReplaceAll(result, "%s", separator)
	return result
}

// checkAdminPermissions checks if the user has admin permissions.
func checkAdminPermissions(ctx context.Context, c client.C1Client) (bool, error) {
	// Get user's identity
	userIntro, err := c.AuthIntrospect(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get user identity: %w", err)
	}

	// Check if user is a super admin or app admin
	for _, role := range userIntro.Roles {
		// Check for roles that indicate super admin access
		if strings.HasPrefix(role, "role/c1.api.tenant.v1.Tenant:owner") ||
			strings.HasPrefix(role, "role/c1.api.auth.v1.Auth:owner") ||
			// Check for app admin roles
			strings.HasPrefix(role, "role/c1.api.app.v1.Apps:owner") ||
			strings.HasPrefix(role, "role/c1.api.app.v1.AppEntitlements:owner") {
			return true, nil
		}
	}

	// Check if user is an app admin for any app
	apps, err := c.ListApps(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to list apps: %w", err)
	}

	for _, app := range apps {
		if app.ID == nil {
			continue
		}

		// Get app users
		appUsers, err := c.ListAppUsers(ctx, *app.ID)
		if err != nil {
			continue // Skip this app if we can't get users
		}

		// Check if user is an admin for this app
		for _, appUser := range appUsers {
			if appUser.IdentityUserID != nil && *appUser.IdentityUserID == *userIntro.UserID {
				// Check if user has admin role in their profile
				if profile, ok := appUser.Profile["roles"]; ok {
					if roles, ok := profile.([]interface{}); ok {
						for _, role := range roles {
							if roleStr, ok := role.(string); ok && roleStr == "admin" {
								return true, nil
							}
						}
					}
				}
			}
		}
	}

	return false, nil
}
