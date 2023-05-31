# C1ApiPolicyV1ProvisionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to **interface{}** | The ConnectorProvision message. | [optional] 
**Delegated** | Pointer to [**C1ApiPolicyV1DelegatedProvision**](C1ApiPolicyV1DelegatedProvision.md) |  | [optional] 
**Manual** | Pointer to [**C1ApiPolicyV1ManualProvision**](C1ApiPolicyV1ManualProvision.md) |  | [optional] 

## Methods

### NewC1ApiPolicyV1ProvisionPolicy

`func NewC1ApiPolicyV1ProvisionPolicy() *C1ApiPolicyV1ProvisionPolicy`

NewC1ApiPolicyV1ProvisionPolicy instantiates a new C1ApiPolicyV1ProvisionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1ProvisionPolicyWithDefaults

`func NewC1ApiPolicyV1ProvisionPolicyWithDefaults() *C1ApiPolicyV1ProvisionPolicy`

NewC1ApiPolicyV1ProvisionPolicyWithDefaults instantiates a new C1ApiPolicyV1ProvisionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *C1ApiPolicyV1ProvisionPolicy) GetConnector() interface{}`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *C1ApiPolicyV1ProvisionPolicy) GetConnectorOk() (*interface{}, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *C1ApiPolicyV1ProvisionPolicy) SetConnector(v interface{})`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *C1ApiPolicyV1ProvisionPolicy) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### SetConnectorNil

`func (o *C1ApiPolicyV1ProvisionPolicy) SetConnectorNil(b bool)`

 SetConnectorNil sets the value for Connector to be an explicit nil

### UnsetConnector
`func (o *C1ApiPolicyV1ProvisionPolicy) UnsetConnector()`

UnsetConnector ensures that no value is present for Connector, not even an explicit nil
### GetDelegated

`func (o *C1ApiPolicyV1ProvisionPolicy) GetDelegated() C1ApiPolicyV1DelegatedProvision`

GetDelegated returns the Delegated field if non-nil, zero value otherwise.

### GetDelegatedOk

`func (o *C1ApiPolicyV1ProvisionPolicy) GetDelegatedOk() (*C1ApiPolicyV1DelegatedProvision, bool)`

GetDelegatedOk returns a tuple with the Delegated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegated

`func (o *C1ApiPolicyV1ProvisionPolicy) SetDelegated(v C1ApiPolicyV1DelegatedProvision)`

SetDelegated sets Delegated field to given value.

### HasDelegated

`func (o *C1ApiPolicyV1ProvisionPolicy) HasDelegated() bool`

HasDelegated returns a boolean if a field has been set.

### GetManual

`func (o *C1ApiPolicyV1ProvisionPolicy) GetManual() C1ApiPolicyV1ManualProvision`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *C1ApiPolicyV1ProvisionPolicy) GetManualOk() (*C1ApiPolicyV1ManualProvision, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *C1ApiPolicyV1ProvisionPolicy) SetManual(v C1ApiPolicyV1ManualProvision)`

SetManual sets Manual field to given value.

### HasManual

`func (o *C1ApiPolicyV1ProvisionPolicy) HasManual() bool`

HasManual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


