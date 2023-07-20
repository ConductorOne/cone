# AppReportAction

### Available Operations

* [GenerateReport](#generatereport) - Invokes the c1.api.app.v1.AppReportActionService.GenerateReport method.

## GenerateReport

Invokes the c1.api.app.v1.AppReportActionService.GenerateReport method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppReportAction.GenerateReport(ctx, operations.C1APIAppV1AppReportActionServiceGenerateReportRequest{
        AppActionsServiceGenerateReportRequest: &shared.AppActionsServiceGenerateReportRequest{},
        AppID: "aliquid",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppActionsServiceGenerateReportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.C1APIAppV1AppReportActionServiceGenerateReportRequest](../../models/operations/c1apiappv1appreportactionservicegeneratereportrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.C1APIAppV1AppReportActionServiceGenerateReportResponse](../../models/operations/c1apiappv1appreportactionservicegeneratereportresponse.md), error**

