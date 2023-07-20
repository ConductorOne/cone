# Roles

### Available Operations

* [Get](#get) - Invokes the c1.api.iam.v1.Roles.Get method.
* [List](#list) - Invokes the c1.api.iam.v1.Roles.List method.
* [Update](#update) - Invokes the c1.api.iam.v1.Roles.Update method.

## Get

Invokes the c1.api.iam.v1.Roles.Get method.

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
    res, err := s.Roles.Get(ctx, operations.C1APIIamV1RolesGetRequest{
        RoleID: "tempore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRolesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.C1APIIamV1RolesGetRequest](../../models/operations/c1apiiamv1rolesgetrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.C1APIIamV1RolesGetResponse](../../models/operations/c1apiiamv1rolesgetresponse.md), error**


## List

Invokes the c1.api.iam.v1.Roles.List method.

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
    res, err := s.Roles.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListRolesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.C1APIIamV1RolesListResponse](../../models/operations/c1apiiamv1roleslistresponse.md), error**


## Update

Invokes the c1.api.iam.v1.Roles.Update method.

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
    res, err := s.Roles.Update(ctx, operations.C1APIIamV1RolesUpdateRequest{
        UpdateRoleRequestInput: &shared.UpdateRoleRequestInput{
            Role: &shared.RoleInput{
                DisplayName: conductoroneapi.String("vero"),
                ID: conductoroneapi.String("2fd57075-7792-4917-bdea-c646ecb57340"),
                Name: conductoroneapi.String("Ignacio Dietrich"),
                Permissions: []string{
                    "earum",
                },
                ServiceRoles: []string{
                    "animi",
                    "dolores",
                },
                SystemBuiltin: conductoroneapi.Bool(false),
            },
            UpdateMask: conductoroneapi.String("nam"),
        },
        RoleID: "dicta",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateRolesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.C1APIIamV1RolesUpdateRequest](../../models/operations/c1apiiamv1rolesupdaterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.C1APIIamV1RolesUpdateResponse](../../models/operations/c1apiiamv1rolesupdateresponse.md), error**

