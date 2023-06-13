package client

import (
	"context"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func (c *client) GetGrantsForIdentity(ctx context.Context, appID string, appEntitlementID string, identityID string) ([]shared.AppEntitlementUserBinding, error) {
	resp, err := c.sdk.AppEntitlementUserBinding.ListAppUsersForIdentityWithGrant(ctx, operations.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantRequest{
		AppEntitlementID: appEntitlementID,
		AppID:            appID,
		IdentityUserID:   identityID,
	})
	if err != nil {
		return nil, err
	}
	defer resp.RawResponse.Body.Close()

	return resp.ListAppUsersForIdentityWithGrantResponse.Bindings, nil
}
