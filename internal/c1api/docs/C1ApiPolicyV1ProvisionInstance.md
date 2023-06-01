# C1ApiPolicyV1ProvisionInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancelled** | Pointer to [**NullableC1ApiPolicyV1CancelledAction**](C1ApiPolicyV1CancelledAction.md) |  | [optional] 
**Completed** | Pointer to [**NullableC1ApiPolicyV1CompletedAction**](C1ApiPolicyV1CompletedAction.md) |  | [optional] 
**Errored** | Pointer to [**NullableC1ApiPolicyV1ErroredAction**](C1ApiPolicyV1ErroredAction.md) |  | [optional] 
**NotificationId** | Pointer to **string** | The notificationId field. | [optional] 
**Provision** | Pointer to [**NullableC1ApiPolicyV1Provision**](C1ApiPolicyV1Provision.md) |  | [optional] 
**ReassignedByError** | Pointer to [**NullableC1ApiPolicyV1ReassignedByErrorAction**](C1ApiPolicyV1ReassignedByErrorAction.md) |  | [optional] 
**State** | Pointer to **string** | The state field. | [optional] 

## Methods

### NewC1ApiPolicyV1ProvisionInstance

`func NewC1ApiPolicyV1ProvisionInstance() *C1ApiPolicyV1ProvisionInstance`

NewC1ApiPolicyV1ProvisionInstance instantiates a new C1ApiPolicyV1ProvisionInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1ProvisionInstanceWithDefaults

`func NewC1ApiPolicyV1ProvisionInstanceWithDefaults() *C1ApiPolicyV1ProvisionInstance`

NewC1ApiPolicyV1ProvisionInstanceWithDefaults instantiates a new C1ApiPolicyV1ProvisionInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelled

`func (o *C1ApiPolicyV1ProvisionInstance) GetCancelled() C1ApiPolicyV1CancelledAction`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *C1ApiPolicyV1ProvisionInstance) GetCancelledOk() (*C1ApiPolicyV1CancelledAction, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *C1ApiPolicyV1ProvisionInstance) SetCancelled(v C1ApiPolicyV1CancelledAction)`

SetCancelled sets Cancelled field to given value.

### HasCancelled

`func (o *C1ApiPolicyV1ProvisionInstance) HasCancelled() bool`

HasCancelled returns a boolean if a field has been set.

### SetCancelledNil

`func (o *C1ApiPolicyV1ProvisionInstance) SetCancelledNil(b bool)`

 SetCancelledNil sets the value for Cancelled to be an explicit nil

### UnsetCancelled
`func (o *C1ApiPolicyV1ProvisionInstance) UnsetCancelled()`

UnsetCancelled ensures that no value is present for Cancelled, not even an explicit nil
### GetCompleted

`func (o *C1ApiPolicyV1ProvisionInstance) GetCompleted() C1ApiPolicyV1CompletedAction`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *C1ApiPolicyV1ProvisionInstance) GetCompletedOk() (*C1ApiPolicyV1CompletedAction, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *C1ApiPolicyV1ProvisionInstance) SetCompleted(v C1ApiPolicyV1CompletedAction)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *C1ApiPolicyV1ProvisionInstance) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### SetCompletedNil

`func (o *C1ApiPolicyV1ProvisionInstance) SetCompletedNil(b bool)`

 SetCompletedNil sets the value for Completed to be an explicit nil

### UnsetCompleted
`func (o *C1ApiPolicyV1ProvisionInstance) UnsetCompleted()`

UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
### GetErrored

`func (o *C1ApiPolicyV1ProvisionInstance) GetErrored() C1ApiPolicyV1ErroredAction`

GetErrored returns the Errored field if non-nil, zero value otherwise.

### GetErroredOk

`func (o *C1ApiPolicyV1ProvisionInstance) GetErroredOk() (*C1ApiPolicyV1ErroredAction, bool)`

GetErroredOk returns a tuple with the Errored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrored

`func (o *C1ApiPolicyV1ProvisionInstance) SetErrored(v C1ApiPolicyV1ErroredAction)`

SetErrored sets Errored field to given value.

### HasErrored

`func (o *C1ApiPolicyV1ProvisionInstance) HasErrored() bool`

HasErrored returns a boolean if a field has been set.

### SetErroredNil

`func (o *C1ApiPolicyV1ProvisionInstance) SetErroredNil(b bool)`

 SetErroredNil sets the value for Errored to be an explicit nil

### UnsetErrored
`func (o *C1ApiPolicyV1ProvisionInstance) UnsetErrored()`

UnsetErrored ensures that no value is present for Errored, not even an explicit nil
### GetNotificationId

`func (o *C1ApiPolicyV1ProvisionInstance) GetNotificationId() string`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *C1ApiPolicyV1ProvisionInstance) GetNotificationIdOk() (*string, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *C1ApiPolicyV1ProvisionInstance) SetNotificationId(v string)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *C1ApiPolicyV1ProvisionInstance) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetProvision

`func (o *C1ApiPolicyV1ProvisionInstance) GetProvision() C1ApiPolicyV1Provision`

GetProvision returns the Provision field if non-nil, zero value otherwise.

### GetProvisionOk

`func (o *C1ApiPolicyV1ProvisionInstance) GetProvisionOk() (*C1ApiPolicyV1Provision, bool)`

GetProvisionOk returns a tuple with the Provision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvision

`func (o *C1ApiPolicyV1ProvisionInstance) SetProvision(v C1ApiPolicyV1Provision)`

SetProvision sets Provision field to given value.

### HasProvision

`func (o *C1ApiPolicyV1ProvisionInstance) HasProvision() bool`

HasProvision returns a boolean if a field has been set.

### SetProvisionNil

`func (o *C1ApiPolicyV1ProvisionInstance) SetProvisionNil(b bool)`

 SetProvisionNil sets the value for Provision to be an explicit nil

### UnsetProvision
`func (o *C1ApiPolicyV1ProvisionInstance) UnsetProvision()`

UnsetProvision ensures that no value is present for Provision, not even an explicit nil
### GetReassignedByError

`func (o *C1ApiPolicyV1ProvisionInstance) GetReassignedByError() C1ApiPolicyV1ReassignedByErrorAction`

GetReassignedByError returns the ReassignedByError field if non-nil, zero value otherwise.

### GetReassignedByErrorOk

`func (o *C1ApiPolicyV1ProvisionInstance) GetReassignedByErrorOk() (*C1ApiPolicyV1ReassignedByErrorAction, bool)`

GetReassignedByErrorOk returns a tuple with the ReassignedByError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassignedByError

`func (o *C1ApiPolicyV1ProvisionInstance) SetReassignedByError(v C1ApiPolicyV1ReassignedByErrorAction)`

SetReassignedByError sets ReassignedByError field to given value.

### HasReassignedByError

`func (o *C1ApiPolicyV1ProvisionInstance) HasReassignedByError() bool`

HasReassignedByError returns a boolean if a field has been set.

### SetReassignedByErrorNil

`func (o *C1ApiPolicyV1ProvisionInstance) SetReassignedByErrorNil(b bool)`

 SetReassignedByErrorNil sets the value for ReassignedByError to be an explicit nil

### UnsetReassignedByError
`func (o *C1ApiPolicyV1ProvisionInstance) UnsetReassignedByError()`

UnsetReassignedByError ensures that no value is present for ReassignedByError, not even an explicit nil
### GetState

`func (o *C1ApiPolicyV1ProvisionInstance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *C1ApiPolicyV1ProvisionInstance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *C1ApiPolicyV1ProvisionInstance) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *C1ApiPolicyV1ProvisionInstance) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


