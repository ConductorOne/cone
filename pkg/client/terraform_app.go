package client

import (
	"context"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/cone/pkg/resource"
)

type AppTemplate struct {
	App shared.App
}

func (a AppTemplate) GetIds() map[string]string {
	ids := make(map[string]string)
	if a.App.ID != nil {
		ids["id"] = *a.App.ID
	}
	return ids
}

func (a AppTemplate) GetResourceType() string {
	return "app"
}

func (a AppTemplate) GetType() string {
	return "conductorone_app" // Assuming the type is "App"
}

func (a AppTemplate) GetPk() string {
	return resource.GeneratePK(a)
}

func (c *client) ListApps(ctx context.Context) ([]shared.App, error) {
	apps := make([]shared.App, 0)
	pageSize := float64(100)
	resp, err := c.sdk.Apps.List(ctx, operations.C1APIAppV1AppsListRequest{
		PageSize: &pageSize,
	})
	if err != nil {
		return nil, err
	}
	apps = append(apps, resp.ListAppsResponse.List...)

	for resp.ListAppsResponse.NextPageToken != nil && *resp.ListAppsResponse.NextPageToken != "" {
		resp, err := c.sdk.Apps.List(ctx, operations.C1APIAppV1AppsListRequest{
			PageToken: resp.ListAppsResponse.NextPageToken,
			PageSize:  &pageSize,
		})
		if err != nil {
			return nil, err
		}
		apps = append(apps, resp.ListAppsResponse.List...)
	}

	if err := handleBadStatus(resp.RawResponse); err != nil {
		return nil, err
	}
	return apps, nil
}
