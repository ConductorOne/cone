// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIIamV1PersonalClientServiceGetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIIamV1PersonalClientServiceGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIIamV1PersonalClientServiceGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	PersonalClientServiceGetResponse *shared.PersonalClientServiceGetResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIIamV1PersonalClientServiceGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIIamV1PersonalClientServiceGetResponse) GetPersonalClientServiceGetResponse() *shared.PersonalClientServiceGetResponse {
	if o == nil {
		return nil
	}
	return o.PersonalClientServiceGetResponse
}

func (o *C1APIIamV1PersonalClientServiceGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIIamV1PersonalClientServiceGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}