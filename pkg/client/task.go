package client

import (
	"context"

	"github.com/conductorone/cone/internal/c1api"
)

func (c *client) CreateGrantTask(ctx context.Context, appId string, appEntitlementId string, identityUserId string) (*c1api.C1ApiTaskV1TaskServiceGetResponse, error) {
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
