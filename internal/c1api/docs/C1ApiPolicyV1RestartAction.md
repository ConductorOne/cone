# C1ApiPolicyV1RestartAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldPolicyStepId** | Pointer to **string** | The oldPolicyStepId field. | [optional] 
**RestartedAt** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **string** | The userId field. | [optional] 

## Methods

### NewC1ApiPolicyV1RestartAction

`func NewC1ApiPolicyV1RestartAction() *C1ApiPolicyV1RestartAction`

NewC1ApiPolicyV1RestartAction instantiates a new C1ApiPolicyV1RestartAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1RestartActionWithDefaults

`func NewC1ApiPolicyV1RestartActionWithDefaults() *C1ApiPolicyV1RestartAction`

NewC1ApiPolicyV1RestartActionWithDefaults instantiates a new C1ApiPolicyV1RestartAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldPolicyStepId

`func (o *C1ApiPolicyV1RestartAction) GetOldPolicyStepId() string`

GetOldPolicyStepId returns the OldPolicyStepId field if non-nil, zero value otherwise.

### GetOldPolicyStepIdOk

`func (o *C1ApiPolicyV1RestartAction) GetOldPolicyStepIdOk() (*string, bool)`

GetOldPolicyStepIdOk returns a tuple with the OldPolicyStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPolicyStepId

`func (o *C1ApiPolicyV1RestartAction) SetOldPolicyStepId(v string)`

SetOldPolicyStepId sets OldPolicyStepId field to given value.

### HasOldPolicyStepId

`func (o *C1ApiPolicyV1RestartAction) HasOldPolicyStepId() bool`

HasOldPolicyStepId returns a boolean if a field has been set.

### GetRestartedAt

`func (o *C1ApiPolicyV1RestartAction) GetRestartedAt() time.Time`

GetRestartedAt returns the RestartedAt field if non-nil, zero value otherwise.

### GetRestartedAtOk

`func (o *C1ApiPolicyV1RestartAction) GetRestartedAtOk() (*time.Time, bool)`

GetRestartedAtOk returns a tuple with the RestartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartedAt

`func (o *C1ApiPolicyV1RestartAction) SetRestartedAt(v time.Time)`

SetRestartedAt sets RestartedAt field to given value.

### HasRestartedAt

`func (o *C1ApiPolicyV1RestartAction) HasRestartedAt() bool`

HasRestartedAt returns a boolean if a field has been set.

### GetUserId

`func (o *C1ApiPolicyV1RestartAction) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *C1ApiPolicyV1RestartAction) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *C1ApiPolicyV1RestartAction) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *C1ApiPolicyV1RestartAction) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


