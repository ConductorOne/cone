package client

import (
	"context"

	"github.com/conductorone/cone/internal/c1api"
)

func (c *client) GetTask(ctx context.Context, userID string) (*c1api.C1, error) {
	userResp, resp, err := c.apiClient.DefaultAPI.C1ApiUserV1UserServiceGet(ctx, userID).Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return userResp, nil
}
