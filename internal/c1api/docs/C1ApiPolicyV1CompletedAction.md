# C1ApiPolicyV1CompletedAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**Entitlements** | Pointer to [**[]C1ApiPolicyV1AppEntitlementReference**](C1ApiPolicyV1AppEntitlementReference.md) | The entitlements field. | [optional] 
**UserId** | Pointer to **string** | The userId field. | [optional] 

## Methods

### NewC1ApiPolicyV1CompletedAction

`func NewC1ApiPolicyV1CompletedAction() *C1ApiPolicyV1CompletedAction`

NewC1ApiPolicyV1CompletedAction instantiates a new C1ApiPolicyV1CompletedAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1CompletedActionWithDefaults

`func NewC1ApiPolicyV1CompletedActionWithDefaults() *C1ApiPolicyV1CompletedAction`

NewC1ApiPolicyV1CompletedActionWithDefaults instantiates a new C1ApiPolicyV1CompletedAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *C1ApiPolicyV1CompletedAction) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *C1ApiPolicyV1CompletedAction) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *C1ApiPolicyV1CompletedAction) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *C1ApiPolicyV1CompletedAction) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetEntitlements

`func (o *C1ApiPolicyV1CompletedAction) GetEntitlements() []C1ApiPolicyV1AppEntitlementReference`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *C1ApiPolicyV1CompletedAction) GetEntitlementsOk() (*[]C1ApiPolicyV1AppEntitlementReference, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *C1ApiPolicyV1CompletedAction) SetEntitlements(v []C1ApiPolicyV1AppEntitlementReference)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *C1ApiPolicyV1CompletedAction) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### SetEntitlementsNil

`func (o *C1ApiPolicyV1CompletedAction) SetEntitlementsNil(b bool)`

 SetEntitlementsNil sets the value for Entitlements to be an explicit nil

### UnsetEntitlements
`func (o *C1ApiPolicyV1CompletedAction) UnsetEntitlements()`

UnsetEntitlements ensures that no value is present for Entitlements, not even an explicit nil
### GetUserId

`func (o *C1ApiPolicyV1CompletedAction) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *C1ApiPolicyV1CompletedAction) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *C1ApiPolicyV1CompletedAction) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *C1ApiPolicyV1CompletedAction) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


