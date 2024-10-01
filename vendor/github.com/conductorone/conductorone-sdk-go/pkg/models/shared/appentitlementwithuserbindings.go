// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The AppEntitlementWithUserBindings message represents an app entitlement and its associated user bindings.
type AppEntitlementWithUserBindings struct {
	// The app entitlement view contains the serialized app entitlement and paths to objects referenced by the app entitlement.
	AppEntitlementView *AppEntitlementView `json:"entitlement,omitempty"`
	// An array of AppEntitlementUserBinding objects which represent the relationships that give app users access to the specific app entitlement.
	AppEntitlementUserBindings []AppEntitlementUserBinding `json:"appEntitlementUserBindings,omitempty"`
}

func (o *AppEntitlementWithUserBindings) GetAppEntitlementView() *AppEntitlementView {
	if o == nil {
		return nil
	}
	return o.AppEntitlementView
}

func (o *AppEntitlementWithUserBindings) GetAppEntitlementUserBindings() []AppEntitlementUserBinding {
	if o == nil {
		return nil
	}
	return o.AppEntitlementUserBindings
}
