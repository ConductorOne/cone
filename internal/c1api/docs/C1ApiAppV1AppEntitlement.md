# C1ApiAppV1AppEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | The alias field. | [optional] 
**AppId** | Pointer to **string** | The appId field. | [optional] 
**AppResourceId** | Pointer to **string** | The appResourceId field. | [optional] 
**AppResourceTypeId** | Pointer to **string** | The appResourceTypeId field. | [optional] 
**CertifyPolicyId** | Pointer to **string** | The certifyPolicyId field. | [optional] 
**ComplianceFrameworkValueIds** | Pointer to **[]string** | The complianceFrameworkValueIds field. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** | The description field. | [optional] 
**DisplayName** | Pointer to **string** | The displayName field. | [optional] 
**DurationGrant** | Pointer to **string** |  | [optional] 
**DurationUnset** | Pointer to **map[string]interface{}** |  | [optional] 
**GrantCount** | Pointer to **string** | The grantCount field. | [optional] 
**GrantPolicyId** | Pointer to **string** | The grantPolicyId field. | [optional] 
**Id** | Pointer to **string** | The id field. | [optional] 
**ProvisionerPolicy** | Pointer to [**C1ApiPolicyV1ProvisionPolicy**](C1ApiPolicyV1ProvisionPolicy.md) |  | [optional] 
**RevokePolicyId** | Pointer to **string** | The revokePolicyId field. | [optional] 
**RiskLevelValueId** | Pointer to **string** | The riskLevelValueId field. | [optional] 
**Slug** | Pointer to **string** | The slug field. | [optional] 
**SystemBuiltin** | Pointer to **bool** | The systemBuiltin field. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewC1ApiAppV1AppEntitlement

`func NewC1ApiAppV1AppEntitlement() *C1ApiAppV1AppEntitlement`

NewC1ApiAppV1AppEntitlement instantiates a new C1ApiAppV1AppEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiAppV1AppEntitlementWithDefaults

`func NewC1ApiAppV1AppEntitlementWithDefaults() *C1ApiAppV1AppEntitlement`

NewC1ApiAppV1AppEntitlementWithDefaults instantiates a new C1ApiAppV1AppEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *C1ApiAppV1AppEntitlement) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *C1ApiAppV1AppEntitlement) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *C1ApiAppV1AppEntitlement) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *C1ApiAppV1AppEntitlement) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAppId

`func (o *C1ApiAppV1AppEntitlement) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *C1ApiAppV1AppEntitlement) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *C1ApiAppV1AppEntitlement) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *C1ApiAppV1AppEntitlement) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppResourceId

`func (o *C1ApiAppV1AppEntitlement) GetAppResourceId() string`

GetAppResourceId returns the AppResourceId field if non-nil, zero value otherwise.

### GetAppResourceIdOk

`func (o *C1ApiAppV1AppEntitlement) GetAppResourceIdOk() (*string, bool)`

GetAppResourceIdOk returns a tuple with the AppResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppResourceId

`func (o *C1ApiAppV1AppEntitlement) SetAppResourceId(v string)`

SetAppResourceId sets AppResourceId field to given value.

### HasAppResourceId

`func (o *C1ApiAppV1AppEntitlement) HasAppResourceId() bool`

HasAppResourceId returns a boolean if a field has been set.

### GetAppResourceTypeId

`func (o *C1ApiAppV1AppEntitlement) GetAppResourceTypeId() string`

GetAppResourceTypeId returns the AppResourceTypeId field if non-nil, zero value otherwise.

### GetAppResourceTypeIdOk

`func (o *C1ApiAppV1AppEntitlement) GetAppResourceTypeIdOk() (*string, bool)`

GetAppResourceTypeIdOk returns a tuple with the AppResourceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppResourceTypeId

`func (o *C1ApiAppV1AppEntitlement) SetAppResourceTypeId(v string)`

SetAppResourceTypeId sets AppResourceTypeId field to given value.

### HasAppResourceTypeId

`func (o *C1ApiAppV1AppEntitlement) HasAppResourceTypeId() bool`

HasAppResourceTypeId returns a boolean if a field has been set.

### GetCertifyPolicyId

`func (o *C1ApiAppV1AppEntitlement) GetCertifyPolicyId() string`

GetCertifyPolicyId returns the CertifyPolicyId field if non-nil, zero value otherwise.

### GetCertifyPolicyIdOk

`func (o *C1ApiAppV1AppEntitlement) GetCertifyPolicyIdOk() (*string, bool)`

GetCertifyPolicyIdOk returns a tuple with the CertifyPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifyPolicyId

`func (o *C1ApiAppV1AppEntitlement) SetCertifyPolicyId(v string)`

SetCertifyPolicyId sets CertifyPolicyId field to given value.

### HasCertifyPolicyId

`func (o *C1ApiAppV1AppEntitlement) HasCertifyPolicyId() bool`

HasCertifyPolicyId returns a boolean if a field has been set.

### GetComplianceFrameworkValueIds

`func (o *C1ApiAppV1AppEntitlement) GetComplianceFrameworkValueIds() []string`

GetComplianceFrameworkValueIds returns the ComplianceFrameworkValueIds field if non-nil, zero value otherwise.

### GetComplianceFrameworkValueIdsOk

`func (o *C1ApiAppV1AppEntitlement) GetComplianceFrameworkValueIdsOk() (*[]string, bool)`

GetComplianceFrameworkValueIdsOk returns a tuple with the ComplianceFrameworkValueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceFrameworkValueIds

`func (o *C1ApiAppV1AppEntitlement) SetComplianceFrameworkValueIds(v []string)`

SetComplianceFrameworkValueIds sets ComplianceFrameworkValueIds field to given value.

### HasComplianceFrameworkValueIds

`func (o *C1ApiAppV1AppEntitlement) HasComplianceFrameworkValueIds() bool`

HasComplianceFrameworkValueIds returns a boolean if a field has been set.

### SetComplianceFrameworkValueIdsNil

`func (o *C1ApiAppV1AppEntitlement) SetComplianceFrameworkValueIdsNil(b bool)`

 SetComplianceFrameworkValueIdsNil sets the value for ComplianceFrameworkValueIds to be an explicit nil

### UnsetComplianceFrameworkValueIds
`func (o *C1ApiAppV1AppEntitlement) UnsetComplianceFrameworkValueIds()`

UnsetComplianceFrameworkValueIds ensures that no value is present for ComplianceFrameworkValueIds, not even an explicit nil
### GetCreatedAt

`func (o *C1ApiAppV1AppEntitlement) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *C1ApiAppV1AppEntitlement) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *C1ApiAppV1AppEntitlement) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *C1ApiAppV1AppEntitlement) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *C1ApiAppV1AppEntitlement) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *C1ApiAppV1AppEntitlement) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *C1ApiAppV1AppEntitlement) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *C1ApiAppV1AppEntitlement) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *C1ApiAppV1AppEntitlement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *C1ApiAppV1AppEntitlement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *C1ApiAppV1AppEntitlement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *C1ApiAppV1AppEntitlement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *C1ApiAppV1AppEntitlement) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *C1ApiAppV1AppEntitlement) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *C1ApiAppV1AppEntitlement) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *C1ApiAppV1AppEntitlement) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDurationGrant

`func (o *C1ApiAppV1AppEntitlement) GetDurationGrant() string`

GetDurationGrant returns the DurationGrant field if non-nil, zero value otherwise.

### GetDurationGrantOk

`func (o *C1ApiAppV1AppEntitlement) GetDurationGrantOk() (*string, bool)`

GetDurationGrantOk returns a tuple with the DurationGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationGrant

`func (o *C1ApiAppV1AppEntitlement) SetDurationGrant(v string)`

SetDurationGrant sets DurationGrant field to given value.

### HasDurationGrant

`func (o *C1ApiAppV1AppEntitlement) HasDurationGrant() bool`

HasDurationGrant returns a boolean if a field has been set.

### GetDurationUnset

`func (o *C1ApiAppV1AppEntitlement) GetDurationUnset() map[string]interface{}`

GetDurationUnset returns the DurationUnset field if non-nil, zero value otherwise.

### GetDurationUnsetOk

`func (o *C1ApiAppV1AppEntitlement) GetDurationUnsetOk() (*map[string]interface{}, bool)`

GetDurationUnsetOk returns a tuple with the DurationUnset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationUnset

`func (o *C1ApiAppV1AppEntitlement) SetDurationUnset(v map[string]interface{})`

SetDurationUnset sets DurationUnset field to given value.

### HasDurationUnset

`func (o *C1ApiAppV1AppEntitlement) HasDurationUnset() bool`

HasDurationUnset returns a boolean if a field has been set.

### SetDurationUnsetNil

`func (o *C1ApiAppV1AppEntitlement) SetDurationUnsetNil(b bool)`

 SetDurationUnsetNil sets the value for DurationUnset to be an explicit nil

### UnsetDurationUnset
`func (o *C1ApiAppV1AppEntitlement) UnsetDurationUnset()`

UnsetDurationUnset ensures that no value is present for DurationUnset, not even an explicit nil
### GetGrantCount

`func (o *C1ApiAppV1AppEntitlement) GetGrantCount() string`

GetGrantCount returns the GrantCount field if non-nil, zero value otherwise.

### GetGrantCountOk

`func (o *C1ApiAppV1AppEntitlement) GetGrantCountOk() (*string, bool)`

GetGrantCountOk returns a tuple with the GrantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantCount

`func (o *C1ApiAppV1AppEntitlement) SetGrantCount(v string)`

SetGrantCount sets GrantCount field to given value.

### HasGrantCount

`func (o *C1ApiAppV1AppEntitlement) HasGrantCount() bool`

HasGrantCount returns a boolean if a field has been set.

### GetGrantPolicyId

`func (o *C1ApiAppV1AppEntitlement) GetGrantPolicyId() string`

GetGrantPolicyId returns the GrantPolicyId field if non-nil, zero value otherwise.

### GetGrantPolicyIdOk

`func (o *C1ApiAppV1AppEntitlement) GetGrantPolicyIdOk() (*string, bool)`

GetGrantPolicyIdOk returns a tuple with the GrantPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantPolicyId

`func (o *C1ApiAppV1AppEntitlement) SetGrantPolicyId(v string)`

SetGrantPolicyId sets GrantPolicyId field to given value.

### HasGrantPolicyId

`func (o *C1ApiAppV1AppEntitlement) HasGrantPolicyId() bool`

HasGrantPolicyId returns a boolean if a field has been set.

### GetId

`func (o *C1ApiAppV1AppEntitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *C1ApiAppV1AppEntitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *C1ApiAppV1AppEntitlement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *C1ApiAppV1AppEntitlement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProvisionerPolicy

`func (o *C1ApiAppV1AppEntitlement) GetProvisionerPolicy() C1ApiPolicyV1ProvisionPolicy`

GetProvisionerPolicy returns the ProvisionerPolicy field if non-nil, zero value otherwise.

### GetProvisionerPolicyOk

`func (o *C1ApiAppV1AppEntitlement) GetProvisionerPolicyOk() (*C1ApiPolicyV1ProvisionPolicy, bool)`

GetProvisionerPolicyOk returns a tuple with the ProvisionerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionerPolicy

`func (o *C1ApiAppV1AppEntitlement) SetProvisionerPolicy(v C1ApiPolicyV1ProvisionPolicy)`

SetProvisionerPolicy sets ProvisionerPolicy field to given value.

### HasProvisionerPolicy

`func (o *C1ApiAppV1AppEntitlement) HasProvisionerPolicy() bool`

HasProvisionerPolicy returns a boolean if a field has been set.

### GetRevokePolicyId

`func (o *C1ApiAppV1AppEntitlement) GetRevokePolicyId() string`

GetRevokePolicyId returns the RevokePolicyId field if non-nil, zero value otherwise.

### GetRevokePolicyIdOk

`func (o *C1ApiAppV1AppEntitlement) GetRevokePolicyIdOk() (*string, bool)`

GetRevokePolicyIdOk returns a tuple with the RevokePolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokePolicyId

`func (o *C1ApiAppV1AppEntitlement) SetRevokePolicyId(v string)`

SetRevokePolicyId sets RevokePolicyId field to given value.

### HasRevokePolicyId

`func (o *C1ApiAppV1AppEntitlement) HasRevokePolicyId() bool`

HasRevokePolicyId returns a boolean if a field has been set.

### GetRiskLevelValueId

`func (o *C1ApiAppV1AppEntitlement) GetRiskLevelValueId() string`

GetRiskLevelValueId returns the RiskLevelValueId field if non-nil, zero value otherwise.

### GetRiskLevelValueIdOk

`func (o *C1ApiAppV1AppEntitlement) GetRiskLevelValueIdOk() (*string, bool)`

GetRiskLevelValueIdOk returns a tuple with the RiskLevelValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevelValueId

`func (o *C1ApiAppV1AppEntitlement) SetRiskLevelValueId(v string)`

SetRiskLevelValueId sets RiskLevelValueId field to given value.

### HasRiskLevelValueId

`func (o *C1ApiAppV1AppEntitlement) HasRiskLevelValueId() bool`

HasRiskLevelValueId returns a boolean if a field has been set.

### GetSlug

`func (o *C1ApiAppV1AppEntitlement) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *C1ApiAppV1AppEntitlement) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *C1ApiAppV1AppEntitlement) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *C1ApiAppV1AppEntitlement) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSystemBuiltin

`func (o *C1ApiAppV1AppEntitlement) GetSystemBuiltin() bool`

GetSystemBuiltin returns the SystemBuiltin field if non-nil, zero value otherwise.

### GetSystemBuiltinOk

`func (o *C1ApiAppV1AppEntitlement) GetSystemBuiltinOk() (*bool, bool)`

GetSystemBuiltinOk returns a tuple with the SystemBuiltin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemBuiltin

`func (o *C1ApiAppV1AppEntitlement) SetSystemBuiltin(v bool)`

SetSystemBuiltin sets SystemBuiltin field to given value.

### HasSystemBuiltin

`func (o *C1ApiAppV1AppEntitlement) HasSystemBuiltin() bool`

HasSystemBuiltin returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *C1ApiAppV1AppEntitlement) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *C1ApiAppV1AppEntitlement) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *C1ApiAppV1AppEntitlement) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *C1ApiAppV1AppEntitlement) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


