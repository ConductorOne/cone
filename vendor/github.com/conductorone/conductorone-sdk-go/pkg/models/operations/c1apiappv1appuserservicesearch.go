// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppUserServiceSearchResponse struct {
	// Successful response
	AppUserServiceSearchResponse *shared.AppUserServiceSearchResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppUserServiceSearchResponse) GetAppUserServiceSearchResponse() *shared.AppUserServiceSearchResponse {
	if o == nil {
		return nil
	}
	return o.AppUserServiceSearchResponse
}

func (o *C1APIAppV1AppUserServiceSearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppUserServiceSearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppUserServiceSearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
