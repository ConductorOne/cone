package client

import (
	"context"
	"errors"

	"github.com/conductorone/cone/internal/c1api"
)

func (c *client) GetUser(ctx context.Context, userID string) (*c1api.C1ApiUserV1User, error) {
	userResp, resp, err := c.apiClient.UserAPI.C1ApiUserV1UserServiceGet(ctx, userID).Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	v, ok := userResp.GetUserViewOk()
	if !ok {
		return nil, errors.New("get-user: view is nil")
	}

	u, ok := v.GetUserOk()
	if !ok {
		return nil, errors.New("get-user: user is nil")
	}

	return u, nil
}
