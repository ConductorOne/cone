# C1ApiAppV1AppResourceServiceGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppResourceView** | Pointer to [**C1ApiAppV1AppResourceView**](C1ApiAppV1AppResourceView.md) |  | [optional] 
**Expanded** | Pointer to [**[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner**](C1ApiAppV1AppResourceServiceGetResponseExpandedInner.md) | The expanded field. | [optional] 

## Methods

### NewC1ApiAppV1AppResourceServiceGetResponse

`func NewC1ApiAppV1AppResourceServiceGetResponse() *C1ApiAppV1AppResourceServiceGetResponse`

NewC1ApiAppV1AppResourceServiceGetResponse instantiates a new C1ApiAppV1AppResourceServiceGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiAppV1AppResourceServiceGetResponseWithDefaults

`func NewC1ApiAppV1AppResourceServiceGetResponseWithDefaults() *C1ApiAppV1AppResourceServiceGetResponse`

NewC1ApiAppV1AppResourceServiceGetResponseWithDefaults instantiates a new C1ApiAppV1AppResourceServiceGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppResourceView

`func (o *C1ApiAppV1AppResourceServiceGetResponse) GetAppResourceView() C1ApiAppV1AppResourceView`

GetAppResourceView returns the AppResourceView field if non-nil, zero value otherwise.

### GetAppResourceViewOk

`func (o *C1ApiAppV1AppResourceServiceGetResponse) GetAppResourceViewOk() (*C1ApiAppV1AppResourceView, bool)`

GetAppResourceViewOk returns a tuple with the AppResourceView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppResourceView

`func (o *C1ApiAppV1AppResourceServiceGetResponse) SetAppResourceView(v C1ApiAppV1AppResourceView)`

SetAppResourceView sets AppResourceView field to given value.

### HasAppResourceView

`func (o *C1ApiAppV1AppResourceServiceGetResponse) HasAppResourceView() bool`

HasAppResourceView returns a boolean if a field has been set.

### GetExpanded

`func (o *C1ApiAppV1AppResourceServiceGetResponse) GetExpanded() []C1ApiAppV1AppResourceServiceGetResponseExpandedInner`

GetExpanded returns the Expanded field if non-nil, zero value otherwise.

### GetExpandedOk

`func (o *C1ApiAppV1AppResourceServiceGetResponse) GetExpandedOk() (*[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner, bool)`

GetExpandedOk returns a tuple with the Expanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanded

`func (o *C1ApiAppV1AppResourceServiceGetResponse) SetExpanded(v []C1ApiAppV1AppResourceServiceGetResponseExpandedInner)`

SetExpanded sets Expanded field to given value.

### HasExpanded

`func (o *C1ApiAppV1AppResourceServiceGetResponse) HasExpanded() bool`

HasExpanded returns a boolean if a field has been set.

### SetExpandedNil

`func (o *C1ApiAppV1AppResourceServiceGetResponse) SetExpandedNil(b bool)`

 SetExpandedNil sets the value for Expanded to be an explicit nil

### UnsetExpanded
`func (o *C1ApiAppV1AppResourceServiceGetResponse) UnsetExpanded()`

UnsetExpanded ensures that no value is present for Expanded, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


