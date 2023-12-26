package client

import (
	"context"
	"errors"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func (c *client) GetResourceType(ctx context.Context, appID string, resourceTypeID string) (*shared.AppResourceType, error) {
	resp, err := c.sdk.AppResourceType.Get(ctx, operations.C1APIAppV1AppResourceTypeServiceGetRequest{
		AppID: appID,
		ID:    resourceTypeID,
	})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	v := resp.AppResourceTypeServiceGetResponse.AppResourceTypeView
	if v == nil {
		return nil, errors.New("get-resource-type: view is nil")
	}

	r := v.AppResourceType
	if r == nil {
		return nil, errors.New("get-resource-type: resource type is nil")
	}

	return r, nil
}

func (c *client) GetResource(ctx context.Context, appID string, resourceTypeID string, resourceID string) (*shared.AppResource, error) {
	resp, err := c.sdk.AppResource.Get(ctx, operations.C1APIAppV1AppResourceServiceGetRequest{
		AppID:             appID,
		AppResourceTypeID: resourceTypeID,
		ID:                resourceID,
	})
	if err != nil {
		return nil, err
	}

	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	v := resp.AppResourceServiceGetResponse.AppResourceView
	if v == nil {
		return nil, errors.New("get-resource: view is nil")
	}

	r := v.AppResource
	if r == nil {
		return nil, errors.New("get-resource: resource type is nil")
	}

	return r, nil
}
