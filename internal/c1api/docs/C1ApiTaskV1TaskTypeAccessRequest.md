# C1ApiTaskV1TaskTypeAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppProfile** | Pointer to [**C1ApiAppV1AppProfile**](C1ApiAppV1AppProfile.md) |  | [optional] 
**AppUserId** | Pointer to **string** | The appUserId field. | [optional] 
**IdentityUserId** | Pointer to **string** | The identityUserId field. | [optional] 
**Outcome** | Pointer to **string** | The outcome field. | [optional] 
**OutcomeTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewC1ApiTaskV1TaskTypeAccessRequest

`func NewC1ApiTaskV1TaskTypeAccessRequest() *C1ApiTaskV1TaskTypeAccessRequest`

NewC1ApiTaskV1TaskTypeAccessRequest instantiates a new C1ApiTaskV1TaskTypeAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiTaskV1TaskTypeAccessRequestWithDefaults

`func NewC1ApiTaskV1TaskTypeAccessRequestWithDefaults() *C1ApiTaskV1TaskTypeAccessRequest`

NewC1ApiTaskV1TaskTypeAccessRequestWithDefaults instantiates a new C1ApiTaskV1TaskTypeAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppProfile

`func (o *C1ApiTaskV1TaskTypeAccessRequest) GetAppProfile() C1ApiAppV1AppProfile`

GetAppProfile returns the AppProfile field if non-nil, zero value otherwise.

### GetAppProfileOk

`func (o *C1ApiTaskV1TaskTypeAccessRequest) GetAppProfileOk() (*C1ApiAppV1AppProfile, bool)`

GetAppProfileOk returns a tuple with the AppProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProfile

`func (o *C1ApiTaskV1TaskTypeAccessRequest) SetAppProfile(v C1ApiAppV1AppProfile)`

SetAppProfile sets AppProfile field to given value.

### HasAppProfile

`func (o *C1ApiTaskV1TaskTypeAccessRequest) HasAppProfile() bool`

HasAppProfile returns a boolean if a field has been set.

### GetAppUserId

`func (o *C1ApiTaskV1TaskTypeAccessRequest) GetAppUserId() string`

GetAppUserId returns the AppUserId field if non-nil, zero value otherwise.

### GetAppUserIdOk

`func (o *C1ApiTaskV1TaskTypeAccessRequest) GetAppUserIdOk() (*string, bool)`

GetAppUserIdOk returns a tuple with the AppUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUserId

`func (o *C1ApiTaskV1TaskTypeAccessRequest) SetAppUserId(v string)`

SetAppUserId sets AppUserId field to given value.

### HasAppUserId

`func (o *C1ApiTaskV1TaskTypeAccessRequest) HasAppUserId() bool`

HasAppUserId returns a boolean if a field has been set.

### GetIdentityUserId

`func (o *C1ApiTaskV1TaskTypeAccessRequest) GetIdentityUserId() string`

GetIdentityUserId returns the IdentityUserId field if non-nil, zero value otherwise.

### GetIdentityUserIdOk

`func (o *C1ApiTaskV1TaskTypeAccessRequest) GetIdentityUserIdOk() (*string, bool)`

GetIdentityUserIdOk returns a tuple with the IdentityUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityUserId

`func (o *C1ApiTaskV1TaskTypeAccessRequest) SetIdentityUserId(v string)`

SetIdentityUserId sets IdentityUserId field to given value.

### HasIdentityUserId

`func (o *C1ApiTaskV1TaskTypeAccessRequest) HasIdentityUserId() bool`

HasIdentityUserId returns a boolean if a field has been set.

### GetOutcome

`func (o *C1ApiTaskV1TaskTypeAccessRequest) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *C1ApiTaskV1TaskTypeAccessRequest) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *C1ApiTaskV1TaskTypeAccessRequest) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *C1ApiTaskV1TaskTypeAccessRequest) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetOutcomeTime

`func (o *C1ApiTaskV1TaskTypeAccessRequest) GetOutcomeTime() time.Time`

GetOutcomeTime returns the OutcomeTime field if non-nil, zero value otherwise.

### GetOutcomeTimeOk

`func (o *C1ApiTaskV1TaskTypeAccessRequest) GetOutcomeTimeOk() (*time.Time, bool)`

GetOutcomeTimeOk returns a tuple with the OutcomeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeTime

`func (o *C1ApiTaskV1TaskTypeAccessRequest) SetOutcomeTime(v time.Time)`

SetOutcomeTime sets OutcomeTime field to given value.

### HasOutcomeTime

`func (o *C1ApiTaskV1TaskTypeAccessRequest) HasOutcomeTime() bool`

HasOutcomeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


