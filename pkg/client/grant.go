package client

import (
	"context"
)

func (c *client) GetGrantsForIdentity(ctx context.Context, appID string, appEntitlementID string, appUserID string) ([]c1api.C1ApiAppV1AppEntitlementUserBinding, error) {
	resp, httpResp, err := c.apiClient.DefaultAPI.C1ApiAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant(ctx, appID, appEntitlementID, appUserID).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	return resp.Bindings, nil
}
