package client

import (
	"context"

	"github.com/conductorone/cone/internal/c1api"
)

func (c *client) GetTask(ctx context.Context, taskId string) (*c1api.C1ApiTaskV1TaskServiceGetResponse, error) {
	task, resp, err := c.apiClient.DefaultAPI.C1ApiTaskV1TaskServiceGet(ctx, taskId).Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return task, err
}

func (c *client) CreateGrantTask(
	ctx context.Context,
	appId string,
	appEntitlementId string,
	identityUserId string,
	justification string,
	duration string,
) (*c1api.C1ApiTaskV1TaskServiceCreateGrantResponse, error) {
	api := c.apiClient.DefaultAPI.C1ApiTaskV1TaskServiceCreateGrantTask(ctx)
	grantReq := c1api.C1ApiTaskV1TaskServiceCreateGrantRequest{
		AppEntitlementId: &appEntitlementId,
		IdentityUserId:   &identityUserId,
		AppId:            &appId,
		Description:      &justification,
	}
	if duration != "" {
		grantReq.GrantDuration = &duration
	}
	req := api.C1ApiTaskV1TaskServiceCreateGrantRequest(grantReq)

	cgtResp, resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return cgtResp, nil
}

func (c *client) CreateRevokeTask(
	ctx context.Context,
	appId string,
	appEntitlementId string,
	identityUserId string,
	justification string,
) (*c1api.C1ApiTaskV1TaskServiceCreateRevokeResponse, error) {
	api := c.apiClient.DefaultAPI.C1ApiTaskV1TaskServiceCreateRevokeTask(ctx)
	req := api.C1ApiTaskV1TaskServiceCreateRevokeRequest(c1api.C1ApiTaskV1TaskServiceCreateRevokeRequest{
		AppEntitlementId: &appEntitlementId,
		IdentityUserId:   &identityUserId,
		AppId:            &appId,
		Description:      &justification,
	})
	cgtResp, resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return cgtResp, nil
}

func (c *client) SearchTasks(ctx context.Context, taskFilter c1api.C1ApiTaskV1TaskSearchRequest) (*c1api.C1ApiTaskV1TaskSearchResponse, error) {
	api := c.apiClient.DefaultAPI.C1ApiTaskV1TaskSearchServiceSearch(ctx)
	req := api.C1ApiTaskV1TaskSearchRequest(taskFilter)
	apiResp, resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return apiResp, nil
}

func (c *client) CommentOnTask(ctx context.Context, taskID string, comment string) (*c1api.C1ApiTaskV1TaskActionsServiceCommentResponse, error) {
	api := c.apiClient.DefaultAPI.C1ApiTaskV1TaskActionsServiceComment(ctx, taskID)
	req := api.C1ApiTaskV1TaskActionsServiceCommentRequestInput(c1api.C1ApiTaskV1TaskActionsServiceCommentRequestInput{
		Comment: &comment,
	})
	cmntResp, resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return cmntResp, nil
}

func (c *client) ApproveTask(ctx context.Context, taskId string, comment string, policyId string) (*c1api.C1ApiTaskV1TaskActionsServiceApproveResponse, error) {
	api := c.apiClient.DefaultAPI.C1ApiTaskV1TaskActionsServiceApprove(ctx, taskId)
	approveReq := c1api.C1ApiTaskV1TaskActionsServiceApproveRequestInput{
		PolicyStepId: &policyId,
	}
	if comment != "" {
		approveReq.Comment = &comment
	}
	req := api.C1ApiTaskV1TaskActionsServiceApproveRequestInput(approveReq)
	approveResp, resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return approveResp, err
}

func (c *client) DenyTask(ctx context.Context, taskId string, comment string, policyId string) (*c1api.C1ApiTaskV1TaskActionsServiceDenyResponse, error) {
	api := c.apiClient.DefaultAPI.C1ApiTaskV1TaskActionsServiceDeny(ctx, taskId)
	denyReq := c1api.C1ApiTaskV1TaskActionsServiceDenyRequestInput{
		PolicyStepId: &policyId,
	}
	if comment != "" {
		denyReq.Comment = &comment
	}
	req := api.C1ApiTaskV1TaskActionsServiceDenyRequestInput(denyReq)
	denyResp, resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return denyResp, err
}
