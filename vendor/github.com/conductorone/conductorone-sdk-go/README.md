# conductorone-api

<!-- Start Summary [summary] -->
## Summary

ConductorOne API: The ConductorOne API is a HTTP API for managing ConductorOne resources.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [conductorone-api](#conductorone-api)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Error Handling](#error-handling)
  * [SDK Example Usage with Custom Server/Tenant](#sdk-example-usage-with-custom-servertenant)
  * [Retries](#retries)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/ConductorOne/conductorone-sdk-go
```
<!-- End SDK Installation [installation] -->

## SDK Example Usage

### Example

```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := NewWithCredentials(ctx, &ClientCredentials{
		ClientID:     "",
		ClientSecret: "",
	} )

	res, err := s.Apps.Create(ctx, &shared.CreateAppRequest{
		Owners: []string{
			"string",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateAppResponse != nil {
		// handle response
	}
}

```
<!-- No SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md)

* [CancelAppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md#cancelappaccessrequestsdefaults) - Cancel App Access Requests Defaults
* [CreateAppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md#createappaccessrequestsdefaults) - Create App Access Requests Defaults
* [GetAppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md#getappaccessrequestsdefaults) - Get App Access Requests Defaults

### [AppEntitlementOwners](docs/sdks/appentitlementowners/README.md)

* [Add](docs/sdks/appentitlementowners/README.md#add) - Add
* [List](docs/sdks/appentitlementowners/README.md#list) - List
* [Remove](docs/sdks/appentitlementowners/README.md#remove) - Remove
* [Set](docs/sdks/appentitlementowners/README.md#set) - Set

### [AppEntitlements](docs/sdks/appentitlements/README.md)

* [AddManuallyManagedMembers](docs/sdks/appentitlements/README.md#addmanuallymanagedmembers) - Add Manually Managed Members
* [Create](docs/sdks/appentitlements/README.md#create) - Create
* [Delete](docs/sdks/appentitlements/README.md#delete) - Delete
* [Get](docs/sdks/appentitlements/README.md#get) - Get
* [List](docs/sdks/appentitlements/README.md#list) - List
* [ListForAppResource](docs/sdks/appentitlements/README.md#listforappresource) - List For App Resource
* [ListForAppUser](docs/sdks/appentitlements/README.md#listforappuser) - List For App User
* [~~ListUsers~~](docs/sdks/appentitlements/README.md#listusers) - List Users :warning: **Deprecated**
* [RemoveEntitlementMembership](docs/sdks/appentitlements/README.md#removeentitlementmembership) - Remove Entitlement Membership
* [Update](docs/sdks/appentitlements/README.md#update) - Update

### [AppEntitlementSearch](docs/sdks/appentitlementsearch/README.md)

* [Search](docs/sdks/appentitlementsearch/README.md#search) - Search
* [SearchAppEntitlementsForAppUser](docs/sdks/appentitlementsearch/README.md#searchappentitlementsforappuser) - Search App Entitlements For App User
* [SearchAppEntitlementsWithExpired](docs/sdks/appentitlementsearch/README.md#searchappentitlementswithexpired) - Search App Entitlements With Expired

### [AppEntitlementsProxy](docs/sdks/appentitlementsproxy/README.md)

* [Create](docs/sdks/appentitlementsproxy/README.md#create) - Create
* [Delete](docs/sdks/appentitlementsproxy/README.md#delete) - Delete
* [Get](docs/sdks/appentitlementsproxy/README.md#get) - Get

### [AppEntitlementUserBinding](docs/sdks/appentitlementuserbinding/README.md)

* [ListAppUsersForIdentityWithGrant](docs/sdks/appentitlementuserbinding/README.md#listappusersforidentitywithgrant) - List App Users For Identity With Grant
* [SearchGrantFeed](docs/sdks/appentitlementuserbinding/README.md#searchgrantfeed) - Search Grant Feed
* [SearchPastGrants](docs/sdks/appentitlementuserbinding/README.md#searchpastgrants) - Search Past Grants

### [AppOwners](docs/sdks/appowners/README.md)

* [Add](docs/sdks/appowners/README.md#add) - Add
* [List](docs/sdks/appowners/README.md#list) - List
* [Remove](docs/sdks/appowners/README.md#remove) - Remove
* [Set](docs/sdks/appowners/README.md#set) - Set

### [AppReport](docs/sdks/appreport/README.md)

* [List](docs/sdks/appreport/README.md#list) - List

### [AppReportAction](docs/sdks/appreportaction/README.md)

* [GenerateReport](docs/sdks/appreportaction/README.md#generatereport) - Generate Report

### [AppResource](docs/sdks/appresource/README.md)

* [CreateManuallyManagedAppResource](docs/sdks/appresource/README.md#createmanuallymanagedappresource) - Create Manually Managed App Resource
* [DeleteManuallyManagedAppResource](docs/sdks/appresource/README.md#deletemanuallymanagedappresource) - Delete Manually Managed App Resource
* [Get](docs/sdks/appresource/README.md#get) - Get
* [List](docs/sdks/appresource/README.md#list) - List
* [Update](docs/sdks/appresource/README.md#update) - Update

### [AppResourceOwners](docs/sdks/appresourceowners/README.md)

* [Add](docs/sdks/appresourceowners/README.md#add) - Add
* [List](docs/sdks/appresourceowners/README.md#list) - List
* [Remove](docs/sdks/appresourceowners/README.md#remove) - Remove

### [AppResourceSearch](docs/sdks/appresourcesearch/README.md)

* [SearchAppResourceTypes](docs/sdks/appresourcesearch/README.md#searchappresourcetypes) - Search App Resource Types

### [AppResourceType](docs/sdks/appresourcetype/README.md)

* [CreateManuallyManagedResourceType](docs/sdks/appresourcetype/README.md#createmanuallymanagedresourcetype) - Create Manually Managed Resource Type
* [DeleteManuallyManagedResourceType](docs/sdks/appresourcetype/README.md#deletemanuallymanagedresourcetype) - Delete Manually Managed Resource Type
* [Get](docs/sdks/appresourcetype/README.md#get) - Get
* [List](docs/sdks/appresourcetype/README.md#list) - List
* [UpdateManuallyManagedResourceType](docs/sdks/appresourcetype/README.md#updatemanuallymanagedresourcetype) - Update Manually Managed Resource Type

### [Apps](docs/sdks/apps/README.md)

* [Create](docs/sdks/apps/README.md#create) - Create
* [Delete](docs/sdks/apps/README.md#delete) - Delete
* [Get](docs/sdks/apps/README.md#get) - Get
* [List](docs/sdks/apps/README.md#list) - List
* [Update](docs/sdks/apps/README.md#update) - Update

### [AppSearch](docs/sdks/appsearch/README.md)

* [Search](docs/sdks/appsearch/README.md#search) - Search

### [AppUsageControls](docs/sdks/appusagecontrols/README.md)

* [Get](docs/sdks/appusagecontrols/README.md#get) - Get
* [Update](docs/sdks/appusagecontrols/README.md#update) - Update

### [AppUser](docs/sdks/appuser/README.md)

* [List](docs/sdks/appuser/README.md#list) - List
* [ListAppUserCredentials](docs/sdks/appuser/README.md#listappusercredentials) - List App User Credentials
* [ListAppUsersForUser](docs/sdks/appuser/README.md#listappusersforuser) - List App Users For User
* [Search](docs/sdks/appuser/README.md#search) - Search
* [Update](docs/sdks/appuser/README.md#update) - Update

### [Attributes](docs/sdks/attributes/README.md)

* [CreateAttributeValue](docs/sdks/attributes/README.md#createattributevalue) - Create Attribute Value
* [DeleteAttributeValue](docs/sdks/attributes/README.md#deleteattributevalue) - Delete Attribute Value
* [GetAttributeValue](docs/sdks/attributes/README.md#getattributevalue) - Get Attribute Value
* [ListAttributeTypes](docs/sdks/attributes/README.md#listattributetypes) - List Attribute Types
* [ListAttributeValues](docs/sdks/attributes/README.md#listattributevalues) - List Attribute Values

### [AttributeSearch](docs/sdks/attributesearch/README.md)

* [SearchAttributeValues](docs/sdks/attributesearch/README.md#searchattributevalues) - Search Attribute Values

### [Auth](docs/sdks/auth/README.md)

* [Introspect](docs/sdks/auth/README.md#introspect) - Introspect

### [AWSExternalIDSettings](docs/sdks/awsexternalidsettings/README.md)

* [Get](docs/sdks/awsexternalidsettings/README.md#get) - Get


### [Connector](docs/sdks/connector/README.md)

* [Create](docs/sdks/connector/README.md#create) - Create
* [CreateDelegated](docs/sdks/connector/README.md#createdelegated) - Create Delegated
* [Delete](docs/sdks/connector/README.md#delete) - Delete
* [ForceSync](docs/sdks/connector/README.md#forcesync) - Force Sync
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

### [Export](docs/sdks/export/README.md)

* [Create](docs/sdks/export/README.md#create) - Create
* [Delete](docs/sdks/export/README.md#delete) - Delete
* [Get](docs/sdks/export/README.md#get) - Get
* [List](docs/sdks/export/README.md#list) - List
* [ListEvents](docs/sdks/export/README.md#listevents) - List Events
* [Update](docs/sdks/export/README.md#update) - Update

### [ExportsSearch](docs/sdks/exportssearch/README.md)

* [Search](docs/sdks/exportssearch/README.md#search) - Search

### [PersonalClient](docs/sdks/personalclient/README.md)

* [Create](docs/sdks/personalclient/README.md#create) - Create
* [Delete](docs/sdks/personalclient/README.md#delete) - Delete
* [Get](docs/sdks/personalclient/README.md#get) - Get
* [List](docs/sdks/personalclient/README.md#list) - NOTE: Only shows personal clients for the current user.
* [Update](docs/sdks/personalclient/README.md#update) - Update

### [PersonalClientSearch](docs/sdks/personalclientsearch/README.md)

* [Search](docs/sdks/personalclientsearch/README.md#search) - NOTE: Searches personal clients for all users

### [Policies](docs/sdks/policies/README.md)

* [Create](docs/sdks/policies/README.md#create) - Create
* [Delete](docs/sdks/policies/README.md#delete) - Delete
* [Get](docs/sdks/policies/README.md#get) - Get
* [List](docs/sdks/policies/README.md#list) - List
* [Update](docs/sdks/policies/README.md#update) - Update

### [PolicySearch](docs/sdks/policysearch/README.md)

* [Search](docs/sdks/policysearch/README.md#search) - Search

### [PolicyValidate](docs/sdks/policyvalidate/README.md)

* [ValidateCEL](docs/sdks/policyvalidate/README.md#validatecel) - Validate Cel

### [RequestCatalogManagement](docs/sdks/requestcatalogmanagement/README.md)

* [AddAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#addaccessentitlements) - Add Access Entitlements
* [AddAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#addappentitlements) - Add App Entitlements
* [Create](docs/sdks/requestcatalogmanagement/README.md#create) - Create
* [Delete](docs/sdks/requestcatalogmanagement/README.md#delete) - Delete
* [ForceRunBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#forcerunbundleautomation) - Force Run Bundle Automation
* [Get](docs/sdks/requestcatalogmanagement/README.md#get) - Get
* [GetBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#getbundleautomation) - Get Bundle Automation
* [List](docs/sdks/requestcatalogmanagement/README.md#list) - List
* [ListEntitlementsForAccess](docs/sdks/requestcatalogmanagement/README.md#listentitlementsforaccess) - List Entitlements For Access
* [ListEntitlementsPerCatalog](docs/sdks/requestcatalogmanagement/README.md#listentitlementspercatalog) - List Entitlements Per Catalog
* [RemoveAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeaccessentitlements) - Remove Access Entitlements
* [RemoveAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeappentitlements) - Remove App Entitlements
* [SetBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#setbundleautomation) - Set Bundle Automation
* [Update](docs/sdks/requestcatalogmanagement/README.md#update) - Update

### [RequestCatalogSearch](docs/sdks/requestcatalogsearch/README.md)

* [SearchEntitlements](docs/sdks/requestcatalogsearch/README.md#searchentitlements) - Search Entitlements

### [Roles](docs/sdks/roles/README.md)

* [Get](docs/sdks/roles/README.md#get) - Get
* [List](docs/sdks/roles/README.md#list) - List
* [Update](docs/sdks/roles/README.md#update) - Update

### [SessionSettings](docs/sdks/sessionsettings/README.md)

* [Get](docs/sdks/sessionsettings/README.md#get) - Get
* [TestSourceIP](docs/sdks/sessionsettings/README.md#testsourceip) - Test Source Ip
* [Update](docs/sdks/sessionsettings/README.md#update) - Update

### [SystemLog](docs/sdks/systemlog/README.md)

* [ListEvents](docs/sdks/systemlog/README.md#listevents) - List Events

### [Task](docs/sdks/task/README.md)

* [CreateGrantTask](docs/sdks/task/README.md#creategranttask) - Create Grant Task
* [CreateOffboardingTask](docs/sdks/task/README.md#createoffboardingtask) - Create Offboarding Task
* [CreateRevokeTask](docs/sdks/task/README.md#createrevoketask) - Create Revoke Task
* [Get](docs/sdks/task/README.md#get) - Get

### [TaskActions](docs/sdks/taskactions/README.md)

* [Approve](docs/sdks/taskactions/README.md#approve) - Approve
* [Comment](docs/sdks/taskactions/README.md#comment) - Comment
* [Deny](docs/sdks/taskactions/README.md#deny) - Deny
* [EscalateToEmergencyAccess](docs/sdks/taskactions/README.md#escalatetoemergencyaccess) - Escalate To Emergency Access
* [HardReset](docs/sdks/taskactions/README.md#hardreset) - Hard Reset
* [ProcessNow](docs/sdks/taskactions/README.md#processnow) - Process Now
* [Reassign](docs/sdks/taskactions/README.md#reassign) - Reassign
* [Restart](docs/sdks/taskactions/README.md#restart) - Restart

### [TaskSearch](docs/sdks/tasksearch/README.md)

* [Search](docs/sdks/tasksearch/README.md#search) - Search

### [User](docs/sdks/user/README.md)

* [Get](docs/sdks/user/README.md#get) - Get
* [List](docs/sdks/user/README.md#list) - List

### [UserSearch](docs/sdks/usersearch/README.md)

* [Search](docs/sdks/usersearch/README.md#search) - Search

### [Webhooks](docs/sdks/webhooks/README.md)

* [Create](docs/sdks/webhooks/README.md#create) - Create
* [Delete](docs/sdks/webhooks/README.md#delete) - Delete
* [Get](docs/sdks/webhooks/README.md#get) - Get
* [List](docs/sdks/webhooks/README.md#list) - List
* [Test](docs/sdks/webhooks/README.md#test) - Test
* [Update](docs/sdks/webhooks/README.md#update) - Update

### [WebhooksSearch](docs/sdks/webhookssearch/README.md)

* [Search](docs/sdks/webhookssearch/README.md#search) - Search

### [Workflow](docs/sdks/workflow/README.md)

* [CreateWorkflow](docs/sdks/workflow/README.md#createworkflow) - Create Workflow
* [DeleteWorkflow](docs/sdks/workflow/README.md#deleteworkflow) - Delete Workflow
* [ExecuteWorkflow](docs/sdks/workflow/README.md#executeworkflow) - Execute Workflow
* [GetWorkflow](docs/sdks/workflow/README.md#getworkflow) - Get Workflow
* [ListWorkflows](docs/sdks/workflow/README.md#listworkflows) - List Workflows
* [UpdateWorkflow](docs/sdks/workflow/README.md#updateworkflow) - Update Workflow

### [WorkflowExecution](docs/sdks/workflowexecution/README.md)

* [GetWorkflowExecution](docs/sdks/workflowexecution/README.md#getworkflowexecution) - Get Workflow Execution
* [ListWorkflowExecutions](docs/sdks/workflowexecution/README.md#listworkflowexecutions) - List Workflow Executions

### [WorkflowExecutionActions](docs/sdks/workflowexecutionactions/README.md)

* [TerminateWorkflow](docs/sdks/workflowexecutionactions/README.md#terminateworkflow) - Terminate Workflow

### [WorkflowExecutionSearch](docs/sdks/workflowexecutionsearch/README.md)

* [SearchWorkflowExecutions](docs/sdks/workflowexecutionsearch/README.md#searchworkflowexecutions) - Search Workflow Executions

### [WorkflowSearch](docs/sdks/workflowsearch/README.md)

* [SearchWorkflows](docs/sdks/workflowsearch/README.md#searchworkflows) - Search Workflows

</details>
<!-- End Available Resources and Operations [operations] -->


<!-- No Special Types [types] -->


## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/sdkerrors"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := NewWithCredentials(ctx, &ClientCredentials{
		ClientID:     "",
		ClientSecret: "",
	})
	res, err := s.Apps.Create(ctx, &shared.CreateAppRequest{
		Owners: []string{
			"string",
		},
	})
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- No Error Handling [errors] -->


## SDK Example Usage with Custom Server/Tenant

### Example

```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go/v2"
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	/* Optional Override 
	* Server URL will be extracted from client, optionally, you can
	* provide a server URL or a tenant domain (will create URL https://{tenant_domain}.conductor.one) 
	*/
	opts := []sdk.CustomSDKOption{}
	opt, _ := sdk.WithTenantCustom("Server URL or Tenant Domain")
	opts = append(opts, opt)

	s := NewWithCredentials(ctx, &ClientCredentials{
		ClientID:     "",
		ClientSecret: "",
	} opts...)

	res, err := s.Apps.Create(ctx, &shared.CreateAppRequest{
		Owners: []string{
			"string",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateAppResponse != nil {
		// handle response
	}
}

```
<!-- No Server Selection [server] -->

<!-- No Custom HTTP Client [http-client] -->

<!-- No Authentication [security] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/retry"
	"log"
	"pkg/models/operations"
)

func main() {
	ctx := context.Background()

	s := conductoronesdkgo.New(
		conductoronesdkgo.WithSecurity(shared.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "<YOUR_OAUTH_HERE>",
		}),
	)

	res, err := s.Apps.Create(ctx, nil, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.CreateAppResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"github.com/conductorone/conductorone-sdk-go/pkg/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := conductoronesdkgo.New(
		conductoronesdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		conductoronesdkgo.WithSecurity(shared.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "<YOUR_OAUTH_HERE>",
		}),
	)

	res, err := s.Apps.Create(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.CreateAppResponse != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
