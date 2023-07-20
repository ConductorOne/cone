# TaskSearch

### Available Operations

* [Search](#search) - Invokes the c1.api.task.v1.TaskSearchService.Search method.

## Search

Invokes the c1.api.task.v1.TaskSearchService.Search method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/types"
)

func main() {
    s := conductoroneapi.New(
        conductoroneapi.WithSecurity(shared.Security{
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TaskSearch.Search(ctx, shared.TaskSearchRequest{
        TaskExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "nisi",
                "voluptatibus",
                "quaerat",
            },
        },
        AccessReviewIds: []string{
            "distinctio",
            "nisi",
            "quis",
        },
        AccountOwnerIds: []string{
            "libero",
            "minus",
        },
        ActorID: conductoroneapi.String("facere"),
        AppEntitlementIds: []string{
            "ipsum",
            "ad",
            "voluptatibus",
        },
        AppResourceIds: []string{
            "consequuntur",
            "debitis",
            "labore",
            "rerum",
        },
        AppResourceTypeIds: []string{
            "reprehenderit",
        },
        AppUserSubjectIds: []string{
            "neque",
            "iusto",
        },
        ApplicationIds: []string{
            "rem",
            "eligendi",
            "fugiat",
        },
        AssigneesInIds: []string{
            "officiis",
            "ducimus",
            "dolor",
        },
        CreatedAfter: types.MustTimeFromString("2022-05-18T15:26:53.505Z"),
        CreatedBefore: types.MustTimeFromString("2022-08-31T21:53:00.029Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepApproval.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusAll.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "ad",
            "aspernatur",
            "enim",
            "delectus",
        },
        ExcludeIds: []string{
            "dignissimos",
            "libero",
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "ab",
        },
        OpenerIds: []string{
            "accusamus",
            "saepe",
        },
        PageSize: conductoroneapi.Float64(7348.14),
        PageToken: conductoroneapi.String("veniam"),
        PreviouslyActedOnIds: []string{
            "reiciendis",
        },
        Query: conductoroneapi.String("earum"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("85fc3781-4d4c-498e-8c2b-b89eb75dad63"),
            },
            shared.TaskRef{
                ID: conductoroneapi.String("6c600503-d8bb-4311-80f7-39ae9e057eb8"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByUnspecified.ToPointer(),
        SubjectIds: []string{
            "accusamus",
            "eos",
            "totam",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
        },
        TaskTypes: []shared.TaskType{
            shared.TaskType{
                TaskTypeCertify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("dolor"),
                    AccessReviewSelection: conductoroneapi.String("sunt"),
                    AppEntitlementID: conductoroneapi.String("a"),
                    AppID: conductoroneapi.String("dolor"),
                    AppUserID: conductoroneapi.String("occaecati"),
                    IdentityUserID: conductoroneapi.String("atque"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-02-20T03:19:49.048Z"),
                },
                TaskTypeGrant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("minus"),
                    AppID: conductoroneapi.String("esse"),
                    AppUserID: conductoroneapi.String("voluptatem"),
                    GrantDuration: conductoroneapi.String("perferendis"),
                    IdentityUserID: conductoroneapi.String("rerum"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-07-06T08:20:32.050Z"),
                },
                TaskTypeRevoke: &shared.TaskTypeRevoke{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-04-09T04:08:42.865Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2021-03-27T10:34:01.168Z"),
                            LastLogin: types.MustTimeFromString("2022-04-01T01:36:59.317Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("dignissimos"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("consectetur"),
                            CertTicketID: conductoroneapi.String("soluta"),
                        },
                    },
                    AppEntitlementID: conductoroneapi.String("natus"),
                    AppID: conductoroneapi.String("temporibus"),
                    AppUserID: conductoroneapi.String("officia"),
                    IdentityUserID: conductoroneapi.String("amet"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-03-22T22:18:02.609Z"),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.TaskSearchRequest](../../models/shared/tasksearchrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.C1APITaskV1TaskSearchServiceSearchResponse](../../models/operations/c1apitaskv1tasksearchservicesearchresponse.md), error**

