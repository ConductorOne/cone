// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The PayloadWorkflowStep message.
type PayloadWorkflowStep struct {
	Context map[string]any `json:"context,omitempty"`
	// The workflow execution ID
	WorkflowExecutionID *int64 `integer:"string" json:"workflowExecutionId,omitempty"`
	// The workflow execution step ID
	WorkflowExecutionStepID *string `json:"workflowExecutionStepId,omitempty"`
	// The workflow template ID
	WorkflowID *string `json:"workflowId,omitempty"`
}

func (o *PayloadWorkflowStep) GetContext() map[string]any {
	if o == nil {
		return nil
	}
	return o.Context
}

func (o *PayloadWorkflowStep) GetWorkflowExecutionID() *int64 {
	if o == nil {
		return nil
	}
	return o.WorkflowExecutionID
}

func (o *PayloadWorkflowStep) GetWorkflowExecutionStepID() *string {
	if o == nil {
		return nil
	}
	return o.WorkflowExecutionStepID
}

func (o *PayloadWorkflowStep) GetWorkflowID() *string {
	if o == nil {
		return nil
	}
	return o.WorkflowID
}
