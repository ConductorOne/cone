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

		appResource := client.GetExpanded[shared.AppResource](ent, client.ExpandedAppResource)
		profileName, err := createAWSProfile(&sdkEnt, appResource, ssoURL, ssoRegion, defaultRegion)
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
func createAWSProfile(entitlement *shared.AppEntitlement, resource *shared.AppResource, ssoURL string, ssoRegion string, defaultRegion string) (string, error) {
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
	configContent, err := os.ReadFile(configPath)
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
cone_sso_account_id = %s
cone_sso_role_name = %s
cone_sso_region = %s
cone_sso_start_url = %s
cone_sso_registration_scopes = sso:account:access
sso_session = cone-sso
region = %s
output = json
`, profileName, profileName, accountID, roleName, ssoRegion, ssoURL, defaultRegion)

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

	accessResult, err := checkC1Access(ctx, c, profileName)
	if err != nil {
		return fmt.Errorf("failed to check access: %w", err)
	}

	if !accessResult.hasAccess {
		if accessResult.appID == "" || accessResult.entitlementID == "" {
			return fmt.Errorf("no entitlement found matching profile %q", profileName)
		}

		// Fetch the entitlement to get its max grant duration.
		entitlement, err := c.GetEntitlement(ctx, accessResult.appID, accessResult.entitlementID)
		if err != nil {
			return fmt.Errorf("failed to get entitlement details: %w", err)
		}

		// Use the entitlement's max duration if set, otherwise leave empty (unlimited).
		duration := ""
		if entitlement.DurationGrant != nil {
			duration = *entitlement.DurationGrant
		}

		fmt.Fprintf(os.Stderr, "No active grant for %q — submitting access request...\n", profileName)

		grantResp, err := c.CreateGrantTask(
			ctx,
			accessResult.appID,
			accessResult.entitlementID,
			accessResult.userID,
			"", // appUserId
			"Auto-requested via cone aws credentials", // justification
			duration, // use entitlement's max duration
			false,    // emergencyAccess
			nil,      // requestData
		)
		if err != nil {
			return fmt.Errorf("failed to submit access request: %w", err)
		}

		taskID := client.StringFromPtr(grantResp.TaskView.Task.ID)
		fmt.Fprintf(os.Stderr, "Access request submitted (task: %s). Checking for auto-approval...\n", taskID)

		// Poll for approval.
		granted := false
		deadline := time.Now().Add(autoRequestPollTimeout)
		for time.Now().Before(deadline) {
			time.Sleep(autoRequestPollInterval)
			fmt.Fprintf(os.Stderr, ".")

			taskResp, err := c.GetTask(ctx, taskID)
			if err != nil {
				break
			}
			task := taskResp.TaskView.Task
			if task.State == nil {
				continue
			}
			if *task.State != shared.TaskStateTaskStateClosed {
				continue
			}
			// Task is closed — check if it was granted.
			if task.TaskType.TaskTypeGrant != nil && task.TaskType.TaskTypeGrant.Outcome != nil {
				if *task.TaskType.TaskTypeGrant.Outcome == shared.TaskTypeGrantOutcomeGrantOutcomeGranted {
					granted = true
				}
			}
			break
		}
		fmt.Fprintf(os.Stderr, "\n")

		if !granted {
			fmt.Fprintf(os.Stderr, "Request is pending approval. Retry the command after it's approved.\n")
			return fmt.Errorf("access request pending for %q", profileName)
		}

		fmt.Fprintf(os.Stderr, "Access granted!\n")
	}

	// We have access — fetch credentials.
	awsConfigDir := filepath.Join(os.Getenv("HOME"), ".aws")
	configPath := filepath.Join(awsConfigDir, "config")

	configContent, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read AWS config: %w", err)
	}

	profileConfig := extractProfileConfig(string(configContent), fmt.Sprintf("[profile %s]", profileName))

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

	creds, err := getTemporaryCredentials(accountID, roleName, ssoStartURL, ssoRegion)
	if err != nil {
		return err
	}

	jsonOutput, err := json.Marshal(creds)
	if err != nil {
		return fmt.Errorf("failed to marshal credentials: %w", err)
	}

	fmt.Fprintln(os.Stdout, string(jsonOutput))
	return nil
}

// --- helpers ---

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
		content, err := os.ReadFile(filepath.Join(cacheDir, file.Name()))
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

func ssoLogin() error {
	fmt.Fprintf(os.Stderr, "AWS SSO session expired. Logging in...\n")
	loginCmd := exec.Command("aws", "sso", "login", "--sso-session", "cone-sso")
	loginCmd.Stdin = os.Stdin
	loginCmd.Stdout = os.Stderr
	loginCmd.Stderr = os.Stderr
	return loginCmd.Run()
}

func getRoleCredentials(token, accountID, roleName, ssoRegion string) ([]byte, error) {
	cmd := exec.Command("aws", "sso", "get-role-credentials",
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

func getTemporaryCredentials(accountID, roleName, ssoStartURL, ssoRegion string) (*AWSCredentials, error) {
	if err := requireAWSCLI(); err != nil {
		return nil, err
	}

	token, err := getSSOToken(ssoStartURL)
	if err != nil {
		if loginErr := ssoLogin(); loginErr != nil {
			return nil, fmt.Errorf("SSO login failed: %w", loginErr)
		}
		token, err = getSSOToken(ssoStartURL)
		if err != nil {
			return nil, fmt.Errorf("failed to get SSO token after login: %w", err)
		}
	}

	output, err := getRoleCredentials(token, accountID, roleName, ssoRegion)
	if err != nil {
		// Token might be cached but invalid — retry with fresh login.
		if strings.Contains(err.Error(), "UnauthorizedException") || strings.Contains(err.Error(), "Session token not found") {
			if loginErr := ssoLogin(); loginErr != nil {
				return nil, fmt.Errorf("SSO login failed: %w", loginErr)
			}
			token, err = getSSOToken(ssoStartURL)
			if err != nil {
				return nil, fmt.Errorf("failed to get SSO token after login: %w", err)
			}
			output, err = getRoleCredentials(token, accountID, roleName, ssoRegion)
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
	hasAccess     bool
	appID         string
	entitlementID string
	userID        string
}

func checkC1Access(ctx context.Context, c client.C1Client, profileName string) (*accessCheckResult, error) {
	userInfo, err := c.AuthIntrospect(ctx)
	if err != nil {
		return nil, err
	}
	userID := client.StringFromPtr(userInfo.UserID)

	entitlements, err := c.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
		EntitlementAlias: profileName,
		GrantedStatus:    shared.GrantedStatusAll,
	})
	if err != nil {
		return nil, err
	}

	result := &accessCheckResult{
		userID: userID,
	}

	for _, ent := range entitlements {
		appID := client.StringFromPtr(ent.Entitlement.AppID)
		entID := client.StringFromPtr(ent.Entitlement.ID)

		// Track the first matching entitlement for request submission.
		if result.appID == "" {
			result.appID = appID
			result.entitlementID = entID
		}

		grants, err := c.GetGrantsForIdentity(ctx, appID, entID, userID)
		if err != nil {
			continue
		}
		for _, grant := range grants {
			if grant.CreatedAt != nil && grant.DeletedAt == nil {
				result.hasAccess = true
				result.appID = appID
				result.entitlementID = entID
				return result, nil
			}
		}
	}

	return result, nil
}
