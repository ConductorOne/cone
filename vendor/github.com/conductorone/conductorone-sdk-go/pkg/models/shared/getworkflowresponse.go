// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The GetWorkflowResponse message.
type GetWorkflowResponse struct {
	// The Workflow message.
	Workflow *Workflow `json:"workflow,omitempty"`
}

func (o *GetWorkflowResponse) GetWorkflow() *Workflow {
	if o == nil {
		return nil
	}
	return o.Workflow
}