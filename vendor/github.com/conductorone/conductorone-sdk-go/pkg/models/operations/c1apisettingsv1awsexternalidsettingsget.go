// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APISettingsV1AWSExternalIDSettingsGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	GetAWSExternalIDResponse *shared.GetAWSExternalIDResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APISettingsV1AWSExternalIDSettingsGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APISettingsV1AWSExternalIDSettingsGetResponse) GetGetAWSExternalIDResponse() *shared.GetAWSExternalIDResponse {
	if o == nil {
		return nil
	}
	return o.GetAWSExternalIDResponse
}

func (o *C1APISettingsV1AWSExternalIDSettingsGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APISettingsV1AWSExternalIDSettingsGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
