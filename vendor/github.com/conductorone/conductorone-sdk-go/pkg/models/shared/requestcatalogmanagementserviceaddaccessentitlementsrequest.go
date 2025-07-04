// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The RequestCatalogManagementServiceAddAccessEntitlementsRequest message is used to add access entitlements to a request
//
//	catalog to determine which users can view the request catalog.
type RequestCatalogManagementServiceAddAccessEntitlementsRequest struct {
	// List of entitlements to add to the request catalog as access entitlements.
	AccessEntitlements []AppEntitlementRef `json:"accessEntitlements"`
}

func (o *RequestCatalogManagementServiceAddAccessEntitlementsRequest) GetAccessEntitlements() []AppEntitlementRef {
	if o == nil {
		return nil
	}
	return o.AccessEntitlements
}
