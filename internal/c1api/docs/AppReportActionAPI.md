# \AppReportActionAPI

All URIs are relative to *https://invalid-example.conductor.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**C1ApiAppV1AppReportActionServiceGenerateReport**](AppReportActionAPI.md#C1ApiAppV1AppReportActionServiceGenerateReport) | **Post** /api/v1/apps/{app_id}/report | 



## C1ApiAppV1AppReportActionServiceGenerateReport

> map[string]interface{} C1ApiAppV1AppReportActionServiceGenerateReport(ctx, appId).Body(body).Execute()





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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppReportActionAPI.C1ApiAppV1AppReportActionServiceGenerateReport(context.Background(), appId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppReportActionAPI.C1ApiAppV1AppReportActionServiceGenerateReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `C1ApiAppV1AppReportActionServiceGenerateReport`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AppReportActionAPI.C1ApiAppV1AppReportActionServiceGenerateReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiC1ApiAppV1AppReportActionServiceGenerateReportRequest struct via the builder pattern


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

