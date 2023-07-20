# Policies

### Available Operations

* [Create](#create) - Invokes the c1.api.policy.v1.Policies.Create method.
* [Delete](#delete) - Invokes the c1.api.policy.v1.Policies.Delete method.
* [Get](#get) - Invokes the c1.api.policy.v1.Policies.Get method.
* [List](#list) - Invokes the c1.api.policy.v1.Policies.List method.
* [Update](#update) - Invokes the c1.api.policy.v1.Policies.Update method.

## Create

Invokes the c1.api.policy.v1.Policies.Create method.

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
    res, err := s.Policies.Create(ctx, shared.CreatePolicyRequest{
        Description: conductoroneapi.String("eius"),
        DisplayName: conductoroneapi.String("libero"),
        PolicySteps: map[string]shared.PolicySteps{
            "soluta": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("aliquam"),
                                AppID: conductoroneapi.String("sapiente"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "ullam",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "ullam",
                                    "nisi",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "voluptatum",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "quibusdam",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "deleniti",
                                    "itaque",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "architecto",
                                    "omnis",
                                    "tenetur",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "at",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("et"),
                                    EntitlementID: conductoroneapi.String("voluptate"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("ipsa"),
                                    UserIds: []string{
                                        "veritatis",
                                        "consectetur",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
            "adipisci": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("temporibus"),
                                AppID: conductoroneapi.String("accusantium"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "aut",
                                    "laudantium",
                                    "eum",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "ab",
                                    "corrupti",
                                    "non",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "dolor",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "numquam",
                                    "impedit",
                                    "explicabo",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "aut",
                                    "dignissimos",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "maiores",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "velit",
                                    "voluptatibus",
                                    "voluptas",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("asperiores"),
                                    EntitlementID: conductoroneapi.String("aperiam"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("ea"),
                                    UserIds: []string{
                                        "consequuntur",
                                        "repellendus",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("officia"),
                                AppID: conductoroneapi.String("maxime"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "officia",
                                    "asperiores",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "quae",
                                    "quaerat",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "quod",
                                    "labore",
                                    "ab",
                                    "adipisci",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "id",
                                    "suscipit",
                                    "velit",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "est",
                                    "recusandae",
                                    "totam",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "vel",
                                    "ducimus",
                                    "quos",
                                    "vel",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "possimus",
                                    "facilis",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("cum"),
                                    EntitlementID: conductoroneapi.String("commodi"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("in"),
                                    UserIds: []string{
                                        "reiciendis",
                                        "assumenda",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("nemo"),
                                AppID: conductoroneapi.String("recusandae"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "aperiam",
                                    "cum",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "in",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "earum",
                                    "facere",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "doloribus",
                                    "suscipit",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "quidem",
                                    "saepe",
                                    "necessitatibus",
                                    "dolore",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "asperiores",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "non",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("amet"),
                                    EntitlementID: conductoroneapi.String("beatae"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("dignissimos"),
                                    UserIds: []string{
                                        "debitis",
                                        "consectetur",
                                        "corporis",
                                        "harum",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
            "laboriosam": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("voluptates"),
                                AppID: conductoroneapi.String("libero"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "accusamus",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "tempora",
                                    "aspernatur",
                                    "voluptas",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "voluptas",
                                    "minima",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "dolorum",
                                    "adipisci",
                                    "minus",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "blanditiis",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "dolore",
                                    "aliquam",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "temporibus",
                                    "ullam",
                                    "adipisci",
                                    "cum",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("blanditiis"),
                                    EntitlementID: conductoroneapi.String("quas"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("hic"),
                                    UserIds: []string{
                                        "culpa",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
            "corrupti": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("totam"),
                                AppID: conductoroneapi.String("hic"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "nobis",
                                    "sit",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "sed",
                                    "reiciendis",
                                    "explicabo",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "facilis",
                                    "voluptate",
                                    "expedita",
                                    "ab",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "dolore",
                                    "laborum",
                                    "sed",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "commodi",
                                    "quidem",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "voluptas",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "architecto",
                                    "suscipit",
                                    "sapiente",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("debitis"),
                                    EntitlementID: conductoroneapi.String("illo"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("reiciendis"),
                                    UserIds: []string{
                                        "corrupti",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("maiores"),
                                AppID: conductoroneapi.String("incidunt"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "provident",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "necessitatibus",
                                    "ipsum",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "occaecati",
                                    "quos",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "tempora",
                                    "tempora",
                                    "voluptate",
                                    "reiciendis",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "sit",
                                    "non",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "praesentium",
                                    "facilis",
                                    "quaerat",
                                    "incidunt",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "debitis",
                                    "rem",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("sit"),
                                    EntitlementID: conductoroneapi.String("nobis"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("error"),
                                    UserIds: []string{
                                        "minima",
                                        "recusandae",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("reiciendis"),
                                AppID: conductoroneapi.String("nulla"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "aperiam",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "numquam",
                                    "veniam",
                                    "in",
                                    "officiis",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "laudantium",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "praesentium",
                                    "cum",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "dolorum",
                                    "voluptatum",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "hic",
                                    "expedita",
                                    "debitis",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "dolorum",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("nostrum"),
                                    EntitlementID: conductoroneapi.String("officia"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("dolorum"),
                                    UserIds: []string{
                                        "accusamus",
                                        "tempora",
                                        "atque",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("fugit"),
                                AppID: conductoroneapi.String("ut"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "voluptatem",
                                    "culpa",
                                    "expedita",
                                    "magnam",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "esse",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "sit",
                                    "voluptatum",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "repudiandae",
                                    "corporis",
                                    "et",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "ex",
                                    "sed",
                                    "sit",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "nostrum",
                                    "saepe",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "consequatur",
                                    "incidunt",
                                    "reiciendis",
                                },
                            },
                            AllowReassignment: conductoroneapi.Bool(false),
                            Assigned: conductoroneapi.Bool(false),
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("dolorem"),
                                    EntitlementID: conductoroneapi.String("harum"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("dicta"),
                                    UserIds: []string{
                                        "occaecati",
                                    },
                                },
                            },
                            Assigned: conductoroneapi.Bool(false),
                        },
                    },
                },
            },
        },
        PolicyType: shared.CreatePolicyRequestPolicyTypePolicyTypeGrant.ToPointer(),
        PostActions: []shared.PolicyPostActions{
            shared.PolicyPostActions{
                CertifyRemediateImmediately: conductoroneapi.Bool(false),
            },
            shared.PolicyPostActions{
                CertifyRemediateImmediately: conductoroneapi.Bool(false),
            },
            shared.PolicyPostActions{
                CertifyRemediateImmediately: conductoroneapi.Bool(false),
            },
        },
        ReassignTasksToDelegates: conductoroneapi.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.CreatePolicyRequest](../../models/shared/createpolicyrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.C1APIPolicyV1PoliciesCreateResponse](../../models/operations/c1apipolicyv1policiescreateresponse.md), error**


## Delete

Invokes the c1.api.policy.v1.Policies.Delete method.

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
    res, err := s.Policies.Delete(ctx, operations.C1APIPolicyV1PoliciesDeleteRequest{
        DeletePolicyRequest: &shared.DeletePolicyRequest{},
        ID: "8abf603a-79f9-4dfe-8ab7-da8a50ce187f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeletePolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.C1APIPolicyV1PoliciesDeleteRequest](../../models/operations/c1apipolicyv1policiesdeleterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.C1APIPolicyV1PoliciesDeleteResponse](../../models/operations/c1apipolicyv1policiesdeleteresponse.md), error**


## Get

Invokes the c1.api.policy.v1.Policies.Get method.

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
    res, err := s.Policies.Get(ctx, operations.C1APIPolicyV1PoliciesGetRequest{
        ID: "86bc173d-689e-4ee9-926f-8d986e881ead",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.C1APIPolicyV1PoliciesGetRequest](../../models/operations/c1apipolicyv1policiesgetrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.C1APIPolicyV1PoliciesGetResponse](../../models/operations/c1apipolicyv1policiesgetresponse.md), error**


## List

Invokes the c1.api.policy.v1.Policies.List method.

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
    res, err := s.Policies.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.C1APIPolicyV1PoliciesListResponse](../../models/operations/c1apipolicyv1policieslistresponse.md), error**


## Update

Invokes the c1.api.policy.v1.Policies.Update method.

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
    res, err := s.Policies.Update(ctx, operations.C1APIPolicyV1PoliciesUpdateRequest{
        UpdatePolicyRequestInput: &shared.UpdatePolicyRequestInput{
            Policy: &shared.PolicyInput{
                Description: conductoroneapi.String("labore"),
                DisplayName: conductoroneapi.String("reiciendis"),
                ID: conductoroneapi.String("0e101256-3f94-4e29-a973-e922a57a15be"),
                PolicySteps: map[string]shared.PolicySteps{
                    "vero": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("iure"),
                                        AppID: conductoroneapi.String("ipsa"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "quae",
                                            "molestiae",
                                            "eveniet",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "cum",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "necessitatibus",
                                            "ratione",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "distinctio",
                                            "voluptatum",
                                            "rem",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "ad",
                                            "repellat",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "corporis",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "nihil",
                                            "mollitia",
                                            "voluptas",
                                        },
                                    },
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    Assigned: conductoroneapi.Bool(false),
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("alias"),
                                            EntitlementID: conductoroneapi.String("maiores"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("reiciendis"),
                                            UserIds: []string{
                                                "id",
                                            },
                                        },
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                },
                            },
                        },
                    },
                },
                PolicyType: shared.PolicyPolicyTypePolicyTypeGrant.ToPointer(),
                PostActions: []shared.PolicyPostActions{
                    shared.PolicyPostActions{
                        CertifyRemediateImmediately: conductoroneapi.Bool(false),
                    },
                    shared.PolicyPostActions{
                        CertifyRemediateImmediately: conductoroneapi.Bool(false),
                    },
                },
                ReassignTasksToDelegates: conductoroneapi.Bool(false),
                SystemBuiltin: conductoroneapi.Bool(false),
            },
            UpdateMask: conductoroneapi.String("dolorum"),
        },
        ID: "31e94764-a3e8-465e-b956-f9251a5a9da6",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdatePolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.C1APIPolicyV1PoliciesUpdateRequest](../../models/operations/c1apipolicyv1policiesupdaterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.C1APIPolicyV1PoliciesUpdateResponse](../../models/operations/c1apipolicyv1policiesupdateresponse.md), error**

