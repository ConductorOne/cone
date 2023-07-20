# RequestCatalogSearch

### Available Operations

* [SearchEntitlements](#searchentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.

## SearchEntitlements

Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.

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
    res, err := s.RequestCatalogSearch.SearchEntitlements(ctx, shared.RequestCatalogSearchServiceSearchEntitlementsRequest{
        AppEntitlementExpandMask: &shared.AppEntitlementExpandMask{
            Paths: []string{
                "minus",
            },
        },
        EntitlementAlias: conductoroneapi.String("fuga"),
        GrantedStatus: shared.RequestCatalogSearchServiceSearchEntitlementsRequestGrantedStatusAll.ToPointer(),
        PageSize: conductoroneapi.Float64(6692.37),
        PageToken: conductoroneapi.String("impedit"),
        Query: conductoroneapi.String("delectus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogSearchServiceSearchEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [shared.RequestCatalogSearchServiceSearchEntitlementsRequest](../../models/shared/requestcatalogsearchservicesearchentitlementsrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogsearchservicesearchentitlementsresponse.md), error**

