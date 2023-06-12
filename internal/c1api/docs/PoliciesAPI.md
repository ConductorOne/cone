# \PoliciesAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiPolicyV1PoliciesCreate**](PoliciesAPI.md#C1ApiPolicyV1PoliciesCreate) | **Post** /api/v1/policies | 
[**C1ApiPolicyV1PoliciesDelete**](PoliciesAPI.md#C1ApiPolicyV1PoliciesDelete) | **Delete** /api/v1/policies/{id} | 
[**C1ApiPolicyV1PoliciesGet**](PoliciesAPI.md#C1ApiPolicyV1PoliciesGet) | **Get** /api/v1/policies/{id} | 
[**C1ApiPolicyV1PoliciesList**](PoliciesAPI.md#C1ApiPolicyV1PoliciesList) | **Get** /api/v1/policies | 



## C1ApiPolicyV1PoliciesCreate

> C1ApiPolicyV1Policy C1ApiPolicyV1PoliciesCreate(ctx).C1ApiPolicyV1CreatePolicyRequest(c1ApiPolicyV1CreatePolicyRequest).Execute()





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
    c1ApiPolicyV1CreatePolicyRequest := *openapiclient.NewC1ApiPolicyV1CreatePolicyRequest() // C1ApiPolicyV1CreatePolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesAPI.C1ApiPolicyV1PoliciesCreate(context.Background()).C1ApiPolicyV1CreatePolicyRequest(c1ApiPolicyV1CreatePolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.C1ApiPolicyV1PoliciesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiPolicyV1PoliciesCreate`: C1ApiPolicyV1Policy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.C1ApiPolicyV1PoliciesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiPolicyV1PoliciesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **c1ApiPolicyV1CreatePolicyRequest** | [**C1ApiPolicyV1CreatePolicyRequest**](C1ApiPolicyV1CreatePolicyRequest.md) |  | 

### Return type

[**C1ApiPolicyV1Policy**](C1ApiPolicyV1Policy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## C1ApiPolicyV1PoliciesDelete

> map[string]interface{} C1ApiPolicyV1PoliciesDelete(ctx, id).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesAPI.C1ApiPolicyV1PoliciesDelete(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.C1ApiPolicyV1PoliciesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiPolicyV1PoliciesDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.C1ApiPolicyV1PoliciesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiPolicyV1PoliciesDeleteRequest struct via the builder pattern


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


## C1ApiPolicyV1PoliciesGet

> C1ApiPolicyV1Policy C1ApiPolicyV1PoliciesGet(ctx, id).Execute()





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
    resp, r, err := apiClient.PoliciesAPI.C1ApiPolicyV1PoliciesGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.C1ApiPolicyV1PoliciesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiPolicyV1PoliciesGet`: C1ApiPolicyV1Policy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.C1ApiPolicyV1PoliciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiPolicyV1PoliciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**C1ApiPolicyV1Policy**](C1ApiPolicyV1Policy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## C1ApiPolicyV1PoliciesList

> C1ApiPolicyV1ListPolicyResponse C1ApiPolicyV1PoliciesList(ctx).Execute()





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
    resp, r, err := apiClient.PoliciesAPI.C1ApiPolicyV1PoliciesList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.C1ApiPolicyV1PoliciesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiPolicyV1PoliciesList`: C1ApiPolicyV1ListPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.C1ApiPolicyV1PoliciesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiPolicyV1PoliciesListRequest struct via the builder pattern


### Return type

[**C1ApiPolicyV1ListPolicyResponse**](C1ApiPolicyV1ListPolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

