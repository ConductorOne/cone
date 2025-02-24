// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The UpdateWorkflowRequest message.
type UpdateWorkflowRequest struct {
	// The Workflow message.
	Workflow   *WorkflowInput `json:"workflow,omitempty"`
	UpdateMask *string        `json:"updateMask,omitempty"`
}

func (o *UpdateWorkflowRequest) GetWorkflow() *WorkflowInput {
	if o == nil {
		return nil
	}
	return o.Workflow
}

func (o *UpdateWorkflowRequest) GetUpdateMask() *string {
	if o == nil {
		return nil
	}
	return o.UpdateMask
}
