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
