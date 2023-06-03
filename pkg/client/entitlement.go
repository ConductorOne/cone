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
	Entitlement *c1api.C1ApiAppV1AppEntitlement
	Bindings    []*c1api.C1ApiAppV1AppEntitlementUserBinding
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

	rv := make([]*EntitlementWithBindings, 0, len(list))
	for _, v := range list {
		ev, ok := v.GetEntitlementOk()
		if !ok {
			return nil, errors.New("search-entitlements: entitlement is nil")
		}

		e, ok := ev.GetAppEntitlementOk()
		if !ok {
			return nil, errors.New("search-entitlements: app-entitlement is nil")
		}

		// TODO(morgabra) Should we be fighting this?
		bv, ok := v.GetAppEntitlementUserBindingsOk()
		var bindings []*c1api.C1ApiAppV1AppEntitlementUserBinding
		if ok {
			for _, b := range bv {
				bindings = append(bindings, &b)
			}
		}

		rv = append(rv, &EntitlementWithBindings{
			Entitlement: e,
			Bindings:    bindings,
		})
	}

	return rv, nil
}

func (c *client) ExpandEntitlements(ctx context.Context, in []*EntitlementWithBindings) (*Expander, error) {
	expander := &Expander{}
	for _, v := range in {
		expander.ExpandApp(v.Entitlement)
		expander.ExpandResourceType(v.Entitlement)
		expander.ExpandResource(v.Entitlement)
	}

	err := expander.Run(ctx, c)
	if err != nil {
		return nil, err
	}

	return expander, nil
}
