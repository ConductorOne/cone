// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppOwnersListRequest struct {
	AppID string `pathParam:"style=simple,explode=false,name=app_id"`
}

type C1APIAppV1AppOwnersListResponse struct {
	ContentType string
	// Successful response
	ListAppOwnersResponse *shared.ListAppOwnersResponse
	StatusCode            int
	RawResponse           *http.Response
}
