// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// CompletedAction - The CompletedAction message.
type CompletedAction struct {
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	// The entitlements field.
	Entitlements []AppEntitlementReference `json:"entitlements,omitempty"`
	// The userId field.
	UserID *string `json:"userId,omitempty"`
}
