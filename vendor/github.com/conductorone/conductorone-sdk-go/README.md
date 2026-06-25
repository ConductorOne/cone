# conductorone-api

<!-- Start Summary [summary] -->
## Summary

C1 API: The C1 API is a HTTP API for managing C1 resources.
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
  * [Pagination](#pagination)
  * [Retries](#retries)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/conductorone/conductorone-sdk-go
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

### [A2Ui](docs/sdks/a2ui/README.md)

* [CreateSurfaceFeedback](docs/sdks/a2ui/README.md#createsurfacefeedback) - Create Surface Feedback
* [ListSurfaceFeedback](docs/sdks/a2ui/README.md#listsurfacefeedback) - List Surface Feedback
* [ListSurfaces](docs/sdks/a2ui/README.md#listsurfaces) - List Surfaces
* [SubmitAction](docs/sdks/a2ui/README.md#submitaction) - Submit Action

### [AccessConflict](docs/sdks/accessconflict/README.md)

* [CreateMonitor](docs/sdks/accessconflict/README.md#createmonitor) - Create Monitor
* [DeleteMonitor](docs/sdks/accessconflict/README.md#deletemonitor) - Delete Monitor
* [GetMonitor](docs/sdks/accessconflict/README.md#getmonitor) - Get Monitor
* [UpdateMonitor](docs/sdks/accessconflict/README.md#updatemonitor) - Update Monitor

### [AccessReview](docs/sdks/accessreview/README.md)

* [Create](docs/sdks/accessreview/README.md#create) - Create
* [Delete](docs/sdks/accessreview/README.md#delete) - Delete
* [Get](docs/sdks/accessreview/README.md#get) - Get
* [List](docs/sdks/accessreview/README.md#list) - List
* [Update](docs/sdks/accessreview/README.md#update) - Update

### [AccessReviewSetupEntitlement](docs/sdks/accessreviewsetupentitlement/README.md)

* [GetCampaignScopeAndEntitlements](docs/sdks/accessreviewsetupentitlement/README.md#getcampaignscopeandentitlements) - Get Campaign Scope And Entitlements
* [SetCampaignScopeAndEntitlements](docs/sdks/accessreviewsetupentitlement/README.md#setcampaignscopeandentitlements) - Set Campaign Scope And Entitlements
* [SetCampaignScopeByResourceType](docs/sdks/accessreviewsetupentitlement/README.md#setcampaignscopebyresourcetype) - Set Campaign Scope By Resource Type

### [AccessReviewTemplate](docs/sdks/accessreviewtemplate/README.md)

* [Create](docs/sdks/accessreviewtemplate/README.md#create) - Create
* [Delete](docs/sdks/accessreviewtemplate/README.md#delete) - Delete
* [Get](docs/sdks/accessreviewtemplate/README.md#get) - Get
* [Update](docs/sdks/accessreviewtemplate/README.md#update) - Update

### [AccessReviewTemplateSetupEntitlement](docs/sdks/accessreviewtemplatesetupentitlement/README.md)

* [GetScopeAndEntitlements](docs/sdks/accessreviewtemplatesetupentitlement/README.md#getscopeandentitlements) - Get Scope And Entitlements
* [SetScopeAndEntitlements](docs/sdks/accessreviewtemplatesetupentitlement/README.md#setscopeandentitlements) - Set Scope And Entitlements
* [SetScopeByResourceType](docs/sdks/accessreviewtemplatesetupentitlement/README.md#setscopebyresourcetype) - Set Scope By Resource Type

### [AccountProvisionPolicyTest](docs/sdks/accountprovisionpolicytest/README.md)

* [Test](docs/sdks/accountprovisionpolicytest/README.md#test) - Test

### [AppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md)

* [CancelAppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md#cancelappaccessrequestsdefaults) - Cancel App Access Requests Defaults
* [CreateAppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md#createappaccessrequestsdefaults) - Create App Access Requests Defaults
* [GetAppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md#getappaccessrequestsdefaults) - Get App Access Requests Defaults

### [AppEntitlementMonitorBinding](docs/sdks/appentitlementmonitorbinding/README.md)

* [CreateAppEntitlementMonitorBinding](docs/sdks/appentitlementmonitorbinding/README.md#createappentitlementmonitorbinding) - Create App Entitlement Monitor Binding
* [DeleteAppEntitlementMonitorBinding](docs/sdks/appentitlementmonitorbinding/README.md#deleteappentitlementmonitorbinding) - Delete App Entitlement Monitor Binding
* [GetAppEntitlementMonitorBinding](docs/sdks/appentitlementmonitorbinding/README.md#getappentitlementmonitorbinding) - Get App Entitlement Monitor Binding

### [AppEntitlementOwners](docs/sdks/appentitlementowners/README.md)

* [Add](docs/sdks/appentitlementowners/README.md#add) - Add
* [Delete](docs/sdks/appentitlementowners/README.md#delete) - Delete
* [List](docs/sdks/appentitlementowners/README.md#list) - List
* [ListOwnerIDs](docs/sdks/appentitlementowners/README.md#listownerids) - List Owner I Ds
* [Remove](docs/sdks/appentitlementowners/README.md#remove) - Remove
* [Set](docs/sdks/appentitlementowners/README.md#set) - Set

### [AppEntitlementOwnersV2](docs/sdks/appentitlementownersv2/README.md)

* [SearchEntitlementOwners](docs/sdks/appentitlementownersv2/README.md#searchentitlementowners) - Search Entitlement Owners
* [SearchUserOwners](docs/sdks/appentitlementownersv2/README.md#searchuserowners) - Search User Owners
* [Set](docs/sdks/appentitlementownersv2/README.md#set) - Set

### [AppEntitlements](docs/sdks/appentitlements/README.md)

* [AddAutomationExclusion](docs/sdks/appentitlements/README.md#addautomationexclusion) - Add Automation Exclusion
* [AddManuallyManagedMembers](docs/sdks/appentitlements/README.md#addmanuallymanagedmembers) - Add Manually Managed Members
* [Create](docs/sdks/appentitlements/README.md#create) - Create
* [CreateAutomation](docs/sdks/appentitlements/README.md#createautomation) - Create Automation
* [Delete](docs/sdks/appentitlements/README.md#delete) - Delete
* [DeleteAutomation](docs/sdks/appentitlements/README.md#deleteautomation) - Delete Automation
* [Get](docs/sdks/appentitlements/README.md#get) - Get
* [GetAutomation](docs/sdks/appentitlements/README.md#getautomation) - Get Automation
* [List](docs/sdks/appentitlements/README.md#list) - List
* [ListAutomationExclusions](docs/sdks/appentitlements/README.md#listautomationexclusions) - List Automation Exclusions
* [ListForAppResource](docs/sdks/appentitlements/README.md#listforappresource) - List For App Resource
* [ListForAppUser](docs/sdks/appentitlements/README.md#listforappuser) - List For App User
* [~~ListUsers~~](docs/sdks/appentitlements/README.md#listusers) - List Users :warning: **Deprecated**
* [RemoveAutomationExclusion](docs/sdks/appentitlements/README.md#removeautomationexclusion) - Remove Automation Exclusion
* [RemoveEntitlementMembership](docs/sdks/appentitlements/README.md#removeentitlementmembership) - Remove Entitlement Membership
* [Update](docs/sdks/appentitlements/README.md#update) - Update
* [UpdateAutomation](docs/sdks/appentitlements/README.md#updateautomation) - Update Automation

### [AppEntitlementSearch](docs/sdks/appentitlementsearch/README.md)

* [Search](docs/sdks/appentitlementsearch/README.md#search) - Search
* [SearchAppEntitlementsForAppUser](docs/sdks/appentitlementsearch/README.md#searchappentitlementsforappuser) - Search App Entitlements For App User
* [SearchAppEntitlementsWithExpired](docs/sdks/appentitlementsearch/README.md#searchappentitlementswithexpired) - Search App Entitlements With Expired
* [SearchGrants](docs/sdks/appentitlementsearch/README.md#searchgrants) - Search Grants

### [AppEntitlementsProxy](docs/sdks/appentitlementsproxy/README.md)

* [Create](docs/sdks/appentitlementsproxy/README.md#create) - Create
* [Delete](docs/sdks/appentitlementsproxy/README.md#delete) - Delete
* [Get](docs/sdks/appentitlementsproxy/README.md#get) - Get

### [AppEntitlementUserBinding](docs/sdks/appentitlementuserbinding/README.md)

* [ListAppUsersForIdentityWithGrant](docs/sdks/appentitlementuserbinding/README.md#listappusersforidentitywithgrant) - List App Users For Identity With Grant
* [RemoveGrantDuration](docs/sdks/appentitlementuserbinding/README.md#removegrantduration) - Remove Grant Duration
* [SearchGrantFeed](docs/sdks/appentitlementuserbinding/README.md#searchgrantfeed) - Search Grant Feed
* [SearchPastGrants](docs/sdks/appentitlementuserbinding/README.md#searchpastgrants) - Search Past Grants
* [UpdateGrantDuration](docs/sdks/appentitlementuserbinding/README.md#updategrantduration) - Update Grant Duration

### [AppOwners](docs/sdks/appowners/README.md)

* [Add](docs/sdks/appowners/README.md#add) - Add
* [Delete](docs/sdks/appowners/README.md#delete) - Delete
* [List](docs/sdks/appowners/README.md#list) - List
* [ListOwnerIDs](docs/sdks/appowners/README.md#listownerids) - List Owner I Ds
* [Remove](docs/sdks/appowners/README.md#remove) - Remove
* [Set](docs/sdks/appowners/README.md#set) - Set

### [AppOwnersV2](docs/sdks/appownersv2/README.md)

* [CreateEntitlementOwner](docs/sdks/appownersv2/README.md#createentitlementowner) - Create Entitlement Owner
* [CreateUserOwner](docs/sdks/appownersv2/README.md#createuserowner) - Create User Owner
* [DeleteEntitlementOwner](docs/sdks/appownersv2/README.md#deleteentitlementowner) - Delete Entitlement Owner
* [DeleteUserOwner](docs/sdks/appownersv2/README.md#deleteuserowner) - Delete User Owner
* [GetEntitlementOwner](docs/sdks/appownersv2/README.md#getentitlementowner) - Get Entitlement Owner
* [GetUserOwner](docs/sdks/appownersv2/README.md#getuserowner) - Get User Owner
* [SearchEntitlementOwners](docs/sdks/appownersv2/README.md#searchentitlementowners) - Search Entitlement Owners
* [SearchUserOwners](docs/sdks/appownersv2/README.md#searchuserowners) - Search User Owners
* [Set](docs/sdks/appownersv2/README.md#set) - Set

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
* [Delete](docs/sdks/appresourceowners/README.md#delete) - Delete
* [List](docs/sdks/appresourceowners/README.md#list) - List
* [ListOwnerIDs](docs/sdks/appresourceowners/README.md#listownerids) - List Owner I Ds
* [Remove](docs/sdks/appresourceowners/README.md#remove) - Remove
* [Set](docs/sdks/appresourceowners/README.md#set) - Set

### [AppResourceSearch](docs/sdks/appresourcesearch/README.md)

* [SearchAppResourceTypes](docs/sdks/appresourcesearch/README.md#searchappresourcetypes) - Search App Resource Types
* [SearchAppResources](docs/sdks/appresourcesearch/README.md#searchappresources) - Search App Resources

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
* [SearchUserOwnership](docs/sdks/appsearch/README.md#searchuserownership) - Search User Ownership

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
* [CreateComplianceFrameworkAttributeValue](docs/sdks/attributes/README.md#createcomplianceframeworkattributevalue) - Create Compliance Framework Attribute Value
* [CreateRiskLevelAttributeValue](docs/sdks/attributes/README.md#createrisklevelattributevalue) - Create Risk Level Attribute Value
* [DeleteAttributeValue](docs/sdks/attributes/README.md#deleteattributevalue) - Delete Attribute Value
* [DeleteComplianceFrameworkAttributeValue](docs/sdks/attributes/README.md#deletecomplianceframeworkattributevalue) - Delete Compliance Framework Attribute Value
* [DeleteRiskLevelAttributeValue](docs/sdks/attributes/README.md#deleterisklevelattributevalue) - Delete Risk Level Attribute Value
* [GetAttributeValue](docs/sdks/attributes/README.md#getattributevalue) - Get Attribute Value
* [GetComplianceFrameworkAttributeValue](docs/sdks/attributes/README.md#getcomplianceframeworkattributevalue) - Get Compliance Framework Attribute Value
* [GetRiskLevelAttributeValue](docs/sdks/attributes/README.md#getrisklevelattributevalue) - Get Risk Level Attribute Value
* [ListAttributeTypes](docs/sdks/attributes/README.md#listattributetypes) - List Attribute Types
* [ListAttributeValues](docs/sdks/attributes/README.md#listattributevalues) - List Attribute Values
* [ListComplianceFrameworks](docs/sdks/attributes/README.md#listcomplianceframeworks) - List Compliance Frameworks
* [ListRiskLevels](docs/sdks/attributes/README.md#listrisklevels) - List Risk Levels

### [AttributeSearch](docs/sdks/attributesearch/README.md)

* [SearchAttributeValues](docs/sdks/attributesearch/README.md#searchattributevalues) - Search Attribute Values

### [Auth](docs/sdks/auth/README.md)

* [Introspect](docs/sdks/auth/README.md#introspect) - Introspect

### [Automation](docs/sdks/automation/README.md)

* [ClearAutomationCircuitBreaker](docs/sdks/automation/README.md#clearautomationcircuitbreaker) - Clear Automation Circuit Breaker
* [CreateAutomation](docs/sdks/automation/README.md#createautomation) - Create Automation
* [DeleteAutomation](docs/sdks/automation/README.md#deleteautomation) - Delete Automation
* [ExecuteAutomation](docs/sdks/automation/README.md#executeautomation) - Execute Automation
* [GetAutomation](docs/sdks/automation/README.md#getautomation) - Get Automation
* [ListAutomations](docs/sdks/automation/README.md#listautomations) - List Automations
* [ResolvePausedAutomationExecutions](docs/sdks/automation/README.md#resolvepausedautomationexecutions) - Resolve Paused Automation Executions
* [UpdateAutomation](docs/sdks/automation/README.md#updateautomation) - Update Automation

### [AutomationExecution](docs/sdks/automationexecution/README.md)

* [GetAutomationExecution](docs/sdks/automationexecution/README.md#getautomationexecution) - Get Automation Execution
* [ListAutomationExecutions](docs/sdks/automationexecution/README.md#listautomationexecutions) - List Automation Executions

### [AutomationExecutionActions](docs/sdks/automationexecutionactions/README.md)

* [TerminateAutomation](docs/sdks/automationexecutionactions/README.md#terminateautomation) - Terminate Automation

### [AutomationExecutionSearch](docs/sdks/automationexecutionsearch/README.md)

* [SearchAllAutomationExecutions](docs/sdks/automationexecutionsearch/README.md#searchallautomationexecutions) - Search All Automation Executions
* [SearchAutomationExecutions](docs/sdks/automationexecutionsearch/README.md#searchautomationexecutions) - Search Automation Executions

### [AutomationSearch](docs/sdks/automationsearch/README.md)

* [SearchAutomationTemplateVersions](docs/sdks/automationsearch/README.md#searchautomationtemplateversions) - Search Automation Template Versions
* [SearchAutomations](docs/sdks/automationsearch/README.md#searchautomations) - Search Automations

### [AWSExternalIDSettings](docs/sdks/awsexternalidsettings/README.md)

* [Get](docs/sdks/awsexternalidsettings/README.md#get) - Get

### [Connector](docs/sdks/connector/README.md)

* [ConfirmSyncValid](docs/sdks/connector/README.md#confirmsyncvalid) - Confirm Sync Valid
* [Create](docs/sdks/connector/README.md#create) - Create
* [CreateDelegated](docs/sdks/connector/README.md#createdelegated) - Create Delegated
* [Delete](docs/sdks/connector/README.md#delete) - Delete
* [ForceSync](docs/sdks/connector/README.md#forcesync) - Force Sync
* [Get](docs/sdks/connector/README.md#get) - Get
* [GetConnectorSyncDownloadURL](docs/sdks/connector/README.md#getconnectorsyncdownloadurl) - Get Connector Sync Download Url
* [GetCredentials](docs/sdks/connector/README.md#getcredentials) - Get Credentials
* [List](docs/sdks/connector/README.md#list) - List
* [PauseSync](docs/sdks/connector/README.md#pausesync) - Pause Sync
* [ResumeSync](docs/sdks/connector/README.md#resumesync) - Resume Sync
* [RevokeCredential](docs/sdks/connector/README.md#revokecredential) - Revoke Credential
* [RotateCredential](docs/sdks/connector/README.md#rotatecredential) - Rotate Credential
* [Update](docs/sdks/connector/README.md#update) - Update
* [UpdateConnectorSchedule](docs/sdks/connector/README.md#updateconnectorschedule) - Update Connector Schedule
* [UpdateDelegated](docs/sdks/connector/README.md#updatedelegated) - Update Delegated
* [ValidateHTTPConnectorConfig](docs/sdks/connector/README.md#validatehttpconnectorconfig) - Validate Http Connector Config

### [ConnectorCatalog](docs/sdks/connectorcatalog/README.md)

* [ConfigurationSchema](docs/sdks/connectorcatalog/README.md#configurationschema) - Configuration Schema

### [ConnectorOwnersV2](docs/sdks/connectorownersv2/README.md)

* [SearchEntitlementOwners](docs/sdks/connectorownersv2/README.md#searchentitlementowners) - Search Entitlement Owners
* [SearchUserOwners](docs/sdks/connectorownersv2/README.md#searchuserowners) - Search User Owners
* [Set](docs/sdks/connectorownersv2/README.md#set) - Set

### [Contacts](docs/sdks/contacts/README.md)

* [GetContacts](docs/sdks/contacts/README.md#getcontacts) - Get Contacts
* [UpdateContacts](docs/sdks/contacts/README.md#updatecontacts) - Update Contacts

### [Directory](docs/sdks/directory/README.md)

* [Create](docs/sdks/directory/README.md#create) - Create
* [Delete](docs/sdks/directory/README.md#delete) - Delete
* [Get](docs/sdks/directory/README.md#get) - Get
* [List](docs/sdks/directory/README.md#list) - List
* [Update](docs/sdks/directory/README.md#update) - Update

### [Export](docs/sdks/export/README.md)

* [Create](docs/sdks/export/README.md#create) - Create
* [Delete](docs/sdks/export/README.md#delete) - Delete
* [Get](docs/sdks/export/README.md#get) - Get
* [List](docs/sdks/export/README.md#list) - List
* [ListEvents](docs/sdks/export/README.md#listevents) - List Events
* [Update](docs/sdks/export/README.md#update) - Update

### [ExportsSearch](docs/sdks/exportssearch/README.md)

* [Search](docs/sdks/exportssearch/README.md#search) - Search

### [ExternalClientSearch](docs/sdks/externalclientsearch/README.md)

* [Search](docs/sdks/externalclientsearch/README.md#search) - NOTE: Searches external client grants for all users

### [Finding](docs/sdks/finding/README.md)

* [BulkCreateFindingTasks](docs/sdks/finding/README.md#bulkcreatefindingtasks) - Bulk Create Finding Tasks
* [BulkUpdateFindingState](docs/sdks/finding/README.md#bulkupdatefindingstate) - Bulk Update Finding State
* [CreateFindingTask](docs/sdks/finding/README.md#createfindingtask) - Create Finding Task
* [GetFinding](docs/sdks/finding/README.md#getfinding) - Get Finding
* [UpdateFindingState](docs/sdks/finding/README.md#updatefindingstate) - Update Finding State

### [FindingRoutingRule](docs/sdks/findingroutingrule/README.md)

* [CreateFindingRoutingRule](docs/sdks/findingroutingrule/README.md#createfindingroutingrule) - Create Finding Routing Rule
* [DeleteFindingRoutingRule](docs/sdks/findingroutingrule/README.md#deletefindingroutingrule) - Delete Finding Routing Rule
* [GetFindingRoutingRule](docs/sdks/findingroutingrule/README.md#getfindingroutingrule) - Get Finding Routing Rule
* [ListFindingRoutingRules](docs/sdks/findingroutingrule/README.md#listfindingroutingrules) - List Finding Routing Rules
* [UpdateFindingRoutingRule](docs/sdks/findingroutingrule/README.md#updatefindingroutingrule) - Update Finding Routing Rule

### [FindingSearch](docs/sdks/findingsearch/README.md)

* [Search](docs/sdks/findingsearch/README.md#search) - Search

### [Functions](docs/sdks/functions/README.md)

* [CreateFinalCommit](docs/sdks/functions/README.md#createfinalcommit) - Create Final Commit
* [CreateFunction](docs/sdks/functions/README.md#createfunction) - Create Function
* [CreateInitialCommit](docs/sdks/functions/README.md#createinitialcommit) - Create Initial Commit
* [CreateTag](docs/sdks/functions/README.md#createtag) - Create Tag
* [DeleteFunction](docs/sdks/functions/README.md#deletefunction) - Delete Function
* [GetCommitContent](docs/sdks/functions/README.md#getcommitcontent) - Get Commit Content
* [GetFunction](docs/sdks/functions/README.md#getfunction) - Get Function
* [GetLockFile](docs/sdks/functions/README.md#getlockfile) - Get Lock File
* [Invoke](docs/sdks/functions/README.md#invoke) - Invoke
* [ListCommits](docs/sdks/functions/README.md#listcommits) - List Commits
* [ListFunctions](docs/sdks/functions/README.md#listfunctions) - List Functions
* [ListTags](docs/sdks/functions/README.md#listtags) - List Tags
* [Test](docs/sdks/functions/README.md#test) - Test
* [UpdateFunction](docs/sdks/functions/README.md#updatefunction) - Update Function

### [FunctionsInvocation](docs/sdks/functionsinvocation/README.md)

* [Get](docs/sdks/functionsinvocation/README.md#get) - Get
* [List](docs/sdks/functionsinvocation/README.md#list) - List

### [FunctionsInvocationSearch](docs/sdks/functionsinvocationsearch/README.md)

* [Search](docs/sdks/functionsinvocationsearch/README.md#search) - Search

### [FunctionsSearch](docs/sdks/functionssearch/README.md)

* [Search](docs/sdks/functionssearch/README.md#search) - Search

### [Hooks](docs/sdks/hooks/README.md)

* [Create](docs/sdks/hooks/README.md#create) - Create
* [Delete](docs/sdks/hooks/README.md#delete) - Delete
* [Get](docs/sdks/hooks/README.md#get) - Get
* [List](docs/sdks/hooks/README.md#list) - List
* [Update](docs/sdks/hooks/README.md#update) - Update

### [HooksSearch](docs/sdks/hookssearch/README.md)

* [Search](docs/sdks/hookssearch/README.md#search) - Search

### [LocalDirectoryConfig](docs/sdks/localdirectoryconfig/README.md)

* [Create](docs/sdks/localdirectoryconfig/README.md#create) - Create
* [Delete](docs/sdks/localdirectoryconfig/README.md#delete) - Delete
* [Get](docs/sdks/localdirectoryconfig/README.md#get) - Get
* [List](docs/sdks/localdirectoryconfig/README.md#list) - List
* [Update](docs/sdks/localdirectoryconfig/README.md#update) - Update

### [LocalUserInvitation](docs/sdks/localuserinvitation/README.md)

* [Create](docs/sdks/localuserinvitation/README.md#create) - Create
* [Get](docs/sdks/localuserinvitation/README.md#get) - Get
* [Revoke](docs/sdks/localuserinvitation/README.md#revoke) - Revoke
* [Search](docs/sdks/localuserinvitation/README.md#search) - Search

### [OnboardingSettings](docs/sdks/onboardingsettings/README.md)

* [Get](docs/sdks/onboardingsettings/README.md#get) - Get
* [Update](docs/sdks/onboardingsettings/README.md#update) - Update

### [OrgDomain](docs/sdks/orgdomain/README.md)

* [List](docs/sdks/orgdomain/README.md#list) - List
* [Update](docs/sdks/orgdomain/README.md#update) - Update

### [OrgNotificationSettings](docs/sdks/orgnotificationsettings/README.md)

* [Get](docs/sdks/orgnotificationsettings/README.md#get) - Get
* [Update](docs/sdks/orgnotificationsettings/README.md#update) - Update

### [PaperSecret](docs/sdks/papersecret/README.md)

* [CreateExternal](docs/sdks/papersecret/README.md#createexternal) - Create External
* [CreateInternal](docs/sdks/papersecret/README.md#createinternal) - Create Internal
* [Get](docs/sdks/papersecret/README.md#get) - Get
* [GetByShareCode](docs/sdks/papersecret/README.md#getbysharecode) - Get By Share Code
* [GetContent](docs/sdks/papersecret/README.md#getcontent) - Get Content
* [Revoke](docs/sdks/papersecret/README.md#revoke) - Revoke
* [SearchAuditEvents](docs/sdks/papersecret/README.md#searchauditevents) - Search Audit Events
* [SearchMySecrets](docs/sdks/papersecret/README.md#searchmysecrets) - Search My Secrets
* [SetTextContent](docs/sdks/papersecret/README.md#settextcontent) - Set Text Content

### [PaperSecretAdmin](docs/sdks/papersecretadmin/README.md)

* [Get](docs/sdks/papersecretadmin/README.md#get) - Get
* [Revoke](docs/sdks/papersecretadmin/README.md#revoke) - Revoke
* [Search](docs/sdks/papersecretadmin/README.md#search) - Search
* [SearchAuditEvents](docs/sdks/papersecretadmin/README.md#searchauditevents) - Search Audit Events

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

### [Principal](docs/sdks/principal/README.md)

* [AddBinding](docs/sdks/principal/README.md#addbinding) - Add Binding
* [Create](docs/sdks/principal/README.md#create) - Create
* [CreateCredential](docs/sdks/principal/README.md#createcredential) - Create Credential
* [Delete](docs/sdks/principal/README.md#delete) - Delete
* [DeleteBinding](docs/sdks/principal/README.md#deletebinding) - Delete Binding
* [Get](docs/sdks/principal/README.md#get) - Get
* [GetCredential](docs/sdks/principal/README.md#getcredential) - Get Credential
* [List](docs/sdks/principal/README.md#list) - List
* [ListBindings](docs/sdks/principal/README.md#listbindings) - List Bindings
* [ListCredentials](docs/sdks/principal/README.md#listcredentials) - List Credentials
* [RevokeCredential](docs/sdks/principal/README.md#revokecredential) - Revoke Credential
* [Update](docs/sdks/principal/README.md#update) - Update
* [UpdateCredential](docs/sdks/principal/README.md#updatecredential) - Update Credential

### [RequestCatalogManagement](docs/sdks/requestcatalogmanagement/README.md)

* [AddAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#addaccessentitlements) - Add Access Entitlements
* [AddAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#addappentitlements) - Add App Entitlements
* [Create](docs/sdks/requestcatalogmanagement/README.md#create) - Create
* [CreateBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#createbundleautomation) - Create Bundle Automation
* [CreateRequestableEntry](docs/sdks/requestcatalogmanagement/README.md#createrequestableentry) - Create Requestable Entry
* [Delete](docs/sdks/requestcatalogmanagement/README.md#delete) - Delete
* [DeleteBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#deletebundleautomation) - Delete Bundle Automation
* [DeleteRequestableEntry](docs/sdks/requestcatalogmanagement/README.md#deleterequestableentry) - Delete Requestable Entry
* [ForceRunBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#forcerunbundleautomation) - Force Run Bundle Automation
* [Get](docs/sdks/requestcatalogmanagement/README.md#get) - Get
* [GetBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#getbundleautomation) - Get Bundle Automation
* [GetRequestableEntry](docs/sdks/requestcatalogmanagement/README.md#getrequestableentry) - Get Requestable Entry
* [List](docs/sdks/requestcatalogmanagement/README.md#list) - List
* [ListAllEntitlementIdsPerApp](docs/sdks/requestcatalogmanagement/README.md#listallentitlementidsperapp) - List All Entitlement Ids Per App
* [ListEntitlementsForAccess](docs/sdks/requestcatalogmanagement/README.md#listentitlementsforaccess) - List Entitlements For Access
* [ListEntitlementsPerCatalog](docs/sdks/requestcatalogmanagement/README.md#listentitlementspercatalog) - List Entitlements Per Catalog
* [RemoveAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeaccessentitlements) - Remove Access Entitlements
* [RemoveAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeappentitlements) - Remove App Entitlements
* [ResumePausedBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#resumepausedbundleautomation) - Resume Paused Bundle Automation
* [SetBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#setbundleautomation) - Set Bundle Automation
* [Update](docs/sdks/requestcatalogmanagement/README.md#update) - Update
* [UpdateAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#updateappentitlements) - Update App Entitlements

### [RequestCatalogSearch](docs/sdks/requestcatalogsearch/README.md)

* [SearchEntitlements](docs/sdks/requestcatalogsearch/README.md#searchentitlements) - Search Entitlements

### [RequestSchema](docs/sdks/requestschema/README.md)

* [Create](docs/sdks/requestschema/README.md#create) - Create
* [CreateEntitlementBinding](docs/sdks/requestschema/README.md#createentitlementbinding) - Create Entitlement Binding
* [Delete](docs/sdks/requestschema/README.md#delete) - Delete
* [FindBindingForAppEntitlement](docs/sdks/requestschema/README.md#findbindingforappentitlement) - Find Binding For App Entitlement
* [Get](docs/sdks/requestschema/README.md#get) - Get
* [RemoveEntitlementBinding](docs/sdks/requestschema/README.md#removeentitlementbinding) - Remove Entitlement Binding
* [Update](docs/sdks/requestschema/README.md#update) - Update

### [RoleMiningManagement](docs/sdks/roleminingmanagement/README.md)

* [CreateAccessProfileFromCohort](docs/sdks/roleminingmanagement/README.md#createaccessprofilefromcohort) - Create Access Profile From Cohort
* [GetCustomAnalysisResult](docs/sdks/roleminingmanagement/README.md#getcustomanalysisresult) - Get Custom Analysis Result
* [GetLatestRun](docs/sdks/roleminingmanagement/README.md#getlatestrun) - Get Latest Run
* [GetRoleMiningConfig](docs/sdks/roleminingmanagement/README.md#getroleminingconfig) - Get Role Mining Config
* [GetSuggestion](docs/sdks/roleminingmanagement/README.md#getsuggestion) - Get Suggestion
* [ListRuns](docs/sdks/roleminingmanagement/README.md#listruns) - List Runs
* [ListSuggestions](docs/sdks/roleminingmanagement/README.md#listsuggestions) - List Suggestions
* [SearchCohortUsers](docs/sdks/roleminingmanagement/README.md#searchcohortusers) - Search Cohort Users
* [TriggerAnalysis](docs/sdks/roleminingmanagement/README.md#triggeranalysis) - Trigger Analysis
* [TriggerCustomAnalysis](docs/sdks/roleminingmanagement/README.md#triggercustomanalysis) - Trigger Custom Analysis
* [UpdateRoleMiningConfig](docs/sdks/roleminingmanagement/README.md#updateroleminingconfig) - Update Role Mining Config
* [UpdateSuggestionState](docs/sdks/roleminingmanagement/README.md#updatesuggestionstate) - Update Suggestion State

### [RoleMiningManagementSearch](docs/sdks/roleminingmanagementsearch/README.md)

* [Search](docs/sdks/roleminingmanagementsearch/README.md#search) - Search

### [Roles](docs/sdks/roles/README.md)

* [Get](docs/sdks/roles/README.md#get) - Get
* [List](docs/sdks/roles/README.md#list) - List
* [Update](docs/sdks/roles/README.md#update) - Update

### [SessionSettings](docs/sdks/sessionsettings/README.md)

* [Get](docs/sdks/sessionsettings/README.md#get) - Get
* [TestSourceIP](docs/sdks/sessionsettings/README.md#testsourceip) - Test Source Ip
* [Update](docs/sdks/sessionsettings/README.md#update) - Update

### [SSFReceiverEvent](docs/sdks/ssfreceiverevent/README.md)

* [List](docs/sdks/ssfreceiverevent/README.md#list) - List

### [SSFReceiverEventSearch](docs/sdks/ssfreceivereventsearch/README.md)

* [Search](docs/sdks/ssfreceivereventsearch/README.md#search) - Search

### [SSFReceiverStream](docs/sdks/ssfreceiverstream/README.md)

* [Create](docs/sdks/ssfreceiverstream/README.md#create) - Create
* [Delete](docs/sdks/ssfreceiverstream/README.md#delete) - Delete
* [Get](docs/sdks/ssfreceiverstream/README.md#get) - Get
* [GetStats](docs/sdks/ssfreceiverstream/README.md#getstats) - Get Stats
* [List](docs/sdks/ssfreceiverstream/README.md#list) - List
* [Test](docs/sdks/ssfreceiverstream/README.md#test) - Test
* [Update](docs/sdks/ssfreceiverstream/README.md#update) - Update

### [StepUpProvider](docs/sdks/stepupprovider/README.md)

* [Create](docs/sdks/stepupprovider/README.md#create) - Create
* [Delete](docs/sdks/stepupprovider/README.md#delete) - Delete
* [Get](docs/sdks/stepupprovider/README.md#get) - Get
* [List](docs/sdks/stepupprovider/README.md#list) - List
* [Search](docs/sdks/stepupprovider/README.md#search) - Search
* [Test](docs/sdks/stepupprovider/README.md#test) - Test
* [Update](docs/sdks/stepupprovider/README.md#update) - Update
* [UpdateSecret](docs/sdks/stepupprovider/README.md#updatesecret) - Update Secret

### [StepUpTransaction](docs/sdks/stepuptransaction/README.md)

* [Get](docs/sdks/stepuptransaction/README.md#get) - Get
* [Search](docs/sdks/stepuptransaction/README.md#search) - Search

### [SystemLog](docs/sdks/systemlog/README.md)

* [ListEvents](docs/sdks/systemlog/README.md#listevents) - List Events

### [Task](docs/sdks/task/README.md)

* [CreateGrantTask](docs/sdks/task/README.md#creategranttask) - Create Grant Task
* [CreateOffboardingTask](docs/sdks/task/README.md#createoffboardingtask) - Create Offboarding Task
* [CreateRevokeTask](docs/sdks/task/README.md#createrevoketask) - Create Revoke Task
* [Get](docs/sdks/task/README.md#get) - Get

### [TaskActions](docs/sdks/taskactions/README.md)

* [Approve](docs/sdks/taskactions/README.md#approve) - Approve
* [ApproveWithStepUp](docs/sdks/taskactions/README.md#approvewithstepup) - Approve With Step Up
* [Close](docs/sdks/taskactions/README.md#close) - Close
* [Comment](docs/sdks/taskactions/README.md#comment) - Comment
* [Deny](docs/sdks/taskactions/README.md#deny) - Deny
* [EscalateToEmergencyAccess](docs/sdks/taskactions/README.md#escalatetoemergencyaccess) - Escalate To Emergency Access
* [HardReset](docs/sdks/taskactions/README.md#hardreset) - Hard Reset
* [ProcessNow](docs/sdks/taskactions/README.md#processnow) - Process Now
* [Reassign](docs/sdks/taskactions/README.md#reassign) - Reassign
* [Restart](docs/sdks/taskactions/README.md#restart) - Restart
* [SkipStep](docs/sdks/taskactions/README.md#skipstep) - Skip Step
* [UpdateGrantDuration](docs/sdks/taskactions/README.md#updategrantduration) - Update Grant Duration
* [UpdateRequestData](docs/sdks/taskactions/README.md#updaterequestdata) - Update Request Data

### [TaskAudit](docs/sdks/taskaudit/README.md)

* [List](docs/sdks/taskaudit/README.md#list) - List

### [TaskSearch](docs/sdks/tasksearch/README.md)

* [Search](docs/sdks/tasksearch/README.md#search) - Search

### [TenantAuthConfig](docs/sdks/tenantauthconfig/README.md)

* [Create](docs/sdks/tenantauthconfig/README.md#create) - Create
* [Delete](docs/sdks/tenantauthconfig/README.md#delete) - Delete
* [Get](docs/sdks/tenantauthconfig/README.md#get) - Get
* [List](docs/sdks/tenantauthconfig/README.md#list) - List
* [Update](docs/sdks/tenantauthconfig/README.md#update) - Update

### [TenantEmailProvider](docs/sdks/tenantemailprovider/README.md)

* [Get](docs/sdks/tenantemailprovider/README.md#get) - Get
* [GetEmailCapabilities](docs/sdks/tenantemailprovider/README.md#getemailcapabilities) - Get Email Capabilities
* [SearchAuditEvents](docs/sdks/tenantemailprovider/README.md#searchauditevents) - Search Audit Events
* [Test](docs/sdks/tenantemailprovider/README.md#test) - Test
* [Update](docs/sdks/tenantemailprovider/README.md#update) - Update

### [User](docs/sdks/user/README.md)

* [Get](docs/sdks/user/README.md#get) - Get
* [GetUserProfileTypes](docs/sdks/user/README.md#getuserprofiletypes) - Get User Profile Types
* [List](docs/sdks/user/README.md#list) - List
* [SetExpiringUserDelegationBindingByAdmin](docs/sdks/user/README.md#setexpiringuserdelegationbindingbyadmin) - Set Expiring User Delegation Binding By Admin

### [UserNotificationSettings](docs/sdks/usernotificationsettings/README.md)

* [Get](docs/sdks/usernotificationsettings/README.md#get) - Get
* [Update](docs/sdks/usernotificationsettings/README.md#update) - Update

### [UserSearch](docs/sdks/usersearch/README.md)

* [Search](docs/sdks/usersearch/README.md#search) - Search

### [Vault](docs/sdks/vault/README.md)

* [Create](docs/sdks/vault/README.md#create) - Create
* [Delete](docs/sdks/vault/README.md#delete) - Delete
* [Get](docs/sdks/vault/README.md#get) - Get
* [Update](docs/sdks/vault/README.md#update) - Update

### [Webhooks](docs/sdks/webhooks/README.md)

* [Create](docs/sdks/webhooks/README.md#create) - Create
* [Delete](docs/sdks/webhooks/README.md#delete) - Delete
* [Get](docs/sdks/webhooks/README.md#get) - Get
* [List](docs/sdks/webhooks/README.md#list) - List
* [Test](docs/sdks/webhooks/README.md#test) - Test
* [Update](docs/sdks/webhooks/README.md#update) - Update

### [WebhooksSearch](docs/sdks/webhookssearch/README.md)

* [Search](docs/sdks/webhookssearch/README.md#search) - Search

### [WorkloadFederation](docs/sdks/workloadfederation/README.md)

* [CreateProvider](docs/sdks/workloadfederation/README.md#createprovider) - Create Provider
* [CreateTrust](docs/sdks/workloadfederation/README.md#createtrust) - Create Trust
* [DeleteProvider](docs/sdks/workloadfederation/README.md#deleteprovider) - Delete Provider
* [DeleteTrust](docs/sdks/workloadfederation/README.md#deletetrust) - Delete Trust
* [GetProvider](docs/sdks/workloadfederation/README.md#getprovider) - Get Provider
* [GetTrust](docs/sdks/workloadfederation/README.md#gettrust) - Get Trust
* [ListProviders](docs/sdks/workloadfederation/README.md#listproviders) - List Providers
* [ListTrusts](docs/sdks/workloadfederation/README.md#listtrusts) - List Trusts
* [SearchTrusts](docs/sdks/workloadfederation/README.md#searchtrusts) - Search Trusts
* [TestCEL](docs/sdks/workloadfederation/README.md#testcel) - Test Cel
* [TestToken](docs/sdks/workloadfederation/README.md#testtoken) - Test Token
* [UpdateProvider](docs/sdks/workloadfederation/README.md#updateprovider) - Update Provider
* [UpdateTrust](docs/sdks/workloadfederation/README.md#updatetrust) - Update Trust

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

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"log"
)

func main() {
	ctx := context.Background()

	s := conductoronesdkgo.New(
		conductoronesdkgo.WithSecurity(shared.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "<YOUR_OAUTH_HERE>",
		}),
	)

	res, err := s.AppEntitlementSearch.Search(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.AppEntitlementSearchServiceSearchResponse != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	conductoronesdkgo "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
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

	res, err := s.A2UI.CreateSurfaceFeedback(ctx, operations.C1APIA2uiV1A2UIServiceCreateSurfaceFeedbackRequest{
		SurfaceID: "<id>",
	}, operations.WithRetries(
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
	if res.A2UIServiceCreateSurfaceFeedbackResponse != nil {
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
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
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

	res, err := s.A2UI.CreateSurfaceFeedback(ctx, operations.C1APIA2uiV1A2UIServiceCreateSurfaceFeedbackRequest{
		SurfaceID: "<id>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.A2UIServiceCreateSurfaceFeedbackResponse != nil {
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
