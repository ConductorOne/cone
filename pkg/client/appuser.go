package client

import (
	"context"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func (c *client) ListAppUsers(ctx context.Context, appID string) ([]shared.AppUser, error) {
	resp, err := c.sdk.AppUser.List(ctx, operations.C1APIAppV1AppUserServiceListRequest{AppID: appID})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	appUsers := make([]shared.AppUser, 0)

	for _, v := range resp.AppUserServiceListResponse.List {
		if v.AppUser == nil {
			continue
		}
		appUsers = append(appUsers, *v.AppUser)
	}

	return appUsers, nil
}

func (c *client) ListAppUserCredentials(ctx context.Context, appID string, appUserID string) ([]shared.AppUserCredential, error) {
	resp, err := c.sdk.AppUser.ListAppUserCredentials(ctx, operations.C1APIAppV1AppUserServiceListAppUserCredentialsRequest{AppID: appID, AppUserID: appUserID})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	// appUsers := make([]shared.AppUser, 0)

	return resp.AppUserServiceListCredentialsResponse.List, nil

	// return appUsers, nil
}
