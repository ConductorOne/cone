// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIIamV1RolesGetRequest struct {
	RoleID string `pathParam:"style=simple,explode=false,name=role_id"`
}

func (o *C1APIIamV1RolesGetRequest) GetRoleID() string {
	if o == nil {
		return ""
	}
	return o.RoleID
}

type C1APIIamV1RolesGetResponse struct {
	ContentType string
	// Successful response
	GetRolesResponse *shared.GetRolesResponse
	StatusCode       int
	RawResponse      *http.Response
}

func (o *C1APIIamV1RolesGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIIamV1RolesGetResponse) GetGetRolesResponse() *shared.GetRolesResponse {
	if o == nil {
		return nil
	}
	return o.GetRolesResponse
}

func (o *C1APIIamV1RolesGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIIamV1RolesGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
