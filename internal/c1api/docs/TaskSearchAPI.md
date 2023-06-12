# \TaskSearchAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiTaskV1TaskSearchServiceSearch**](TaskSearchAPI.md#C1ApiTaskV1TaskSearchServiceSearch) | **Post** /api/v1/search/tasks | 



## C1ApiTaskV1TaskSearchServiceSearch

> C1ApiTaskV1TaskSearchResponse C1ApiTaskV1TaskSearchServiceSearch(ctx).C1ApiTaskV1TaskSearchRequest(c1ApiTaskV1TaskSearchRequest).Execute()





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
    c1ApiTaskV1TaskSearchRequest := *openapiclient.NewC1ApiTaskV1TaskSearchRequest() // C1ApiTaskV1TaskSearchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskSearchAPI.C1ApiTaskV1TaskSearchServiceSearch(context.Background()).C1ApiTaskV1TaskSearchRequest(c1ApiTaskV1TaskSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskSearchAPI.C1ApiTaskV1TaskSearchServiceSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiTaskV1TaskSearchServiceSearch`: C1ApiTaskV1TaskSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskSearchAPI.C1ApiTaskV1TaskSearchServiceSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiTaskV1TaskSearchServiceSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **c1ApiTaskV1TaskSearchRequest** | [**C1ApiTaskV1TaskSearchRequest**](C1ApiTaskV1TaskSearchRequest.md) |  | 

### Return type

[**C1ApiTaskV1TaskSearchResponse**](C1ApiTaskV1TaskSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

