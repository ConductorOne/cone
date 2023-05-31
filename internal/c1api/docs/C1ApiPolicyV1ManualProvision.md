# C1ApiPolicyV1ManualProvision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instructions** | Pointer to **string** | The instructions field. | [optional] 
**UserIds** | Pointer to **[]string** | The userIds field. | [optional] 

## Methods

### NewC1ApiPolicyV1ManualProvision

`func NewC1ApiPolicyV1ManualProvision() *C1ApiPolicyV1ManualProvision`

NewC1ApiPolicyV1ManualProvision instantiates a new C1ApiPolicyV1ManualProvision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1ManualProvisionWithDefaults

`func NewC1ApiPolicyV1ManualProvisionWithDefaults() *C1ApiPolicyV1ManualProvision`

NewC1ApiPolicyV1ManualProvisionWithDefaults instantiates a new C1ApiPolicyV1ManualProvision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructions

`func (o *C1ApiPolicyV1ManualProvision) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *C1ApiPolicyV1ManualProvision) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *C1ApiPolicyV1ManualProvision) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *C1ApiPolicyV1ManualProvision) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetUserIds

`func (o *C1ApiPolicyV1ManualProvision) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *C1ApiPolicyV1ManualProvision) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *C1ApiPolicyV1ManualProvision) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *C1ApiPolicyV1ManualProvision) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *C1ApiPolicyV1ManualProvision) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *C1ApiPolicyV1ManualProvision) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


