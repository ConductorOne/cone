package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

const (
	GrantedStatusGranted     = shared.GrantedStatusGranted
	GrantedStatusUnspecified = shared.GrantedStatusUnspecified
	GrantedStatusNotGranted  = shared.GrantedStatusNotGranted
	GrantedStatusAll         = shared.GrantedStatusAll
)

type SearchEntitlementsFilter struct {
	Query                    string
	EntitlementAlias         string
	AppDisplayName           string
	GrantedStatus            shared.GrantedStatus
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

func (e *EntitlementWithBindings) GetExpanded() map[string]*any {
	if e == nil {
		return nil
	}
	return e.expanded
}

type ExpandableEntitlementWithBindings struct {
	shared.AppEntitlementWithUserBindings
	ExpandedMap map[string]int
}

func NewExpandableEntitlementWithBindings(v shared.AppEntitlementWithUserBindings) *ExpandableEntitlementWithBindings {
	if v.Entitlement == nil {
		return nil
	}
	return &ExpandableEntitlementWithBindings{
		AppEntitlementWithUserBindings: v,
	}
}

func (e *ExpandableEntitlementWithBindings) GetPaths() []PathDetails {
	if e == nil {
		return nil
	}
	view := *e.Entitlement
	return []PathDetails{
		{
			Name: ExpandedApp,
			Path: view.GetAppPath(),
		},
		{
			Name: ExpandedAppResource,
			Path: view.GetAppResourcePath(),
		},
		{
			Name: ExpandedAppResourceType,
			Path: view.GetAppResourceTypePath(),
		},
	}
}

func (e *ExpandableEntitlementWithBindings) SetPath(pathname string, value int) {
	if e == nil {
		return
	}
	if e.ExpandedMap == nil {
		e.ExpandedMap = make(map[string]int)
	}
	e.ExpandedMap[pathname] = value
}

// maxRepeatedSearchEntitlementsPageToken bounds how many times the server may
// hand back a page token identical to the one just sent before we treat paging
// as stuck. Without it a server that never advances the token spins forever.
const maxRepeatedSearchEntitlementsPageToken = 8

func (c *client) SearchEntitlements(ctx context.Context, filter *SearchEntitlementsFilter) ([]*EntitlementWithBindings, error) {
	// TODO(morgabra) Should we abstract the OpenAPI objects from the rest of cone? Kinda... no? But they aren't typed...
	rv := make([]*EntitlementWithBindings, 0)
	pageToken := ""
	repeatedToken := 0

	for {
		req := shared.RequestCatalogSearchServiceSearchEntitlementsRequest{
			EntitlementAlias: stringPtr(filter.EntitlementAlias),
			GrantedStatus:    filter.GrantedStatus.ToPointer(),
			PageSize:         intPtr(100),
			PageToken:        stringPtr(pageToken),
			Query:            stringPtr(filter.Query),
			AppDisplayName:   stringPtr(filter.AppDisplayName),
			IncludeDeleted:   &filter.IncludeDeleted,
			ExpandMask:       &filter.AppEntitlementExpandMask,
		}
		resp, err := c.sdk.RequestCatalogSearch.SearchEntitlements(ctx, &req)
		if err != nil {
			return nil, err
		}

		if err := NewHTTPError(resp.RawResponse); err != nil {
			return nil, err
		}

		page, err := convertSearchEntitlementsPage(resp.RequestCatalogSearchServiceSearchEntitlementsResponse)
		if err != nil {
			return nil, err
		}
		rv = append(rv, page...)

		// Stop on an empty token, never on an empty page: the server filters
		// granted status after it cuts the page, so a page can come back short or
		// empty while later pages still hold results.
		nextPageToken := StringFromPtr(resp.RequestCatalogSearchServiceSearchEntitlementsResponse.NextPageToken)
		if nextPageToken == "" {
			break
		}
		if nextPageToken == pageToken {
			repeatedToken++
			if repeatedToken >= maxRepeatedSearchEntitlementsPageToken {
				return nil, fmt.Errorf("search-entitlements: page token repeated %d times, pagination is not advancing", repeatedToken)
			}
		} else {
			repeatedToken = 0
		}
		pageToken = nextPageToken
	}

	return rv, nil
}

// convertSearchEntitlementsPage expands and converts one response page.
//
// This runs per page rather than once over a concatenated list because the
// response's expanded array is indexed per response: ExpandedMap holds offsets
// into THIS page's expanded objects, so merging raw pages first would resolve
// those offsets against the wrong array.
func convertSearchEntitlementsPage(
	resp *shared.RequestCatalogSearchServiceSearchEntitlementsResponse,
) ([]*EntitlementWithBindings, error) {
	if resp == nil {
		return nil, errors.New("search-entitlements: response is nil")
	}

	// Unmarshal the expanded fields
	expanded := make([]any, 0, len(resp.Expanded))
	for _, x := range resp.Expanded {
		x := x
		converted, err := UnmarshalAnyType[shared.RequestCatalogSearchServiceSearchEntitlementsResponseExpanded](&x)
		if err != nil {
			return nil, err
		}
		expanded = append(expanded, converted)
	}

	// Convert the list of entitlements to a list of expandable entitlements. A
	// nil list is an empty page, not a fault -- the wire format omits an empty
	// repeated field entirely.
	expandableList := make([]*ExpandableEntitlementWithBindings, 0, len(resp.List))
	for _, v := range resp.List {
		ent := NewExpandableEntitlementWithBindings(v)
		if ent == nil {
			return nil, errors.New("search-entitlements: entitlement is nil")
		}

		expandableList = append(expandableList, ent)
	}

	// Populate the expandable objects with the indexes of related objects
	err := ExpandableReponse[*ExpandableEntitlementWithBindings]{
		List: expandableList,
	}.PopulateExpandedIndexes()

	if err != nil {
		return nil, err
	}

	// Iterate over the expandable objects and convert them to the final response
	rv := make([]*EntitlementWithBindings, 0, len(expandableList))
	for _, v := range expandableList {
		rv = append(rv, &EntitlementWithBindings{
			Entitlement: AppEntitlement(*v.Entitlement.AppEntitlement),
			Bindings:    v.AppEntitlementUserBindings,
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
	if err := NewHTTPError(resp.RawResponse); err != nil {
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

func (c *client) ListEntitlements(ctx context.Context, appId string) ([]shared.AppEntitlement, error) {
	entitlements := make([]shared.AppEntitlement, 0)
	pageSize := 100
	pageToken := ""
	for {
		resp, err := c.sdk.AppEntitlements.List(ctx, operations.C1APIAppV1AppEntitlementsListRequest{
			PageSize:  &pageSize,
			AppID:     appId,
			PageToken: &pageToken,
		})
		if err != nil {
			return nil, err
		}
		if err := NewHTTPError(resp.RawResponse); err != nil {
			return nil, err
		}

		for _, v := range resp.ListAppEntitlementsResponse.List {
			if v.AppEntitlement == nil {
				continue
			}
			entitlements = append(entitlements, *v.AppEntitlement)
		}

		if resp.ListAppEntitlementsResponse.NextPageToken == nil || *resp.ListAppEntitlementsResponse.NextPageToken == "" {
			break
		}
		pageToken = *resp.ListAppEntitlementsResponse.NextPageToken
	}

	return entitlements, nil
}

// HasRequestForm checks if an entitlement has a request form (custom fields) bound to it.
func (c *client) HasRequestForm(ctx context.Context, appID string, entitlementID string) (bool, error) {
	resp, err := c.sdk.RequestSchema.FindBindingForAppEntitlement(
		ctx,
		&shared.RequestSchemaServiceFindBindingForAppEntitlementRequest{
			EntitlementRef: &shared.AppEntitlementRef{
				AppID: &appID,
				ID:    &entitlementID,
			},
		},
	)
	if err != nil {
		return false, err
	}
	// 404 means no binding — no form.
	if resp.RawResponse.StatusCode == http.StatusNotFound {
		return false, nil
	}
	if httpErr := NewHTTPError(resp.RawResponse); httpErr != nil {
		return false, httpErr
	}

	result := resp.RequestSchemaServiceFindBindingForAppEntitlementResponse
	return result != nil && result.RequestSchemaID != nil && *result.RequestSchemaID != "", nil
}

func (c *client) UpdateEntitlement(ctx context.Context, appID, entitlementID string, req *shared.UpdateAppEntitlementRequest) error {
	resp, err := c.sdk.AppEntitlements.Update(ctx, operations.C1APIAppV1AppEntitlementsUpdateRequest{
		AppID:                       appID,
		ID:                          entitlementID,
		UpdateAppEntitlementRequest: req,
	})
	if err != nil {
		return err
	}
	if err := NewHTTPError(resp.RawResponse); err != nil {
		return err
	}
	return nil
}
