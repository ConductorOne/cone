// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIUserV1UserSearchSearchResponse struct {
	ContentType string
	// Successful response
	SearchUsersResponse *shared.SearchUsersResponse
	StatusCode          int
	RawResponse         *http.Response
}

func (o *C1APIUserV1UserSearchSearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIUserV1UserSearchSearchResponse) GetSearchUsersResponse() *shared.SearchUsersResponse {
	if o == nil {
		return nil
	}
	return o.SearchUsersResponse
}

func (o *C1APIUserV1UserSearchSearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIUserV1UserSearchSearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
