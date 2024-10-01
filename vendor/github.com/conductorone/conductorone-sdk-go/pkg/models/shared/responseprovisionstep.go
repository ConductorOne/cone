// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ResponseProvisionStep message.
//
// This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
//   - complete
//   - errored
type ResponseProvisionStep struct {
	// The ResponseProvisionStepComplete message.
	ResponseProvisionStepComplete *ResponseProvisionStepComplete `json:"complete,omitempty"`
	// The ResponseProvisionStepErrored message.
	ResponseProvisionStepErrored *ResponseProvisionStepErrored `json:"errored,omitempty"`
	// version contains the constant value "v1". Future versions of the Webhook Response
	//  will use a different string.
	Version *string `json:"version,omitempty"`
}

func (o *ResponseProvisionStep) GetResponseProvisionStepComplete() *ResponseProvisionStepComplete {
	if o == nil {
		return nil
	}
	return o.ResponseProvisionStepComplete
}

func (o *ResponseProvisionStep) GetResponseProvisionStepErrored() *ResponseProvisionStepErrored {
	if o == nil {
		return nil
	}
	return o.ResponseProvisionStepErrored
}

func (o *ResponseProvisionStep) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}
