# AppEntitlementSearch

### Available Operations

* [Search](#search) - Invokes the c1.api.app.v1.AppEntitlementSearchService.Search method.

## Search

Invokes the c1.api.app.v1.AppEntitlementSearchService.Search method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppEntitlementSearch.Search(ctx, shared.AppEntitlementSearchServiceSearchRequest{
        AppEntitlementExpandMask: &shared.AppEntitlementExpandMask{
            Paths: []string{
                "sapiente",
                "quo",
                "odit",
                "at",
            },
        },
        AccessReviewID: conductoroneapi.String("at"),
        Alias: conductoroneapi.String("maiores"),
        AppIds: []string{
            "quod",
            "quod",
        },
        AppUserIds: []string{
            "totam",
            "porro",
        },
        ComplianceFrameworkIds: []string{
            "dicta",
            "nam",
            "officia",
        },
        ExcludeAppIds: []string{
            "fugit",
            "deleniti",
            "hic",
        },
        ExcludeAppUserIds: []string{
            "totam",
            "beatae",
            "commodi",
            "molestiae",
        },
        OnlyGetExpiring: conductoroneapi.Bool(false),
        PageSize: conductoroneapi.Float64(2645.55),
        PageToken: conductoroneapi.String("qui"),
        Query: conductoroneapi.String("impedit"),
        ResourceTypeIds: []string{
            "esse",
            "ipsum",
            "excepturi",
        },
        RiskLevelIds: []string{
            "perferendis",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppEntitlementSearchServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [shared.AppEntitlementSearchServiceSearchRequest](../../models/shared/appentitlementsearchservicesearchrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchResponse](../../models/operations/c1apiappv1appentitlementsearchservicesearchresponse.md), error**

