# \UserAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiUserV1UserServiceGet**](UserAPI.md#C1ApiUserV1UserServiceGet) | **Get** /api/v1/users/{id} | 
[**C1ApiUserV1UserServiceList**](UserAPI.md#C1ApiUserV1UserServiceList) | **Get** /api/v1/users | 



## C1ApiUserV1UserServiceGet

> C1ApiUserV1UserServiceGetResponse C1ApiUserV1UserServiceGet(ctx, id).Execute()





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
    resp, r, err := apiClient.UserAPI.C1ApiUserV1UserServiceGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.C1ApiUserV1UserServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiUserV1UserServiceGet`: C1ApiUserV1UserServiceGetResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.C1ApiUserV1UserServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiUserV1UserServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**C1ApiUserV1UserServiceGetResponse**](C1ApiUserV1UserServiceGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## C1ApiUserV1UserServiceList

> C1ApiUserV1UserServiceListResponse C1ApiUserV1UserServiceList(ctx).Execute()





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
    resp, r, err := apiClient.UserAPI.C1ApiUserV1UserServiceList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.C1ApiUserV1UserServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiUserV1UserServiceList`: C1ApiUserV1UserServiceListResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.C1ApiUserV1UserServiceList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiUserV1UserServiceListRequest struct via the builder pattern


### Return type

[**C1ApiUserV1UserServiceListResponse**](C1ApiUserV1UserServiceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

