// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AppEntitlementReference - The AppEntitlementReference message.
type AppEntitlementReference struct {
	// The appEntitlementId field.
	AppEntitlementID *string `json:"appEntitlementId,omitempty"`
	// The appId field.
	AppID *string `json:"appId,omitempty"`
}

func (o *AppEntitlementReference) GetAppEntitlementID() *string {
	if o == nil {
		return nil
	}
	return o.AppEntitlementID
}

func (o *AppEntitlementReference) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}
