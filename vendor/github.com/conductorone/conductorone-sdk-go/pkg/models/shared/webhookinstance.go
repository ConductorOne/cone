// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// WebhookInstanceState - The state field.
type WebhookInstanceState string

const (
	WebhookInstanceStateWebhookStateUnspecified     WebhookInstanceState = "WEBHOOK_STATE_UNSPECIFIED"
	WebhookInstanceStateWebhookStatePending         WebhookInstanceState = "WEBHOOK_STATE_PENDING"
	WebhookInstanceStateWebhookStateRunning         WebhookInstanceState = "WEBHOOK_STATE_RUNNING"
	WebhookInstanceStateWebhookStateError           WebhookInstanceState = "WEBHOOK_STATE_ERROR"
	WebhookInstanceStateWebhookStateWaitingCallback WebhookInstanceState = "WEBHOOK_STATE_WAITING_CALLBACK"
	WebhookInstanceStateWebhookStateProcessResponse WebhookInstanceState = "WEBHOOK_STATE_PROCESS_RESPONSE"
	WebhookInstanceStateWebhookStateSuccess         WebhookInstanceState = "WEBHOOK_STATE_SUCCESS"
	WebhookInstanceStateWebhookStateFatalError      WebhookInstanceState = "WEBHOOK_STATE_FATAL_ERROR"
)

func (e WebhookInstanceState) ToPointer() *WebhookInstanceState {
	return &e
}
func (e *WebhookInstanceState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "WEBHOOK_STATE_UNSPECIFIED":
		fallthrough
	case "WEBHOOK_STATE_PENDING":
		fallthrough
	case "WEBHOOK_STATE_RUNNING":
		fallthrough
	case "WEBHOOK_STATE_ERROR":
		fallthrough
	case "WEBHOOK_STATE_WAITING_CALLBACK":
		fallthrough
	case "WEBHOOK_STATE_PROCESS_RESPONSE":
		fallthrough
	case "WEBHOOK_STATE_SUCCESS":
		fallthrough
	case "WEBHOOK_STATE_FATAL_ERROR":
		*e = WebhookInstanceState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WebhookInstanceState: %v", v)
	}
}

// The WebhookInstance message.
type WebhookInstance struct {
	// The WebhookSource message.
	//
	// This message contains a oneof named source. Only a single field of the following list may be set at a time:
	//   - test
	//   - policyPostAction
	//   - approvalStep
	//   - provisionStep
	//
	WebhookSource *WebhookSource `json:"source,omitempty"`
	// The WebhookSpec message.
	WebhookSpec *WebhookSpec `json:"spec,omitempty"`
	// The attempts field.
	Attempts    *int       `json:"attempts,omitempty"`
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	ExpiresAt   *time.Time `json:"expiresAt,omitempty"`
	// The id field.
	ID              *string    `json:"id,omitempty"`
	LastAttemptedAt *time.Time `json:"lastAttemptedAt,omitempty"`
	// The state field.
	State     *WebhookInstanceState `json:"state,omitempty"`
	UpdatedAt *time.Time            `json:"updatedAt,omitempty"`
	// The webhookId field.
	WebhookID *string `json:"webhookId,omitempty"`
}

func (w WebhookInstance) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WebhookInstance) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WebhookInstance) GetWebhookSource() *WebhookSource {
	if o == nil {
		return nil
	}
	return o.WebhookSource
}

func (o *WebhookInstance) GetWebhookSpec() *WebhookSpec {
	if o == nil {
		return nil
	}
	return o.WebhookSpec
}

func (o *WebhookInstance) GetAttempts() *int {
	if o == nil {
		return nil
	}
	return o.Attempts
}

func (o *WebhookInstance) GetCompletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *WebhookInstance) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *WebhookInstance) GetExpiresAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *WebhookInstance) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WebhookInstance) GetLastAttemptedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastAttemptedAt
}

func (o *WebhookInstance) GetState() *WebhookInstanceState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *WebhookInstance) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *WebhookInstance) GetWebhookID() *string {
	if o == nil {
		return nil
	}
	return o.WebhookID
}