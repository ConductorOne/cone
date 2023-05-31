# C1ApiUserV1UserServiceGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expanded** | Pointer to [**[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner**](C1ApiAppV1AppResourceServiceGetResponseExpandedInner.md) | The expanded field. | [optional] 
**UserView** | Pointer to [**C1ApiUserV1UserView**](C1ApiUserV1UserView.md) |  | [optional] 

## Methods

### NewC1ApiUserV1UserServiceGetResponse

`func NewC1ApiUserV1UserServiceGetResponse() *C1ApiUserV1UserServiceGetResponse`

NewC1ApiUserV1UserServiceGetResponse instantiates a new C1ApiUserV1UserServiceGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiUserV1UserServiceGetResponseWithDefaults

`func NewC1ApiUserV1UserServiceGetResponseWithDefaults() *C1ApiUserV1UserServiceGetResponse`

NewC1ApiUserV1UserServiceGetResponseWithDefaults instantiates a new C1ApiUserV1UserServiceGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpanded

`func (o *C1ApiUserV1UserServiceGetResponse) GetExpanded() []C1ApiAppV1AppResourceServiceGetResponseExpandedInner`

GetExpanded returns the Expanded field if non-nil, zero value otherwise.

### GetExpandedOk

`func (o *C1ApiUserV1UserServiceGetResponse) GetExpandedOk() (*[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner, bool)`

GetExpandedOk returns a tuple with the Expanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanded

`func (o *C1ApiUserV1UserServiceGetResponse) SetExpanded(v []C1ApiAppV1AppResourceServiceGetResponseExpandedInner)`

SetExpanded sets Expanded field to given value.

### HasExpanded

`func (o *C1ApiUserV1UserServiceGetResponse) HasExpanded() bool`

HasExpanded returns a boolean if a field has been set.

### SetExpandedNil

`func (o *C1ApiUserV1UserServiceGetResponse) SetExpandedNil(b bool)`

 SetExpandedNil sets the value for Expanded to be an explicit nil

### UnsetExpanded
`func (o *C1ApiUserV1UserServiceGetResponse) UnsetExpanded()`

UnsetExpanded ensures that no value is present for Expanded, not even an explicit nil
### GetUserView

`func (o *C1ApiUserV1UserServiceGetResponse) GetUserView() C1ApiUserV1UserView`

GetUserView returns the UserView field if non-nil, zero value otherwise.

### GetUserViewOk

`func (o *C1ApiUserV1UserServiceGetResponse) GetUserViewOk() (*C1ApiUserV1UserView, bool)`

GetUserViewOk returns a tuple with the UserView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserView

`func (o *C1ApiUserV1UserServiceGetResponse) SetUserView(v C1ApiUserV1UserView)`

SetUserView sets UserView field to given value.

### HasUserView

`func (o *C1ApiUserV1UserServiceGetResponse) HasUserView() bool`

HasUserView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


