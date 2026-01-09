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

### [AccessReviewTemplate](docs/sdks/accessreviewtemplate/README.md)

* [Create](docs/sdks/accessreviewtemplate/README.md#create) - Create
* [Delete](docs/sdks/accessreviewtemplate/README.md#delete) - Delete
* [Get](docs/sdks/accessreviewtemplate/README.md#get) - Get
* [Update](docs/sdks/accessreviewtemplate/README.md#update) - Update

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

* [CreateAutomation](docs/sdks/automation/README.md#createautomation) - Create Automation
* [DeleteAutomation](docs/sdks/automation/README.md#deleteautomation) - Delete Automation
* [ExecuteAutomation](docs/sdks/automation/README.md#executeautomation) - Execute Automation
* [GetAutomation](docs/sdks/automation/README.md#getautomation) - Get Automation
* [ListAutomations](docs/sdks/automation/README.md#listautomations) - List Automations
* [UpdateAutomation](docs/sdks/automation/README.md#updateautomation) - Update Automation

### [AutomationExecution](docs/sdks/automationexecution/README.md)

* [GetAutomationExecution](docs/sdks/automationexecution/README.md#getautomationexecution) - Get Automation Execution
* [ListAutomationExecutions](docs/sdks/automationexecution/README.md#listautomationexecutions) - List Automation Executions

### [AutomationExecutionActions](docs/sdks/automationexecutionactions/README.md)

* [TerminateAutomation](docs/sdks/automationexecutionactions/README.md#terminateautomation) - Terminate Automation

### [AutomationExecutionSearch](docs/sdks/automationexecutionsearch/README.md)

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
* [UpdateDelegated](docs/sdks/connector/README.md#updatedelegated) - Update Delegated
* [ValidateHTTPConnectorConfig](docs/sdks/connector/README.md#validatehttpconnectorconfig) - Validate Http Connector Config

### [ConnectorCatalog](docs/sdks/connectorcatalog/README.md)

* [ConfigurationSchema](docs/sdks/connectorcatalog/README.md#configurationschema) - Configuration Schema

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

### [Functions](docs/sdks/functions/README.md)

* [CreateFunction](docs/sdks/functions/README.md#createfunction) - Create Function
* [CreateTag](docs/sdks/functions/README.md#createtag) - Create Tag
* [DeleteFunction](docs/sdks/functions/README.md#deletefunction) - Delete Function
* [GetFunction](docs/sdks/functions/README.md#getfunction) - Get Function
* [GetFunctionSecretEncryptionKey](docs/sdks/functions/README.md#getfunctionsecretencryptionkey) - Get Function Secret Encryption Key
* [Invoke](docs/sdks/functions/README.md#invoke) - Invoke
* [ListCommits](docs/sdks/functions/README.md#listcommits) - List Commits
* [ListFunctions](docs/sdks/functions/README.md#listfunctions) - List Functions
* [ListTags](docs/sdks/functions/README.md#listtags) - List Tags
* [UpdateFunction](docs/sdks/functions/README.md#updatefunction) - Update Function

### [FunctionsInvocation](docs/sdks/functionsinvocation/README.md)

* [Get](docs/sdks/functionsinvocation/README.md#get) - Get
* [List](docs/sdks/functionsinvocation/README.md#list) - List

### [FunctionsSearch](docs/sdks/functionssearch/README.md)

* [Search](docs/sdks/functionssearch/README.md#search) - Search

### [OrgDomain](docs/sdks/orgdomain/README.md)

* [List](docs/sdks/orgdomain/README.md#list) - List
* [Update](docs/sdks/orgdomain/README.md#update) - Update

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

### [Roles](docs/sdks/roles/README.md)

* [Get](docs/sdks/roles/README.md#get) - Get
* [List](docs/sdks/roles/README.md#list) - List
* [Update](docs/sdks/roles/README.md#update) - Update

### [SessionSettings](docs/sdks/sessionsettings/README.md)

* [Get](docs/sdks/sessionsettings/README.md#get) - Get
* [TestSourceIP](docs/sdks/sessionsettings/README.md#testsourceip) - Test Source Ip
* [Update](docs/sdks/sessionsettings/README.md#update) - Update

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

### [User](docs/sdks/user/README.md)

* [Get](docs/sdks/user/README.md#get) - Get
* [GetUserProfileTypes](docs/sdks/user/README.md#getuserprofiletypes) - Get User Profile Types
* [List](docs/sdks/user/README.md#list) - List
* [SetExpiringUserDelegationBindingByAdmin](docs/sdks/user/README.md#setexpiringuserdelegationbindingbyadmin) - Set Expiring User Delegation Binding By Admin

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

	res, err := s.AccessReview.Create(ctx, nil, operations.WithRetries(
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
	if res.AccessReviewServiceCreateResponse != nil {
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

	res, err := s.AccessReview.Create(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.AccessReviewServiceCreateResponse != nil {
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
