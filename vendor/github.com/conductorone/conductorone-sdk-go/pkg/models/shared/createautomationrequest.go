// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The CreateAutomationRequest message.
type CreateAutomationRequest struct {
	// The AppEntitlementAutomation message.
	//
	// This message contains a oneof named conditions. Only a single field of the following list may be set at a time:
	//   - none
	//   - entitlements
	//   - cel
	//   - basic
	//
	AppEntitlementAutomation *AppEntitlementAutomationInput `json:"automation,omitempty"`
}

func (o *CreateAutomationRequest) GetAppEntitlementAutomation() *AppEntitlementAutomationInput {
	if o == nil {
		return nil
	}
	return o.AppEntitlementAutomation
}
