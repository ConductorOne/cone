// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// TaskActionsServiceRestartResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type TaskActionsServiceRestartResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (t TaskActionsServiceRestartResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskActionsServiceRestartResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskActionsServiceRestartResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *TaskActionsServiceRestartResponseExpanded) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The TaskActionsServiceRestartResponse message.
type TaskActionsServiceRestartResponse struct {
	// Contains a task and JSONPATH expressions that describe where in the expanded array related objects are located. This view can be used to display a fully-detailed dashboard of task information.
	TaskView *TaskView `json:"taskView,omitempty"`
	// The expanded field.
	Expanded []TaskActionsServiceRestartResponseExpanded `json:"expanded,omitempty"`
	// The ticketActionId field.
	TicketActionID *string `json:"ticketActionId,omitempty"`
}

func (o *TaskActionsServiceRestartResponse) GetTaskView() *TaskView {
	if o == nil {
		return nil
	}
	return o.TaskView
}

func (o *TaskActionsServiceRestartResponse) GetExpanded() []TaskActionsServiceRestartResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *TaskActionsServiceRestartResponse) GetTicketActionID() *string {
	if o == nil {
		return nil
	}
	return o.TicketActionID
}
