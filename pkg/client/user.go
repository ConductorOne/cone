package client

import (
	"context"
	"errors"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func (c *client) GetUser(ctx context.Context, userID string) (*shared.User, error) {
	resp, err := c.sdk.User.Get(ctx, operations.C1APIUserV1UserServiceGetRequest{ID: userID})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	view := resp.UserServiceGetResponse.UserView
	if resp.UserServiceGetResponse.UserView == nil {
		return nil, errors.New("get-user: view is nil")
	}

	if view.User == nil {
		return nil, errors.New("get-user: user is nil")
	}

	return view.User, nil
}

func (c *client) SearchUsers(ctx context.Context, req *shared.SearchUsersRequest) ([]*shared.User, error) {
	resp, err := c.sdk.UserSearch.Search(ctx, req)
	if err != nil {
		return nil, err
	}
	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	if resp.SearchUsersResponse == nil {
		return nil, nil
	}
	users := make([]*shared.User, 0, len(resp.SearchUsersResponse.List))
	for _, view := range resp.SearchUsersResponse.List {
		if view.User != nil {
			users = append(users, view.User)
		}
	}
	return users, nil
}
