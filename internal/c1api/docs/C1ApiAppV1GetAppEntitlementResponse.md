# C1ApiAppV1GetAppEntitlementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppEntitlementView** | Pointer to [**C1ApiAppV1AppEntitlementView**](C1ApiAppV1AppEntitlementView.md) |  | [optional] 
**Expanded** | Pointer to [**[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner**](C1ApiAppV1AppResourceServiceGetResponseExpandedInner.md) | The expanded field. | [optional] 

## Methods

### NewC1ApiAppV1GetAppEntitlementResponse

`func NewC1ApiAppV1GetAppEntitlementResponse() *C1ApiAppV1GetAppEntitlementResponse`

NewC1ApiAppV1GetAppEntitlementResponse instantiates a new C1ApiAppV1GetAppEntitlementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiAppV1GetAppEntitlementResponseWithDefaults

`func NewC1ApiAppV1GetAppEntitlementResponseWithDefaults() *C1ApiAppV1GetAppEntitlementResponse`

NewC1ApiAppV1GetAppEntitlementResponseWithDefaults instantiates a new C1ApiAppV1GetAppEntitlementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppEntitlementView

`func (o *C1ApiAppV1GetAppEntitlementResponse) GetAppEntitlementView() C1ApiAppV1AppEntitlementView`

GetAppEntitlementView returns the AppEntitlementView field if non-nil, zero value otherwise.

### GetAppEntitlementViewOk

`func (o *C1ApiAppV1GetAppEntitlementResponse) GetAppEntitlementViewOk() (*C1ApiAppV1AppEntitlementView, bool)`

GetAppEntitlementViewOk returns a tuple with the AppEntitlementView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEntitlementView

`func (o *C1ApiAppV1GetAppEntitlementResponse) SetAppEntitlementView(v C1ApiAppV1AppEntitlementView)`

SetAppEntitlementView sets AppEntitlementView field to given value.

### HasAppEntitlementView

`func (o *C1ApiAppV1GetAppEntitlementResponse) HasAppEntitlementView() bool`

HasAppEntitlementView returns a boolean if a field has been set.

### GetExpanded

`func (o *C1ApiAppV1GetAppEntitlementResponse) GetExpanded() []C1ApiAppV1AppResourceServiceGetResponseExpandedInner`

GetExpanded returns the Expanded field if non-nil, zero value otherwise.

### GetExpandedOk

`func (o *C1ApiAppV1GetAppEntitlementResponse) GetExpandedOk() (*[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner, bool)`

GetExpandedOk returns a tuple with the Expanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanded

`func (o *C1ApiAppV1GetAppEntitlementResponse) SetExpanded(v []C1ApiAppV1AppResourceServiceGetResponseExpandedInner)`

SetExpanded sets Expanded field to given value.

### HasExpanded

`func (o *C1ApiAppV1GetAppEntitlementResponse) HasExpanded() bool`

HasExpanded returns a boolean if a field has been set.

### SetExpandedNil

`func (o *C1ApiAppV1GetAppEntitlementResponse) SetExpandedNil(b bool)`

 SetExpandedNil sets the value for Expanded to be an explicit nil

### UnsetExpanded
`func (o *C1ApiAppV1GetAppEntitlementResponse) UnsetExpanded()`

UnsetExpanded ensures that no value is present for Expanded, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


