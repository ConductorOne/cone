// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ProvisionInstanceState - This property indicates the current state of this step.
type ProvisionInstanceState string

const (
	ProvisionInstanceStateProvisionInstanceStateUnspecified                     ProvisionInstanceState = "PROVISION_INSTANCE_STATE_UNSPECIFIED"
	ProvisionInstanceStateProvisionInstanceStateInit                            ProvisionInstanceState = "PROVISION_INSTANCE_STATE_INIT"
	ProvisionInstanceStateProvisionInstanceStateCreateConnectorActionsForTarget ProvisionInstanceState = "PROVISION_INSTANCE_STATE_CREATE_CONNECTOR_ACTIONS_FOR_TARGET"
	ProvisionInstanceStateProvisionInstanceStateSendingNotifications            ProvisionInstanceState = "PROVISION_INSTANCE_STATE_SENDING_NOTIFICATIONS"
	ProvisionInstanceStateProvisionInstanceStateWaiting                         ProvisionInstanceState = "PROVISION_INSTANCE_STATE_WAITING"
	ProvisionInstanceStateProvisionInstanceStateWebhook                         ProvisionInstanceState = "PROVISION_INSTANCE_STATE_WEBHOOK"
	ProvisionInstanceStateProvisionInstanceStateWebhookWaiting                  ProvisionInstanceState = "PROVISION_INSTANCE_STATE_WEBHOOK_WAITING"
	ProvisionInstanceStateProvisionInstanceStateExternalTicket                  ProvisionInstanceState = "PROVISION_INSTANCE_STATE_EXTERNAL_TICKET"
	ProvisionInstanceStateProvisionInstanceStateExternalTicketWaiting           ProvisionInstanceState = "PROVISION_INSTANCE_STATE_EXTERNAL_TICKET_WAITING"
	ProvisionInstanceStateProvisionInstanceStateDone                            ProvisionInstanceState = "PROVISION_INSTANCE_STATE_DONE"
)

func (e ProvisionInstanceState) ToPointer() *ProvisionInstanceState {
	return &e
}
func (e *ProvisionInstanceState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PROVISION_INSTANCE_STATE_UNSPECIFIED":
		fallthrough
	case "PROVISION_INSTANCE_STATE_INIT":
		fallthrough
	case "PROVISION_INSTANCE_STATE_CREATE_CONNECTOR_ACTIONS_FOR_TARGET":
		fallthrough
	case "PROVISION_INSTANCE_STATE_SENDING_NOTIFICATIONS":
		fallthrough
	case "PROVISION_INSTANCE_STATE_WAITING":
		fallthrough
	case "PROVISION_INSTANCE_STATE_WEBHOOK":
		fallthrough
	case "PROVISION_INSTANCE_STATE_WEBHOOK_WAITING":
		fallthrough
	case "PROVISION_INSTANCE_STATE_EXTERNAL_TICKET":
		fallthrough
	case "PROVISION_INSTANCE_STATE_EXTERNAL_TICKET_WAITING":
		fallthrough
	case "PROVISION_INSTANCE_STATE_DONE":
		*e = ProvisionInstanceState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProvisionInstanceState: %v", v)
	}
}

// ProvisionInstance - A provision instance describes the specific configuration of an executing provision policy step including actions taken and notification id.
//
// This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
//   - completed
//   - cancelled
//   - errored
//   - reassignedByError
type ProvisionInstance struct {
	// The outcome of a provision instance that is cancelled.
	CancelledAction *CancelledAction `json:"cancelled,omitempty"`
	// The outcome of a provision instance that has been completed succesfully.
	CompletedAction *CompletedAction `json:"completed,omitempty"`
	// The outcome of a provision instance that has errored.
	ErroredAction *ErroredAction `json:"errored,omitempty"`
	// The provision step references a provision policy for this step.
	Provision *Provision `json:"provision,omitempty"`
	// The ReassignedByErrorAction object describes the outcome of a policy step that has been reassigned because it had an error provisioning.
	ReassignedByErrorAction *ReassignedByErrorAction `json:"reassignedByError,omitempty"`
	// This indicates the notification id for this step.
	NotificationID *string `json:"notificationId,omitempty"`
	// This property indicates the current state of this step.
	State *ProvisionInstanceState `json:"state,omitempty"`
}

func (o *ProvisionInstance) GetCancelledAction() *CancelledAction {
	if o == nil {
		return nil
	}
	return o.CancelledAction
}

func (o *ProvisionInstance) GetCompletedAction() *CompletedAction {
	if o == nil {
		return nil
	}
	return o.CompletedAction
}

func (o *ProvisionInstance) GetErroredAction() *ErroredAction {
	if o == nil {
		return nil
	}
	return o.ErroredAction
}

func (o *ProvisionInstance) GetProvision() *Provision {
	if o == nil {
		return nil
	}
	return o.Provision
}

func (o *ProvisionInstance) GetReassignedByErrorAction() *ReassignedByErrorAction {
	if o == nil {
		return nil
	}
	return o.ReassignedByErrorAction
}

func (o *ProvisionInstance) GetNotificationID() *string {
	if o == nil {
		return nil
	}
	return o.NotificationID
}

func (o *ProvisionInstance) GetState() *ProvisionInstanceState {
	if o == nil {
		return nil
	}
	return o.State
}
