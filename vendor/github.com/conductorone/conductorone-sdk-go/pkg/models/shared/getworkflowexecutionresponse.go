// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The GetWorkflowExecutionResponse message.
type GetWorkflowExecutionResponse struct {
	// The WorkflowExecution message.
	WorkflowExecution *WorkflowExecution `json:"workflowExecution,omitempty"`
}

func (o *GetWorkflowExecutionResponse) GetWorkflowExecution() *WorkflowExecution {
	if o == nil {
		return nil
	}
	return o.WorkflowExecution
}
