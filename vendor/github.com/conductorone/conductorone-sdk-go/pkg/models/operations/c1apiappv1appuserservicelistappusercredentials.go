// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppUserServiceListAppUserCredentialsRequest struct {
	AppID     string  `pathParam:"style=simple,explode=false,name=app_id"`
	AppUserID string  `pathParam:"style=simple,explode=false,name=app_user_id"`
	PageSize  *int    `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1AppUserServiceListAppUserCredentialsRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppUserServiceListAppUserCredentialsRequest) GetAppUserID() string {
	if o == nil {
		return ""
	}
	return o.AppUserID
}

func (o *C1APIAppV1AppUserServiceListAppUserCredentialsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1AppUserServiceListAppUserCredentialsRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1AppUserServiceListAppUserCredentialsResponse struct {
	// Successful response
	AppUserServiceListCredentialsResponse *shared.AppUserServiceListCredentialsResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppUserServiceListAppUserCredentialsResponse) GetAppUserServiceListCredentialsResponse() *shared.AppUserServiceListCredentialsResponse {
	if o == nil {
		return nil
	}
	return o.AppUserServiceListCredentialsResponse
}

func (o *C1APIAppV1AppUserServiceListAppUserCredentialsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppUserServiceListAppUserCredentialsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppUserServiceListAppUserCredentialsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
