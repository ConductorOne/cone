# \AppResourceSearchAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiAppV1AppResourceSearchSearchAppResourceTypes**](AppResourceSearchAPI.md#C1ApiAppV1AppResourceSearchSearchAppResourceTypes) | **Post** /api/v1/search/app_resource_types | 



## C1ApiAppV1AppResourceSearchSearchAppResourceTypes

> C1ApiAppV1SearchAppResourceTypesResponse C1ApiAppV1AppResourceSearchSearchAppResourceTypes(ctx).C1ApiAppV1SearchAppResourceTypesRequest(c1ApiAppV1SearchAppResourceTypesRequest).Execute()





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
    c1ApiAppV1SearchAppResourceTypesRequest := *openapiclient.NewC1ApiAppV1SearchAppResourceTypesRequest() // C1ApiAppV1SearchAppResourceTypesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppResourceSearchAPI.C1ApiAppV1AppResourceSearchSearchAppResourceTypes(context.Background()).C1ApiAppV1SearchAppResourceTypesRequest(c1ApiAppV1SearchAppResourceTypesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppResourceSearchAPI.C1ApiAppV1AppResourceSearchSearchAppResourceTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAppV1AppResourceSearchSearchAppResourceTypes`: C1ApiAppV1SearchAppResourceTypesResponse
    fmt.Fprintf(os.Stdout, "Response from `AppResourceSearchAPI.C1ApiAppV1AppResourceSearchSearchAppResourceTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiAppV1AppResourceSearchSearchAppResourceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **c1ApiAppV1SearchAppResourceTypesRequest** | [**C1ApiAppV1SearchAppResourceTypesRequest**](C1ApiAppV1SearchAppResourceTypesRequest.md) |  | 

### Return type

[**C1ApiAppV1SearchAppResourceTypesResponse**](C1ApiAppV1SearchAppResourceTypesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

