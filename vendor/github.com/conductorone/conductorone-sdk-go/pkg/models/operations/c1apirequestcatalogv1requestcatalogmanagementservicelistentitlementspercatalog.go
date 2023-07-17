// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalog_id"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogResponse struct {
	ContentType string
	// Successful response
	RequestCatalogManagementServiceListEntitlementsPerCatalogResponse *shared.RequestCatalogManagementServiceListEntitlementsPerCatalogResponse
	StatusCode                                                        int
	RawResponse                                                       *http.Response
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogResponse) GetRequestCatalogManagementServiceListEntitlementsPerCatalogResponse() *shared.RequestCatalogManagementServiceListEntitlementsPerCatalogResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceListEntitlementsPerCatalogResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
