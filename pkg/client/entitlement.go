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
	Query                    string
	EntitlementAlias         string
	AppDisplayName           string
	GrantedStatus            shared.RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus
	IncludeDeleted           bool
	AppEntitlementExpandMask shared.AppEntitlementExpandMask
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
	expanded    map[string]*any
}

func (e EntitlementWithBindings) GetExpanded() map[string]*any {
	if e.expanded == nil {
		return make(map[string]*any)
	}
	return e.expanded
}

type ExpandableEntitlementWithBindings struct {
	shared.AppEntitlementWithUserBindings
	ExpandedMap map[string]int
}

func NewExpandableEntitlementWithBindings(v shared.AppEntitlementWithUserBindings) *ExpandableEntitlementWithBindings {
	if v.AppEntitlementView == nil {
		return nil
	}
	return &ExpandableEntitlementWithBindings{
		AppEntitlementWithUserBindings: v,
	}
}

func (e *ExpandableEntitlementWithBindings) GetPaths() []PathDetails {
	view := *e.AppEntitlementWithUserBindings.AppEntitlementView
	return []PathDetails{
		{
			Name: "App",
			Path: view.GetAppPath(),
		},
		{
			Name: "AppResource",
			Path: view.GetAppResourcePath(),
		},
		{
			Name: "AppResourceType",
			Path: view.GetAppResourceTypePath(),
		},
	}
}

func (e *ExpandableEntitlementWithBindings) SetPath(pathname string, value int) {
	if e.ExpandedMap == nil {
		e.ExpandedMap = make(map[string]int)
	}
	e.ExpandedMap[pathname] = value
}

func (c *client) SearchEntitlements(ctx context.Context, filter *SearchEntitlementsFilter) ([]*EntitlementWithBindings, error) {
	// TODO(morgabra) Pagination
	// TODO(morgabra) Should we abstract the OpenAPI objects from the rest of cone? Kinda... no? But they aren't typed...
	req := shared.RequestCatalogSearchServiceSearchEntitlementsRequest{
		EntitlementAlias:         stringPtr(filter.EntitlementAlias),
		GrantedStatus:            filter.GrantedStatus.ToPointer(),
		PageSize:                 float64Ptr(100),
		PageToken:                nil,
		Query:                    stringPtr(filter.Query),
		AppDisplayName:           stringPtr(filter.AppDisplayName),
		IncludeDeleted:           &filter.IncludeDeleted,
		AppEntitlementExpandMask: &filter.AppEntitlementExpandMask,
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

	// Unmarshal the expanded fields
	expanded := make([]any, 0, len(resp.RequestCatalogSearchServiceSearchEntitlementsResponse.Expanded))
	for _, x := range resp.RequestCatalogSearchServiceSearchEntitlementsResponse.Expanded {
		x := x
		converted, err := UnmarshalAnyType[shared.RequestCatalogSearchServiceSearchEntitlementsResponseExpanded](&x)
		if err != nil {
			return nil, err
		}
		expanded = append(expanded, converted)
	}

	// Convert the list of entitlements to a list of expandable entitlements
	expandableList := make([]*ExpandableEntitlementWithBindings, 0, len(list))
	for _, v := range list {
		ent := NewExpandableEntitlementWithBindings(v)
		if ent == nil {
			return nil, errors.New("search-entitlements: entitlement is nil")
		}

		expandableList = append(expandableList, ent)
	}

	// Populate the expandable objects with the indexes of related objects
	err = ExpandableReponse[*ExpandableEntitlementWithBindings]{
		List: expandableList,
	}.PopulateExpandedIndexes()

	if err != nil {
		return nil, err
	}

	// Iterate over the expandable objects and convert them to the final response
	rv := make([]*EntitlementWithBindings, 0, len(list))
	for _, v := range expandableList {
		rv = append(rv, &EntitlementWithBindings{
			Entitlement: AppEntitlement(*v.AppEntitlementWithUserBindings.AppEntitlementView.AppEntitlement),
			Bindings:    v.AppEntitlementWithUserBindings.AppEntitlementUserBindings,
			expanded:    PopulateExpandedMap(v.ExpandedMap, expanded),
		})
	}
	return rv, nil
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
