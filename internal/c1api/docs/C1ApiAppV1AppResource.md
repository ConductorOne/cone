# C1ApiAppV1AppResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | The appId field. | [optional] 
**AppResourceTypeId** | Pointer to **string** | The appResourceTypeId field. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CustomDescription** | Pointer to **string** | The customDescription field. | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** | The description field. | [optional] 
**DisplayName** | Pointer to **string** | The displayName field. | [optional] 
**GrantCount** | Pointer to **string** | The grantCount field. | [optional] 
**Id** | Pointer to **string** | The id field. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewC1ApiAppV1AppResource

`func NewC1ApiAppV1AppResource() *C1ApiAppV1AppResource`

NewC1ApiAppV1AppResource instantiates a new C1ApiAppV1AppResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiAppV1AppResourceWithDefaults

`func NewC1ApiAppV1AppResourceWithDefaults() *C1ApiAppV1AppResource`

NewC1ApiAppV1AppResourceWithDefaults instantiates a new C1ApiAppV1AppResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *C1ApiAppV1AppResource) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *C1ApiAppV1AppResource) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *C1ApiAppV1AppResource) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *C1ApiAppV1AppResource) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppResourceTypeId

`func (o *C1ApiAppV1AppResource) GetAppResourceTypeId() string`

GetAppResourceTypeId returns the AppResourceTypeId field if non-nil, zero value otherwise.

### GetAppResourceTypeIdOk

`func (o *C1ApiAppV1AppResource) GetAppResourceTypeIdOk() (*string, bool)`

GetAppResourceTypeIdOk returns a tuple with the AppResourceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppResourceTypeId

`func (o *C1ApiAppV1AppResource) SetAppResourceTypeId(v string)`

SetAppResourceTypeId sets AppResourceTypeId field to given value.

### HasAppResourceTypeId

`func (o *C1ApiAppV1AppResource) HasAppResourceTypeId() bool`

HasAppResourceTypeId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *C1ApiAppV1AppResource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *C1ApiAppV1AppResource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *C1ApiAppV1AppResource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *C1ApiAppV1AppResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomDescription

`func (o *C1ApiAppV1AppResource) GetCustomDescription() string`

GetCustomDescription returns the CustomDescription field if non-nil, zero value otherwise.

### GetCustomDescriptionOk

`func (o *C1ApiAppV1AppResource) GetCustomDescriptionOk() (*string, bool)`

GetCustomDescriptionOk returns a tuple with the CustomDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDescription

`func (o *C1ApiAppV1AppResource) SetCustomDescription(v string)`

SetCustomDescription sets CustomDescription field to given value.

### HasCustomDescription

`func (o *C1ApiAppV1AppResource) HasCustomDescription() bool`

HasCustomDescription returns a boolean if a field has been set.

### GetDeletedAt

`func (o *C1ApiAppV1AppResource) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *C1ApiAppV1AppResource) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *C1ApiAppV1AppResource) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *C1ApiAppV1AppResource) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *C1ApiAppV1AppResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *C1ApiAppV1AppResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *C1ApiAppV1AppResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *C1ApiAppV1AppResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *C1ApiAppV1AppResource) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *C1ApiAppV1AppResource) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *C1ApiAppV1AppResource) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *C1ApiAppV1AppResource) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGrantCount

`func (o *C1ApiAppV1AppResource) GetGrantCount() string`

GetGrantCount returns the GrantCount field if non-nil, zero value otherwise.

### GetGrantCountOk

`func (o *C1ApiAppV1AppResource) GetGrantCountOk() (*string, bool)`

GetGrantCountOk returns a tuple with the GrantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantCount

`func (o *C1ApiAppV1AppResource) SetGrantCount(v string)`

SetGrantCount sets GrantCount field to given value.

### HasGrantCount

`func (o *C1ApiAppV1AppResource) HasGrantCount() bool`

HasGrantCount returns a boolean if a field has been set.

### GetId

`func (o *C1ApiAppV1AppResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *C1ApiAppV1AppResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *C1ApiAppV1AppResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *C1ApiAppV1AppResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *C1ApiAppV1AppResource) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *C1ApiAppV1AppResource) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *C1ApiAppV1AppResource) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *C1ApiAppV1AppResource) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


