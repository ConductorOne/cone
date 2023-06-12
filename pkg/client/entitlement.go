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

type EntitlementWithBindings struct {
	Entitlement c1api.C1ApiAppV1AppEntitlement
	Bindings    []c1api.C1ApiAppV1AppEntitlementUserBinding
}

func (c *client) SearchEntitlements(ctx context.Context, filter *SearchEntitlementsFilter) ([]*EntitlementWithBindings, error) {
	// TODO(morgabra) Pagination
	// TODO(morgabra) Should we abstract the OpenAPI objects from the rest of cone? Kinda... no? But they aren't typed...
	req := &c1api.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest{
		EntitlementAlias: stringPtr(filter.EntitlementAlias),
		PageSize:         float32Ptr(100),
		PageToken:        nil,
		Query:            stringPtr(filter.Query),
	}
	api := c.apiClient.RequestCatalogSearchAPI.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements(ctx)
	resp, httpResp, err := api.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest(*req).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	list, ok := resp.GetListOk()
	if !ok {
		return nil, errors.New("search-entitlements: list is nil")
	}

	rv := make([]*EntitlementWithBindings, 0, len(list))
	for _, v := range list {
		ev, ok := v.GetEntitlementOk()
		if !ok {
			return nil, errors.New("search-entitlements: entitlement is nil")
		}

		rv = append(rv, &EntitlementWithBindings{
			Entitlement: ev.GetAppEntitlement(),
			Bindings:    v.GetAppEntitlementUserBindings(),
		})
	}

	return rv, nil
}

func (c *client) ExpandEntitlements(ctx context.Context, in []*EntitlementWithBindings) (*Expander, error) {
	expander := &Expander{}
	for _, v := range in {
		expander.ExpandApp(&v.Entitlement)
		expander.ExpandResourceType(&v.Entitlement)
		expander.ExpandResource(&v.Entitlement)
	}

	err := expander.Run(ctx, c)
	if err != nil {
		return nil, err
	}

	return expander, nil
}

func (c *client) GetEntitlement(ctx context.Context, appId string, entitlementId string) (*c1api.C1ApiAppV1AppEntitlement, error) {
	resp, httpResp, err := c.apiClient.AppEntitlementsAPI.C1ApiAppV1AppEntitlementsGet(ctx, appId, entitlementId).Execute()
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	v, ok := resp.GetAppEntitlementViewOk()
	if !ok {
		return nil, errors.New("get-entitlement: view is nil")
	}

	r, ok := v.GetAppEntitlementOk()
	if !ok {
		return nil, errors.New("get-entitlement: entitlement is nil")
	}

	return r, nil
}
