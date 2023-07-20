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

* [Search](docs/sdks/appentitlementsearch/README.md#search) - Invokes the c1.api.app.v1.AppEntitlementSearchService.Search method.

### [AppEntitlementUserBinding](docs/sdks/appentitlementuserbinding/README.md)

* [ListAppUsersForIdentityWithGrant](docs/sdks/appentitlementuserbinding/README.md#listappusersforidentitywithgrant) - Invokes the c1.api.app.v1.AppEntitlementUserBindingService.ListAppUsersForIdentityWithGrant method.

### [AppEntitlements](docs/sdks/appentitlements/README.md)

* [Get](docs/sdks/appentitlements/README.md#get) - Invokes the c1.api.app.v1.AppEntitlements.Get method.
* [List](docs/sdks/appentitlements/README.md#list) - Invokes the c1.api.app.v1.AppEntitlements.List method.
* [ListForAppResource](docs/sdks/appentitlements/README.md#listforappresource) - Invokes the c1.api.app.v1.AppEntitlements.ListForAppResource method.
* [ListForAppUser](docs/sdks/appentitlements/README.md#listforappuser) - Invokes the c1.api.app.v1.AppEntitlements.ListForAppUser method.
* [ListGroups](docs/sdks/appentitlements/README.md#listgroups) - Invokes the c1.api.app.v1.AppEntitlements.ListGroups method.
* [ListUsers](docs/sdks/appentitlements/README.md#listusers) - Invokes the c1.api.app.v1.AppEntitlements.ListUsers method.
* [Update](docs/sdks/appentitlements/README.md#update) - Invokes the c1.api.app.v1.AppEntitlements.Update method.

### [AppOwners](docs/sdks/appowners/README.md)

* [Add](docs/sdks/appowners/README.md#add) - Invokes the c1.api.app.v1.AppOwners.Add method.
* [List](docs/sdks/appowners/README.md#list) - Invokes the c1.api.app.v1.AppOwners.List method.
* [Remove](docs/sdks/appowners/README.md#remove) - Invokes the c1.api.app.v1.AppOwners.Remove method.

### [AppReport](docs/sdks/appreport/README.md)

* [List](docs/sdks/appreport/README.md#list) - Invokes the c1.api.app.v1.AppReportService.List method.

### [AppReportAction](docs/sdks/appreportaction/README.md)

* [GenerateReport](docs/sdks/appreportaction/README.md#generatereport) - Invokes the c1.api.app.v1.AppReportActionService.GenerateReport method.

### [AppResource](docs/sdks/appresource/README.md)

* [Get](docs/sdks/appresource/README.md#get) - Invokes the c1.api.app.v1.AppResourceService.Get method.
* [List](docs/sdks/appresource/README.md#list) - Invokes the c1.api.app.v1.AppResourceService.List method.

### [AppResourceOwners](docs/sdks/appresourceowners/README.md)

* [List](docs/sdks/appresourceowners/README.md#list) - Invokes the c1.api.app.v1.AppResourceOwners.List method.

### [AppResourceSearch](docs/sdks/appresourcesearch/README.md)

* [SearchAppResourceTypes](docs/sdks/appresourcesearch/README.md#searchappresourcetypes) - Invokes the c1.api.app.v1.AppResourceSearch.SearchAppResourceTypes method.

### [AppResourceType](docs/sdks/appresourcetype/README.md)

* [Get](docs/sdks/appresourcetype/README.md#get) - Invokes the c1.api.app.v1.AppResourceTypeService.Get method.
* [List](docs/sdks/appresourcetype/README.md#list) - Invokes the c1.api.app.v1.AppResourceTypeService.List method.

### [AppSearch](docs/sdks/appsearch/README.md)

* [Search](docs/sdks/appsearch/README.md#search) - Invokes the c1.api.app.v1.AppSearch.Search method.

### [AppUsageControls](docs/sdks/appusagecontrols/README.md)

* [Get](docs/sdks/appusagecontrols/README.md#get) - Invokes the c1.api.app.v1.AppUsageControlsService.Get method.
* [Update](docs/sdks/appusagecontrols/README.md#update) - Invokes the c1.api.app.v1.AppUsageControlsService.Update method.

### [Apps](docs/sdks/apps/README.md)

* [Create](docs/sdks/apps/README.md#create) - Invokes the c1.api.app.v1.Apps.Create method.
* [Delete](docs/sdks/apps/README.md#delete) - Invokes the c1.api.app.v1.Apps.Delete method.
* [Get](docs/sdks/apps/README.md#get) - Invokes the c1.api.app.v1.Apps.Get method.
* [List](docs/sdks/apps/README.md#list) - Invokes the c1.api.app.v1.Apps.List method.
* [Update](docs/sdks/apps/README.md#update) - Invokes the c1.api.app.v1.Apps.Update method.

### [Auth](docs/sdks/auth/README.md)

* [Introspect](docs/sdks/auth/README.md#introspect) - Invokes the c1.api.auth.v1.Auth.Introspect method.

### [Connector](docs/sdks/connector/README.md)

* [CreateDelegated](docs/sdks/connector/README.md#createdelegated) - Invokes the c1.api.app.v1.ConnectorService.CreateDelegated method.
* [Delete](docs/sdks/connector/README.md#delete) - Invokes the c1.api.app.v1.ConnectorService.Delete method.
* [Get](docs/sdks/connector/README.md#get) - Invokes the c1.api.app.v1.ConnectorService.Get method.
* [GetCredentials](docs/sdks/connector/README.md#getcredentials) - Invokes the c1.api.app.v1.ConnectorService.GetCredentials method.
* [List](docs/sdks/connector/README.md#list) - Invokes the c1.api.app.v1.ConnectorService.List method.
* [RevokeCredential](docs/sdks/connector/README.md#revokecredential) - Invokes the c1.api.app.v1.ConnectorService.RevokeCredential method.
* [RotateCredential](docs/sdks/connector/README.md#rotatecredential) - Invokes the c1.api.app.v1.ConnectorService.RotateCredential method.
* [Update](docs/sdks/connector/README.md#update) - Invokes the c1.api.app.v1.ConnectorService.Update method.
* [UpdateDelegated](docs/sdks/connector/README.md#updatedelegated) - Invokes the c1.api.app.v1.ConnectorService.UpdateDelegated method.

### [Directory](docs/sdks/directory/README.md)

* [Create](docs/sdks/directory/README.md#create) - Invokes the c1.api.directory.v1.DirectoryService.Create method.
* [Delete](docs/sdks/directory/README.md#delete) - Invokes the c1.api.directory.v1.DirectoryService.Delete method.
* [Get](docs/sdks/directory/README.md#get) - Invokes the c1.api.directory.v1.DirectoryService.Get method.
* [List](docs/sdks/directory/README.md#list) - Invokes the c1.api.directory.v1.DirectoryService.List method.

### [PersonalClient](docs/sdks/personalclient/README.md)

* [Create](docs/sdks/personalclient/README.md#create) - Invokes the c1.api.iam.v1.PersonalClientService.Create method.

### [Policies](docs/sdks/policies/README.md)

* [Create](docs/sdks/policies/README.md#create) - Invokes the c1.api.policy.v1.Policies.Create method.
* [Delete](docs/sdks/policies/README.md#delete) - Invokes the c1.api.policy.v1.Policies.Delete method.
* [Get](docs/sdks/policies/README.md#get) - Invokes the c1.api.policy.v1.Policies.Get method.
* [List](docs/sdks/policies/README.md#list) - Invokes the c1.api.policy.v1.Policies.List method.
* [Update](docs/sdks/policies/README.md#update) - Invokes the c1.api.policy.v1.Policies.Update method.

### [PolicySearch](docs/sdks/policysearch/README.md)

* [Search](docs/sdks/policysearch/README.md#search) - Invokes the c1.api.policy.v1.PolicySearch.Search method.

### [RequestCatalogManagement](docs/sdks/requestcatalogmanagement/README.md)

* [AddAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#addaccessentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.AddAccessEntitlements method.
* [AddAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#addappentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.AddAppEntitlements method.
* [Create](docs/sdks/requestcatalogmanagement/README.md#create) - Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.Create method.
* [Delete](docs/sdks/requestcatalogmanagement/README.md#delete) - Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.Delete method.
* [Get](docs/sdks/requestcatalogmanagement/README.md#get) - Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.Get method.
* [ListEntitlementsForAccess](docs/sdks/requestcatalogmanagement/README.md#listentitlementsforaccess) - Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.ListEntitlementsForAccess method.
* [ListEntitlementsPerCatalog](docs/sdks/requestcatalogmanagement/README.md#listentitlementspercatalog) - Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.ListEntitlementsPerCatalog method.
* [RemoveAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeaccessentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.RemoveAccessEntitlements method.
* [RemoveAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeappentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.RemoveAppEntitlements method.
* [Update](docs/sdks/requestcatalogmanagement/README.md#update) - Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.Update method.

### [RequestCatalogSearch](docs/sdks/requestcatalogsearch/README.md)

* [SearchEntitlements](docs/sdks/requestcatalogsearch/README.md#searchentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.

### [Roles](docs/sdks/roles/README.md)

* [Get](docs/sdks/roles/README.md#get) - Invokes the c1.api.iam.v1.Roles.Get method.
* [List](docs/sdks/roles/README.md#list) - Invokes the c1.api.iam.v1.Roles.List method.
* [Update](docs/sdks/roles/README.md#update) - Invokes the c1.api.iam.v1.Roles.Update method.

### [Task](docs/sdks/task/README.md)

* [CreateGrantTask](docs/sdks/task/README.md#creategranttask) - Invokes the c1.api.task.v1.TaskService.CreateGrantTask method.
* [CreateRevokeTask](docs/sdks/task/README.md#createrevoketask) - Invokes the c1.api.task.v1.TaskService.CreateRevokeTask method.
* [Get](docs/sdks/task/README.md#get) - Invokes the c1.api.task.v1.TaskService.Get method.

### [TaskActions](docs/sdks/taskactions/README.md)

* [Approve](docs/sdks/taskactions/README.md#approve) - Invokes the c1.api.task.v1.TaskActionsService.Approve method.
* [Comment](docs/sdks/taskactions/README.md#comment) - Invokes the c1.api.task.v1.TaskActionsService.Comment method.
* [Deny](docs/sdks/taskactions/README.md#deny) - Invokes the c1.api.task.v1.TaskActionsService.Deny method.
* [EscalateToEmergencyAccess](docs/sdks/taskactions/README.md#escalatetoemergencyaccess) - Invokes the c1.api.task.v1.TaskActionsService.EscalateToEmergencyAccess method.

### [TaskSearch](docs/sdks/tasksearch/README.md)

* [Search](docs/sdks/tasksearch/README.md#search) - Invokes the c1.api.task.v1.TaskSearchService.Search method.

### [User](docs/sdks/user/README.md)

* [Get](docs/sdks/user/README.md#get) - Invokes the c1.api.user.v1.UserService.Get method.
* [List](docs/sdks/user/README.md#list) - Invokes the c1.api.user.v1.UserService.List method.

### [UserSearch](docs/sdks/usersearch/README.md)

* [Search](docs/sdks/usersearch/README.md#search) - Invokes the c1.api.user.v1.UserSearch.Search method.
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
