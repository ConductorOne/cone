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
	defer resp.RawResponse.Body.Close()

	if err := handleBadStatus(resp.RawResponse); err != nil {
		return nil, err
	}
	return resp.GetAppResponse.App, nil
}
