// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppUserServiceListAppUsersForUserRequest struct {
	AppID  string `pathParam:"style=simple,explode=false,name=app_id"`
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

func (o *C1APIAppV1AppUserServiceListAppUsersForUserRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppUserServiceListAppUsersForUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type C1APIAppV1AppUserServiceListAppUsersForUserResponse struct {
	// Successful response
	AppUsersForUserServiceListResponse *shared.AppUsersForUserServiceListResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppUserServiceListAppUsersForUserResponse) GetAppUsersForUserServiceListResponse() *shared.AppUsersForUserServiceListResponse {
	if o == nil {
		return nil
	}
	return o.AppUsersForUserServiceListResponse
}

func (o *C1APIAppV1AppUserServiceListAppUsersForUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppUserServiceListAppUsersForUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppUserServiceListAppUsersForUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
