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
	defer resp.RawResponse.Body.Close()

	view := resp.UserServiceGetResponse.UserView
	if resp.UserServiceGetResponse.UserView == nil {
		return nil, errors.New("get-user: view is nil")
	}

	if view.User == nil {
		return nil, errors.New("get-user: user is nil")
	}

	return view.User, nil
}
