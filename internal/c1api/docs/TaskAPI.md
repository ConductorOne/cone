# \TaskAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiTaskV1TaskServiceCreateGrantTask**](TaskAPI.md#C1ApiTaskV1TaskServiceCreateGrantTask) | **Post** /api/v1/task/grant | 
[**C1ApiTaskV1TaskServiceCreateRevokeTask**](TaskAPI.md#C1ApiTaskV1TaskServiceCreateRevokeTask) | **Post** /api/v1/task/revoke | 
[**C1ApiTaskV1TaskServiceGet**](TaskAPI.md#C1ApiTaskV1TaskServiceGet) | **Get** /api/v1/tasks/{id} | 



## C1ApiTaskV1TaskServiceCreateGrantTask

> C1ApiTaskV1TaskServiceCreateGrantResponse C1ApiTaskV1TaskServiceCreateGrantTask(ctx).C1ApiTaskV1TaskServiceCreateGrantRequest(c1ApiTaskV1TaskServiceCreateGrantRequest).Execute()





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
    c1ApiTaskV1TaskServiceCreateGrantRequest := *openapiclient.NewC1ApiTaskV1TaskServiceCreateGrantRequest() // C1ApiTaskV1TaskServiceCreateGrantRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskAPI.C1ApiTaskV1TaskServiceCreateGrantTask(context.Background()).C1ApiTaskV1TaskServiceCreateGrantRequest(c1ApiTaskV1TaskServiceCreateGrantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.C1ApiTaskV1TaskServiceCreateGrantTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiTaskV1TaskServiceCreateGrantTask`: C1ApiTaskV1TaskServiceCreateGrantResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskAPI.C1ApiTaskV1TaskServiceCreateGrantTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiTaskV1TaskServiceCreateGrantTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **c1ApiTaskV1TaskServiceCreateGrantRequest** | [**C1ApiTaskV1TaskServiceCreateGrantRequest**](C1ApiTaskV1TaskServiceCreateGrantRequest.md) |  | 

### Return type

[**C1ApiTaskV1TaskServiceCreateGrantResponse**](C1ApiTaskV1TaskServiceCreateGrantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## C1ApiTaskV1TaskServiceCreateRevokeTask

> C1ApiTaskV1TaskServiceCreateRevokeResponse C1ApiTaskV1TaskServiceCreateRevokeTask(ctx).C1ApiTaskV1TaskServiceCreateRevokeRequest(c1ApiTaskV1TaskServiceCreateRevokeRequest).Execute()





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
    c1ApiTaskV1TaskServiceCreateRevokeRequest := *openapiclient.NewC1ApiTaskV1TaskServiceCreateRevokeRequest() // C1ApiTaskV1TaskServiceCreateRevokeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskAPI.C1ApiTaskV1TaskServiceCreateRevokeTask(context.Background()).C1ApiTaskV1TaskServiceCreateRevokeRequest(c1ApiTaskV1TaskServiceCreateRevokeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.C1ApiTaskV1TaskServiceCreateRevokeTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiTaskV1TaskServiceCreateRevokeTask`: C1ApiTaskV1TaskServiceCreateRevokeResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskAPI.C1ApiTaskV1TaskServiceCreateRevokeTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiTaskV1TaskServiceCreateRevokeTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **c1ApiTaskV1TaskServiceCreateRevokeRequest** | [**C1ApiTaskV1TaskServiceCreateRevokeRequest**](C1ApiTaskV1TaskServiceCreateRevokeRequest.md) |  | 

### Return type

[**C1ApiTaskV1TaskServiceCreateRevokeResponse**](C1ApiTaskV1TaskServiceCreateRevokeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## C1ApiTaskV1TaskServiceGet

> C1ApiTaskV1TaskServiceGetResponse C1ApiTaskV1TaskServiceGet(ctx, id).Execute()





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
    resp, r, err := apiClient.TaskAPI.C1ApiTaskV1TaskServiceGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.C1ApiTaskV1TaskServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiTaskV1TaskServiceGet`: C1ApiTaskV1TaskServiceGetResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskAPI.C1ApiTaskV1TaskServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiTaskV1TaskServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**C1ApiTaskV1TaskServiceGetResponse**](C1ApiTaskV1TaskServiceGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

