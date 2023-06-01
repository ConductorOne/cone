# C1ApiAppV1App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppAccountId** | Pointer to **string** | The appAccountId field. | [optional] 
**AppAccountName** | Pointer to **string** | The appAccountName field. | [optional] 
**CertifyPolicyId** | Pointer to **string** | The certifyPolicyId field. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** | The description field. | [optional] 
**DisplayName** | Pointer to **string** | The displayName field. | [optional] 
**FieldMask** | Pointer to **NullableString** |  | [optional] 
**GrantPolicyId** | Pointer to **string** | The grantPolicyId field. | [optional] 
**IconUrl** | Pointer to **string** | The iconUrl field. | [optional] 
**Id** | Pointer to **string** | The id field. | [optional] 
**LogoUri** | Pointer to **string** | The logoUri field. | [optional] 
**MonthlyCostUsd** | Pointer to **float32** | The monthlyCostUsd field. | [optional] 
**ParentAppId** | Pointer to **string** | The parentAppId field. | [optional] 
**RevokePolicyId** | Pointer to **string** | The revokePolicyId field. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UserCount** | Pointer to **string** | The userCount field. | [optional] 

## Methods

### NewC1ApiAppV1App

`func NewC1ApiAppV1App() *C1ApiAppV1App`

NewC1ApiAppV1App instantiates a new C1ApiAppV1App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiAppV1AppWithDefaults

`func NewC1ApiAppV1AppWithDefaults() *C1ApiAppV1App`

NewC1ApiAppV1AppWithDefaults instantiates a new C1ApiAppV1App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppAccountId

`func (o *C1ApiAppV1App) GetAppAccountId() string`

GetAppAccountId returns the AppAccountId field if non-nil, zero value otherwise.

### GetAppAccountIdOk

`func (o *C1ApiAppV1App) GetAppAccountIdOk() (*string, bool)`

GetAppAccountIdOk returns a tuple with the AppAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAccountId

`func (o *C1ApiAppV1App) SetAppAccountId(v string)`

SetAppAccountId sets AppAccountId field to given value.

### HasAppAccountId

`func (o *C1ApiAppV1App) HasAppAccountId() bool`

HasAppAccountId returns a boolean if a field has been set.

### GetAppAccountName

`func (o *C1ApiAppV1App) GetAppAccountName() string`

GetAppAccountName returns the AppAccountName field if non-nil, zero value otherwise.

### GetAppAccountNameOk

`func (o *C1ApiAppV1App) GetAppAccountNameOk() (*string, bool)`

GetAppAccountNameOk returns a tuple with the AppAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAccountName

`func (o *C1ApiAppV1App) SetAppAccountName(v string)`

SetAppAccountName sets AppAccountName field to given value.

### HasAppAccountName

`func (o *C1ApiAppV1App) HasAppAccountName() bool`

HasAppAccountName returns a boolean if a field has been set.

### GetCertifyPolicyId

`func (o *C1ApiAppV1App) GetCertifyPolicyId() string`

GetCertifyPolicyId returns the CertifyPolicyId field if non-nil, zero value otherwise.

### GetCertifyPolicyIdOk

`func (o *C1ApiAppV1App) GetCertifyPolicyIdOk() (*string, bool)`

GetCertifyPolicyIdOk returns a tuple with the CertifyPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifyPolicyId

`func (o *C1ApiAppV1App) SetCertifyPolicyId(v string)`

SetCertifyPolicyId sets CertifyPolicyId field to given value.

### HasCertifyPolicyId

`func (o *C1ApiAppV1App) HasCertifyPolicyId() bool`

HasCertifyPolicyId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *C1ApiAppV1App) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *C1ApiAppV1App) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *C1ApiAppV1App) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *C1ApiAppV1App) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *C1ApiAppV1App) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *C1ApiAppV1App) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *C1ApiAppV1App) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *C1ApiAppV1App) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *C1ApiAppV1App) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *C1ApiAppV1App) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *C1ApiAppV1App) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *C1ApiAppV1App) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *C1ApiAppV1App) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *C1ApiAppV1App) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *C1ApiAppV1App) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *C1ApiAppV1App) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFieldMask

`func (o *C1ApiAppV1App) GetFieldMask() string`

GetFieldMask returns the FieldMask field if non-nil, zero value otherwise.

### GetFieldMaskOk

`func (o *C1ApiAppV1App) GetFieldMaskOk() (*string, bool)`

GetFieldMaskOk returns a tuple with the FieldMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMask

`func (o *C1ApiAppV1App) SetFieldMask(v string)`

SetFieldMask sets FieldMask field to given value.

### HasFieldMask

`func (o *C1ApiAppV1App) HasFieldMask() bool`

HasFieldMask returns a boolean if a field has been set.

### SetFieldMaskNil

`func (o *C1ApiAppV1App) SetFieldMaskNil(b bool)`

 SetFieldMaskNil sets the value for FieldMask to be an explicit nil

### UnsetFieldMask
`func (o *C1ApiAppV1App) UnsetFieldMask()`

UnsetFieldMask ensures that no value is present for FieldMask, not even an explicit nil
### GetGrantPolicyId

`func (o *C1ApiAppV1App) GetGrantPolicyId() string`

GetGrantPolicyId returns the GrantPolicyId field if non-nil, zero value otherwise.

### GetGrantPolicyIdOk

`func (o *C1ApiAppV1App) GetGrantPolicyIdOk() (*string, bool)`

GetGrantPolicyIdOk returns a tuple with the GrantPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantPolicyId

`func (o *C1ApiAppV1App) SetGrantPolicyId(v string)`

SetGrantPolicyId sets GrantPolicyId field to given value.

### HasGrantPolicyId

`func (o *C1ApiAppV1App) HasGrantPolicyId() bool`

HasGrantPolicyId returns a boolean if a field has been set.

### GetIconUrl

`func (o *C1ApiAppV1App) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *C1ApiAppV1App) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *C1ApiAppV1App) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *C1ApiAppV1App) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetId

`func (o *C1ApiAppV1App) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *C1ApiAppV1App) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *C1ApiAppV1App) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *C1ApiAppV1App) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogoUri

`func (o *C1ApiAppV1App) GetLogoUri() string`

GetLogoUri returns the LogoUri field if non-nil, zero value otherwise.

### GetLogoUriOk

`func (o *C1ApiAppV1App) GetLogoUriOk() (*string, bool)`

GetLogoUriOk returns a tuple with the LogoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUri

`func (o *C1ApiAppV1App) SetLogoUri(v string)`

SetLogoUri sets LogoUri field to given value.

### HasLogoUri

`func (o *C1ApiAppV1App) HasLogoUri() bool`

HasLogoUri returns a boolean if a field has been set.

### GetMonthlyCostUsd

`func (o *C1ApiAppV1App) GetMonthlyCostUsd() float32`

GetMonthlyCostUsd returns the MonthlyCostUsd field if non-nil, zero value otherwise.

### GetMonthlyCostUsdOk

`func (o *C1ApiAppV1App) GetMonthlyCostUsdOk() (*float32, bool)`

GetMonthlyCostUsdOk returns a tuple with the MonthlyCostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCostUsd

`func (o *C1ApiAppV1App) SetMonthlyCostUsd(v float32)`

SetMonthlyCostUsd sets MonthlyCostUsd field to given value.

### HasMonthlyCostUsd

`func (o *C1ApiAppV1App) HasMonthlyCostUsd() bool`

HasMonthlyCostUsd returns a boolean if a field has been set.

### GetParentAppId

`func (o *C1ApiAppV1App) GetParentAppId() string`

GetParentAppId returns the ParentAppId field if non-nil, zero value otherwise.

### GetParentAppIdOk

`func (o *C1ApiAppV1App) GetParentAppIdOk() (*string, bool)`

GetParentAppIdOk returns a tuple with the ParentAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAppId

`func (o *C1ApiAppV1App) SetParentAppId(v string)`

SetParentAppId sets ParentAppId field to given value.

### HasParentAppId

`func (o *C1ApiAppV1App) HasParentAppId() bool`

HasParentAppId returns a boolean if a field has been set.

### GetRevokePolicyId

`func (o *C1ApiAppV1App) GetRevokePolicyId() string`

GetRevokePolicyId returns the RevokePolicyId field if non-nil, zero value otherwise.

### GetRevokePolicyIdOk

`func (o *C1ApiAppV1App) GetRevokePolicyIdOk() (*string, bool)`

GetRevokePolicyIdOk returns a tuple with the RevokePolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokePolicyId

`func (o *C1ApiAppV1App) SetRevokePolicyId(v string)`

SetRevokePolicyId sets RevokePolicyId field to given value.

### HasRevokePolicyId

`func (o *C1ApiAppV1App) HasRevokePolicyId() bool`

HasRevokePolicyId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *C1ApiAppV1App) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *C1ApiAppV1App) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *C1ApiAppV1App) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *C1ApiAppV1App) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserCount

`func (o *C1ApiAppV1App) GetUserCount() string`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *C1ApiAppV1App) GetUserCountOk() (*string, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *C1ApiAppV1App) SetUserCount(v string)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *C1ApiAppV1App) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


