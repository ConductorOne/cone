package client

import (
	"context"

	"github.com/conductorone/cone/internal/c1api"
)

type SearchEntitlementsFilter struct {
	Query            string
	EntitlementAlias string
}

func (c *client) SearchEntitlements(ctx context.Context, filter *SearchEntitlementsFilter) (*c1api.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse, error) {
	// TODO(morgabra) Pagination
	// TODO(morgabra) Should we abstract the OpenAPI objects from the rest of cone? Kinda... no? But they aren't typed...
	req := &c1api.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest{
		EntitlementAlias: stringPtr(filter.EntitlementAlias),
		PageSize:         float32Ptr(100),
		PageToken:        nil,
		Query:            stringPtr(filter.Query),
	}
	api := c.apiClient.DefaultAPI.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements(ctx)
	resp, httpResp, err := api.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest(*req).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	return resp, nil
}
