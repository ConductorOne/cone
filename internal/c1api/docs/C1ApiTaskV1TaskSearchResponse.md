# C1ApiTaskV1TaskSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expanded** | Pointer to [**[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner**](C1ApiAppV1AppResourceServiceGetResponseExpandedInner.md) | The expanded field. | [optional] 
**List** | Pointer to [**[]C1ApiTaskV1TaskView**](C1ApiTaskV1TaskView.md) | The list field. | [optional] 
**NextPageToken** | Pointer to **string** | The nextPageToken field. | [optional] 

## Methods

### NewC1ApiTaskV1TaskSearchResponse

`func NewC1ApiTaskV1TaskSearchResponse() *C1ApiTaskV1TaskSearchResponse`

NewC1ApiTaskV1TaskSearchResponse instantiates a new C1ApiTaskV1TaskSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiTaskV1TaskSearchResponseWithDefaults

`func NewC1ApiTaskV1TaskSearchResponseWithDefaults() *C1ApiTaskV1TaskSearchResponse`

NewC1ApiTaskV1TaskSearchResponseWithDefaults instantiates a new C1ApiTaskV1TaskSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpanded

`func (o *C1ApiTaskV1TaskSearchResponse) GetExpanded() []C1ApiAppV1AppResourceServiceGetResponseExpandedInner`

GetExpanded returns the Expanded field if non-nil, zero value otherwise.

### GetExpandedOk

`func (o *C1ApiTaskV1TaskSearchResponse) GetExpandedOk() (*[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner, bool)`

GetExpandedOk returns a tuple with the Expanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanded

`func (o *C1ApiTaskV1TaskSearchResponse) SetExpanded(v []C1ApiAppV1AppResourceServiceGetResponseExpandedInner)`

SetExpanded sets Expanded field to given value.

### HasExpanded

`func (o *C1ApiTaskV1TaskSearchResponse) HasExpanded() bool`

HasExpanded returns a boolean if a field has been set.

### SetExpandedNil

`func (o *C1ApiTaskV1TaskSearchResponse) SetExpandedNil(b bool)`

 SetExpandedNil sets the value for Expanded to be an explicit nil

### UnsetExpanded
`func (o *C1ApiTaskV1TaskSearchResponse) UnsetExpanded()`

UnsetExpanded ensures that no value is present for Expanded, not even an explicit nil
### GetList

`func (o *C1ApiTaskV1TaskSearchResponse) GetList() []C1ApiTaskV1TaskView`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *C1ApiTaskV1TaskSearchResponse) GetListOk() (*[]C1ApiTaskV1TaskView, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *C1ApiTaskV1TaskSearchResponse) SetList(v []C1ApiTaskV1TaskView)`

SetList sets List field to given value.

### HasList

`func (o *C1ApiTaskV1TaskSearchResponse) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *C1ApiTaskV1TaskSearchResponse) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *C1ApiTaskV1TaskSearchResponse) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetNextPageToken

`func (o *C1ApiTaskV1TaskSearchResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *C1ApiTaskV1TaskSearchResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *C1ApiTaskV1TaskSearchResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *C1ApiTaskV1TaskSearchResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


