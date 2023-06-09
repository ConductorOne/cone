// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// ReassignedAction - The ReassignedAction message.
type ReassignedAction struct {
	// The newPolicyStepId field.
	NewPolicyStepID *string    `json:"newPolicyStepId,omitempty"`
	ReassignedAt    *time.Time `json:"reassignedAt,omitempty"`
	// The userId field.
	UserID *string `json:"userId,omitempty"`
}
