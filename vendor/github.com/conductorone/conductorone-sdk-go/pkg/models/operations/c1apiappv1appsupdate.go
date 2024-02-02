// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppsUpdateRequest struct {
	UpdateAppRequest *shared.UpdateAppRequest `request:"mediaType=application/json"`
	ID               string                   `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAppV1AppsUpdateRequest) GetUpdateAppRequest() *shared.UpdateAppRequest {
	if o == nil {
		return nil
	}
	return o.UpdateAppRequest
}

func (o *C1APIAppV1AppsUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAppV1AppsUpdateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns the updated app's new values.
	UpdateAppResponse *shared.UpdateAppResponse
}

func (o *C1APIAppV1AppsUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppsUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppsUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1AppsUpdateResponse) GetUpdateAppResponse() *shared.UpdateAppResponse {
	if o == nil {
		return nil
	}
	return o.UpdateAppResponse
}
