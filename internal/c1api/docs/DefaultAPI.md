# \DefaultAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiAuthV1AuthIntrospect**](DefaultAPI.md#C1ApiAuthV1AuthIntrospect) | **Get** /api/v1/auth/introspect | 
[**C1ApiUserV1UserServiceGet**](DefaultAPI.md#C1ApiUserV1UserServiceGet) | **Get** /api/v1/user/get/{id} | 



## C1ApiAuthV1AuthIntrospect

> C1ApiAuthV1IntrospectResponse C1ApiAuthV1AuthIntrospect(ctx).Execute()





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
    resp, r, err := apiClient.DefaultAPI.C1ApiAuthV1AuthIntrospect(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.C1ApiAuthV1AuthIntrospect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAuthV1AuthIntrospect`: C1ApiAuthV1IntrospectResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.C1ApiAuthV1AuthIntrospect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiAuthV1AuthIntrospectRequest struct via the builder pattern


### Return type

[**C1ApiAuthV1IntrospectResponse**](C1ApiAuthV1IntrospectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    id := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.C1ApiUserV1UserServiceGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.C1ApiUserV1UserServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiUserV1UserServiceGet`: C1ApiUserV1UserServiceGetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.C1ApiUserV1UserServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

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

