# C1ApiPolicyV1PolicyStepInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approval** | Pointer to [**NullableC1ApiPolicyV1ApprovalInstance**](C1ApiPolicyV1ApprovalInstance.md) |  | [optional] 
**Id** | Pointer to **string** | The id field. | [optional] 
**Provision** | Pointer to [**NullableC1ApiPolicyV1ProvisionInstance**](C1ApiPolicyV1ProvisionInstance.md) |  | [optional] 
**State** | Pointer to **string** | The state field. | [optional] 

## Methods

### NewC1ApiPolicyV1PolicyStepInstance

`func NewC1ApiPolicyV1PolicyStepInstance() *C1ApiPolicyV1PolicyStepInstance`

NewC1ApiPolicyV1PolicyStepInstance instantiates a new C1ApiPolicyV1PolicyStepInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1PolicyStepInstanceWithDefaults

`func NewC1ApiPolicyV1PolicyStepInstanceWithDefaults() *C1ApiPolicyV1PolicyStepInstance`

NewC1ApiPolicyV1PolicyStepInstanceWithDefaults instantiates a new C1ApiPolicyV1PolicyStepInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproval

`func (o *C1ApiPolicyV1PolicyStepInstance) GetApproval() C1ApiPolicyV1ApprovalInstance`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *C1ApiPolicyV1PolicyStepInstance) GetApprovalOk() (*C1ApiPolicyV1ApprovalInstance, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *C1ApiPolicyV1PolicyStepInstance) SetApproval(v C1ApiPolicyV1ApprovalInstance)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *C1ApiPolicyV1PolicyStepInstance) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### SetApprovalNil

`func (o *C1ApiPolicyV1PolicyStepInstance) SetApprovalNil(b bool)`

 SetApprovalNil sets the value for Approval to be an explicit nil

### UnsetApproval
`func (o *C1ApiPolicyV1PolicyStepInstance) UnsetApproval()`

UnsetApproval ensures that no value is present for Approval, not even an explicit nil
### GetId

`func (o *C1ApiPolicyV1PolicyStepInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *C1ApiPolicyV1PolicyStepInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *C1ApiPolicyV1PolicyStepInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *C1ApiPolicyV1PolicyStepInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProvision

`func (o *C1ApiPolicyV1PolicyStepInstance) GetProvision() C1ApiPolicyV1ProvisionInstance`

GetProvision returns the Provision field if non-nil, zero value otherwise.

### GetProvisionOk

`func (o *C1ApiPolicyV1PolicyStepInstance) GetProvisionOk() (*C1ApiPolicyV1ProvisionInstance, bool)`

GetProvisionOk returns a tuple with the Provision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvision

`func (o *C1ApiPolicyV1PolicyStepInstance) SetProvision(v C1ApiPolicyV1ProvisionInstance)`

SetProvision sets Provision field to given value.

### HasProvision

`func (o *C1ApiPolicyV1PolicyStepInstance) HasProvision() bool`

HasProvision returns a boolean if a field has been set.

### SetProvisionNil

`func (o *C1ApiPolicyV1PolicyStepInstance) SetProvisionNil(b bool)`

 SetProvisionNil sets the value for Provision to be an explicit nil

### UnsetProvision
`func (o *C1ApiPolicyV1PolicyStepInstance) UnsetProvision()`

UnsetProvision ensures that no value is present for Provision, not even an explicit nil
### GetState

`func (o *C1ApiPolicyV1PolicyStepInstance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *C1ApiPolicyV1PolicyStepInstance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *C1ApiPolicyV1PolicyStepInstance) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *C1ApiPolicyV1PolicyStepInstance) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


