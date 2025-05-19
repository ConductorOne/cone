package client

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/spf13/viper"
)

func (c *client) GetTask(ctx context.Context, taskId string) (*shared.TaskServiceGetResponse, error) {
	resp, err := c.sdk.Task.Get(ctx, operations.C1APITaskV1TaskServiceGetRequest{ID: taskId})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	return resp.TaskServiceGetResponse, err
}

func (c *client) CreateGrantTask(
	ctx context.Context,
	appId string,
	appEntitlementId string,
	identityUserId string,
	appuserId string,
	justification string,
	duration string,
	emergencyAccess bool,
) (*shared.TaskServiceCreateGrantResponse, error) {
	req := shared.TaskServiceCreateGrantRequest{
		AppEntitlementID: appEntitlementId,
		IdentityUserID:   &identityUserId,
		AppID:            appId,
		AppUserID:        &appuserId,
		Description:      &justification,
		EmergencyAccess:  &emergencyAccess,
	}
	if duration != "" {
		req.GrantDuration = &duration
	}
	resp, err := c.sdk.Task.CreateGrantTask(ctx, &req)
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	return resp.TaskServiceCreateGrantResponse, nil
}

func (c *client) CreateRevokeTask(
	ctx context.Context,
	appId string,
	appEntitlementId string,
	identityUserId string,
	justification string,
) (*shared.TaskServiceCreateRevokeResponse, error) {
	req := shared.TaskServiceCreateRevokeRequest{
		AppEntitlementID: appEntitlementId,
		IdentityUserID:   &identityUserId,
		AppID:            appId,
		Description:      &justification,
	}
	resp, err := c.sdk.Task.CreateRevokeTask(ctx, &req)
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	return resp.TaskServiceCreateRevokeResponse, nil
}

func (c *client) SearchTasks(ctx context.Context, taskFilter shared.TaskSearchRequest) (*shared.TaskSearchResponse, error) {
	resp, err := c.sdk.TaskSearch.Search(ctx, &taskFilter)
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	return resp.TaskSearchResponse, nil
}

func (c *client) CommentOnTask(ctx context.Context, taskID string, comment string) (*shared.TaskActionsServiceCommentResponse, error) {
	resp, err := c.sdk.TaskActions.Comment(ctx, operations.C1APITaskV1TaskActionsServiceCommentRequest{
		TaskActionsServiceCommentRequest: &shared.TaskActionsServiceCommentRequest{
			Comment: &comment,
		},
		TaskID: taskID,
	})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	return resp.TaskActionsServiceCommentResponse, nil
}

func (c *client) ApproveTask(ctx context.Context, taskId string, comment string, policyId string) (*shared.TaskActionsServiceApproveResponse, error) {
	resp, err := c.sdk.TaskActions.Approve(ctx, operations.C1APITaskV1TaskActionsServiceApproveRequest{
		TaskActionsServiceApproveRequest: &shared.TaskActionsServiceApproveRequest{
			Comment:      &comment,
			PolicyStepID: policyId,
		},
		TaskID: taskId,
	})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	return resp.TaskActionsServiceApproveResponse, nil
}

func (c *client) DenyTask(ctx context.Context, taskId string, comment string, policyId string) (*shared.TaskActionsServiceDenyResponse, error) {
	resp, err := c.sdk.TaskActions.Deny(ctx, operations.C1APITaskV1TaskActionsServiceDenyRequest{
		TaskActionsServiceDenyRequest: &shared.TaskActionsServiceDenyRequest{
			Comment:      &comment,
			PolicyStepID: &policyId,
		},
		TaskID: taskId,
	})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	return resp.TaskActionsServiceDenyResponse, nil
}

func (c *client) EscalateTask(ctx context.Context, taskID string) (*shared.TaskServiceActionResponse, error) {
	resp, err := c.sdk.TaskActions.EscalateToEmergencyAccess(ctx, operations.C1APITaskV1TaskActionsServiceEscalateToEmergencyAccessRequest{
		TaskActionsServiceEscalateToEmergencyAccessRequest: &shared.TaskActionsServiceEscalateToEmergencyAccessRequest{},
		TaskID: taskID,
	})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	return resp.TaskServiceActionResponse, nil
}

func IsAWSPermissionSet(entitlement *shared.AppEntitlement, resourceType *shared.AppResourceType) bool {
	if entitlement == nil || resourceType == nil {
		return false
	}

	// Check resource type display name
	if resourceType.DisplayName != nil {
		if strings.Contains(strings.ToLower(*resourceType.DisplayName), "aws permission set") {
			return true
		}
	}

	// Check SourceConnectorIds for AWS SSO permission set ARNs
	if entitlement.SourceConnectorIds != nil {
		for _, value := range entitlement.SourceConnectorIds {
			if strings.Contains(value, "arn:aws:sso:::permissionSet/") {
				return true
			}
		}
	}

	return false
}

// CreateAWSSSOProfile creates an AWS SSO profile for a permission set.
func CreateAWSSSOProfile(entitlement *shared.AppEntitlement, resource *shared.AppResource) error {
	if entitlement == nil || resource == nil {
		return errors.New("entitlement and resource are required")
	}

	// Get AWS account ID and permission set ARN from sourceConnectorIds
	var accountID, permissionSetARN string
	for _, value := range entitlement.SourceConnectorIds {
		parts := strings.Split(value, "|")
		if len(parts) == 2 {
			accountID = parts[0]
			permissionSetARN = parts[1]
			break
		}
	}

	if accountID == "" || permissionSetARN == "" {
		return errors.New("could not find AWS account ID or permission set ARN in sourceConnectorIds")
	}

	// Extract role name from entitlement display name (everything before first space)
	roleName := strings.Split(*entitlement.DisplayName, " ")[0]

	// Get AWS account name from resource display name
	accountName := "aws"
	if resource.DisplayName != nil {
		accountName = strings.ToLower(strings.ReplaceAll(*resource.DisplayName, " ", "-"))
	}

	// Create a profile/session name based on the account name and role name
	profileName := fmt.Sprintf("%s-%s", accountName, strings.ToLower(roleName))

	// Get the SSO start URL from Viper config
	ssoStartURL := viper.GetString("aws_sso_start_url")
	if ssoStartURL == "" {
		return errors.New("AWS SSO start URL is not configured. Please run 'cone config-aws set-sso-url <url>'")
	}

	// Create the AWS config directory if it doesn't exist
	awsConfigDir := filepath.Join(os.Getenv("HOME"), ".aws")
	if err := os.MkdirAll(awsConfigDir, 0700); err != nil {
		return fmt.Errorf("failed to create AWS config directory: %w", err)
	}

	// Read existing config file
	configPath := filepath.Join(awsConfigDir, "config")
	configContent, err := os.ReadFile(configPath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read AWS config file: %w", err)
	}

	// Check if profile already exists
	configStr := string(configContent)
	if strings.Contains(configStr, fmt.Sprintf("[profile %s]", profileName)) {
		return fmt.Errorf("AWS profile '%s' already exists", profileName)
	}

	// Check if SSO session already exists
	ssoSessionExists := strings.Contains(configStr, "[sso-session cone-sso]")

	// Create new profile configuration
	newProfile := fmt.Sprintf(`
[profile %s]
credential_process = cone aws-credentials "%s"
cone_sso_account_id = %s
cone_sso_role_name = %s
cone_sso_region = us-east-1
cone_sso_start_url = %s
cone_sso_registration_scopes = sso:account:access
sso_session = cone-sso
region = us-east-1
output = json
`, profileName, profileName, accountID, roleName, ssoStartURL)

	// Add SSO session configuration if it doesn't exist
	if !ssoSessionExists {
		newProfile += fmt.Sprintf(`
[sso-session cone-sso]
sso_start_url = %s
sso_region = us-east-1
sso_registration_scopes = sso:account:access
`, ssoStartURL)
	}

	// Append new profile to config file
	if err := os.WriteFile(configPath, append(configContent, []byte(newProfile)...), 0600); err != nil {
		return fmt.Errorf("failed to write AWS config file: %w", err)
	}

	return nil
}
