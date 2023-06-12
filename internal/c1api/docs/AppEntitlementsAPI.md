# \AppEntitlementsAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiAppV1AppEntitlementsGet**](AppEntitlementsAPI.md#C1ApiAppV1AppEntitlementsGet) | **Get** /api/v1/apps/{app_id}/entitlements/{id} | 



## C1ApiAppV1AppEntitlementsGet

> C1ApiAppV1GetAppEntitlementResponse C1ApiAppV1AppEntitlementsGet(ctx, appId, id).Execute()





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
    resp, r, err := apiClient.AppEntitlementsAPI.C1ApiAppV1AppEntitlementsGet(context.Background(), appId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppEntitlementsAPI.C1ApiAppV1AppEntitlementsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAppV1AppEntitlementsGet`: C1ApiAppV1GetAppEntitlementResponse
    fmt.Fprintf(os.Stdout, "Response from `AppEntitlementsAPI.C1ApiAppV1AppEntitlementsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiAppV1AppEntitlementsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**C1ApiAppV1GetAppEntitlementResponse**](C1ApiAppV1GetAppEntitlementResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

