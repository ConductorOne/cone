package client

import (
	"context"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
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
	justification string,
	duration string,
	emergencyAccess bool,
) (*shared.TaskServiceCreateGrantResponse, error) {
	req := shared.TaskServiceCreateGrantRequest{
		AppEntitlementID: appEntitlementId,
		IdentityUserID:   &identityUserId,
		AppID:            appId,
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

func (c *client) SearchTasks(ctx context.Context, taskFilter *shared.TaskSearchRequest) (*shared.TaskSearchResponse, error) {
	resp, err := c.sdk.TaskSearch.Search(ctx, taskFilter)
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
