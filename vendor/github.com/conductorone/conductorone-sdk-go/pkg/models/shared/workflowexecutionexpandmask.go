// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The WorkflowExecutionExpandMask message.
type WorkflowExecutionExpandMask struct {
	// The paths field.
	Paths []string `json:"paths,omitempty"`
}

func (o *WorkflowExecutionExpandMask) GetPaths() []string {
	if o == nil {
		return nil
	}
	return o.Paths
}
