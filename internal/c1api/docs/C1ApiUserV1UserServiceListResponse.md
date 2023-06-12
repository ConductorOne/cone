# C1ApiUserV1UserServiceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expanded** | Pointer to [**[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner**](C1ApiAppV1AppResourceServiceGetResponseExpandedInner.md) | The expanded field. | [optional] 
**List** | Pointer to [**[]C1ApiUserV1UserView**](C1ApiUserV1UserView.md) | The list field. | [optional] 
**NextPageToken** | Pointer to **string** | The nextPageToken field. | [optional] 
**NotificationToken** | Pointer to **string** | The notificationToken field. | [optional] 

## Methods

### NewC1ApiUserV1UserServiceListResponse

`func NewC1ApiUserV1UserServiceListResponse() *C1ApiUserV1UserServiceListResponse`

NewC1ApiUserV1UserServiceListResponse instantiates a new C1ApiUserV1UserServiceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiUserV1UserServiceListResponseWithDefaults

`func NewC1ApiUserV1UserServiceListResponseWithDefaults() *C1ApiUserV1UserServiceListResponse`

NewC1ApiUserV1UserServiceListResponseWithDefaults instantiates a new C1ApiUserV1UserServiceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpanded

`func (o *C1ApiUserV1UserServiceListResponse) GetExpanded() []C1ApiAppV1AppResourceServiceGetResponseExpandedInner`

GetExpanded returns the Expanded field if non-nil, zero value otherwise.

### GetExpandedOk

`func (o *C1ApiUserV1UserServiceListResponse) GetExpandedOk() (*[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner, bool)`

GetExpandedOk returns a tuple with the Expanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanded

`func (o *C1ApiUserV1UserServiceListResponse) SetExpanded(v []C1ApiAppV1AppResourceServiceGetResponseExpandedInner)`

SetExpanded sets Expanded field to given value.

### HasExpanded

`func (o *C1ApiUserV1UserServiceListResponse) HasExpanded() bool`

HasExpanded returns a boolean if a field has been set.

### SetExpandedNil

`func (o *C1ApiUserV1UserServiceListResponse) SetExpandedNil(b bool)`

 SetExpandedNil sets the value for Expanded to be an explicit nil

### UnsetExpanded
`func (o *C1ApiUserV1UserServiceListResponse) UnsetExpanded()`

UnsetExpanded ensures that no value is present for Expanded, not even an explicit nil
### GetList

`func (o *C1ApiUserV1UserServiceListResponse) GetList() []C1ApiUserV1UserView`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *C1ApiUserV1UserServiceListResponse) GetListOk() (*[]C1ApiUserV1UserView, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *C1ApiUserV1UserServiceListResponse) SetList(v []C1ApiUserV1UserView)`

SetList sets List field to given value.

### HasList

`func (o *C1ApiUserV1UserServiceListResponse) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *C1ApiUserV1UserServiceListResponse) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *C1ApiUserV1UserServiceListResponse) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetNextPageToken

`func (o *C1ApiUserV1UserServiceListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *C1ApiUserV1UserServiceListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *C1ApiUserV1UserServiceListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *C1ApiUserV1UserServiceListResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetNotificationToken

`func (o *C1ApiUserV1UserServiceListResponse) GetNotificationToken() string`

GetNotificationToken returns the NotificationToken field if non-nil, zero value otherwise.

### GetNotificationTokenOk

`func (o *C1ApiUserV1UserServiceListResponse) GetNotificationTokenOk() (*string, bool)`

GetNotificationTokenOk returns a tuple with the NotificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationToken

`func (o *C1ApiUserV1UserServiceListResponse) SetNotificationToken(v string)`

SetNotificationToken sets NotificationToken field to given value.

### HasNotificationToken

`func (o *C1ApiUserV1UserServiceListResponse) HasNotificationToken() bool`

HasNotificationToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


