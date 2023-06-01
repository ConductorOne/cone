# C1ApiTaskV1TaskTypeRevoke

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppEntitlementIds** | Pointer to **[]string** | The appEntitlementIds field. | [optional] 
**AppId** | Pointer to **string** | The appId field. | [optional] 
**AppUserId** | Pointer to **string** | The appUserId field. | [optional] 
**IdentityUserId** | Pointer to **string** | The identityUserId field. | [optional] 
**Outcome** | Pointer to **string** | The outcome field. | [optional] 
**OutcomeTime** | Pointer to **time.Time** |  | [optional] 
**Source** | Pointer to [**C1ApiTaskV1TaskRevokeSource**](C1ApiTaskV1TaskRevokeSource.md) |  | [optional] 

## Methods

### NewC1ApiTaskV1TaskTypeRevoke

`func NewC1ApiTaskV1TaskTypeRevoke() *C1ApiTaskV1TaskTypeRevoke`

NewC1ApiTaskV1TaskTypeRevoke instantiates a new C1ApiTaskV1TaskTypeRevoke object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiTaskV1TaskTypeRevokeWithDefaults

`func NewC1ApiTaskV1TaskTypeRevokeWithDefaults() *C1ApiTaskV1TaskTypeRevoke`

NewC1ApiTaskV1TaskTypeRevokeWithDefaults instantiates a new C1ApiTaskV1TaskTypeRevoke object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppEntitlementIds

`func (o *C1ApiTaskV1TaskTypeRevoke) GetAppEntitlementIds() []string`

GetAppEntitlementIds returns the AppEntitlementIds field if non-nil, zero value otherwise.

### GetAppEntitlementIdsOk

`func (o *C1ApiTaskV1TaskTypeRevoke) GetAppEntitlementIdsOk() (*[]string, bool)`

GetAppEntitlementIdsOk returns a tuple with the AppEntitlementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEntitlementIds

`func (o *C1ApiTaskV1TaskTypeRevoke) SetAppEntitlementIds(v []string)`

SetAppEntitlementIds sets AppEntitlementIds field to given value.

### HasAppEntitlementIds

`func (o *C1ApiTaskV1TaskTypeRevoke) HasAppEntitlementIds() bool`

HasAppEntitlementIds returns a boolean if a field has been set.

### SetAppEntitlementIdsNil

`func (o *C1ApiTaskV1TaskTypeRevoke) SetAppEntitlementIdsNil(b bool)`

 SetAppEntitlementIdsNil sets the value for AppEntitlementIds to be an explicit nil

### UnsetAppEntitlementIds
`func (o *C1ApiTaskV1TaskTypeRevoke) UnsetAppEntitlementIds()`

UnsetAppEntitlementIds ensures that no value is present for AppEntitlementIds, not even an explicit nil
### GetAppId

`func (o *C1ApiTaskV1TaskTypeRevoke) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *C1ApiTaskV1TaskTypeRevoke) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *C1ApiTaskV1TaskTypeRevoke) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *C1ApiTaskV1TaskTypeRevoke) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppUserId

`func (o *C1ApiTaskV1TaskTypeRevoke) GetAppUserId() string`

GetAppUserId returns the AppUserId field if non-nil, zero value otherwise.

### GetAppUserIdOk

`func (o *C1ApiTaskV1TaskTypeRevoke) GetAppUserIdOk() (*string, bool)`

GetAppUserIdOk returns a tuple with the AppUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUserId

`func (o *C1ApiTaskV1TaskTypeRevoke) SetAppUserId(v string)`

SetAppUserId sets AppUserId field to given value.

### HasAppUserId

`func (o *C1ApiTaskV1TaskTypeRevoke) HasAppUserId() bool`

HasAppUserId returns a boolean if a field has been set.

### GetIdentityUserId

`func (o *C1ApiTaskV1TaskTypeRevoke) GetIdentityUserId() string`

GetIdentityUserId returns the IdentityUserId field if non-nil, zero value otherwise.

### GetIdentityUserIdOk

`func (o *C1ApiTaskV1TaskTypeRevoke) GetIdentityUserIdOk() (*string, bool)`

GetIdentityUserIdOk returns a tuple with the IdentityUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityUserId

`func (o *C1ApiTaskV1TaskTypeRevoke) SetIdentityUserId(v string)`

SetIdentityUserId sets IdentityUserId field to given value.

### HasIdentityUserId

`func (o *C1ApiTaskV1TaskTypeRevoke) HasIdentityUserId() bool`

HasIdentityUserId returns a boolean if a field has been set.

### GetOutcome

`func (o *C1ApiTaskV1TaskTypeRevoke) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *C1ApiTaskV1TaskTypeRevoke) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *C1ApiTaskV1TaskTypeRevoke) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *C1ApiTaskV1TaskTypeRevoke) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetOutcomeTime

`func (o *C1ApiTaskV1TaskTypeRevoke) GetOutcomeTime() time.Time`

GetOutcomeTime returns the OutcomeTime field if non-nil, zero value otherwise.

### GetOutcomeTimeOk

`func (o *C1ApiTaskV1TaskTypeRevoke) GetOutcomeTimeOk() (*time.Time, bool)`

GetOutcomeTimeOk returns a tuple with the OutcomeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeTime

`func (o *C1ApiTaskV1TaskTypeRevoke) SetOutcomeTime(v time.Time)`

SetOutcomeTime sets OutcomeTime field to given value.

### HasOutcomeTime

`func (o *C1ApiTaskV1TaskTypeRevoke) HasOutcomeTime() bool`

HasOutcomeTime returns a boolean if a field has been set.

### GetSource

`func (o *C1ApiTaskV1TaskTypeRevoke) GetSource() C1ApiTaskV1TaskRevokeSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *C1ApiTaskV1TaskTypeRevoke) GetSourceOk() (*C1ApiTaskV1TaskRevokeSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *C1ApiTaskV1TaskTypeRevoke) SetSource(v C1ApiTaskV1TaskRevokeSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *C1ApiTaskV1TaskTypeRevoke) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


