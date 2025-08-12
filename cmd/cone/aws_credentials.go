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

// AWSCredentials represents the structure of AWS temporary credentials.
// that will be output in JSON format.
type AWSCredentials struct {
	Version         int    `json:"Version"`
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
	Expiration      string `json:"Expiration"`
}

// RoleCredentialsResponse represents the response from AWS SSO get-role-credentials API.
type RoleCredentialsResponse struct {
	RoleCredentials struct {
		AccessKeyID     string `json:"accessKeyId"`
		SecretAccessKey string `json:"secretAccessKey"`
		SessionToken    string `json:"sessionToken"`
		Expiration      int64  `json:"expiration"`
	} `json:"roleCredentials"`
}

// awsCredentialsCmd creates the cobra command for getting AWS credentials.
// Usage: cone aws-credentials <profile-name>.
func awsCredentialsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aws-credentials <profile-name>",
		Short: "Get AWS credentials for a profile",
		RunE:  awsCredentialsRun,
	}
	return cmd
}

// awsCredentialsRun is the main function that handles getting AWS credentials.
// It verifies access, reads AWS config, and retrieves temporary credentials.
func awsCredentialsRun(cmd *cobra.Command, args []string) error {
	ctx, _, _, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}

	profileName := args[0]

	// Check if user has access to this permission set
	hasAccess, err := checkC1Access(ctx, profileName)
	if err != nil {
		return fmt.Errorf("failed to check C1 access: %w", err)
	}

	if !hasAccess {
		fmt.Fprintf(os.Stderr, "You do not have access to this permission set.\n")
		fmt.Fprintf(os.Stderr, "To request access, run: cone get %s --wait\n", profileName)
		fmt.Fprintf(os.Stderr, "This will allow you to specify justification and duration for your access request.\n")
		return fmt.Errorf("access denied: please request access using 'cone get %s --wait'", profileName)
	}

	awsConfigDir := filepath.Join(os.Getenv("HOME"), ".aws")
	configPath := filepath.Join(awsConfigDir, "config")

	configContent, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read AWS config: %w", err)
	}

	configStr := string(configContent)
	profileSection := fmt.Sprintf("[profile %s]", profileName)
	profileConfig := extractProfileConfig(configStr, profileSection)

	accountID := extractConeSSOAccountID(profileConfig)
	roleName := extractConeSSORoleName(profileConfig)
	ssoStartURL := extractConeSSOStartURL(profileConfig)
	ssoRegion := extractConeSSORegion(profileConfig)

	if accountID == "" || roleName == "" || ssoStartURL == "" {
		return fmt.Errorf("missing required SSO configuration for profile %s", profileName)
	}

	if err := verifySSOSession(ssoStartURL, ssoRegion); err != nil {
		return fmt.Errorf("SSO session verification failed: %w", err)
	}

	creds, err := getTemporaryCredentials(accountID, roleName)
	if err != nil {
		return fmt.Errorf("failed to get temporary credentials: %w", err)
	}

	output := AWSCredentials{
		Version:         1,
		AccessKeyID:     creds.AccessKeyID,
		SecretAccessKey: creds.SecretAccessKey,
		SessionToken:    creds.SessionToken,
		Expiration:      creds.Expiration,
	}

	jsonOutput, err := json.Marshal(output)
	if err != nil {
		return fmt.Errorf("failed to marshal credentials: %w", err)
	}

	pterm.Println(string(jsonOutput))
	return nil
}

// extractProfileConfig extracts the configuration section for a specific AWS profile.
// from the AWS config file.
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

// extractConeSSOAccountID extracts the AWS account ID from the profile configuration.
// This is used to identify which AWS account to get credentials for.
func extractConeSSOAccountID(profileConfig string) string {
	for _, line := range strings.Split(profileConfig, "\n") {
		if strings.HasPrefix(line, "cone_sso_account_id") {
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}

// extractConeSSORoleName extracts the AWS role name from the profile configuration.
// This is the role that will be assumed when getting credentials.
func extractConeSSORoleName(profileConfig string) string {
	for _, line := range strings.Split(profileConfig, "\n") {
		if strings.HasPrefix(line, "cone_sso_role_name") {
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}

// extractConeSSOStartURL extracts the AWS SSO start URL from the profile configuration.
// This is the URL used to initiate the SSO login process.
func extractConeSSOStartURL(profileConfig string) string {
	for _, line := range strings.Split(profileConfig, "\n") {
		if strings.HasPrefix(line, "cone_sso_start_url") {
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}

// extractConeSSORegion extracts the AWS region from the profile configuration.
// Defaults to us-east-1 if not specified.
func extractConeSSORegion(profileConfig string) string {
	for _, line := range strings.Split(profileConfig, "\n") {
		if strings.HasPrefix(line, "cone_sso_region") {
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return "us-east-1" // Default region
}

// getSSOToken retrieves a valid SSO token from the AWS SSO cache.
// It looks for a token that matches the given start URL and hasn't expired.
func getSSOToken(ssoStartURL string) (string, error) {
	cacheDir := filepath.Join(os.Getenv("HOME"), ".aws", "sso", "cache")
	files, err := os.ReadDir(cacheDir)
	if err != nil {
		return "", fmt.Errorf("failed to read SSO cache directory: %w", err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
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
	}

	return "", fmt.Errorf("no valid SSO token found for %s", ssoStartURL)
}

// getTemporaryCredentials retrieves temporary AWS credentials using AWS SSO.
// It handles the SSO login process if needed and returns the credentials.
func getTemporaryCredentials(accountID, roleName string) (*AWSCredentials, error) {
	ssoStartURL := viper.GetString("aws_sso_start_url")
	if ssoStartURL == "" {
		return nil, fmt.Errorf("missing AWS SSO URL. Please run 'cone config-aws set-sso-url <url>' first")
	}

	token, err := getSSOToken(ssoStartURL)
	if err != nil {
		loginCmd := exec.Command("aws", "sso", "login", "--sso-session", "cone-sso")
		loginCmd.Stdout = nil
		loginCmd.Stderr = nil
		_ = loginCmd.Run() // ignore output, just try to login
		token, err = getSSOToken(ssoStartURL)
		if err != nil {
			return nil, fmt.Errorf("failed to get token after login: %w", err)
		}
	}

	cmd := exec.Command("aws", "sso", "get-role-credentials",
		"--access-token", token,
		"--account-id", accountID,
		"--role-name", roleName,
		"--region", "us-east-1",
		"--output", "json")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		if strings.Contains(stderr.String(), "AccessDenied") {
			return nil, fmt.Errorf("access denied: you don't have access to this role")
		}
		return nil, fmt.Errorf("failed to get credentials: %w\nCommand output: %s\nError output: %s",
			err, stdout.String(), stderr.String())
	}

	var response RoleCredentialsResponse
	if err := json.Unmarshal(stdout.Bytes(), &response); err != nil {
		return nil, fmt.Errorf("failed to parse credentials: %w\nCommand output: %s",
			err, stdout.String())
	}

	creds := &AWSCredentials{
		Version:         1,
		AccessKeyID:     response.RoleCredentials.AccessKeyID,
		SecretAccessKey: response.RoleCredentials.SecretAccessKey,
		SessionToken:    response.RoleCredentials.SessionToken,
		Expiration:      time.UnixMilli(response.RoleCredentials.Expiration).Format(time.RFC3339),
	}

	return creds, nil
}

// checkC1Access verifies if the user has access to the requested AWS profile.
// by checking their grants in ConductorOne.
func checkC1Access(ctx context.Context, profileName string) (bool, error) {
	// Create a temporary command with the necessary flags for cmdContext
	cmd := &cobra.Command{
		Use: "temp",
	}
	cmd.PersistentFlags().StringP("profile", "p", "default", "The config profile to use.")
	cmd.PersistentFlags().BoolP("non-interactive", "i", false, "Disable prompts.")
	cmd.PersistentFlags().String("client-id", "", "Client ID")
	cmd.PersistentFlags().String("client-secret", "", "Client secret")
	cmd.PersistentFlags().String("api-endpoint", "", "Override the API endpoint")
	cmd.PersistentFlags().StringP("output", "o", "table", "Output format. Valid values: table, json, json-pretty, wide.")
	cmd.PersistentFlags().Bool("debug", false, "Enable debug logging")
	cmd.SetContext(ctx)
	_, c1Client, _, err := cmdContext(cmd)
	if err != nil {
		return false, fmt.Errorf("failed to get C1 client: %w", err)
	}

	// Get current user ID
	userIntro, err := c1Client.AuthIntrospect(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get user info: %w", err)
	}
	userID := client.StringFromPtr(userIntro.UserID)

	// Search for the entitlement by alias (profile name)
	entitlements, err := c1Client.SearchEntitlements(ctx, &client.SearchEntitlementsFilter{
		EntitlementAlias: profileName,
		GrantedStatus:    shared.GrantedStatusAll,
	})
	if err != nil {
		return false, fmt.Errorf("failed to search entitlements: %w", err)
	}

	if len(entitlements) == 0 {
		return false, fmt.Errorf("no entitlements found matching profile name: %s", profileName)
	}

	// Check grants for each matching entitlement
	for _, entitlement := range entitlements {
		grants, err := c1Client.GetGrantsForIdentity(ctx, client.StringFromPtr(entitlement.Entitlement.AppID), client.StringFromPtr(entitlement.Entitlement.ID), userID)
		if err != nil {
			return false, fmt.Errorf("failed to check grants: %w", err)
		}

		// Check if user has an active grant
		for _, grant := range grants {
			if grant.CreatedAt != nil && grant.DeletedAt == nil {
				return true, nil
			}
		}
	}

	return false, nil
}

// verifySSOSession checks if the AWS SSO session is properly configured.
// in the AWS config file.
func verifySSOSession(ssoStartURL, ssoRegion string) error {
	awsConfigDir := filepath.Join(os.Getenv("HOME"), ".aws")
	configPath := filepath.Join(awsConfigDir, "config")
	configContent, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read AWS config: %w", err)
	}

	configStr := string(configContent)
	sessionSection := "[sso-session cone-sso]"
	if !strings.Contains(configStr, sessionSection) {
		return fmt.Errorf("SSO session configuration not found in AWS config")
	}

	return nil
}
