# C1ApiPolicyV1SelfApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedUserIds** | Pointer to **[]string** | The assignedUserIds field. | [optional] 
**Fallback** | Pointer to **bool** | The fallback field. | [optional] 
**FallbackUserIds** | Pointer to **[]string** |  Self approval is the target of the ticket  | [optional] 

## Methods

### NewC1ApiPolicyV1SelfApproval

`func NewC1ApiPolicyV1SelfApproval() *C1ApiPolicyV1SelfApproval`

NewC1ApiPolicyV1SelfApproval instantiates a new C1ApiPolicyV1SelfApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1SelfApprovalWithDefaults

`func NewC1ApiPolicyV1SelfApprovalWithDefaults() *C1ApiPolicyV1SelfApproval`

NewC1ApiPolicyV1SelfApprovalWithDefaults instantiates a new C1ApiPolicyV1SelfApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedUserIds

`func (o *C1ApiPolicyV1SelfApproval) GetAssignedUserIds() []string`

GetAssignedUserIds returns the AssignedUserIds field if non-nil, zero value otherwise.

### GetAssignedUserIdsOk

`func (o *C1ApiPolicyV1SelfApproval) GetAssignedUserIdsOk() (*[]string, bool)`

GetAssignedUserIdsOk returns a tuple with the AssignedUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserIds

`func (o *C1ApiPolicyV1SelfApproval) SetAssignedUserIds(v []string)`

SetAssignedUserIds sets AssignedUserIds field to given value.

### HasAssignedUserIds

`func (o *C1ApiPolicyV1SelfApproval) HasAssignedUserIds() bool`

HasAssignedUserIds returns a boolean if a field has been set.

### SetAssignedUserIdsNil

`func (o *C1ApiPolicyV1SelfApproval) SetAssignedUserIdsNil(b bool)`

 SetAssignedUserIdsNil sets the value for AssignedUserIds to be an explicit nil

### UnsetAssignedUserIds
`func (o *C1ApiPolicyV1SelfApproval) UnsetAssignedUserIds()`

UnsetAssignedUserIds ensures that no value is present for AssignedUserIds, not even an explicit nil
### GetFallback

`func (o *C1ApiPolicyV1SelfApproval) GetFallback() bool`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *C1ApiPolicyV1SelfApproval) GetFallbackOk() (*bool, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *C1ApiPolicyV1SelfApproval) SetFallback(v bool)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *C1ApiPolicyV1SelfApproval) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetFallbackUserIds

`func (o *C1ApiPolicyV1SelfApproval) GetFallbackUserIds() []string`

GetFallbackUserIds returns the FallbackUserIds field if non-nil, zero value otherwise.

### GetFallbackUserIdsOk

`func (o *C1ApiPolicyV1SelfApproval) GetFallbackUserIdsOk() (*[]string, bool)`

GetFallbackUserIdsOk returns a tuple with the FallbackUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUserIds

`func (o *C1ApiPolicyV1SelfApproval) SetFallbackUserIds(v []string)`

SetFallbackUserIds sets FallbackUserIds field to given value.

### HasFallbackUserIds

`func (o *C1ApiPolicyV1SelfApproval) HasFallbackUserIds() bool`

HasFallbackUserIds returns a boolean if a field has been set.

### SetFallbackUserIdsNil

`func (o *C1ApiPolicyV1SelfApproval) SetFallbackUserIdsNil(b bool)`

 SetFallbackUserIdsNil sets the value for FallbackUserIds to be an explicit nil

### UnsetFallbackUserIds
`func (o *C1ApiPolicyV1SelfApproval) UnsetFallbackUserIds()`

UnsetFallbackUserIds ensures that no value is present for FallbackUserIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


