# \UserSearchAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiUserV1UserSearchSearch**](UserSearchAPI.md#C1ApiUserV1UserSearchSearch) | **Post** /api/v1/search/users | 



## C1ApiUserV1UserSearchSearch

> C1ApiUserV1SearchUsersResponse C1ApiUserV1UserSearchSearch(ctx).C1ApiUserV1SearchUsersRequest(c1ApiUserV1SearchUsersRequest).Execute()





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
    c1ApiUserV1SearchUsersRequest := *openapiclient.NewC1ApiUserV1SearchUsersRequest() // C1ApiUserV1SearchUsersRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSearchAPI.C1ApiUserV1UserSearchSearch(context.Background()).C1ApiUserV1SearchUsersRequest(c1ApiUserV1SearchUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSearchAPI.C1ApiUserV1UserSearchSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiUserV1UserSearchSearch`: C1ApiUserV1SearchUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `UserSearchAPI.C1ApiUserV1UserSearchSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiUserV1UserSearchSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **c1ApiUserV1SearchUsersRequest** | [**C1ApiUserV1SearchUsersRequest**](C1ApiUserV1SearchUsersRequest.md) |  | 

### Return type

[**C1ApiUserV1SearchUsersResponse**](C1ApiUserV1SearchUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

