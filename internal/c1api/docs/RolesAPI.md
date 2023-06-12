# \RolesAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiIamV1RolesGet**](RolesAPI.md#C1ApiIamV1RolesGet) | **Get** /api/v1/iam/roles/{role_id} | 
[**C1ApiIamV1RolesList**](RolesAPI.md#C1ApiIamV1RolesList) | **Get** /api/v1/iam/roles | 



## C1ApiIamV1RolesGet

> C1ApiIamV1Role C1ApiIamV1RolesGet(ctx, roleId).Execute()





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
    roleId := "roleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesAPI.C1ApiIamV1RolesGet(context.Background(), roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.C1ApiIamV1RolesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiIamV1RolesGet`: C1ApiIamV1Role
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.C1ApiIamV1RolesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiIamV1RolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**C1ApiIamV1Role**](C1ApiIamV1Role.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## C1ApiIamV1RolesList

> C1ApiIamV1ListRolesResponse C1ApiIamV1RolesList(ctx).Execute()





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
    resp, r, err := apiClient.RolesAPI.C1ApiIamV1RolesList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.C1ApiIamV1RolesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiIamV1RolesList`: C1ApiIamV1ListRolesResponse
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.C1ApiIamV1RolesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiIamV1RolesListRequest struct via the builder pattern


### Return type

[**C1ApiIamV1ListRolesResponse**](C1ApiIamV1ListRolesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

