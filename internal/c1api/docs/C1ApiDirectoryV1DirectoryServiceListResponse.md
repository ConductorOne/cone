# C1ApiDirectoryV1DirectoryServiceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expanded** | Pointer to [**[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner**](C1ApiAppV1AppResourceServiceGetResponseExpandedInner.md) | The expanded field. | [optional] 
**List** | Pointer to [**[]C1ApiDirectoryV1DirectoryView**](C1ApiDirectoryV1DirectoryView.md) | The list field. | [optional] 
**NextPageToken** | Pointer to **string** | The nextPageToken field. | [optional] 

## Methods

### NewC1ApiDirectoryV1DirectoryServiceListResponse

`func NewC1ApiDirectoryV1DirectoryServiceListResponse() *C1ApiDirectoryV1DirectoryServiceListResponse`

NewC1ApiDirectoryV1DirectoryServiceListResponse instantiates a new C1ApiDirectoryV1DirectoryServiceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiDirectoryV1DirectoryServiceListResponseWithDefaults

`func NewC1ApiDirectoryV1DirectoryServiceListResponseWithDefaults() *C1ApiDirectoryV1DirectoryServiceListResponse`

NewC1ApiDirectoryV1DirectoryServiceListResponseWithDefaults instantiates a new C1ApiDirectoryV1DirectoryServiceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpanded

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) GetExpanded() []C1ApiAppV1AppResourceServiceGetResponseExpandedInner`

GetExpanded returns the Expanded field if non-nil, zero value otherwise.

### GetExpandedOk

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) GetExpandedOk() (*[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner, bool)`

GetExpandedOk returns a tuple with the Expanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanded

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) SetExpanded(v []C1ApiAppV1AppResourceServiceGetResponseExpandedInner)`

SetExpanded sets Expanded field to given value.

### HasExpanded

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) HasExpanded() bool`

HasExpanded returns a boolean if a field has been set.

### SetExpandedNil

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) SetExpandedNil(b bool)`

 SetExpandedNil sets the value for Expanded to be an explicit nil

### UnsetExpanded
`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) UnsetExpanded()`

UnsetExpanded ensures that no value is present for Expanded, not even an explicit nil
### GetList

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) GetList() []C1ApiDirectoryV1DirectoryView`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) GetListOk() (*[]C1ApiDirectoryV1DirectoryView, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) SetList(v []C1ApiDirectoryV1DirectoryView)`

SetList sets List field to given value.

### HasList

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetNextPageToken

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *C1ApiDirectoryV1DirectoryServiceListResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


