package client

import (
	"context"
	"errors"

	"github.com/conductorone/cone/internal/c1api"
)

func (c *client) GetResourceType(ctx context.Context, appID string, resourceTypeID string) (*c1api.C1ApiAppV1AppResourceType, error) {
	resp, httpResp, err := c.apiClient.DefaultAPI.C1ApiAppV1AppResourceTypeServiceGet(ctx, appID, resourceTypeID).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	v, ok := resp.GetAppResourceTypeViewOk()
	if !ok {
		return nil, errors.New("get-resource-type: view is nil")
	}

	rt, ok := v.GetAppResourceTypeOk()
	if !ok {
		return nil, errors.New("get-resource-type: resource type is nil")
	}

	return rt, nil
}

func (c *client) GetResource(ctx context.Context, appID string, resourceTypeID string, resourceID string) (*c1api.C1ApiAppV1AppResource, error) {
	resp, httpResp, err := c.apiClient.DefaultAPI.C1ApiAppV1AppResourceServiceGet(ctx, appID, resourceTypeID, resourceID).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	v, ok := resp.GetAppResourceViewOk()
	if !ok {
		return nil, errors.New("get-resource: view is nil")
	}

	r, ok := v.GetAppResourceOk()
	if !ok {
		return nil, errors.New("get-resource: resource type is nil")
	}

	return r, nil
}
