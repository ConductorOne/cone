package client

import (
	"context"
	"fmt"

	"github.com/conductorone/cone/internal/c1api"
	"github.com/conductorone/cone/pkg/output"
)

func (c *client) GetResourceType(ctx context.Context, appID string, resourceTypeID string) (*C1ApiAppV1AppResourceType, error) {
	resp, httpResp, err := c.apiClient.DefaultAPI.C1ApiAppV1AppResourceTypeServiceGet(ctx, appID, resourceTypeID).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	resourceType := C1ApiAppV1AppResourceType(*resp.AppResourceTypeView.AppResourceType)
	return &resourceType, nil
}

type C1ApiAppV1AppResourceType c1api.C1ApiAppV1AppResourceType

func (a *C1ApiAppV1AppResourceType) CacheKey() string {
	return ResourceTypeCacheKey(a.AppId, a.Id)
}

func ResourceTypeCacheKey(appID *string, resourceTypeID *string) string {
	return fmt.Sprintf("%s:%s", output.FromPtr(appID), output.FromPtr(resourceTypeID))
}

func (c *client) GetResource(ctx context.Context, appID string, resourceTypeID string, resourceID string) (*C1ApiAppV1AppResource, error) {
	resp, httpResp, err := c.apiClient.DefaultAPI.C1ApiAppV1AppResourceServiceGet(ctx, appID, resourceTypeID, resourceID).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	resource := C1ApiAppV1AppResource(*resp.AppResourceView.AppResource)
	return &resource, nil
}

func ResourceCacheKey(appID *string, resourceTypeID *string, resourceID *string) string {
	return fmt.Sprintf("%s:%s:%s", output.FromPtr(appID), output.FromPtr(resourceTypeID), output.FromPtr(resourceID))
}

type C1ApiAppV1AppResource c1api.C1ApiAppV1AppResource

func (a *C1ApiAppV1AppResource) CacheKey() string {
	return ResourceCacheKey(a.AppId,
		a.AppResourceTypeId,
		a.Id)
}
