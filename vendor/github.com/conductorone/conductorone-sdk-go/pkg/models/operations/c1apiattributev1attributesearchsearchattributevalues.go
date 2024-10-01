// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// SearchAttributeValuesResponse is the response for searching AttributeValues.
	SearchAttributeValuesResponse *shared.SearchAttributeValuesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse) GetSearchAttributeValuesResponse() *shared.SearchAttributeValuesResponse {
	if o == nil {
		return nil
	}
	return o.SearchAttributeValuesResponse
}

func (o *C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
