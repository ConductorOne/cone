// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceCreateBundleAutomationRequest struct {
	CreateBundleAutomationRequest *shared.CreateBundleAutomationRequest `request:"mediaType=application/json"`
	RequestCatalogID              string                                `pathParam:"style=simple,explode=false,name=request_catalog_id"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateBundleAutomationRequest) GetCreateBundleAutomationRequest() *shared.CreateBundleAutomationRequest {
	if o == nil {
		return nil
	}
	return o.CreateBundleAutomationRequest
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateBundleAutomationRequest) GetRequestCatalogID() string {
	if o == nil {
		return ""
	}
	return o.RequestCatalogID
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceCreateBundleAutomationResponse struct {
	// Successful response
	BundleAutomation *shared.BundleAutomation
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateBundleAutomationResponse) GetBundleAutomation() *shared.BundleAutomation {
	if o == nil {
		return nil
	}
	return o.BundleAutomation
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateBundleAutomationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateBundleAutomationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateBundleAutomationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
