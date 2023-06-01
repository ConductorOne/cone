# C1ApiPolicyV1Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** | The description field. | [optional] 
**DisplayName** | Pointer to **string** | The displayName field. | [optional] 
**Id** | Pointer to **string** |  Properties  | [optional] 
**PolicySteps** | Pointer to [**map[string]C1ApiPolicyV1PolicySteps**](C1ApiPolicyV1PolicySteps.md) | The policySteps field. | [optional] 
**PolicyType** | Pointer to **string** | The policyType field. | [optional] 
**PostActions** | Pointer to [**[]C1ApiPolicyV1PolicyPostActions**](C1ApiPolicyV1PolicyPostActions.md) | The postActions field. | [optional] 
**ReassignTasksToDelegates** | Pointer to **bool** | The reassignTasksToDelegates field. | [optional] 
**SystemBuiltin** | Pointer to **bool** | The systemBuiltin field. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewC1ApiPolicyV1Policy

`func NewC1ApiPolicyV1Policy() *C1ApiPolicyV1Policy`

NewC1ApiPolicyV1Policy instantiates a new C1ApiPolicyV1Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1PolicyWithDefaults

`func NewC1ApiPolicyV1PolicyWithDefaults() *C1ApiPolicyV1Policy`

NewC1ApiPolicyV1PolicyWithDefaults instantiates a new C1ApiPolicyV1Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *C1ApiPolicyV1Policy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *C1ApiPolicyV1Policy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *C1ApiPolicyV1Policy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *C1ApiPolicyV1Policy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *C1ApiPolicyV1Policy) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *C1ApiPolicyV1Policy) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *C1ApiPolicyV1Policy) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *C1ApiPolicyV1Policy) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *C1ApiPolicyV1Policy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *C1ApiPolicyV1Policy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *C1ApiPolicyV1Policy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *C1ApiPolicyV1Policy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *C1ApiPolicyV1Policy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *C1ApiPolicyV1Policy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *C1ApiPolicyV1Policy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *C1ApiPolicyV1Policy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *C1ApiPolicyV1Policy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *C1ApiPolicyV1Policy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *C1ApiPolicyV1Policy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *C1ApiPolicyV1Policy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicySteps

`func (o *C1ApiPolicyV1Policy) GetPolicySteps() map[string]C1ApiPolicyV1PolicySteps`

GetPolicySteps returns the PolicySteps field if non-nil, zero value otherwise.

### GetPolicyStepsOk

`func (o *C1ApiPolicyV1Policy) GetPolicyStepsOk() (*map[string]C1ApiPolicyV1PolicySteps, bool)`

GetPolicyStepsOk returns a tuple with the PolicySteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySteps

`func (o *C1ApiPolicyV1Policy) SetPolicySteps(v map[string]C1ApiPolicyV1PolicySteps)`

SetPolicySteps sets PolicySteps field to given value.

### HasPolicySteps

`func (o *C1ApiPolicyV1Policy) HasPolicySteps() bool`

HasPolicySteps returns a boolean if a field has been set.

### GetPolicyType

`func (o *C1ApiPolicyV1Policy) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *C1ApiPolicyV1Policy) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *C1ApiPolicyV1Policy) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *C1ApiPolicyV1Policy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetPostActions

`func (o *C1ApiPolicyV1Policy) GetPostActions() []C1ApiPolicyV1PolicyPostActions`

GetPostActions returns the PostActions field if non-nil, zero value otherwise.

### GetPostActionsOk

`func (o *C1ApiPolicyV1Policy) GetPostActionsOk() (*[]C1ApiPolicyV1PolicyPostActions, bool)`

GetPostActionsOk returns a tuple with the PostActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostActions

`func (o *C1ApiPolicyV1Policy) SetPostActions(v []C1ApiPolicyV1PolicyPostActions)`

SetPostActions sets PostActions field to given value.

### HasPostActions

`func (o *C1ApiPolicyV1Policy) HasPostActions() bool`

HasPostActions returns a boolean if a field has been set.

### SetPostActionsNil

`func (o *C1ApiPolicyV1Policy) SetPostActionsNil(b bool)`

 SetPostActionsNil sets the value for PostActions to be an explicit nil

### UnsetPostActions
`func (o *C1ApiPolicyV1Policy) UnsetPostActions()`

UnsetPostActions ensures that no value is present for PostActions, not even an explicit nil
### GetReassignTasksToDelegates

`func (o *C1ApiPolicyV1Policy) GetReassignTasksToDelegates() bool`

GetReassignTasksToDelegates returns the ReassignTasksToDelegates field if non-nil, zero value otherwise.

### GetReassignTasksToDelegatesOk

`func (o *C1ApiPolicyV1Policy) GetReassignTasksToDelegatesOk() (*bool, bool)`

GetReassignTasksToDelegatesOk returns a tuple with the ReassignTasksToDelegates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassignTasksToDelegates

`func (o *C1ApiPolicyV1Policy) SetReassignTasksToDelegates(v bool)`

SetReassignTasksToDelegates sets ReassignTasksToDelegates field to given value.

### HasReassignTasksToDelegates

`func (o *C1ApiPolicyV1Policy) HasReassignTasksToDelegates() bool`

HasReassignTasksToDelegates returns a boolean if a field has been set.

### GetSystemBuiltin

`func (o *C1ApiPolicyV1Policy) GetSystemBuiltin() bool`

GetSystemBuiltin returns the SystemBuiltin field if non-nil, zero value otherwise.

### GetSystemBuiltinOk

`func (o *C1ApiPolicyV1Policy) GetSystemBuiltinOk() (*bool, bool)`

GetSystemBuiltinOk returns a tuple with the SystemBuiltin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemBuiltin

`func (o *C1ApiPolicyV1Policy) SetSystemBuiltin(v bool)`

SetSystemBuiltin sets SystemBuiltin field to given value.

### HasSystemBuiltin

`func (o *C1ApiPolicyV1Policy) HasSystemBuiltin() bool`

HasSystemBuiltin returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *C1ApiPolicyV1Policy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *C1ApiPolicyV1Policy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *C1ApiPolicyV1Policy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *C1ApiPolicyV1Policy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


