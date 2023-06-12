# \AuthAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiAuthV1AuthIntrospect**](AuthAPI.md#C1ApiAuthV1AuthIntrospect) | **Get** /api/v1/auth/introspect | 



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
    resp, r, err := apiClient.AuthAPI.C1ApiAuthV1AuthIntrospect(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.C1ApiAuthV1AuthIntrospect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAuthV1AuthIntrospect`: C1ApiAuthV1IntrospectResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthAPI.C1ApiAuthV1AuthIntrospect`: %v\n", resp)
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

