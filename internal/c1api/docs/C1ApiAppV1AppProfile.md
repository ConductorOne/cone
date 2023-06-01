# C1ApiAppV1AppProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppEntitlementIds** | Pointer to **[]string** | The appEntitlementIds field. | [optional] 
**AppId** | Pointer to **string** | The appId field. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** | The description field. | [optional] 
**DisplayName** | Pointer to **string** | The displayName field. | [optional] 
**Id** | Pointer to **string** | The id field. | [optional] 
**NotificationToken** | Pointer to **string** | The notificationToken field. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewC1ApiAppV1AppProfile

`func NewC1ApiAppV1AppProfile() *C1ApiAppV1AppProfile`

NewC1ApiAppV1AppProfile instantiates a new C1ApiAppV1AppProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiAppV1AppProfileWithDefaults

`func NewC1ApiAppV1AppProfileWithDefaults() *C1ApiAppV1AppProfile`

NewC1ApiAppV1AppProfileWithDefaults instantiates a new C1ApiAppV1AppProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppEntitlementIds

`func (o *C1ApiAppV1AppProfile) GetAppEntitlementIds() []string`

GetAppEntitlementIds returns the AppEntitlementIds field if non-nil, zero value otherwise.

### GetAppEntitlementIdsOk

`func (o *C1ApiAppV1AppProfile) GetAppEntitlementIdsOk() (*[]string, bool)`

GetAppEntitlementIdsOk returns a tuple with the AppEntitlementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEntitlementIds

`func (o *C1ApiAppV1AppProfile) SetAppEntitlementIds(v []string)`

SetAppEntitlementIds sets AppEntitlementIds field to given value.

### HasAppEntitlementIds

`func (o *C1ApiAppV1AppProfile) HasAppEntitlementIds() bool`

HasAppEntitlementIds returns a boolean if a field has been set.

### SetAppEntitlementIdsNil

`func (o *C1ApiAppV1AppProfile) SetAppEntitlementIdsNil(b bool)`

 SetAppEntitlementIdsNil sets the value for AppEntitlementIds to be an explicit nil

### UnsetAppEntitlementIds
`func (o *C1ApiAppV1AppProfile) UnsetAppEntitlementIds()`

UnsetAppEntitlementIds ensures that no value is present for AppEntitlementIds, not even an explicit nil
### GetAppId

`func (o *C1ApiAppV1AppProfile) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *C1ApiAppV1AppProfile) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *C1ApiAppV1AppProfile) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *C1ApiAppV1AppProfile) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *C1ApiAppV1AppProfile) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *C1ApiAppV1AppProfile) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *C1ApiAppV1AppProfile) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *C1ApiAppV1AppProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *C1ApiAppV1AppProfile) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *C1ApiAppV1AppProfile) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *C1ApiAppV1AppProfile) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *C1ApiAppV1AppProfile) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *C1ApiAppV1AppProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *C1ApiAppV1AppProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *C1ApiAppV1AppProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *C1ApiAppV1AppProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *C1ApiAppV1AppProfile) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *C1ApiAppV1AppProfile) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *C1ApiAppV1AppProfile) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *C1ApiAppV1AppProfile) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *C1ApiAppV1AppProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *C1ApiAppV1AppProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *C1ApiAppV1AppProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *C1ApiAppV1AppProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotificationToken

`func (o *C1ApiAppV1AppProfile) GetNotificationToken() string`

GetNotificationToken returns the NotificationToken field if non-nil, zero value otherwise.

### GetNotificationTokenOk

`func (o *C1ApiAppV1AppProfile) GetNotificationTokenOk() (*string, bool)`

GetNotificationTokenOk returns a tuple with the NotificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationToken

`func (o *C1ApiAppV1AppProfile) SetNotificationToken(v string)`

SetNotificationToken sets NotificationToken field to given value.

### HasNotificationToken

`func (o *C1ApiAppV1AppProfile) HasNotificationToken() bool`

HasNotificationToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *C1ApiAppV1AppProfile) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *C1ApiAppV1AppProfile) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *C1ApiAppV1AppProfile) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *C1ApiAppV1AppProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


