package client

import (
	"context"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func (c *client) GetApp(ctx context.Context, appID string) (*shared.App, error) {
	resp, err := c.sdk.Apps.Get(ctx, operations.C1APIAppV1AppsGetRequest{
		ID: appID,
	})
	if err != nil {
		return nil, err
	}

	if err := handleBadStatus(resp.RawResponse); err != nil {
		return nil, err
	}
	return resp.GetAppResponse.App, nil
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
