# C1ApiPolicyV1EntitlementOwnerApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowSelfApproval** | Pointer to **bool** |  Entitlement owner is based on the current entitlement&#39;s id and doesn&#39;t need to have self-contained data  | [optional] 
**Fallback** | Pointer to **bool** | The fallback field. | [optional] 
**FallbackUserIds** | Pointer to **[]string** | The fallbackUserIds field. | [optional] 

## Methods

### NewC1ApiPolicyV1EntitlementOwnerApproval

`func NewC1ApiPolicyV1EntitlementOwnerApproval() *C1ApiPolicyV1EntitlementOwnerApproval`

NewC1ApiPolicyV1EntitlementOwnerApproval instantiates a new C1ApiPolicyV1EntitlementOwnerApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1EntitlementOwnerApprovalWithDefaults

`func NewC1ApiPolicyV1EntitlementOwnerApprovalWithDefaults() *C1ApiPolicyV1EntitlementOwnerApproval`

NewC1ApiPolicyV1EntitlementOwnerApprovalWithDefaults instantiates a new C1ApiPolicyV1EntitlementOwnerApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowSelfApproval

`func (o *C1ApiPolicyV1EntitlementOwnerApproval) GetAllowSelfApproval() bool`

GetAllowSelfApproval returns the AllowSelfApproval field if non-nil, zero value otherwise.

### GetAllowSelfApprovalOk

`func (o *C1ApiPolicyV1EntitlementOwnerApproval) GetAllowSelfApprovalOk() (*bool, bool)`

GetAllowSelfApprovalOk returns a tuple with the AllowSelfApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSelfApproval

`func (o *C1ApiPolicyV1EntitlementOwnerApproval) SetAllowSelfApproval(v bool)`

SetAllowSelfApproval sets AllowSelfApproval field to given value.

### HasAllowSelfApproval

`func (o *C1ApiPolicyV1EntitlementOwnerApproval) HasAllowSelfApproval() bool`

HasAllowSelfApproval returns a boolean if a field has been set.

### GetFallback

`func (o *C1ApiPolicyV1EntitlementOwnerApproval) GetFallback() bool`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *C1ApiPolicyV1EntitlementOwnerApproval) GetFallbackOk() (*bool, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *C1ApiPolicyV1EntitlementOwnerApproval) SetFallback(v bool)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *C1ApiPolicyV1EntitlementOwnerApproval) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetFallbackUserIds

`func (o *C1ApiPolicyV1EntitlementOwnerApproval) GetFallbackUserIds() []string`

GetFallbackUserIds returns the FallbackUserIds field if non-nil, zero value otherwise.

### GetFallbackUserIdsOk

`func (o *C1ApiPolicyV1EntitlementOwnerApproval) GetFallbackUserIdsOk() (*[]string, bool)`

GetFallbackUserIdsOk returns a tuple with the FallbackUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUserIds

`func (o *C1ApiPolicyV1EntitlementOwnerApproval) SetFallbackUserIds(v []string)`

SetFallbackUserIds sets FallbackUserIds field to given value.

### HasFallbackUserIds

`func (o *C1ApiPolicyV1EntitlementOwnerApproval) HasFallbackUserIds() bool`

HasFallbackUserIds returns a boolean if a field has been set.

### SetFallbackUserIdsNil

`func (o *C1ApiPolicyV1EntitlementOwnerApproval) SetFallbackUserIdsNil(b bool)`

 SetFallbackUserIdsNil sets the value for FallbackUserIds to be an explicit nil

### UnsetFallbackUserIds
`func (o *C1ApiPolicyV1EntitlementOwnerApproval) UnsetFallbackUserIds()`

UnsetFallbackUserIds ensures that no value is present for FallbackUserIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


