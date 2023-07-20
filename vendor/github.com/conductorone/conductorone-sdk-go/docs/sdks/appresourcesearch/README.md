# AppResourceSearch

### Available Operations

* [SearchAppResourceTypes](#searchappresourcetypes) - Invokes the c1.api.app.v1.AppResourceSearch.SearchAppResourceTypes method.

## SearchAppResourceTypes

Invokes the c1.api.app.v1.AppResourceSearch.SearchAppResourceTypes method.

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
    res, err := s.AppResourceSearch.SearchAppResourceTypes(ctx, shared.SearchAppResourceTypesRequest{
        AppIds: []string{
            "molestiae",
            "perferendis",
            "nihil",
        },
        ExcludeResourceTypeIds: []string{
            "distinctio",
            "id",
        },
        ExcludeResourceTypeTraitIds: []string{
            "labore",
            "suscipit",
        },
        PageSize: conductoroneapi.Float64(6180.16),
        PageToken: conductoroneapi.String("nobis"),
        Query: conductoroneapi.String("eum"),
        ResourceTypeIds: []string{
            "aspernatur",
            "architecto",
            "magnam",
            "et",
        },
        ResourceTypeTraitIds: []string{
            "ullam",
            "provident",
            "quos",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchAppResourceTypesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.SearchAppResourceTypesRequest](../../models/shared/searchappresourcetypesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.C1APIAppV1AppResourceSearchSearchAppResourceTypesResponse](../../models/operations/c1apiappv1appresourcesearchsearchappresourcetypesresponse.md), error**

