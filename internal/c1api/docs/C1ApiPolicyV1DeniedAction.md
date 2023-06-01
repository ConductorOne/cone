# C1ApiPolicyV1DeniedAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeniedAt** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **string** | The userId field. | [optional] 

## Methods

### NewC1ApiPolicyV1DeniedAction

`func NewC1ApiPolicyV1DeniedAction() *C1ApiPolicyV1DeniedAction`

NewC1ApiPolicyV1DeniedAction instantiates a new C1ApiPolicyV1DeniedAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1DeniedActionWithDefaults

`func NewC1ApiPolicyV1DeniedActionWithDefaults() *C1ApiPolicyV1DeniedAction`

NewC1ApiPolicyV1DeniedActionWithDefaults instantiates a new C1ApiPolicyV1DeniedAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeniedAt

`func (o *C1ApiPolicyV1DeniedAction) GetDeniedAt() time.Time`

GetDeniedAt returns the DeniedAt field if non-nil, zero value otherwise.

### GetDeniedAtOk

`func (o *C1ApiPolicyV1DeniedAction) GetDeniedAtOk() (*time.Time, bool)`

GetDeniedAtOk returns a tuple with the DeniedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedAt

`func (o *C1ApiPolicyV1DeniedAction) SetDeniedAt(v time.Time)`

SetDeniedAt sets DeniedAt field to given value.

### HasDeniedAt

`func (o *C1ApiPolicyV1DeniedAction) HasDeniedAt() bool`

HasDeniedAt returns a boolean if a field has been set.

### GetUserId

`func (o *C1ApiPolicyV1DeniedAction) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *C1ApiPolicyV1DeniedAction) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *C1ApiPolicyV1DeniedAction) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *C1ApiPolicyV1DeniedAction) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


