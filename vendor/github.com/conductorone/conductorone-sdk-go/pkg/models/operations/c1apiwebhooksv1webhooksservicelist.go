// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIWebhooksV1WebhooksServiceListRequest struct {
	PageSize  *int    `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIWebhooksV1WebhooksServiceListRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIWebhooksV1WebhooksServiceListRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIWebhooksV1WebhooksServiceListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	WebhooksServiceListResponse *shared.WebhooksServiceListResponse
}

func (o *C1APIWebhooksV1WebhooksServiceListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIWebhooksV1WebhooksServiceListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIWebhooksV1WebhooksServiceListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIWebhooksV1WebhooksServiceListResponse) GetWebhooksServiceListResponse() *shared.WebhooksServiceListResponse {
	if o == nil {
		return nil
	}
	return o.WebhooksServiceListResponse
}
