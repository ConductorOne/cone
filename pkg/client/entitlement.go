package client

import (
	"context"
	"errors"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

const (
	GrantedStatusGranted     = shared.RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusGranted
	GrantedStatusUnspecified = shared.RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusUnspecified
	GrantedStatusNotGranted  = shared.RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusNotGranted
	GrantedStatusAll         = shared.RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusAll
)

type SearchEntitlementsFilter struct {
	Query            string
	EntitlementAlias string
	AppDisplayName   string
	GrantedStatus    shared.RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus
	IncludeDeleted   bool
}

type AppEntitlement shared.AppEntitlement

func (a AppEntitlement) GetAppResourceId() string {
	return StringFromPtr(a.AppResourceID)
}

func (a AppEntitlement) GetAppResourceTypeId() string {
	return StringFromPtr(a.AppResourceTypeID)
}

func (a AppEntitlement) GetAppId() string {
	return StringFromPtr(a.AppID)
}

type EntitlementWithBindings struct {
	Entitlement AppEntitlement
	Bindings    []shared.AppEntitlementUserBinding
}

func (c *client) SearchEntitlements(ctx context.Context, filter *SearchEntitlementsFilter) ([]*EntitlementWithBindings, error) {
	// TODO(morgabra) Pagination
	// TODO(morgabra) Should we abstract the OpenAPI objects from the rest of cone? Kinda... no? But they aren't typed...
	req := shared.RequestCatalogSearchServiceSearchEntitlementsRequest{
		EntitlementAlias: stringPtr(filter.EntitlementAlias),
		GrantedStatus:    filter.GrantedStatus.ToPointer(),
		PageSize:         float64Ptr(100),
		PageToken:        nil,
		Query:            stringPtr(filter.Query),
		AppDisplayName:   stringPtr(filter.AppDisplayName),
		IncludeDeleted:   &filter.IncludeDeleted,
	}
	resp, err := c.sdk.RequestCatalogSearch.SearchEntitlements(ctx, &req)
	if err != nil {
		return nil, err
	}

	if err := handleBadStatus(resp.RawResponse); err != nil {
		return nil, err
	}

	list := resp.RequestCatalogSearchServiceSearchEntitlementsResponse.List
	if list == nil {
		return nil, errors.New("search-entitlements: list is nil")
	}

	rv := make([]*EntitlementWithBindings, 0, len(list))
	for _, v := range list {
		ent := v.AppEntitlementView
		if ent == nil {
			return nil, errors.New("search-entitlements: entitlement is nil")
		}

		rv = append(rv, &EntitlementWithBindings{
			Entitlement: AppEntitlement(*ent.AppEntitlement),
			Bindings:    v.AppEntitlementUserBindings,
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

func (c *client) GetEntitlement(ctx context.Context, appId string, entitlementId string) (*shared.AppEntitlement, error) {
	resp, err := c.sdk.AppEntitlements.Get(ctx, operations.C1APIAppV1AppEntitlementsGetRequest{
		AppID: appId,
		ID:    entitlementId,
	})
	if err != nil {
		return nil, err
	}

	if err := handleBadStatus(resp.RawResponse); err != nil {
		return nil, err
	}

	if resp.GetAppEntitlementResponse.AppEntitlementView == nil {
		return nil, errors.New("get-entitlement: view is nil")
	}

	if resp.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement == nil {
		return nil, errors.New("get-entitlement: entitlement is nil")
	}

	return resp.GetAppEntitlementResponse.AppEntitlementView.AppEntitlement, nil
}
