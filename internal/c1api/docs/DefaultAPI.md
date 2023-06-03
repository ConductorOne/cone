# \DefaultAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiAppV1AppResourceServiceGet**](DefaultAPI.md#C1ApiAppV1AppResourceServiceGet) | **Get** /api/v1/apps/{app_id}/resource_types/{app_resource_type_id}/resource/{id} | 
[**C1ApiAppV1AppResourceTypeServiceGet**](DefaultAPI.md#C1ApiAppV1AppResourceTypeServiceGet) | **Get** /api/v1/apps/{app_id}/resource_types/{id} | 
[**C1ApiAppV1AppsGet**](DefaultAPI.md#C1ApiAppV1AppsGet) | **Get** /api/v1/apps/{id} | 
[**C1ApiAuthV1AuthIntrospect**](DefaultAPI.md#C1ApiAuthV1AuthIntrospect) | **Get** /api/v1/auth/introspect | 
[**C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements**](DefaultAPI.md#C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements) | **Post** /api/v1/search/request_catalog/entitlements | 
[**C1ApiTaskV1TaskServiceCreateGrantTask**](DefaultAPI.md#C1ApiTaskV1TaskServiceCreateGrantTask) | **Post** /api/v1/task/grant | 
[**C1ApiTaskV1TaskServiceCreateRevokeTask**](DefaultAPI.md#C1ApiTaskV1TaskServiceCreateRevokeTask) | **Post** /api/v1/task/revoke | 
[**C1ApiTaskV1TaskServiceGet**](DefaultAPI.md#C1ApiTaskV1TaskServiceGet) | **Get** /api/v1/tasks/{id} | 
[**C1ApiUserV1UserServiceGet**](DefaultAPI.md#C1ApiUserV1UserServiceGet) | **Get** /api/v1/users/{id} | 



## C1ApiAppV1AppResourceServiceGet

> C1ApiAppV1AppResourceServiceGetResponse C1ApiAppV1AppResourceServiceGet(ctx, appId, appResourceTypeId, id).Execute()





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
    appResourceTypeId := "appResourceTypeId_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.C1ApiAppV1AppResourceServiceGet(context.Background(), appId, appResourceTypeId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.C1ApiAppV1AppResourceServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAppV1AppResourceServiceGet`: C1ApiAppV1AppResourceServiceGetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.C1ApiAppV1AppResourceServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**appResourceTypeId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiAppV1AppResourceServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**C1ApiAppV1AppResourceServiceGetResponse**](C1ApiAppV1AppResourceServiceGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## C1ApiAppV1AppResourceTypeServiceGet

> C1ApiAppV1AppResourceTypeServiceGetResponse C1ApiAppV1AppResourceTypeServiceGet(ctx, appId, id).Execute()





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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.C1ApiAppV1AppResourceTypeServiceGet(context.Background(), appId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.C1ApiAppV1AppResourceTypeServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAppV1AppResourceTypeServiceGet`: C1ApiAppV1AppResourceTypeServiceGetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.C1ApiAppV1AppResourceTypeServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiAppV1AppResourceTypeServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**C1ApiAppV1AppResourceTypeServiceGetResponse**](C1ApiAppV1AppResourceTypeServiceGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.DefaultAPI.C1ApiAppV1AppsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.C1ApiAppV1AppsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAppV1AppsGet`: C1ApiAppV1GetAppResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.C1ApiAppV1AppsGet`: %v\n", resp)
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


## C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements

> C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements(ctx).C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest(c1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest).Execute()





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
    c1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest := *openapiclient.NewC1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest() // C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements(context.Background()).C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest(c1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements`: C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **c1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest** | [**C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest**](C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest.md) |  | 

### Return type

[**C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse**](C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.DefaultAPI.C1ApiTaskV1TaskServiceCreateGrantTask(context.Background()).C1ApiTaskV1TaskServiceCreateGrantRequest(c1ApiTaskV1TaskServiceCreateGrantRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.C1ApiTaskV1TaskServiceCreateGrantTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiTaskV1TaskServiceCreateGrantTask`: C1ApiTaskV1TaskServiceCreateGrantResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.C1ApiTaskV1TaskServiceCreateGrantTask`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultAPI.C1ApiTaskV1TaskServiceCreateRevokeTask(context.Background()).C1ApiTaskV1TaskServiceCreateRevokeRequest(c1ApiTaskV1TaskServiceCreateRevokeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.C1ApiTaskV1TaskServiceCreateRevokeTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiTaskV1TaskServiceCreateRevokeTask`: C1ApiTaskV1TaskServiceCreateRevokeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.C1ApiTaskV1TaskServiceCreateRevokeTask`: %v\n", resp)
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
    resp, r, err := apiClient.DefaultAPI.C1ApiTaskV1TaskServiceGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.C1ApiTaskV1TaskServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiTaskV1TaskServiceGet`: C1ApiTaskV1TaskServiceGetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.C1ApiTaskV1TaskServiceGet`: %v\n", resp)
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

