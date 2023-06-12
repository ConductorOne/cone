package client

import (
	"context"

	"github.com/conductorone/cone/internal/c1api"
)

func (c *client) GetGrantsForIdentity(ctx context.Context, appID string, appEntitlementID string, appUserID string) ([]c1api.C1ApiAppV1AppEntitlementUserBinding, error) {
	resp, httpResp, err := c.apiClient.AppEntitlementUserBindingAPI.C1ApiAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant(ctx, appID, appEntitlementID, appUserID).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	return resp.Bindings, nil
}
