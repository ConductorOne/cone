# C1ApiDirectoryV1DirectoryServiceGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryView** | Pointer to [**C1ApiDirectoryV1DirectoryView**](C1ApiDirectoryV1DirectoryView.md) |  | [optional] 
**Expanded** | Pointer to [**[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner**](C1ApiAppV1AppResourceServiceGetResponseExpandedInner.md) | The expanded field. | [optional] 

## Methods

### NewC1ApiDirectoryV1DirectoryServiceGetResponse

`func NewC1ApiDirectoryV1DirectoryServiceGetResponse() *C1ApiDirectoryV1DirectoryServiceGetResponse`

NewC1ApiDirectoryV1DirectoryServiceGetResponse instantiates a new C1ApiDirectoryV1DirectoryServiceGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiDirectoryV1DirectoryServiceGetResponseWithDefaults

`func NewC1ApiDirectoryV1DirectoryServiceGetResponseWithDefaults() *C1ApiDirectoryV1DirectoryServiceGetResponse`

NewC1ApiDirectoryV1DirectoryServiceGetResponseWithDefaults instantiates a new C1ApiDirectoryV1DirectoryServiceGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryView

`func (o *C1ApiDirectoryV1DirectoryServiceGetResponse) GetDirectoryView() C1ApiDirectoryV1DirectoryView`

GetDirectoryView returns the DirectoryView field if non-nil, zero value otherwise.

### GetDirectoryViewOk

`func (o *C1ApiDirectoryV1DirectoryServiceGetResponse) GetDirectoryViewOk() (*C1ApiDirectoryV1DirectoryView, bool)`

GetDirectoryViewOk returns a tuple with the DirectoryView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryView

`func (o *C1ApiDirectoryV1DirectoryServiceGetResponse) SetDirectoryView(v C1ApiDirectoryV1DirectoryView)`

SetDirectoryView sets DirectoryView field to given value.

### HasDirectoryView

`func (o *C1ApiDirectoryV1DirectoryServiceGetResponse) HasDirectoryView() bool`

HasDirectoryView returns a boolean if a field has been set.

### GetExpanded

`func (o *C1ApiDirectoryV1DirectoryServiceGetResponse) GetExpanded() []C1ApiAppV1AppResourceServiceGetResponseExpandedInner`

GetExpanded returns the Expanded field if non-nil, zero value otherwise.

### GetExpandedOk

`func (o *C1ApiDirectoryV1DirectoryServiceGetResponse) GetExpandedOk() (*[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner, bool)`

GetExpandedOk returns a tuple with the Expanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanded

`func (o *C1ApiDirectoryV1DirectoryServiceGetResponse) SetExpanded(v []C1ApiAppV1AppResourceServiceGetResponseExpandedInner)`

SetExpanded sets Expanded field to given value.

### HasExpanded

`func (o *C1ApiDirectoryV1DirectoryServiceGetResponse) HasExpanded() bool`

HasExpanded returns a boolean if a field has been set.

### SetExpandedNil

`func (o *C1ApiDirectoryV1DirectoryServiceGetResponse) SetExpandedNil(b bool)`

 SetExpandedNil sets the value for Expanded to be an explicit nil

### UnsetExpanded
`func (o *C1ApiDirectoryV1DirectoryServiceGetResponse) UnsetExpanded()`

UnsetExpanded ensures that no value is present for Expanded, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


