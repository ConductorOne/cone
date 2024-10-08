// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppResourceTypeServiceDeleteManuallyManagedResourceTypeRequest struct {
	DeleteManuallyManagedResourceTypeRequest *shared.DeleteManuallyManagedResourceTypeRequest `request:"mediaType=application/json"`
	AppID                                    string                                           `pathParam:"style=simple,explode=false,name=app_id"`
	ID                                       string                                           `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAppV1AppResourceTypeServiceDeleteManuallyManagedResourceTypeRequest) GetDeleteManuallyManagedResourceTypeRequest() *shared.DeleteManuallyManagedResourceTypeRequest {
	if o == nil {
		return nil
	}
	return o.DeleteManuallyManagedResourceTypeRequest
}

func (o *C1APIAppV1AppResourceTypeServiceDeleteManuallyManagedResourceTypeRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppResourceTypeServiceDeleteManuallyManagedResourceTypeRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAppV1AppResourceTypeServiceDeleteManuallyManagedResourceTypeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	DeleteManuallyManagedResourceTypeResponse *shared.DeleteManuallyManagedResourceTypeResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppResourceTypeServiceDeleteManuallyManagedResourceTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppResourceTypeServiceDeleteManuallyManagedResourceTypeResponse) GetDeleteManuallyManagedResourceTypeResponse() *shared.DeleteManuallyManagedResourceTypeResponse {
	if o == nil {
		return nil
	}
	return o.DeleteManuallyManagedResourceTypeResponse
}

func (o *C1APIAppV1AppResourceTypeServiceDeleteManuallyManagedResourceTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppResourceTypeServiceDeleteManuallyManagedResourceTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
