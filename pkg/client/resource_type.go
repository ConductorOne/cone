package client

import (
	"context"

	"github.com/conductorone/cone/internal/c1api"
)

func (c *client) GetResourceType(ctx context.Context, appID string, resourceTypeID string) (*c1api.C1ApiAppV1AppResourceTypeServiceGetResponse, error) {
	resp, httpResp, err := c.apiClient.DefaultAPI.C1ApiAppV1AppResourceTypeServiceGet(ctx, appID, resourceTypeID).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	return resp, nil
}

func (c *client) GetResource(ctx context.Context, appID string, resourceTypeID string, resourceID string) (*c1api.C1ApiAppV1AppResourceServiceGetResponse, error) {
	resp, httpResp, err := c.apiClient.DefaultAPI.C1ApiAppV1AppResourceServiceGet(ctx, appID, resourceTypeID, resourceID).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	return resp, nil
}
