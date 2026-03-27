package client

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

var resourceTypeMap = map[string]shared.ResourceType{
	"ROLE":         shared.ResourceTypeRole,
	"GROUP":        shared.ResourceTypeGroup,
	"LICENSE":      shared.ResourceTypeLicense,
	"PROJECT":      shared.ResourceTypeProject,
	"CATALOG":      shared.ResourceTypeCatalog,
	"CUSTOM":       shared.ResourceTypeCustom,
	"VAULT":        shared.ResourceTypeVault,
	"PROFILE_TYPE": shared.ResourceTypeProfileType,
}

// ResolveResourceType maps a user-provided type string to (ResourceType enum, display name).
// Known enum names map directly. Anything else uses CUSTOM with the user's string as the display name.
func ResolveResourceType(typeStr string) (shared.ResourceType, string) {
	normalized := strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(typeStr, "-", "_"), " ", "_"))
	if rt, ok := resourceTypeMap[normalized]; ok {
		return rt, strings.ToUpper(normalized[:1]) + strings.ToLower(normalized[1:])
	}
	return shared.ResourceTypeCustom, typeStr
}

func (c *client) CreateManuallyManagedResourceType(
	ctx context.Context,
	appID string,
	resourceType shared.ResourceType,
	displayName string,
) (*shared.AppResourceType, error) {
	resp, err := c.sdk.AppResourceType.CreateManuallyManagedResourceType(
		ctx,
		operations.C1APIAppV1AppResourceTypeServiceCreateManuallyManagedResourceTypeRequest{
			AppID: appID,
			CreateManuallyManagedResourceTypeRequest: &shared.CreateManuallyManagedResourceTypeRequest{
				DisplayName:  displayName,
				ResourceType: resourceType,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	rt := resp.CreateManuallyManagedResourceTypeResponse.GetAppResourceType()
	if rt == nil {
		return nil, errors.New("create-resource-type: response is nil")
	}
	return rt, nil
}

func (c *client) CreateManuallyManagedResource(
	ctx context.Context,
	appID string,
	resourceTypeID string,
	displayName string,
	description string,
) (*shared.AppResource, error) {
	req := shared.CreateManuallyManagedAppResourceRequest{
		DisplayName: displayName,
	}
	if description != "" {
		req.Description = &description
	}

	resp, err := c.sdk.AppResource.CreateManuallyManagedAppResource(
		ctx,
		operations.C1APIAppV1AppResourceServiceCreateManuallyManagedAppResourceRequest{
			AppID:                                  appID,
			AppResourceTypeID:                      resourceTypeID,
			CreateManuallyManagedAppResourceRequest: &req,
		},
	)
	if err != nil {
		return nil, err
	}
	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	resource := resp.CreateManuallyManagedAppResourceResponse.GetAppResource()
	if resource == nil {
		return nil, errors.New("create-resource: response is nil")
	}
	return resource, nil
}

func (c *client) CreateAppEntitlement(
	ctx context.Context,
	appID string,
	resourceTypeID string,
	resourceID string,
	displayName string,
	slug string,
) (*shared.AppEntitlement, error) {
	resp, err := c.sdk.AppEntitlements.Create(
		ctx,
		operations.C1APIAppV1AppEntitlementsCreateRequest{
			AppID: appID,
			CreateAppEntitlementRequest: &shared.CreateAppEntitlementRequest{
				DisplayName:       displayName,
				AppResourceTypeID: &resourceTypeID,
				AppResourceID:     &resourceID,
				Slug:              &slug,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}

	view := resp.CreateAppEntitlementResponse.GetAppEntitlementView()
	if view == nil || view.AppEntitlement == nil {
		return nil, errors.New("create-entitlement: response is nil")
	}
	return view.AppEntitlement, nil
}

func (c *client) SearchApps(ctx context.Context, query string) ([]shared.App, error) {
	pageSize := 50
	resp, err := c.sdk.AppSearch.Search(ctx, &shared.SearchAppsRequest{
		DisplayName: stringPtr(query),
		PageSize:    &pageSize,
	})
	if err != nil {
		return nil, err
	}
	if err := NewHTTPError(resp.RawResponse); err != nil {
		return nil, err
	}
	return resp.SearchAppsResponse.GetList(), nil
}

// ResolveAppByNameOrID looks up an app by display name or ID.
// If the input looks like a KSUID (27 alphanumeric chars), it fetches directly by ID.
// Otherwise, it searches by display name and returns an exact or unique match.
func (c *client) ResolveAppByNameOrID(ctx context.Context, appIDOrName string) (*shared.App, error) {
	if len(appIDOrName) == 27 && isAlphanumeric(appIDOrName) {
		return c.GetApp(ctx, appIDOrName)
	}

	apps, err := c.SearchApps(ctx, appIDOrName)
	if err != nil {
		return nil, err
	}
	if len(apps) == 0 {
		return nil, fmt.Errorf("no app found matching %q", appIDOrName)
	}

	// Exact match first.
	for i := range apps {
		if strings.EqualFold(StringFromPtr(apps[i].DisplayName), appIDOrName) {
			return &apps[i], nil
		}
	}

	if len(apps) == 1 {
		return &apps[0], nil
	}

	var names []string
	for _, app := range apps {
		names = append(names, fmt.Sprintf("  %s  %s", StringFromPtr(app.ID), StringFromPtr(app.DisplayName)))
	}
	return nil, fmt.Errorf("multiple apps match %q, please use the app ID directly:\n%s", appIDOrName, strings.Join(names, "\n"))
}

func isAlphanumeric(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}
