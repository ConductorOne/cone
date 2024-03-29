// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus - Search entitlements with this granted status for your signed in user.
type RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus string

const (
	RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusUnspecified RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus = "UNSPECIFIED"
	RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusAll         RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus = "ALL"
	RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusGranted     RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus = "GRANTED"
	RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusNotGranted  RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus = "NOT_GRANTED"
)

func (e RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus) ToPointer() *RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus {
	return &e
}

func (e *RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNSPECIFIED":
		fallthrough
	case "ALL":
		fallthrough
	case "GRANTED":
		fallthrough
	case "NOT_GRANTED":
		*e = RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus: %v", v)
	}
}

// The RequestCatalogSearchServiceSearchEntitlementsRequest searches entitlements, but only ones that are available to you through the open catalogs.
type RequestCatalogSearchServiceSearchEntitlementsRequest struct {
	// The app entitlement expand mask allows the user to get additional information when getting responses containing app entitlement views.
	AppEntitlementExpandMask *AppEntitlementExpandMask `json:"expandMask,omitempty"`
	// Search entitlements that belong to this app name (exact match).
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	// Search for entitlements with this alias (exact match).
	EntitlementAlias *string `json:"entitlementAlias,omitempty"`
	// Search entitlements with this granted status for your signed in user.
	GrantedStatus *RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus `json:"grantedStatus,omitempty"`
	// Include deleted entitlements
	IncludeDeleted *bool `json:"includeDeleted,omitempty"`
	// The pageSize where 0 <= pageSize <= 100. Values < 10 will be set to 10. A value of 0 returns the default page size (currently 25)
	PageSize *float64 `json:"pageSize,omitempty"`
	// The pageToken field.
	PageToken *string `json:"pageToken,omitempty"`
	// Fuzzy search the display name of resource types.
	Query *string `json:"query,omitempty"`
}

func (o *RequestCatalogSearchServiceSearchEntitlementsRequest) GetAppEntitlementExpandMask() *AppEntitlementExpandMask {
	if o == nil {
		return nil
	}
	return o.AppEntitlementExpandMask
}

func (o *RequestCatalogSearchServiceSearchEntitlementsRequest) GetAppDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.AppDisplayName
}

func (o *RequestCatalogSearchServiceSearchEntitlementsRequest) GetEntitlementAlias() *string {
	if o == nil {
		return nil
	}
	return o.EntitlementAlias
}

func (o *RequestCatalogSearchServiceSearchEntitlementsRequest) GetGrantedStatus() *RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatus {
	if o == nil {
		return nil
	}
	return o.GrantedStatus
}

func (o *RequestCatalogSearchServiceSearchEntitlementsRequest) GetIncludeDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDeleted
}

func (o *RequestCatalogSearchServiceSearchEntitlementsRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *RequestCatalogSearchServiceSearchEntitlementsRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *RequestCatalogSearchServiceSearchEntitlementsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}
