# \AppEntitlementUserBindingAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant**](AppEntitlementUserBindingAPI.md#C1ApiAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant) | **Get** /api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/users/{identity_user_id}/grants | 



## C1ApiAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant

> C1ApiAppV1ListAppUsersForIdentityWithGrantResponse C1ApiAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant(ctx, appId, appEntitlementId, identityUserId).Execute()





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
    appEntitlementId := "appEntitlementId_example" // string | 
    identityUserId := "identityUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppEntitlementUserBindingAPI.C1ApiAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant(context.Background(), appId, appEntitlementId, identityUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppEntitlementUserBindingAPI.C1ApiAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant`: C1ApiAppV1ListAppUsersForIdentityWithGrantResponse
    fmt.Fprintf(os.Stdout, "Response from `AppEntitlementUserBindingAPI.C1ApiAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**appEntitlementId** | **string** |  | 
**identityUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**C1ApiAppV1ListAppUsersForIdentityWithGrantResponse**](C1ApiAppV1ListAppUsersForIdentityWithGrantResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

