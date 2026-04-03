package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/client"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultAWSRegion = "us-east-1"

func awsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aws",
		Short: "AWS SSO integration commands.",
	}
	cmd.AddCommand(awsSetupCmd())
	cmd.AddCommand(awsCredentialsCmd())
	return cmd
}

// --- cone aws setup ---

func awsSetupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Configure AWS SSO and create profiles for available permission sets.",
		Long: `Scans ConductorOne for all AWS permission set entitlements available to you
and creates corresponding profiles in ~/.aws/config.

On first run, provide your SSO start URL and regions:
  cone aws setup --sso-url https://myorg.awsapps.com/start --sso-region us-east-1 --region us-west-2

Subsequent runs reuse the saved config:
  cone aws setup

View current configuration:
  cone aws setup show`,
		RunE: awsSetupRun,
	}
	cmd.Flags().String("sso-url", "", "AWS SSO start URL (saved to config for future runs).")
	cmd.Flags().String("sso-region", "", "AWS region where SSO Identity Center lives (saved to config, default: us-east-1).")
	cmd.Flags().String("region", "", "Default AWS region for CLI profiles, e.g. us-west-2 (saved to config, default: us-east-1).")
	cmd.AddCommand(awsSetupShowCmd())
	return cmd
}

func awsSetupShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show current AWS SSO configuration.",
		RunE: func(cmd *cobra.Command, args []string) error {
			ssoURL := viper.GetString("aws_sso_start_url")
			ssoRegion := viper.GetString("aws_sso_region")
			defaultRegion := viper.GetString("aws_default_region")

			if ssoURL == "" && ssoRegion == "" && defaultRegion == "" {
				pterm.Warning.Println("No AWS SSO configuration found. Run 'cone aws setup --sso-url <url>' first.")
				return nil
			}

			if ssoURL != "" {
				pterm.Info.Printf("SSO URL:        %s\n", ssoURL)
			} else {
				pterm.Warning.Println("SSO URL:        not set")
			}

			if ssoRegion != "" {
				pterm.Info.Printf("SSO Region:     %s\n", ssoRegion)
			} else {
				pterm.Warning.Println("SSO Region:     not set")
			}

			if defaultRegion != "" {
				pterm.Info.Printf("Default Region: %s\n", defaultRegion)
			} else {
				pterm.Warning.Println("Default Region: not set")
			}

			return nil
		},
	}
}

func awsSetupRun(cmd *cobra.Command, args []string) error {
	ctx, c, _, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	// Resolve SSO URL: flag > saved config.
	ssoURL, _ := cmd.Flags().GetString("sso-url")
	if ssoURL == "" {
		ssoURL = viper.GetString("aws_sso_start_url")
	}
	if ssoURL == "" {
		return fmt.Errorf("--sso-url is required on first run")
	}

	// Resolve SSO region (where Identity Center lives): flag > saved config > default.
	ssoRegion, _ := cmd.Flags().GetString("sso-region")
	if ssoRegion == "" {
		ssoRegion = viper.GetString("aws_sso_region")
	}
	if ssoRegion == "" {
		ssoRegion = defaultAWSRegion
	}

	// Resolve default AWS region for profiles: flag > saved config > default.
	defaultRegion, _ := cmd.Flags().GetString("region")
	if defaultRegion == "" {
		defaultRegion = viper.GetString("aws_default_region")
	}
	if defaultRegion == "" {
		defaultRegion = defaultAWSRegion
	}

	// Persist for future runs.
	viper.Set("aws_sso_start_url", ssoURL)
	viper.Set("aws_sso_region", ssoRegion)
	viper.Set("aws_default_region", defaultRegion)
	if err := viper.WriteConfig(); err != nil {
		// If config file doesn't exist yet, create it.
		if err := viper.SafeWriteConfig(); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
	}

	pterm.Info.Printf("SSO URL:        %s\n", ssoURL)
	pterm.Info.Printf("SSO Region:     %s\n", ssoRegion)
	pterm.Info.Printf("Default Region: %s\n", defaultRegion)

	// Search for all entitlements the user can see, expanding resource type info.
	entitlements, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
		GrantedStatus: shared.GrantedStatusAll,
		AppEntitlementExpandMask: shared.AppEntitlementExpandMask{
			Paths: []string{"app_id", "app_resource_type_id", "app_resource_id"},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to search entitlements: %w", err)
	}

	// Filter to AWS permission sets and create profiles.
	created := 0
	skipped := 0
	for _, ent := range entitlements {
		appResourceType := client.GetExpanded[shared.AppResourceType](ent, client.ExpandedAppResourceType)
		sdkEnt := shared.AppEntitlement(ent.Entitlement)
		if !client.IsAWSPermissionSet(&sdkEnt, appResourceType) {
			continue
		}

		appID := client.StringFromPtr(ent.Entitlement.AppID)
		entID := client.StringFromPtr(ent.Entitlement.ID)
		appResource := client.GetExpanded[shared.AppResource](ent, client.ExpandedAppResource)
		profileName, err := createAWSProfile(&sdkEnt, appResource, appID, entID, ssoURL, ssoRegion, defaultRegion)
		if err != nil {
			pterm.Warning.Printf("Skipping %s: %v\n", client.StringFromPtr(ent.Entitlement.DisplayName), err)
			skipped++
			continue
		}
		if profileName == "" {
			skipped++
			continue
		}
		pterm.Success.Printf("Created profile: %s\n", profileName)
		created++
	}

	if created == 0 && skipped == 0 {
		pterm.Info.Println("No AWS permission set entitlements found.")
	} else {
		pterm.Info.Printf("\nDone: %d profiles created, %d skipped (already exist or missing data).\n", created, skipped)
	}
	return nil
}

// createAWSProfile writes a single profile to ~/.aws/config.
// Returns the profile name if created, empty string if it already exists.
func createAWSProfile(entitlement *shared.AppEntitlement, resource *shared.AppResource, appID string, entitlementID string, ssoURL string, ssoRegion string, defaultRegion string) (string, error) {
	if entitlement == nil {
		return "", fmt.Errorf("entitlement is nil")
	}

	var accountID string
	for _, value := range entitlement.SourceConnectorIds {
		parts := strings.Split(value, "|")
		if len(parts) == 2 {
			accountID = parts[0]
			break
		}
	}
	if accountID == "" {
		return "", fmt.Errorf("could not extract AWS account ID from sourceConnectorIds")
	}

	if entitlement.DisplayName == nil {
		return "", fmt.Errorf("entitlement has no display name")
	}
	roleName := strings.Split(*entitlement.DisplayName, " ")[0]

	accountName := "aws"
	if resource != nil && resource.DisplayName != nil {
		accountName = strings.ToLower(strings.ReplaceAll(*resource.DisplayName, " ", "-"))
	}
	profileName := fmt.Sprintf("%s-%s", accountName, strings.ToLower(roleName))

	awsConfigDir := filepath.Join(os.Getenv("HOME"), ".aws")
	if err := os.MkdirAll(awsConfigDir, 0700); err != nil {
		return "", fmt.Errorf("failed to create ~/.aws: %w", err)
	}

	configPath := filepath.Join(awsConfigDir, "config")
	configContent, err := os.ReadFile(configPath) //nolint:gosec // trusted path ~/.aws/config
	if err != nil && !os.IsNotExist(err) {
		return "", fmt.Errorf("failed to read AWS config: %w", err)
	}

	configStr := string(configContent)

	// Skip if profile already exists.
	if strings.Contains(configStr, fmt.Sprintf("[profile %s]", profileName)) {
		return "", nil
	}

	ssoSessionExists := strings.Contains(configStr, "[sso-session cone-sso]")

	newConfig := fmt.Sprintf(`
[profile %s]
credential_process = cone aws credentials "%s"
cone_app_id = %s
cone_entitlement_id = %s
cone_sso_account_id = %s
cone_sso_role_name = %s
cone_sso_region = %s
cone_sso_start_url = %s
cone_sso_registration_scopes = sso:account:access
sso_session = cone-sso
region = %s
output = json
`, profileName, profileName, appID, entitlementID, accountID, roleName, ssoRegion, ssoURL, defaultRegion)

	if !ssoSessionExists {
		newConfig += fmt.Sprintf(`
[sso-session cone-sso]
sso_start_url = %s
sso_region = %s
sso_registration_scopes = sso:account:access
`, ssoURL, ssoRegion)
	}

	if err := os.WriteFile(configPath, append(configContent, []byte(newConfig)...), 0600); err != nil {
		return "", fmt.Errorf("failed to write AWS config: %w", err)
	}

	return profileName, nil
}

// --- cone aws credentials ---

// AWSCredentials is the JSON format expected by AWS credential_process.
type AWSCredentials struct {
	Version         int    `json:"Version"`
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
	Expiration      string `json:"Expiration"`
}

const autoRequestPollInterval = 3 * time.Second
const autoRequestPollTimeout = 90 * time.Second

func awsCredentialsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "credentials <profile-name>",
		Short: "Get temporary AWS credentials for a profile.",
		Long: `Retrieve temporary AWS credentials for an AWS SSO profile managed by cone.

This command checks ConductorOne for an active grant. If you don't have access,
it automatically submits an access request. If the request is auto-approved,
credentials are returned immediately. Otherwise it tells you the request is
pending approval.

Can be used directly or as an AWS credential_process:
  credential_process = cone aws credentials "my-profile"`,
		RunE: awsCredentialsRun,
	}
	return cmd
}

func awsCredentialsRun(cmd *cobra.Command, args []string) error {
	ctx, c, _, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}

	profileName := args[0]

	// Read entitlement IDs from the AWS profile config.
	awsConfigDir := filepath.Join(os.Getenv("HOME"), ".aws")
	configPath := filepath.Join(awsConfigDir, "config")
	configContent, err := os.ReadFile(configPath) //nolint:gosec // trusted path ~/.aws/config
	if err != nil {
		return fmt.Errorf("failed to read AWS config: %w", err)
	}
	profileConfig := extractProfileConfig(string(configContent), fmt.Sprintf("[profile %s]", profileName))

	appID := extractProfileValue(profileConfig, "cone_app_id")
	entitlementID := extractProfileValue(profileConfig, "cone_entitlement_id")
	if appID == "" || entitlementID == "" {
		return fmt.Errorf("profile %q is missing cone_app_id or cone_entitlement_id — re-run 'cone aws setup'", profileName)
	}

	accessResult, err := checkC1Access(ctx, c, appID, entitlementID)
	if err != nil {
		return fmt.Errorf("failed to check access: %w", err)
	}

	if !accessResult.hasAccess {
		// Check if the entitlement requires a request form — if so, the user must use cone get interactively.
		hasForm, err := c.HasRequestForm(ctx, appID, entitlementID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: could not check for request form: %v\n", err)
		}
		if hasForm {
			return fmt.Errorf("no active grant for %q. This entitlement requires a request form — request access with:\n  cone get --app-id %s --entitlement-id %s", profileName, appID, entitlementID)
		}

		// Fetch the entitlement to get its max grant duration.
		entitlement, err := c.GetEntitlement(ctx, appID, entitlementID)
		if err != nil {
			return fmt.Errorf("failed to get entitlement details: %w", err)
		}

		duration := ""
		if entitlement.DurationGrant != nil {
			duration = *entitlement.DurationGrant
		}

		fmt.Fprintf(os.Stderr, "No active grant for %q — submitting access request...\n", profileName)

		grantResp, err := c.CreateGrantTask(
			ctx,
			appID,
			entitlementID,
			accessResult.userID,
			"", // appUserId
			"Auto-requested via cone aws credentials", // justification
			duration,
			false, // emergencyAccess
			nil,   // requestData
		)
		if err != nil {
			// Handle duplicate task — extract the existing task ID from the error.
			if strings.Contains(err.Error(), "duplicate ticket found") {
				existingTaskID := extractDuplicateTaskID(err.Error())
				if existingTaskID != "" {
					dupNum := existingTaskID
					if dupTask, fetchErr := c.GetTask(ctx, existingTaskID); fetchErr == nil {
						dupNum = client.StringFromIntPtr(dupTask.TaskView.Task.NumericID)
					}
					fmt.Fprintf(os.Stderr, "A pending request already exists for %q.\n", profileName)
					fmt.Fprintf(os.Stderr, "Check status: cone task get %s\n", dupNum)
					fmt.Fprintf(os.Stderr, "Once resolved, retry the command.\n")
					return fmt.Errorf("duplicate request (task: %s)", dupNum)
				}
				return fmt.Errorf("a pending request already exists for %q. Check ConductorOne for status, then retry", profileName)
			}
			return fmt.Errorf("failed to submit access request: %w", err)
		}

		task := grantResp.TaskView.Task
		taskID := client.StringFromPtr(task.ID)
		taskNum := client.StringFromIntPtr(task.NumericID)
		fmt.Fprintf(os.Stderr, "Access request submitted (task: %s)\n", taskNum)

		// Poll until the task closes, requires human action, or we time out.
		var outcome string
		deadline := time.Now().Add(autoRequestPollTimeout)
		for time.Now().Before(deadline) {
			time.Sleep(autoRequestPollInterval)
			fmt.Fprintf(os.Stderr, ".")

			taskResp, err := c.GetTask(ctx, taskID)
			if err != nil {
				break
			}
			t := taskResp.TaskView.Task

			// Check the current step — if it needs human action, stop polling.
			if t.PolicyInstance != nil && t.PolicyInstance.PolicyStepInstance != nil {
				step := t.PolicyInstance.PolicyStepInstance
				switch {
				case step.ApprovalInstance != nil:
					fmt.Fprintf(os.Stderr, "\n")
					fmt.Fprintf(os.Stderr, "Request submitted for %q but requires approval.\n", profileName)
					fmt.Fprintf(os.Stderr, "Check status: cone task get %s\n", taskNum)
					fmt.Fprintf(os.Stderr, "Once approved, retry the command.\n")
					return fmt.Errorf("awaiting approval")
				case step.FormInstance != nil:
					fmt.Fprintf(os.Stderr, "\n")
					return fmt.Errorf("request submitted for %q but requires form input. Complete it with:\n  cone get --app-id %s --entitlement-id %s", profileName, appID, entitlementID)
				}
			}

			// Check if closed.
			if t.State == nil || *t.State != shared.TaskStateTaskStateClosed {
				continue
			}
			if t.TaskType.TaskTypeGrant != nil && t.TaskType.TaskTypeGrant.Outcome != nil {
				outcome = string(*t.TaskType.TaskTypeGrant.Outcome)
			}
			break
		}
		fmt.Fprintf(os.Stderr, "\n")

		switch outcome {
		case string(shared.TaskTypeGrantOutcomeGrantOutcomeGranted):
			fmt.Fprintf(os.Stderr, "Access granted!\n")
		case string(shared.TaskTypeGrantOutcomeGrantOutcomeDenied):
			return fmt.Errorf("access request for %q was denied", profileName)
		case string(shared.TaskTypeGrantOutcomeGrantOutcomeError):
			return fmt.Errorf("access request for %q encountered an error. Check ConductorOne for details", profileName)
		case string(shared.TaskTypeGrantOutcomeGrantOutcomeCancelled):
			return fmt.Errorf("access request for %q was cancelled", profileName)
		case "":
			return fmt.Errorf("request for %q is still processing. Retry the command shortly", profileName)
		default:
			return fmt.Errorf("access request for %q completed with outcome: %s", profileName, outcome)
		}
	}

	// We have access — fetch credentials using profile config already read above.
	accountID := extractProfileValue(profileConfig, "cone_sso_account_id")
	roleName := extractProfileValue(profileConfig, "cone_sso_role_name")
	ssoStartURL := extractProfileValue(profileConfig, "cone_sso_start_url")
	ssoRegion := extractProfileValue(profileConfig, "cone_sso_region")
	if ssoRegion == "" {
		ssoRegion = defaultAWSRegion
	}

	if accountID == "" || roleName == "" || ssoStartURL == "" {
		return fmt.Errorf("missing required SSO configuration for profile %s", profileName)
	}

	creds, credErr := getTemporaryCredentials(ctx, accountID, roleName, ssoStartURL, ssoRegion)

	// If credentials fail with a Forbidden/No access error, the AWS permission may still be provisioning.
	// Retry for up to 60 seconds.
	if credErr != nil && (strings.Contains(credErr.Error(), "ForbiddenException") || strings.Contains(credErr.Error(), "No access")) {
		fmt.Fprintf(os.Stderr, "Waiting for AWS permission to propagate...\n")
		credDeadline := time.Now().Add(60 * time.Second)
		for time.Now().Before(credDeadline) {
			time.Sleep(autoRequestPollInterval)
			fmt.Fprintf(os.Stderr, ".")
			creds, credErr = getTemporaryCredentials(ctx, accountID, roleName, ssoStartURL, ssoRegion)
			if credErr == nil {
				break
			}
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
	if credErr != nil {
		return credErr
	}

	jsonOutput, err := json.Marshal(creds)
	if err != nil {
		return fmt.Errorf("failed to marshal credentials: %w", err)
	}

	_, err = fmt.Fprintln(os.Stdout, string(jsonOutput))
	return err
}

// --- helpers ---

// extractDuplicateTaskID parses the task ID from a C1 "duplicate ticket found" error.
// The error JSON contains: "details":[{"@type":"...TaskRef", "id":"<task-id>"}].
func extractDuplicateTaskID(errMsg string) string {
	// Find the JSON body in the error message.
	jsonStart := strings.Index(errMsg, "{")
	if jsonStart == -1 {
		return ""
	}
	var errBody struct {
		Details []struct {
			ID string `json:"id"`
		} `json:"details"`
	}
	if err := json.Unmarshal([]byte(errMsg[jsonStart:]), &errBody); err != nil {
		return ""
	}
	if len(errBody.Details) > 0 && errBody.Details[0].ID != "" {
		return errBody.Details[0].ID
	}
	return ""
}

func extractProfileConfig(config, profileSection string) string {
	lines := strings.Split(config, "\n")
	var profileLines []string
	inProfile := false

	for _, line := range lines {
		if line == profileSection {
			inProfile = true
			continue
		}
		if inProfile {
			if strings.HasPrefix(line, "[") {
				break
			}
			profileLines = append(profileLines, line)
		}
	}

	return strings.Join(profileLines, "\n")
}

func extractProfileValue(profileConfig, key string) string {
	for _, line := range strings.Split(profileConfig, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), key) {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}

func getSSOToken(ssoStartURL string) (string, error) {
	cacheDir := filepath.Join(os.Getenv("HOME"), ".aws", "sso", "cache")
	files, err := os.ReadDir(cacheDir)
	if err != nil {
		return "", fmt.Errorf("failed to read SSO cache: %w", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		content, err := os.ReadFile(filepath.Join(cacheDir, file.Name())) //nolint:gosec // reading AWS SSO cache files
		if err != nil {
			continue
		}

		var cache struct {
			AccessToken string    `json:"accessToken"`
			ExpiresAt   time.Time `json:"expiresAt"`
			StartURL    string    `json:"startUrl"`
		}
		if err := json.Unmarshal(content, &cache); err != nil {
			continue
		}

		if cache.StartURL == ssoStartURL && cache.ExpiresAt.After(time.Now()) {
			return cache.AccessToken, nil
		}
	}

	return "", fmt.Errorf("no valid SSO token found for %s", ssoStartURL)
}

func requireAWSCLI() error {
	if _, err := exec.LookPath("aws"); err != nil {
		return fmt.Errorf("the AWS CLI is required but was not found on PATH — install it from https://aws.amazon.com/cli/")
	}
	return nil
}

func ssoLogin(ctx context.Context) error {
	fmt.Fprintf(os.Stderr, "AWS SSO session expired. Logging in...\n")
	loginCmd := exec.CommandContext(ctx, "aws", "sso", "login", "--sso-session", "cone-sso")
	loginCmd.Stdin = os.Stdin
	loginCmd.Stdout = os.Stderr
	loginCmd.Stderr = os.Stderr
	return loginCmd.Run()
}

func getRoleCredentials(ctx context.Context, token, accountID, roleName, ssoRegion string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "aws", "sso", "get-role-credentials",
		"--access-token", token,
		"--account-id", accountID,
		"--role-name", roleName,
		"--region", ssoRegion,
		"--output", "json")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%s", strings.TrimSpace(stderr.String()))
	}
	return stdout.Bytes(), nil
}

func getTemporaryCredentials(ctx context.Context, accountID, roleName, ssoStartURL, ssoRegion string) (*AWSCredentials, error) {
	if err := requireAWSCLI(); err != nil {
		return nil, err
	}

	token, err := getSSOToken(ssoStartURL)
	if err != nil {
		if loginErr := ssoLogin(ctx); loginErr != nil {
			return nil, fmt.Errorf("SSO login failed: %w", loginErr)
		}
		token, err = getSSOToken(ssoStartURL)
		if err != nil {
			return nil, fmt.Errorf("failed to get SSO token after login: %w", err)
		}
	}

	output, err := getRoleCredentials(ctx, token, accountID, roleName, ssoRegion)
	if err != nil {
		// Token might be cached but invalid — retry with fresh login.
		if strings.Contains(err.Error(), "UnauthorizedException") || strings.Contains(err.Error(), "Session token not found") {
			if loginErr := ssoLogin(ctx); loginErr != nil {
				return nil, fmt.Errorf("SSO login failed: %w", loginErr)
			}
			token, err = getSSOToken(ssoStartURL)
			if err != nil {
				return nil, fmt.Errorf("failed to get SSO token after login: %w", err)
			}
			output, err = getRoleCredentials(ctx, token, accountID, roleName, ssoRegion)
			if err != nil {
				return nil, fmt.Errorf("failed to get credentials after re-login: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to get credentials: %w", err)
		}
	}

	var response struct {
		RoleCredentials struct {
			AccessKeyID     string `json:"accessKeyId"`
			SecretAccessKey string `json:"secretAccessKey"`
			SessionToken    string `json:"sessionToken"`
			Expiration      int64  `json:"expiration"`
		} `json:"roleCredentials"`
	}
	if err := json.Unmarshal(output, &response); err != nil {
		return nil, fmt.Errorf("failed to parse credentials: %w", err)
	}

	return &AWSCredentials{
		Version:         1,
		AccessKeyID:     response.RoleCredentials.AccessKeyID,
		SecretAccessKey: response.RoleCredentials.SecretAccessKey,
		SessionToken:    response.RoleCredentials.SessionToken,
		Expiration:      time.UnixMilli(response.RoleCredentials.Expiration).Format(time.RFC3339),
	}, nil
}

type accessCheckResult struct {
	hasAccess bool
	userID    string
}

func checkC1Access(ctx context.Context, c client.C1Client, appID string, entitlementID string) (*accessCheckResult, error) {
	userInfo, err := c.AuthIntrospect(ctx)
	if err != nil {
		return nil, err
	}
	userID := client.StringFromPtr(userInfo.UserID)

	result := &accessCheckResult{
		userID: userID,
	}

	grants, err := c.GetGrantsForIdentity(ctx, appID, entitlementID, userID)
	if err != nil {
		return result, err
	}

	for _, grant := range grants {
		if grant.CreatedAt != nil && grant.DeletedAt == nil {
			result.hasAccess = true
			return result, nil
		}
	}

	return result, nil
}
