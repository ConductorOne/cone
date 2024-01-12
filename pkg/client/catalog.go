package client

import (
	"context"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func (c *client) ListCatalogs(ctx context.Context) ([]shared.RequestCatalog, error) {
	catalogs := make([]shared.RequestCatalog, 0)
	for {
		resp, err := c.sdk.RequestCatalogManagement.List(ctx)
		if err != nil {
			return nil, err
		}
		if err := NewHTTPError(resp.RawResponse); err != nil {
			return nil, err
		}

		for _, catalog := range resp.RequestCatalogManagementServiceListResponse.List {
			catalogs = append(catalogs, *catalog.RequestCatalog)
		}

		if resp.RequestCatalogManagementServiceListResponse.NextPageToken == nil || *resp.RequestCatalogManagementServiceListResponse.NextPageToken == "" {
			break
		}
	}

	return catalogs, nil
}
