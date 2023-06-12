# C1ApiPolicyV1ListPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]C1ApiPolicyV1Policy**](C1ApiPolicyV1Policy.md) | The list field. | [optional] 
**NextPageToken** | Pointer to **string** | The nextPageToken field. | [optional] 
**NotificationToken** | Pointer to **string** | The notificationToken field. | [optional] 

## Methods

### NewC1ApiPolicyV1ListPolicyResponse

`func NewC1ApiPolicyV1ListPolicyResponse() *C1ApiPolicyV1ListPolicyResponse`

NewC1ApiPolicyV1ListPolicyResponse instantiates a new C1ApiPolicyV1ListPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1ListPolicyResponseWithDefaults

`func NewC1ApiPolicyV1ListPolicyResponseWithDefaults() *C1ApiPolicyV1ListPolicyResponse`

NewC1ApiPolicyV1ListPolicyResponseWithDefaults instantiates a new C1ApiPolicyV1ListPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *C1ApiPolicyV1ListPolicyResponse) GetList() []C1ApiPolicyV1Policy`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *C1ApiPolicyV1ListPolicyResponse) GetListOk() (*[]C1ApiPolicyV1Policy, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *C1ApiPolicyV1ListPolicyResponse) SetList(v []C1ApiPolicyV1Policy)`

SetList sets List field to given value.

### HasList

`func (o *C1ApiPolicyV1ListPolicyResponse) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *C1ApiPolicyV1ListPolicyResponse) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *C1ApiPolicyV1ListPolicyResponse) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetNextPageToken

`func (o *C1ApiPolicyV1ListPolicyResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *C1ApiPolicyV1ListPolicyResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *C1ApiPolicyV1ListPolicyResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *C1ApiPolicyV1ListPolicyResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetNotificationToken

`func (o *C1ApiPolicyV1ListPolicyResponse) GetNotificationToken() string`

GetNotificationToken returns the NotificationToken field if non-nil, zero value otherwise.

### GetNotificationTokenOk

`func (o *C1ApiPolicyV1ListPolicyResponse) GetNotificationTokenOk() (*string, bool)`

GetNotificationTokenOk returns a tuple with the NotificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationToken

`func (o *C1ApiPolicyV1ListPolicyResponse) SetNotificationToken(v string)`

SetNotificationToken sets NotificationToken field to given value.

### HasNotificationToken

`func (o *C1ApiPolicyV1ListPolicyResponse) HasNotificationToken() bool`

HasNotificationToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


