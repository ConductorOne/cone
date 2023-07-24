# conductorone-api

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/conductorone/conductorone-sdk-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AppEntitlementSearch](docs/sdks/appentitlementsearch/README.md)

* [Search](docs/sdks/appentitlementsearch/README.md#search) - Search

### [AppEntitlementUserBinding](docs/sdks/appentitlementuserbinding/README.md)

* [ListAppUsersForIdentityWithGrant](docs/sdks/appentitlementuserbinding/README.md#listappusersforidentitywithgrant) - List App Users For Identity With Grant

### [AppEntitlements](docs/sdks/appentitlements/README.md)

* [Get](docs/sdks/appentitlements/README.md#get) - Get
* [List](docs/sdks/appentitlements/README.md#list) - List
* [ListForAppResource](docs/sdks/appentitlements/README.md#listforappresource) - List For App Resource
* [ListForAppUser](docs/sdks/appentitlements/README.md#listforappuser) - List For App User
* [ListGroups](docs/sdks/appentitlements/README.md#listgroups) - List Groups
* [ListUsers](docs/sdks/appentitlements/README.md#listusers) - List Users
* [Update](docs/sdks/appentitlements/README.md#update) - Update

### [AppOwners](docs/sdks/appowners/README.md)

* [Add](docs/sdks/appowners/README.md#add) - Add
* [List](docs/sdks/appowners/README.md#list) - List
* [Remove](docs/sdks/appowners/README.md#remove) - Remove

### [AppReport](docs/sdks/appreport/README.md)

* [List](docs/sdks/appreport/README.md#list) - List

### [AppReportAction](docs/sdks/appreportaction/README.md)

* [GenerateReport](docs/sdks/appreportaction/README.md#generatereport) - Generate Report

### [AppResource](docs/sdks/appresource/README.md)

* [Get](docs/sdks/appresource/README.md#get) - Get
* [List](docs/sdks/appresource/README.md#list) - List

### [AppResourceOwners](docs/sdks/appresourceowners/README.md)

* [List](docs/sdks/appresourceowners/README.md#list) - List

### [AppResourceSearch](docs/sdks/appresourcesearch/README.md)

* [SearchAppResourceTypes](docs/sdks/appresourcesearch/README.md#searchappresourcetypes) - Search App Resource Types

### [AppResourceType](docs/sdks/appresourcetype/README.md)

* [Get](docs/sdks/appresourcetype/README.md#get) - Get
* [List](docs/sdks/appresourcetype/README.md#list) - List

### [AppSearch](docs/sdks/appsearch/README.md)

* [Search](docs/sdks/appsearch/README.md#search) - Search

### [AppUsageControls](docs/sdks/appusagecontrols/README.md)

* [Get](docs/sdks/appusagecontrols/README.md#get) - Get
* [Update](docs/sdks/appusagecontrols/README.md#update) - Update

### [Apps](docs/sdks/apps/README.md)

* [Create](docs/sdks/apps/README.md#create) - Create
* [Delete](docs/sdks/apps/README.md#delete) - Delete
* [Get](docs/sdks/apps/README.md#get) - Get
* [List](docs/sdks/apps/README.md#list) - List
* [Update](docs/sdks/apps/README.md#update) - Update

### [Auth](docs/sdks/auth/README.md)

* [Introspect](docs/sdks/auth/README.md#introspect) - Introspect

### [Connector](docs/sdks/connector/README.md)

* [CreateDelegated](docs/sdks/connector/README.md#createdelegated) - Create Delegated
* [Delete](docs/sdks/connector/README.md#delete) - Delete
* [Get](docs/sdks/connector/README.md#get) - Get
* [GetCredentials](docs/sdks/connector/README.md#getcredentials) - Get Credentials
* [List](docs/sdks/connector/README.md#list) - List
* [RevokeCredential](docs/sdks/connector/README.md#revokecredential) - Revoke Credential
* [RotateCredential](docs/sdks/connector/README.md#rotatecredential) - Rotate Credential
* [Update](docs/sdks/connector/README.md#update) - Update
* [UpdateDelegated](docs/sdks/connector/README.md#updatedelegated) - Update Delegated

### [Directory](docs/sdks/directory/README.md)

* [Create](docs/sdks/directory/README.md#create) - Create
* [Delete](docs/sdks/directory/README.md#delete) - Delete
* [Get](docs/sdks/directory/README.md#get) - Get
* [List](docs/sdks/directory/README.md#list) - List

### [PersonalClient](docs/sdks/personalclient/README.md)

* [Create](docs/sdks/personalclient/README.md#create) - Create

### [Policies](docs/sdks/policies/README.md)

* [Create](docs/sdks/policies/README.md#create) - Create
* [Delete](docs/sdks/policies/README.md#delete) - Delete
* [Get](docs/sdks/policies/README.md#get) - Get
* [List](docs/sdks/policies/README.md#list) - List
* [Update](docs/sdks/policies/README.md#update) - Update

### [PolicySearch](docs/sdks/policysearch/README.md)

* [Search](docs/sdks/policysearch/README.md#search) - Search

### [RequestCatalogManagement](docs/sdks/requestcatalogmanagement/README.md)

* [AddAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#addaccessentitlements) - Add Access Entitlements
* [AddAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#addappentitlements) - Add App Entitlements
* [Create](docs/sdks/requestcatalogmanagement/README.md#create) - Create
* [Delete](docs/sdks/requestcatalogmanagement/README.md#delete) - Delete
* [Get](docs/sdks/requestcatalogmanagement/README.md#get) - Get
* [ListEntitlementsForAccess](docs/sdks/requestcatalogmanagement/README.md#listentitlementsforaccess) - List Entitlements For Access
* [ListEntitlementsPerCatalog](docs/sdks/requestcatalogmanagement/README.md#listentitlementspercatalog) - List Entitlements Per Catalog
* [RemoveAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeaccessentitlements) - Remove Access Entitlements
* [RemoveAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeappentitlements) - Remove App Entitlements
* [Update](docs/sdks/requestcatalogmanagement/README.md#update) - Update

### [RequestCatalogSearch](docs/sdks/requestcatalogsearch/README.md)

* [SearchEntitlements](docs/sdks/requestcatalogsearch/README.md#searchentitlements) - Search Entitlements

### [Roles](docs/sdks/roles/README.md)

* [Get](docs/sdks/roles/README.md#get) - Get
* [List](docs/sdks/roles/README.md#list) - List
* [Update](docs/sdks/roles/README.md#update) - Update

### [Task](docs/sdks/task/README.md)

* [CreateGrantTask](docs/sdks/task/README.md#creategranttask) - Create Grant Task
* [CreateRevokeTask](docs/sdks/task/README.md#createrevoketask) - Create Revoke Task
* [Get](docs/sdks/task/README.md#get) - Get

### [TaskActions](docs/sdks/taskactions/README.md)

* [Approve](docs/sdks/taskactions/README.md#approve) - Approve
* [Comment](docs/sdks/taskactions/README.md#comment) - Comment
* [Deny](docs/sdks/taskactions/README.md#deny) - Deny
* [EscalateToEmergencyAccess](docs/sdks/taskactions/README.md#escalatetoemergencyaccess) - Escalate To Emergency Access

### [TaskSearch](docs/sdks/tasksearch/README.md)

* [Search](docs/sdks/tasksearch/README.md#search) - Search

### [User](docs/sdks/user/README.md)

* [Get](docs/sdks/user/README.md#get) - Get
* [List](docs/sdks/user/README.md#list) - List

### [UserSearch](docs/sdks/usersearch/README.md)

* [Search](docs/sdks/usersearch/README.md#search) - Search
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
