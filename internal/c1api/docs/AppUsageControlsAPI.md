# \AppUsageControlsAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiAppV1AppUsageControlsServiceGet**](AppUsageControlsAPI.md#C1ApiAppV1AppUsageControlsServiceGet) | **Get** /api/v1/apps/{app_id}/usage_controls | 



## C1ApiAppV1AppUsageControlsServiceGet

> C1ApiAppV1GetAppUsageControlsResponse C1ApiAppV1AppUsageControlsServiceGet(ctx, appId).Execute()





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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppUsageControlsAPI.C1ApiAppV1AppUsageControlsServiceGet(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppUsageControlsAPI.C1ApiAppV1AppUsageControlsServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAppV1AppUsageControlsServiceGet`: C1ApiAppV1GetAppUsageControlsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppUsageControlsAPI.C1ApiAppV1AppUsageControlsServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiAppV1AppUsageControlsServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**C1ApiAppV1GetAppUsageControlsResponse**](C1ApiAppV1GetAppUsageControlsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

