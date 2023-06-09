// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// CancelledAction - The CancelledAction message.
type CancelledAction struct {
	CancelledAt *time.Time `json:"cancelledAt,omitempty"`
	// The cancelledByUserId field.
	CancelledByUserID *string `json:"cancelledByUserId,omitempty"`
}
