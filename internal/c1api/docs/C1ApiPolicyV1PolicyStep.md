# C1ApiPolicyV1PolicyStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approval** | Pointer to [**NullableC1ApiPolicyV1Approval**](C1ApiPolicyV1Approval.md) |  | [optional] 
**Notify** | Pointer to **map[string]interface{}** | The Notification message. | [optional] 
**Provision** | Pointer to [**NullableC1ApiPolicyV1Provision**](C1ApiPolicyV1Provision.md) |  | [optional] 
**Webhook** | Pointer to **map[string]interface{}** | The Webhook message. | [optional] 

## Methods

### NewC1ApiPolicyV1PolicyStep

`func NewC1ApiPolicyV1PolicyStep() *C1ApiPolicyV1PolicyStep`

NewC1ApiPolicyV1PolicyStep instantiates a new C1ApiPolicyV1PolicyStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1PolicyStepWithDefaults

`func NewC1ApiPolicyV1PolicyStepWithDefaults() *C1ApiPolicyV1PolicyStep`

NewC1ApiPolicyV1PolicyStepWithDefaults instantiates a new C1ApiPolicyV1PolicyStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproval

`func (o *C1ApiPolicyV1PolicyStep) GetApproval() C1ApiPolicyV1Approval`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *C1ApiPolicyV1PolicyStep) GetApprovalOk() (*C1ApiPolicyV1Approval, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *C1ApiPolicyV1PolicyStep) SetApproval(v C1ApiPolicyV1Approval)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *C1ApiPolicyV1PolicyStep) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### SetApprovalNil

`func (o *C1ApiPolicyV1PolicyStep) SetApprovalNil(b bool)`

 SetApprovalNil sets the value for Approval to be an explicit nil

### UnsetApproval
`func (o *C1ApiPolicyV1PolicyStep) UnsetApproval()`

UnsetApproval ensures that no value is present for Approval, not even an explicit nil
### GetNotify

`func (o *C1ApiPolicyV1PolicyStep) GetNotify() map[string]interface{}`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *C1ApiPolicyV1PolicyStep) GetNotifyOk() (*map[string]interface{}, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *C1ApiPolicyV1PolicyStep) SetNotify(v map[string]interface{})`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *C1ApiPolicyV1PolicyStep) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### SetNotifyNil

`func (o *C1ApiPolicyV1PolicyStep) SetNotifyNil(b bool)`

 SetNotifyNil sets the value for Notify to be an explicit nil

### UnsetNotify
`func (o *C1ApiPolicyV1PolicyStep) UnsetNotify()`

UnsetNotify ensures that no value is present for Notify, not even an explicit nil
### GetProvision

`func (o *C1ApiPolicyV1PolicyStep) GetProvision() C1ApiPolicyV1Provision`

GetProvision returns the Provision field if non-nil, zero value otherwise.

### GetProvisionOk

`func (o *C1ApiPolicyV1PolicyStep) GetProvisionOk() (*C1ApiPolicyV1Provision, bool)`

GetProvisionOk returns a tuple with the Provision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvision

`func (o *C1ApiPolicyV1PolicyStep) SetProvision(v C1ApiPolicyV1Provision)`

SetProvision sets Provision field to given value.

### HasProvision

`func (o *C1ApiPolicyV1PolicyStep) HasProvision() bool`

HasProvision returns a boolean if a field has been set.

### SetProvisionNil

`func (o *C1ApiPolicyV1PolicyStep) SetProvisionNil(b bool)`

 SetProvisionNil sets the value for Provision to be an explicit nil

### UnsetProvision
`func (o *C1ApiPolicyV1PolicyStep) UnsetProvision()`

UnsetProvision ensures that no value is present for Provision, not even an explicit nil
### GetWebhook

`func (o *C1ApiPolicyV1PolicyStep) GetWebhook() map[string]interface{}`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *C1ApiPolicyV1PolicyStep) GetWebhookOk() (*map[string]interface{}, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *C1ApiPolicyV1PolicyStep) SetWebhook(v map[string]interface{})`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *C1ApiPolicyV1PolicyStep) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### SetWebhookNil

`func (o *C1ApiPolicyV1PolicyStep) SetWebhookNil(b bool)`

 SetWebhookNil sets the value for Webhook to be an explicit nil

### UnsetWebhook
`func (o *C1ApiPolicyV1PolicyStep) UnsetWebhook()`

UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


