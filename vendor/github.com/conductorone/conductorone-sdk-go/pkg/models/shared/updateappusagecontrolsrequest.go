// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateAppUsageControlsRequest - The UpdateAppUsageControlsRequest message.
type UpdateAppUsageControlsRequest struct {
	// The AppUsageControls message.
	AppUsageControls *AppUsageControls `json:"appUsageControls,omitempty"`
	UpdateMask       *string           `json:"updateMask,omitempty"`
}
