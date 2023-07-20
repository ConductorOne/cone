# UserSearch

### Available Operations

* [Search](#search) - Invokes the c1.api.user.v1.UserSearch.Search method.

## Search

Invokes the c1.api.user.v1.UserSearch.Search method.

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
    res, err := s.UserSearch.Search(ctx, shared.SearchUsersRequest{
        UserExpandMask: &shared.UserExpandMask{
            Paths: []string{
                "deleniti",
            },
        },
        Email: conductoroneapi.String("Jabari72@gmail.com"),
        ExcludeIds: []string{
            "sit",
            "modi",
        },
        Ids: []string{
            "nesciunt",
            "mollitia",
        },
        PageSize: conductoroneapi.Float64(4926.32),
        PageToken: conductoroneapi.String("fugiat"),
        Query: conductoroneapi.String("nostrum"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("5f1400e7-64ad-4733-8ec1-b781b36a0808"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("8d100efa-da20-40ef-8422-eb2164cf9ab8"),
            },
        },
        RoleIds: []string{
            "aliquid",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesDeleted,
            shared.SearchUsersRequestUserStatusesEnabled,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchUsersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.SearchUsersRequest](../../models/shared/searchusersrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.C1APIUserV1UserSearchSearchResponse](../../models/operations/c1apiuserv1usersearchsearchresponse.md), error**

