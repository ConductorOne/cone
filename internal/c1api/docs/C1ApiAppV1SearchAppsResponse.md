# C1ApiAppV1SearchAppsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]C1ApiAppV1App**](C1ApiAppV1App.md) | The list field. | [optional] 
**NextPageToken** | Pointer to **string** | The nextPageToken field. | [optional] 
**NotificationToken** | Pointer to **string** | The notificationToken field. | [optional] 

## Methods

### NewC1ApiAppV1SearchAppsResponse

`func NewC1ApiAppV1SearchAppsResponse() *C1ApiAppV1SearchAppsResponse`

NewC1ApiAppV1SearchAppsResponse instantiates a new C1ApiAppV1SearchAppsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiAppV1SearchAppsResponseWithDefaults

`func NewC1ApiAppV1SearchAppsResponseWithDefaults() *C1ApiAppV1SearchAppsResponse`

NewC1ApiAppV1SearchAppsResponseWithDefaults instantiates a new C1ApiAppV1SearchAppsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *C1ApiAppV1SearchAppsResponse) GetList() []C1ApiAppV1App`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *C1ApiAppV1SearchAppsResponse) GetListOk() (*[]C1ApiAppV1App, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *C1ApiAppV1SearchAppsResponse) SetList(v []C1ApiAppV1App)`

SetList sets List field to given value.

### HasList

`func (o *C1ApiAppV1SearchAppsResponse) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *C1ApiAppV1SearchAppsResponse) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *C1ApiAppV1SearchAppsResponse) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetNextPageToken

`func (o *C1ApiAppV1SearchAppsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *C1ApiAppV1SearchAppsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *C1ApiAppV1SearchAppsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *C1ApiAppV1SearchAppsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetNotificationToken

`func (o *C1ApiAppV1SearchAppsResponse) GetNotificationToken() string`

GetNotificationToken returns the NotificationToken field if non-nil, zero value otherwise.

### GetNotificationTokenOk

`func (o *C1ApiAppV1SearchAppsResponse) GetNotificationTokenOk() (*string, bool)`

GetNotificationTokenOk returns a tuple with the NotificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationToken

`func (o *C1ApiAppV1SearchAppsResponse) SetNotificationToken(v string)`

SetNotificationToken sets NotificationToken field to given value.

### HasNotificationToken

`func (o *C1ApiAppV1SearchAppsResponse) HasNotificationToken() bool`

HasNotificationToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


