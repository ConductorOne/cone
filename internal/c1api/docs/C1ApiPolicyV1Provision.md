# C1ApiPolicyV1Provision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assigned** | Pointer to **bool** | The assigned field. | [optional] 
**ProvisionPolicy** | Pointer to [**C1ApiPolicyV1ProvisionPolicy**](C1ApiPolicyV1ProvisionPolicy.md) |  | [optional] 

## Methods

### NewC1ApiPolicyV1Provision

`func NewC1ApiPolicyV1Provision() *C1ApiPolicyV1Provision`

NewC1ApiPolicyV1Provision instantiates a new C1ApiPolicyV1Provision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1ProvisionWithDefaults

`func NewC1ApiPolicyV1ProvisionWithDefaults() *C1ApiPolicyV1Provision`

NewC1ApiPolicyV1ProvisionWithDefaults instantiates a new C1ApiPolicyV1Provision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigned

`func (o *C1ApiPolicyV1Provision) GetAssigned() bool`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *C1ApiPolicyV1Provision) GetAssignedOk() (*bool, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *C1ApiPolicyV1Provision) SetAssigned(v bool)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *C1ApiPolicyV1Provision) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### GetProvisionPolicy

`func (o *C1ApiPolicyV1Provision) GetProvisionPolicy() C1ApiPolicyV1ProvisionPolicy`

GetProvisionPolicy returns the ProvisionPolicy field if non-nil, zero value otherwise.

### GetProvisionPolicyOk

`func (o *C1ApiPolicyV1Provision) GetProvisionPolicyOk() (*C1ApiPolicyV1ProvisionPolicy, bool)`

GetProvisionPolicyOk returns a tuple with the ProvisionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionPolicy

`func (o *C1ApiPolicyV1Provision) SetProvisionPolicy(v C1ApiPolicyV1ProvisionPolicy)`

SetProvisionPolicy sets ProvisionPolicy field to given value.

### HasProvisionPolicy

`func (o *C1ApiPolicyV1Provision) HasProvisionPolicy() bool`

HasProvisionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


