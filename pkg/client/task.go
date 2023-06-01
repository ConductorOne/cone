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

func (c *client) CreateGrantTask(ctx context.Context, appId string, appEntitlementId string, identityUserId string) (*c1api.C1ApiTaskV1TaskServiceCreateGrantResponse, error) {
	api := c.apiClient.DefaultAPI.C1ApiTaskV1TaskServiceCreateGrantTask(ctx)
	req := api.C1ApiTaskV1TaskServiceCreateGrantRequest(c1api.C1ApiTaskV1TaskServiceCreateGrantRequest{
		AppEntitlementId: &appEntitlementId,
		IdentityUserId:   &identityUserId,
		AppId:            &appId,
	})
	cgtResp, resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return cgtResp, nil
}

func (c *client) CreateRevokeTask(ctx context.Context, appId string, appEntitlementId string, identityUserId string) (*c1api.C1ApiTaskV1TaskServiceCreateRevokeResponse, error) {
	api := c.apiClient.DefaultAPI.C1ApiTaskV1TaskServiceCreateRevokeTask(ctx)
	req := api.C1ApiTaskV1TaskServiceCreateRevokeRequest(c1api.C1ApiTaskV1TaskServiceCreateRevokeRequest{
		AppEntitlementId: &appEntitlementId,
		IdentityUserId:   &identityUserId,
		AppId:            &appId,
	})
	cgtResp, resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return cgtResp, nil
}
