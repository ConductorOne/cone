# AppUsageControls

### Available Operations

* [Get](#get) - Invokes the c1.api.app.v1.AppUsageControlsService.Get method.
* [Update](#update) - Invokes the c1.api.app.v1.AppUsageControlsService.Update method.

## Get

Invokes the c1.api.app.v1.AppUsageControlsService.Get method.

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
    res, err := s.AppUsageControls.Get(ctx, operations.C1APIAppV1AppUsageControlsServiceGetRequest{
        AppID: "ab",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAppUsageControlsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIAppV1AppUsageControlsServiceGetRequest](../../models/operations/c1apiappv1appusagecontrolsservicegetrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.C1APIAppV1AppUsageControlsServiceGetResponse](../../models/operations/c1apiappv1appusagecontrolsservicegetresponse.md), error**


## Update

Invokes the c1.api.app.v1.AppUsageControlsService.Update method.

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
    res, err := s.AppUsageControls.Update(ctx, operations.C1APIAppV1AppUsageControlsServiceUpdateRequest{
        UpdateAppUsageControlsRequest: &shared.UpdateAppUsageControlsRequest{
            AppUsageControls: &shared.AppUsageControls{
                AppID: conductoroneapi.String("maiores"),
                Notify: conductoroneapi.Bool(false),
                NotifyAfterDays: conductoroneapi.Float64(6974.29),
                Revoke: conductoroneapi.Bool(false),
                RevokeAfterDays: conductoroneapi.Float64(3732.91),
            },
            UpdateMask: conductoroneapi.String("voluptate"),
        },
        AppID: "autem",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateAppUsageControlsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.C1APIAppV1AppUsageControlsServiceUpdateRequest](../../models/operations/c1apiappv1appusagecontrolsserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.C1APIAppV1AppUsageControlsServiceUpdateResponse](../../models/operations/c1apiappv1appusagecontrolsserviceupdateresponse.md), error**

