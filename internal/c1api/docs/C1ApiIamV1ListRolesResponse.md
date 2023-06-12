# C1ApiIamV1ListRolesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]C1ApiIamV1Role**](C1ApiIamV1Role.md) | The list field. | [optional] 
**NextPageToken** | Pointer to **string** | The nextPageToken field. | [optional] 
**NotificationToken** | Pointer to **string** | The notificationToken field. | [optional] 

## Methods

### NewC1ApiIamV1ListRolesResponse

`func NewC1ApiIamV1ListRolesResponse() *C1ApiIamV1ListRolesResponse`

NewC1ApiIamV1ListRolesResponse instantiates a new C1ApiIamV1ListRolesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiIamV1ListRolesResponseWithDefaults

`func NewC1ApiIamV1ListRolesResponseWithDefaults() *C1ApiIamV1ListRolesResponse`

NewC1ApiIamV1ListRolesResponseWithDefaults instantiates a new C1ApiIamV1ListRolesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *C1ApiIamV1ListRolesResponse) GetList() []C1ApiIamV1Role`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *C1ApiIamV1ListRolesResponse) GetListOk() (*[]C1ApiIamV1Role, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *C1ApiIamV1ListRolesResponse) SetList(v []C1ApiIamV1Role)`

SetList sets List field to given value.

### HasList

`func (o *C1ApiIamV1ListRolesResponse) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *C1ApiIamV1ListRolesResponse) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *C1ApiIamV1ListRolesResponse) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetNextPageToken

`func (o *C1ApiIamV1ListRolesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *C1ApiIamV1ListRolesResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *C1ApiIamV1ListRolesResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *C1ApiIamV1ListRolesResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetNotificationToken

`func (o *C1ApiIamV1ListRolesResponse) GetNotificationToken() string`

GetNotificationToken returns the NotificationToken field if non-nil, zero value otherwise.

### GetNotificationTokenOk

`func (o *C1ApiIamV1ListRolesResponse) GetNotificationTokenOk() (*string, bool)`

GetNotificationTokenOk returns a tuple with the NotificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationToken

`func (o *C1ApiIamV1ListRolesResponse) SetNotificationToken(v string)`

SetNotificationToken sets NotificationToken field to given value.

### HasNotificationToken

`func (o *C1ApiIamV1ListRolesResponse) HasNotificationToken() bool`

HasNotificationToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


