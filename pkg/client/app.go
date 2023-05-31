package client

import (
	"context"

	"github.com/conductorone/cone/internal/c1api"
)

func (c *client) GetApp(ctx context.Context, appID string) (*c1api.C1ApiAppV1App, error) {
	resp, httpResp, err := c.apiClient.DefaultAPI.C1ApiAppV1AppsGet(ctx, appID).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	return resp, nil
}
