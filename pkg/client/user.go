package client

import (
	"context"

	"github.com/conductorone/cone/internal/c1api"
)

func (c *client) GetUser(ctx context.Context, userID string) (*c1api.C1ApiUserV1UserServiceGetResponse, error) {
	userResp, resp, err := c.apiClient.DefaultAPI.C1ApiUserV1UserServiceGet(ctx, userID).Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return userResp, nil
}
