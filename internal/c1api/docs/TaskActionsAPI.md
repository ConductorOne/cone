# \TaskActionsAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiTaskV1TaskActionsServiceApprove**](TaskActionsAPI.md#C1ApiTaskV1TaskActionsServiceApprove) | **Post** /api/v1/tasks/{task_id}/action/approve | 
[**C1ApiTaskV1TaskActionsServiceComment**](TaskActionsAPI.md#C1ApiTaskV1TaskActionsServiceComment) | **Post** /api/v1/tasks/{task_id}/action/comment | 
[**C1ApiTaskV1TaskActionsServiceDeny**](TaskActionsAPI.md#C1ApiTaskV1TaskActionsServiceDeny) | **Post** /api/v1/tasks/{task_id}/action/deny | 



## C1ApiTaskV1TaskActionsServiceApprove

> C1ApiTaskV1TaskActionsServiceApproveResponse C1ApiTaskV1TaskActionsServiceApprove(ctx, taskId).C1ApiTaskV1TaskActionsServiceApproveRequestInput(c1ApiTaskV1TaskActionsServiceApproveRequestInput).Execute()





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
    taskId := "taskId_example" // string | 
    c1ApiTaskV1TaskActionsServiceApproveRequestInput := *openapiclient.NewC1ApiTaskV1TaskActionsServiceApproveRequestInput() // C1ApiTaskV1TaskActionsServiceApproveRequestInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskActionsAPI.C1ApiTaskV1TaskActionsServiceApprove(context.Background(), taskId).C1ApiTaskV1TaskActionsServiceApproveRequestInput(c1ApiTaskV1TaskActionsServiceApproveRequestInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskActionsAPI.C1ApiTaskV1TaskActionsServiceApprove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiTaskV1TaskActionsServiceApprove`: C1ApiTaskV1TaskActionsServiceApproveResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskActionsAPI.C1ApiTaskV1TaskActionsServiceApprove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiTaskV1TaskActionsServiceApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **c1ApiTaskV1TaskActionsServiceApproveRequestInput** | [**C1ApiTaskV1TaskActionsServiceApproveRequestInput**](C1ApiTaskV1TaskActionsServiceApproveRequestInput.md) |  | 

### Return type

[**C1ApiTaskV1TaskActionsServiceApproveResponse**](C1ApiTaskV1TaskActionsServiceApproveResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## C1ApiTaskV1TaskActionsServiceComment

> C1ApiTaskV1TaskActionsServiceCommentResponse C1ApiTaskV1TaskActionsServiceComment(ctx, taskId).C1ApiTaskV1TaskActionsServiceCommentRequestInput(c1ApiTaskV1TaskActionsServiceCommentRequestInput).Execute()





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
    taskId := "taskId_example" // string | 
    c1ApiTaskV1TaskActionsServiceCommentRequestInput := *openapiclient.NewC1ApiTaskV1TaskActionsServiceCommentRequestInput() // C1ApiTaskV1TaskActionsServiceCommentRequestInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskActionsAPI.C1ApiTaskV1TaskActionsServiceComment(context.Background(), taskId).C1ApiTaskV1TaskActionsServiceCommentRequestInput(c1ApiTaskV1TaskActionsServiceCommentRequestInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskActionsAPI.C1ApiTaskV1TaskActionsServiceComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiTaskV1TaskActionsServiceComment`: C1ApiTaskV1TaskActionsServiceCommentResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskActionsAPI.C1ApiTaskV1TaskActionsServiceComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiTaskV1TaskActionsServiceCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **c1ApiTaskV1TaskActionsServiceCommentRequestInput** | [**C1ApiTaskV1TaskActionsServiceCommentRequestInput**](C1ApiTaskV1TaskActionsServiceCommentRequestInput.md) |  | 

### Return type

[**C1ApiTaskV1TaskActionsServiceCommentResponse**](C1ApiTaskV1TaskActionsServiceCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## C1ApiTaskV1TaskActionsServiceDeny

> C1ApiTaskV1TaskActionsServiceDenyResponse C1ApiTaskV1TaskActionsServiceDeny(ctx, taskId).C1ApiTaskV1TaskActionsServiceDenyRequestInput(c1ApiTaskV1TaskActionsServiceDenyRequestInput).Execute()





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
    taskId := "taskId_example" // string | 
    c1ApiTaskV1TaskActionsServiceDenyRequestInput := *openapiclient.NewC1ApiTaskV1TaskActionsServiceDenyRequestInput() // C1ApiTaskV1TaskActionsServiceDenyRequestInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskActionsAPI.C1ApiTaskV1TaskActionsServiceDeny(context.Background(), taskId).C1ApiTaskV1TaskActionsServiceDenyRequestInput(c1ApiTaskV1TaskActionsServiceDenyRequestInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskActionsAPI.C1ApiTaskV1TaskActionsServiceDeny``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiTaskV1TaskActionsServiceDeny`: C1ApiTaskV1TaskActionsServiceDenyResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskActionsAPI.C1ApiTaskV1TaskActionsServiceDeny`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiTaskV1TaskActionsServiceDenyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **c1ApiTaskV1TaskActionsServiceDenyRequestInput** | [**C1ApiTaskV1TaskActionsServiceDenyRequestInput**](C1ApiTaskV1TaskActionsServiceDenyRequestInput.md) |  | 

### Return type

[**C1ApiTaskV1TaskActionsServiceDenyResponse**](C1ApiTaskV1TaskActionsServiceDenyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

