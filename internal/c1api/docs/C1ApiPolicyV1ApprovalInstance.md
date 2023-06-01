# C1ApiPolicyV1ApprovalInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approval** | Pointer to [**NullableC1ApiPolicyV1Approval**](C1ApiPolicyV1Approval.md) |  | [optional] 
**Approved** | Pointer to [**NullableC1ApiPolicyV1ApprovedAction**](C1ApiPolicyV1ApprovedAction.md) |  | [optional] 
**Denied** | Pointer to [**NullableC1ApiPolicyV1DeniedAction**](C1ApiPolicyV1DeniedAction.md) |  | [optional] 
**Reassigned** | Pointer to [**NullableC1ApiPolicyV1ReassignedAction**](C1ApiPolicyV1ReassignedAction.md) |  | [optional] 
**ReassignedByError** | Pointer to [**NullableC1ApiPolicyV1ReassignedByErrorAction**](C1ApiPolicyV1ReassignedByErrorAction.md) |  | [optional] 
**Restarted** | Pointer to [**NullableC1ApiPolicyV1RestartAction**](C1ApiPolicyV1RestartAction.md) |  | [optional] 
**State** | Pointer to **string** | The state field. | [optional] 

## Methods

### NewC1ApiPolicyV1ApprovalInstance

`func NewC1ApiPolicyV1ApprovalInstance() *C1ApiPolicyV1ApprovalInstance`

NewC1ApiPolicyV1ApprovalInstance instantiates a new C1ApiPolicyV1ApprovalInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1ApprovalInstanceWithDefaults

`func NewC1ApiPolicyV1ApprovalInstanceWithDefaults() *C1ApiPolicyV1ApprovalInstance`

NewC1ApiPolicyV1ApprovalInstanceWithDefaults instantiates a new C1ApiPolicyV1ApprovalInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproval

`func (o *C1ApiPolicyV1ApprovalInstance) GetApproval() C1ApiPolicyV1Approval`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *C1ApiPolicyV1ApprovalInstance) GetApprovalOk() (*C1ApiPolicyV1Approval, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *C1ApiPolicyV1ApprovalInstance) SetApproval(v C1ApiPolicyV1Approval)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *C1ApiPolicyV1ApprovalInstance) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### SetApprovalNil

`func (o *C1ApiPolicyV1ApprovalInstance) SetApprovalNil(b bool)`

 SetApprovalNil sets the value for Approval to be an explicit nil

### UnsetApproval
`func (o *C1ApiPolicyV1ApprovalInstance) UnsetApproval()`

UnsetApproval ensures that no value is present for Approval, not even an explicit nil
### GetApproved

`func (o *C1ApiPolicyV1ApprovalInstance) GetApproved() C1ApiPolicyV1ApprovedAction`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *C1ApiPolicyV1ApprovalInstance) GetApprovedOk() (*C1ApiPolicyV1ApprovedAction, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *C1ApiPolicyV1ApprovalInstance) SetApproved(v C1ApiPolicyV1ApprovedAction)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *C1ApiPolicyV1ApprovalInstance) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### SetApprovedNil

`func (o *C1ApiPolicyV1ApprovalInstance) SetApprovedNil(b bool)`

 SetApprovedNil sets the value for Approved to be an explicit nil

### UnsetApproved
`func (o *C1ApiPolicyV1ApprovalInstance) UnsetApproved()`

UnsetApproved ensures that no value is present for Approved, not even an explicit nil
### GetDenied

`func (o *C1ApiPolicyV1ApprovalInstance) GetDenied() C1ApiPolicyV1DeniedAction`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *C1ApiPolicyV1ApprovalInstance) GetDeniedOk() (*C1ApiPolicyV1DeniedAction, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *C1ApiPolicyV1ApprovalInstance) SetDenied(v C1ApiPolicyV1DeniedAction)`

SetDenied sets Denied field to given value.

### HasDenied

`func (o *C1ApiPolicyV1ApprovalInstance) HasDenied() bool`

HasDenied returns a boolean if a field has been set.

### SetDeniedNil

`func (o *C1ApiPolicyV1ApprovalInstance) SetDeniedNil(b bool)`

 SetDeniedNil sets the value for Denied to be an explicit nil

### UnsetDenied
`func (o *C1ApiPolicyV1ApprovalInstance) UnsetDenied()`

UnsetDenied ensures that no value is present for Denied, not even an explicit nil
### GetReassigned

`func (o *C1ApiPolicyV1ApprovalInstance) GetReassigned() C1ApiPolicyV1ReassignedAction`

GetReassigned returns the Reassigned field if non-nil, zero value otherwise.

### GetReassignedOk

`func (o *C1ApiPolicyV1ApprovalInstance) GetReassignedOk() (*C1ApiPolicyV1ReassignedAction, bool)`

GetReassignedOk returns a tuple with the Reassigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassigned

`func (o *C1ApiPolicyV1ApprovalInstance) SetReassigned(v C1ApiPolicyV1ReassignedAction)`

SetReassigned sets Reassigned field to given value.

### HasReassigned

`func (o *C1ApiPolicyV1ApprovalInstance) HasReassigned() bool`

HasReassigned returns a boolean if a field has been set.

### SetReassignedNil

`func (o *C1ApiPolicyV1ApprovalInstance) SetReassignedNil(b bool)`

 SetReassignedNil sets the value for Reassigned to be an explicit nil

### UnsetReassigned
`func (o *C1ApiPolicyV1ApprovalInstance) UnsetReassigned()`

UnsetReassigned ensures that no value is present for Reassigned, not even an explicit nil
### GetReassignedByError

`func (o *C1ApiPolicyV1ApprovalInstance) GetReassignedByError() C1ApiPolicyV1ReassignedByErrorAction`

GetReassignedByError returns the ReassignedByError field if non-nil, zero value otherwise.

### GetReassignedByErrorOk

`func (o *C1ApiPolicyV1ApprovalInstance) GetReassignedByErrorOk() (*C1ApiPolicyV1ReassignedByErrorAction, bool)`

GetReassignedByErrorOk returns a tuple with the ReassignedByError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassignedByError

`func (o *C1ApiPolicyV1ApprovalInstance) SetReassignedByError(v C1ApiPolicyV1ReassignedByErrorAction)`

SetReassignedByError sets ReassignedByError field to given value.

### HasReassignedByError

`func (o *C1ApiPolicyV1ApprovalInstance) HasReassignedByError() bool`

HasReassignedByError returns a boolean if a field has been set.

### SetReassignedByErrorNil

`func (o *C1ApiPolicyV1ApprovalInstance) SetReassignedByErrorNil(b bool)`

 SetReassignedByErrorNil sets the value for ReassignedByError to be an explicit nil

### UnsetReassignedByError
`func (o *C1ApiPolicyV1ApprovalInstance) UnsetReassignedByError()`

UnsetReassignedByError ensures that no value is present for ReassignedByError, not even an explicit nil
### GetRestarted

`func (o *C1ApiPolicyV1ApprovalInstance) GetRestarted() C1ApiPolicyV1RestartAction`

GetRestarted returns the Restarted field if non-nil, zero value otherwise.

### GetRestartedOk

`func (o *C1ApiPolicyV1ApprovalInstance) GetRestartedOk() (*C1ApiPolicyV1RestartAction, bool)`

GetRestartedOk returns a tuple with the Restarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarted

`func (o *C1ApiPolicyV1ApprovalInstance) SetRestarted(v C1ApiPolicyV1RestartAction)`

SetRestarted sets Restarted field to given value.

### HasRestarted

`func (o *C1ApiPolicyV1ApprovalInstance) HasRestarted() bool`

HasRestarted returns a boolean if a field has been set.

### SetRestartedNil

`func (o *C1ApiPolicyV1ApprovalInstance) SetRestartedNil(b bool)`

 SetRestartedNil sets the value for Restarted to be an explicit nil

### UnsetRestarted
`func (o *C1ApiPolicyV1ApprovalInstance) UnsetRestarted()`

UnsetRestarted ensures that no value is present for Restarted, not even an explicit nil
### GetState

`func (o *C1ApiPolicyV1ApprovalInstance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *C1ApiPolicyV1ApprovalInstance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *C1ApiPolicyV1ApprovalInstance) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *C1ApiPolicyV1ApprovalInstance) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


