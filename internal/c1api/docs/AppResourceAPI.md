# \AppResourceAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiAppV1AppResourceServiceGet**](AppResourceAPI.md#C1ApiAppV1AppResourceServiceGet) | **Get** /api/v1/apps/{app_id}/resource_types/{app_resource_type_id}/resource/{id} | 
[**C1ApiAppV1AppResourceServiceList**](AppResourceAPI.md#C1ApiAppV1AppResourceServiceList) | **Get** /api/v1/apps/{app_id}/resource_types/{app_resource_type_id}/resource | 



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
    resp, r, err := apiClient.AppResourceAPI.C1ApiAppV1AppResourceServiceGet(context.Background(), appId, appResourceTypeId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppResourceAPI.C1ApiAppV1AppResourceServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAppV1AppResourceServiceGet`: C1ApiAppV1AppResourceServiceGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AppResourceAPI.C1ApiAppV1AppResourceServiceGet`: %v\n", resp)
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


## C1ApiAppV1AppResourceServiceList

> C1ApiAppV1AppResourceServiceListResponse C1ApiAppV1AppResourceServiceList(ctx, appId, appResourceTypeId).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppResourceAPI.C1ApiAppV1AppResourceServiceList(context.Background(), appId, appResourceTypeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppResourceAPI.C1ApiAppV1AppResourceServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAppV1AppResourceServiceList`: C1ApiAppV1AppResourceServiceListResponse
    fmt.Fprintf(os.Stdout, "Response from `AppResourceAPI.C1ApiAppV1AppResourceServiceList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**appResourceTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiAppV1AppResourceServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**C1ApiAppV1AppResourceServiceListResponse**](C1ApiAppV1AppResourceServiceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

