# C1ApiPolicyV1UserApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowSelfApproval** | Pointer to **bool** | The allowSelfApproval field. | [optional] 
**UserIds** | Pointer to **[]string** | The userIds field. | [optional] 

## Methods

### NewC1ApiPolicyV1UserApproval

`func NewC1ApiPolicyV1UserApproval() *C1ApiPolicyV1UserApproval`

NewC1ApiPolicyV1UserApproval instantiates a new C1ApiPolicyV1UserApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1UserApprovalWithDefaults

`func NewC1ApiPolicyV1UserApprovalWithDefaults() *C1ApiPolicyV1UserApproval`

NewC1ApiPolicyV1UserApprovalWithDefaults instantiates a new C1ApiPolicyV1UserApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowSelfApproval

`func (o *C1ApiPolicyV1UserApproval) GetAllowSelfApproval() bool`

GetAllowSelfApproval returns the AllowSelfApproval field if non-nil, zero value otherwise.

### GetAllowSelfApprovalOk

`func (o *C1ApiPolicyV1UserApproval) GetAllowSelfApprovalOk() (*bool, bool)`

GetAllowSelfApprovalOk returns a tuple with the AllowSelfApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSelfApproval

`func (o *C1ApiPolicyV1UserApproval) SetAllowSelfApproval(v bool)`

SetAllowSelfApproval sets AllowSelfApproval field to given value.

### HasAllowSelfApproval

`func (o *C1ApiPolicyV1UserApproval) HasAllowSelfApproval() bool`

HasAllowSelfApproval returns a boolean if a field has been set.

### GetUserIds

`func (o *C1ApiPolicyV1UserApproval) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *C1ApiPolicyV1UserApproval) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *C1ApiPolicyV1UserApproval) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *C1ApiPolicyV1UserApproval) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *C1ApiPolicyV1UserApproval) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *C1ApiPolicyV1UserApproval) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


