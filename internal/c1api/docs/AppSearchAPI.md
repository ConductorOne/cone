# \AppSearchAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiAppV1AppSearchSearch**](AppSearchAPI.md#C1ApiAppV1AppSearchSearch) | **Post** /api/v1/search/apps | 



## C1ApiAppV1AppSearchSearch

> C1ApiAppV1SearchAppsResponse C1ApiAppV1AppSearchSearch(ctx).C1ApiAppV1SearchAppsRequest(c1ApiAppV1SearchAppsRequest).Execute()





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
    c1ApiAppV1SearchAppsRequest := *openapiclient.NewC1ApiAppV1SearchAppsRequest() // C1ApiAppV1SearchAppsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppSearchAPI.C1ApiAppV1AppSearchSearch(context.Background()).C1ApiAppV1SearchAppsRequest(c1ApiAppV1SearchAppsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppSearchAPI.C1ApiAppV1AppSearchSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAppV1AppSearchSearch`: C1ApiAppV1SearchAppsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppSearchAPI.C1ApiAppV1AppSearchSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiAppV1AppSearchSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **c1ApiAppV1SearchAppsRequest** | [**C1ApiAppV1SearchAppsRequest**](C1ApiAppV1SearchAppsRequest.md) |  | 

### Return type

[**C1ApiAppV1SearchAppsResponse**](C1ApiAppV1SearchAppsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

