// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIWorkflowsV1betaWorkflowServiceListWorkflowsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	ListWorkflowsResponse *shared.ListWorkflowsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIWorkflowsV1betaWorkflowServiceListWorkflowsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIWorkflowsV1betaWorkflowServiceListWorkflowsResponse) GetListWorkflowsResponse() *shared.ListWorkflowsResponse {
	if o == nil {
		return nil
	}
	return o.ListWorkflowsResponse
}

func (o *C1APIWorkflowsV1betaWorkflowServiceListWorkflowsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIWorkflowsV1betaWorkflowServiceListWorkflowsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}