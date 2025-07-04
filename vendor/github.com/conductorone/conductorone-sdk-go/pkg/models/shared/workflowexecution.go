// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// WorkflowExecutionState - The state field.
type WorkflowExecutionState string

const (
	WorkflowExecutionStateWorkflowExecutionStateUnspecified  WorkflowExecutionState = "WORKFLOW_EXECUTION_STATE_UNSPECIFIED"
	WorkflowExecutionStateWorkflowExecutionStatePending      WorkflowExecutionState = "WORKFLOW_EXECUTION_STATE_PENDING"
	WorkflowExecutionStateWorkflowExecutionStateCreating     WorkflowExecutionState = "WORKFLOW_EXECUTION_STATE_CREATING"
	WorkflowExecutionStateWorkflowExecutionStateGetStep      WorkflowExecutionState = "WORKFLOW_EXECUTION_STATE_GET_STEP"
	WorkflowExecutionStateWorkflowExecutionStateProcessStep  WorkflowExecutionState = "WORKFLOW_EXECUTION_STATE_PROCESS_STEP"
	WorkflowExecutionStateWorkflowExecutionStateCompleteStep WorkflowExecutionState = "WORKFLOW_EXECUTION_STATE_COMPLETE_STEP"
	WorkflowExecutionStateWorkflowExecutionStateDone         WorkflowExecutionState = "WORKFLOW_EXECUTION_STATE_DONE"
	WorkflowExecutionStateWorkflowExecutionStateError        WorkflowExecutionState = "WORKFLOW_EXECUTION_STATE_ERROR"
	WorkflowExecutionStateWorkflowExecutionStateTerminate    WorkflowExecutionState = "WORKFLOW_EXECUTION_STATE_TERMINATE"
	WorkflowExecutionStateWorkflowExecutionStateWaiting      WorkflowExecutionState = "WORKFLOW_EXECUTION_STATE_WAITING"
)

func (e WorkflowExecutionState) ToPointer() *WorkflowExecutionState {
	return &e
}

// The WorkflowExecution message.
type WorkflowExecution struct {
	// The WorkflowContext message.
	WorkflowContext *WorkflowContext `json:"context,omitempty"`
	CompletedAt     *time.Time       `json:"completedAt,omitempty"`
	CreatedAt       *time.Time       `json:"createdAt,omitempty"`
	// The currentVersion field.
	CurrentVersion *int       `json:"currentVersion,omitempty"`
	DeletedAt      *time.Time `json:"deletedAt,omitempty"`
	Duration       *string    `json:"duration,omitempty"`
	// The id field.
	ID *int64 `integer:"string" json:"id,omitempty"`
	// The isDraft field.
	IsDraft *bool `json:"isDraft,omitempty"`
	// The state field.
	State     *WorkflowExecutionState `json:"state,omitempty"`
	UpdatedAt *time.Time              `json:"updatedAt,omitempty"`
	// The workflowTemplateId field.
	WorkflowTemplateID *string `json:"workflowTemplateId,omitempty"`
}

func (w WorkflowExecution) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WorkflowExecution) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WorkflowExecution) GetWorkflowContext() *WorkflowContext {
	if o == nil {
		return nil
	}
	return o.WorkflowContext
}

func (o *WorkflowExecution) GetCompletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *WorkflowExecution) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *WorkflowExecution) GetCurrentVersion() *int {
	if o == nil {
		return nil
	}
	return o.CurrentVersion
}

func (o *WorkflowExecution) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *WorkflowExecution) GetDuration() *string {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *WorkflowExecution) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WorkflowExecution) GetIsDraft() *bool {
	if o == nil {
		return nil
	}
	return o.IsDraft
}

func (o *WorkflowExecution) GetState() *WorkflowExecutionState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *WorkflowExecution) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *WorkflowExecution) GetWorkflowTemplateID() *string {
	if o == nil {
		return nil
	}
	return o.WorkflowTemplateID
}
