package client

import (
	"context"
	"fmt"

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

	if resp == nil || resp.AppUserServiceListCredentialsResponse == nil {
		return nil, fmt.Errorf("ListAppUserCredentials: response is nil")
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	return resp.AppUserServiceListCredentialsResponse.List, nil
}

func (c *client) ListAppUsersForUser(ctx context.Context, appID string, userId string) ([]shared.AppUser, error) {
	resp, err := c.sdk.AppUser.ListAppUsersForUser(ctx, operations.C1APIAppV1AppUserServiceListAppUsersForUserRequest{AppID: appID, UserID: userId})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	appUsersForUser := make([]shared.AppUser, 0)

	for _, v := range resp.AppUsersForUserServiceListResponse.List {
		if v.AppUser == nil {
			continue
		}
		appUsersForUser = append(appUsersForUser, *v.AppUser)
	}

	return appUsersForUser, nil
}
