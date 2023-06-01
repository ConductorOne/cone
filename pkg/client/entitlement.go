package client

import (
	"context"
	"errors"

	"github.com/conductorone/cone/internal/c1api"
)

type SearchEntitlementsFilter struct {
	Query            string
	EntitlementAlias string
}

func (c *client) SearchEntitlements(ctx context.Context, filter *SearchEntitlementsFilter) ([]*c1api.C1ApiAppV1AppEntitlement, error) {
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

	list, ok := resp.GetListOk()
	if !ok {
		return nil, errors.New("search-entitlements: list is nil")
	}

	rv := make([]*c1api.C1ApiAppV1AppEntitlement, 0, len(list))
	for _, v := range list {
		e, ok := v.GetAppEntitlementOk()
		if !ok {
			return nil, errors.New("search-entitlements: entitlement is nil")
		}
		rv = append(rv, e)
	}

	return rv, nil
}

func (c *client) ExpandEntitlements(ctx context.Context, in []*c1api.C1ApiAppV1AppEntitlement) (*Expander, error) {
	expander := &Expander{}
	for _, v := range in {
		expander.ExpandApp(v)
		expander.ExpandResourceType(v)
		expander.ExpandResource(v)
	}

	err := expander.Run(ctx, c)
	if err != nil {
		return nil, err
	}

	return expander, nil
}
