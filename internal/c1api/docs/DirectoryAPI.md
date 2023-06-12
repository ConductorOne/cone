# \DirectoryAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiDirectoryV1DirectoryServiceCreate**](DirectoryAPI.md#C1ApiDirectoryV1DirectoryServiceCreate) | **Post** /api/v1/directories | 
[**C1ApiDirectoryV1DirectoryServiceDelete**](DirectoryAPI.md#C1ApiDirectoryV1DirectoryServiceDelete) | **Delete** /api/v1/directories/{app_id} | 
[**C1ApiDirectoryV1DirectoryServiceGet**](DirectoryAPI.md#C1ApiDirectoryV1DirectoryServiceGet) | **Get** /api/v1/directories/{app_id} | 
[**C1ApiDirectoryV1DirectoryServiceList**](DirectoryAPI.md#C1ApiDirectoryV1DirectoryServiceList) | **Get** /api/v1/directories | 



## C1ApiDirectoryV1DirectoryServiceCreate

> C1ApiDirectoryV1DirectoryServiceCreateResponse C1ApiDirectoryV1DirectoryServiceCreate(ctx).C1ApiDirectoryV1DirectoryServiceCreateRequest(c1ApiDirectoryV1DirectoryServiceCreateRequest).Execute()





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
    c1ApiDirectoryV1DirectoryServiceCreateRequest := *openapiclient.NewC1ApiDirectoryV1DirectoryServiceCreateRequest() // C1ApiDirectoryV1DirectoryServiceCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectoryAPI.C1ApiDirectoryV1DirectoryServiceCreate(context.Background()).C1ApiDirectoryV1DirectoryServiceCreateRequest(c1ApiDirectoryV1DirectoryServiceCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAPI.C1ApiDirectoryV1DirectoryServiceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiDirectoryV1DirectoryServiceCreate`: C1ApiDirectoryV1DirectoryServiceCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAPI.C1ApiDirectoryV1DirectoryServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiDirectoryV1DirectoryServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **c1ApiDirectoryV1DirectoryServiceCreateRequest** | [**C1ApiDirectoryV1DirectoryServiceCreateRequest**](C1ApiDirectoryV1DirectoryServiceCreateRequest.md) |  | 

### Return type

[**C1ApiDirectoryV1DirectoryServiceCreateResponse**](C1ApiDirectoryV1DirectoryServiceCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## C1ApiDirectoryV1DirectoryServiceDelete

> map[string]interface{} C1ApiDirectoryV1DirectoryServiceDelete(ctx, appId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DirectoryAPI.C1ApiDirectoryV1DirectoryServiceDelete(context.Background(), appId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAPI.C1ApiDirectoryV1DirectoryServiceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiDirectoryV1DirectoryServiceDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAPI.C1ApiDirectoryV1DirectoryServiceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiDirectoryV1DirectoryServiceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## C1ApiDirectoryV1DirectoryServiceGet

> C1ApiDirectoryV1DirectoryServiceGetResponse C1ApiDirectoryV1DirectoryServiceGet(ctx, appId).Execute()





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
    resp, r, err := apiClient.DirectoryAPI.C1ApiDirectoryV1DirectoryServiceGet(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAPI.C1ApiDirectoryV1DirectoryServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiDirectoryV1DirectoryServiceGet`: C1ApiDirectoryV1DirectoryServiceGetResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAPI.C1ApiDirectoryV1DirectoryServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiDirectoryV1DirectoryServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**C1ApiDirectoryV1DirectoryServiceGetResponse**](C1ApiDirectoryV1DirectoryServiceGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## C1ApiDirectoryV1DirectoryServiceList

> C1ApiDirectoryV1DirectoryServiceListResponse C1ApiDirectoryV1DirectoryServiceList(ctx).Execute()





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
    resp, r, err := apiClient.DirectoryAPI.C1ApiDirectoryV1DirectoryServiceList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryAPI.C1ApiDirectoryV1DirectoryServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiDirectoryV1DirectoryServiceList`: C1ApiDirectoryV1DirectoryServiceListResponse
    fmt.Fprintf(os.Stdout, "Response from `DirectoryAPI.C1ApiDirectoryV1DirectoryServiceList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiDirectoryV1DirectoryServiceListRequest struct via the builder pattern


### Return type

[**C1ApiDirectoryV1DirectoryServiceListResponse**](C1ApiDirectoryV1DirectoryServiceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

