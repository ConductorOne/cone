# C1ApiPolicyV1PolicyInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | Pointer to [**C1ApiPolicyV1PolicyStepInstance**](C1ApiPolicyV1PolicyStepInstance.md) |  | [optional] 
**History** | Pointer to [**[]C1ApiPolicyV1PolicyStepInstance**](C1ApiPolicyV1PolicyStepInstance.md) | The history field. | [optional] 
**Next** | Pointer to [**[]C1ApiPolicyV1PolicyStep**](C1ApiPolicyV1PolicyStep.md) | The next field. | [optional] 
**Policy** | Pointer to [**C1ApiPolicyV1Policy**](C1ApiPolicyV1Policy.md) |  | [optional] 

## Methods

### NewC1ApiPolicyV1PolicyInstance

`func NewC1ApiPolicyV1PolicyInstance() *C1ApiPolicyV1PolicyInstance`

NewC1ApiPolicyV1PolicyInstance instantiates a new C1ApiPolicyV1PolicyInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1PolicyInstanceWithDefaults

`func NewC1ApiPolicyV1PolicyInstanceWithDefaults() *C1ApiPolicyV1PolicyInstance`

NewC1ApiPolicyV1PolicyInstanceWithDefaults instantiates a new C1ApiPolicyV1PolicyInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *C1ApiPolicyV1PolicyInstance) GetCurrent() C1ApiPolicyV1PolicyStepInstance`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *C1ApiPolicyV1PolicyInstance) GetCurrentOk() (*C1ApiPolicyV1PolicyStepInstance, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *C1ApiPolicyV1PolicyInstance) SetCurrent(v C1ApiPolicyV1PolicyStepInstance)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *C1ApiPolicyV1PolicyInstance) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetHistory

`func (o *C1ApiPolicyV1PolicyInstance) GetHistory() []C1ApiPolicyV1PolicyStepInstance`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *C1ApiPolicyV1PolicyInstance) GetHistoryOk() (*[]C1ApiPolicyV1PolicyStepInstance, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *C1ApiPolicyV1PolicyInstance) SetHistory(v []C1ApiPolicyV1PolicyStepInstance)`

SetHistory sets History field to given value.

### HasHistory

`func (o *C1ApiPolicyV1PolicyInstance) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistoryNil

`func (o *C1ApiPolicyV1PolicyInstance) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *C1ApiPolicyV1PolicyInstance) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil
### GetNext

`func (o *C1ApiPolicyV1PolicyInstance) GetNext() []C1ApiPolicyV1PolicyStep`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *C1ApiPolicyV1PolicyInstance) GetNextOk() (*[]C1ApiPolicyV1PolicyStep, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *C1ApiPolicyV1PolicyInstance) SetNext(v []C1ApiPolicyV1PolicyStep)`

SetNext sets Next field to given value.

### HasNext

`func (o *C1ApiPolicyV1PolicyInstance) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *C1ApiPolicyV1PolicyInstance) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *C1ApiPolicyV1PolicyInstance) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPolicy

`func (o *C1ApiPolicyV1PolicyInstance) GetPolicy() C1ApiPolicyV1Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *C1ApiPolicyV1PolicyInstance) GetPolicyOk() (*C1ApiPolicyV1Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *C1ApiPolicyV1PolicyInstance) SetPolicy(v C1ApiPolicyV1Policy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *C1ApiPolicyV1PolicyInstance) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


