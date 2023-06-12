# \AppsAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiAppV1AppsGet**](AppsAPI.md#C1ApiAppV1AppsGet) | **Get** /api/v1/apps/{id} | 
[**C1ApiAppV1AppsList**](AppsAPI.md#C1ApiAppV1AppsList) | **Get** /api/v1/apps | 



## C1ApiAppV1AppsGet

> C1ApiAppV1GetAppResponse C1ApiAppV1AppsGet(ctx, id).Execute()





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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.C1ApiAppV1AppsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.C1ApiAppV1AppsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAppV1AppsGet`: C1ApiAppV1GetAppResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.C1ApiAppV1AppsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiAppV1AppsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**C1ApiAppV1GetAppResponse**](C1ApiAppV1GetAppResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## C1ApiAppV1AppsList

> C1ApiAppV1ListAppsResponse C1ApiAppV1AppsList(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsAPI.C1ApiAppV1AppsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.C1ApiAppV1AppsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAppV1AppsList`: C1ApiAppV1ListAppsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsAPI.C1ApiAppV1AppsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiAppV1AppsListRequest struct via the builder pattern


### Return type

[**C1ApiAppV1ListAppsResponse**](C1ApiAppV1ListAppsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

