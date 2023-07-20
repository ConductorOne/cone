<!-- Start SDK Example Usage -->


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
                "provident",
                "distinctio",
                "quibusdam",
            },
        },
        AccessReviewID: conductoroneapi.String("unde"),
        Alias: conductoroneapi.String("nulla"),
        AppIds: []string{
            "illum",
            "vel",
            "error",
        },
        AppUserIds: []string{
            "suscipit",
            "iure",
            "magnam",
        },
        ComplianceFrameworkIds: []string{
            "ipsa",
            "delectus",
            "tempora",
            "suscipit",
        },
        ExcludeAppIds: []string{
            "minus",
            "placeat",
        },
        ExcludeAppUserIds: []string{
            "iusto",
            "excepturi",
            "nisi",
        },
        OnlyGetExpiring: conductoroneapi.Bool(false),
        PageSize: conductoroneapi.Float64(9255.97),
        PageToken: conductoroneapi.String("temporibus"),
        Query: conductoroneapi.String("ab"),
        ResourceTypeIds: []string{
            "veritatis",
            "deserunt",
        },
        RiskLevelIds: []string{
            "ipsam",
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
<!-- End SDK Example Usage -->