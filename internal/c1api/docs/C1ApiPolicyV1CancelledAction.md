# C1ApiPolicyV1CancelledAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelledAt** | Pointer to **time.Time** |  | [optional] 
**CancelledByUserId** | Pointer to **string** | The cancelledByUserId field. | [optional] 

## Methods

### NewC1ApiPolicyV1CancelledAction

`func NewC1ApiPolicyV1CancelledAction() *C1ApiPolicyV1CancelledAction`

NewC1ApiPolicyV1CancelledAction instantiates a new C1ApiPolicyV1CancelledAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1CancelledActionWithDefaults

`func NewC1ApiPolicyV1CancelledActionWithDefaults() *C1ApiPolicyV1CancelledAction`

NewC1ApiPolicyV1CancelledActionWithDefaults instantiates a new C1ApiPolicyV1CancelledAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelledAt

`func (o *C1ApiPolicyV1CancelledAction) GetCancelledAt() time.Time`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *C1ApiPolicyV1CancelledAction) GetCancelledAtOk() (*time.Time, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *C1ApiPolicyV1CancelledAction) SetCancelledAt(v time.Time)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *C1ApiPolicyV1CancelledAction) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### GetCancelledByUserId

`func (o *C1ApiPolicyV1CancelledAction) GetCancelledByUserId() string`

GetCancelledByUserId returns the CancelledByUserId field if non-nil, zero value otherwise.

### GetCancelledByUserIdOk

`func (o *C1ApiPolicyV1CancelledAction) GetCancelledByUserIdOk() (*string, bool)`

GetCancelledByUserIdOk returns a tuple with the CancelledByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledByUserId

`func (o *C1ApiPolicyV1CancelledAction) SetCancelledByUserId(v string)`

SetCancelledByUserId sets CancelledByUserId field to given value.

### HasCancelledByUserId

`func (o *C1ApiPolicyV1CancelledAction) HasCancelledByUserId() bool`

HasCancelledByUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


