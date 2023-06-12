# \RequestCatalogSearchAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements**](RequestCatalogSearchAPI.md#C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements) | **Post** /api/v1/search/request_catalog/entitlements | 



## C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements

> C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements(ctx).C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest(c1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/conductorone/cone/internal/c1api"
)

func main() {
    c1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest := *openapiclient.NewC1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest() // C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestCatalogSearchAPI.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements(context.Background()).C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest(c1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestCatalogSearchAPI.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements`: C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse
    fmt.Fprintf(os.Stdout, "Response from `RequestCatalogSearchAPI.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **c1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest** | [**C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest**](C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest.md) |  | 

### Return type

[**C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse**](C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

