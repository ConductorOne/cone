// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIWorkflowsV1betaWorkflowServiceGetWorkflowRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIWorkflowsV1betaWorkflowServiceGetWorkflowRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIWorkflowsV1betaWorkflowServiceGetWorkflowResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	GetWorkflowResponse *shared.GetWorkflowResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIWorkflowsV1betaWorkflowServiceGetWorkflowResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIWorkflowsV1betaWorkflowServiceGetWorkflowResponse) GetGetWorkflowResponse() *shared.GetWorkflowResponse {
	if o == nil {
		return nil
	}
	return o.GetWorkflowResponse
}

func (o *C1APIWorkflowsV1betaWorkflowServiceGetWorkflowResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIWorkflowsV1betaWorkflowServiceGetWorkflowResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}