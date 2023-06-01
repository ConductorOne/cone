# C1ApiTaskV1TaskTypeGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppEntitlementIds** | Pointer to **[]string** | The appEntitlementIds field. | [optional] 
**AppId** | Pointer to **string** | The appId field. | [optional] 
**AppUserId** | Pointer to **string** | The appUserId field. | [optional] 
**EntitlementInstances** | Pointer to [**[]C1ApiTaskV1GrantEntitlementInstance**](C1ApiTaskV1GrantEntitlementInstance.md) | The entitlementInstances field. | [optional] 
**IdentityUserId** | Pointer to **string** | The identityUserId field. | [optional] 
**Outcome** | Pointer to **string** | The outcome field. | [optional] 
**OutcomeTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewC1ApiTaskV1TaskTypeGrant

`func NewC1ApiTaskV1TaskTypeGrant() *C1ApiTaskV1TaskTypeGrant`

NewC1ApiTaskV1TaskTypeGrant instantiates a new C1ApiTaskV1TaskTypeGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiTaskV1TaskTypeGrantWithDefaults

`func NewC1ApiTaskV1TaskTypeGrantWithDefaults() *C1ApiTaskV1TaskTypeGrant`

NewC1ApiTaskV1TaskTypeGrantWithDefaults instantiates a new C1ApiTaskV1TaskTypeGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppEntitlementIds

`func (o *C1ApiTaskV1TaskTypeGrant) GetAppEntitlementIds() []string`

GetAppEntitlementIds returns the AppEntitlementIds field if non-nil, zero value otherwise.

### GetAppEntitlementIdsOk

`func (o *C1ApiTaskV1TaskTypeGrant) GetAppEntitlementIdsOk() (*[]string, bool)`

GetAppEntitlementIdsOk returns a tuple with the AppEntitlementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEntitlementIds

`func (o *C1ApiTaskV1TaskTypeGrant) SetAppEntitlementIds(v []string)`

SetAppEntitlementIds sets AppEntitlementIds field to given value.

### HasAppEntitlementIds

`func (o *C1ApiTaskV1TaskTypeGrant) HasAppEntitlementIds() bool`

HasAppEntitlementIds returns a boolean if a field has been set.

### SetAppEntitlementIdsNil

`func (o *C1ApiTaskV1TaskTypeGrant) SetAppEntitlementIdsNil(b bool)`

 SetAppEntitlementIdsNil sets the value for AppEntitlementIds to be an explicit nil

### UnsetAppEntitlementIds
`func (o *C1ApiTaskV1TaskTypeGrant) UnsetAppEntitlementIds()`

UnsetAppEntitlementIds ensures that no value is present for AppEntitlementIds, not even an explicit nil
### GetAppId

`func (o *C1ApiTaskV1TaskTypeGrant) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *C1ApiTaskV1TaskTypeGrant) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *C1ApiTaskV1TaskTypeGrant) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *C1ApiTaskV1TaskTypeGrant) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAppUserId

`func (o *C1ApiTaskV1TaskTypeGrant) GetAppUserId() string`

GetAppUserId returns the AppUserId field if non-nil, zero value otherwise.

### GetAppUserIdOk

`func (o *C1ApiTaskV1TaskTypeGrant) GetAppUserIdOk() (*string, bool)`

GetAppUserIdOk returns a tuple with the AppUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUserId

`func (o *C1ApiTaskV1TaskTypeGrant) SetAppUserId(v string)`

SetAppUserId sets AppUserId field to given value.

### HasAppUserId

`func (o *C1ApiTaskV1TaskTypeGrant) HasAppUserId() bool`

HasAppUserId returns a boolean if a field has been set.

### GetEntitlementInstances

`func (o *C1ApiTaskV1TaskTypeGrant) GetEntitlementInstances() []C1ApiTaskV1GrantEntitlementInstance`

GetEntitlementInstances returns the EntitlementInstances field if non-nil, zero value otherwise.

### GetEntitlementInstancesOk

`func (o *C1ApiTaskV1TaskTypeGrant) GetEntitlementInstancesOk() (*[]C1ApiTaskV1GrantEntitlementInstance, bool)`

GetEntitlementInstancesOk returns a tuple with the EntitlementInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementInstances

`func (o *C1ApiTaskV1TaskTypeGrant) SetEntitlementInstances(v []C1ApiTaskV1GrantEntitlementInstance)`

SetEntitlementInstances sets EntitlementInstances field to given value.

### HasEntitlementInstances

`func (o *C1ApiTaskV1TaskTypeGrant) HasEntitlementInstances() bool`

HasEntitlementInstances returns a boolean if a field has been set.

### SetEntitlementInstancesNil

`func (o *C1ApiTaskV1TaskTypeGrant) SetEntitlementInstancesNil(b bool)`

 SetEntitlementInstancesNil sets the value for EntitlementInstances to be an explicit nil

### UnsetEntitlementInstances
`func (o *C1ApiTaskV1TaskTypeGrant) UnsetEntitlementInstances()`

UnsetEntitlementInstances ensures that no value is present for EntitlementInstances, not even an explicit nil
### GetIdentityUserId

`func (o *C1ApiTaskV1TaskTypeGrant) GetIdentityUserId() string`

GetIdentityUserId returns the IdentityUserId field if non-nil, zero value otherwise.

### GetIdentityUserIdOk

`func (o *C1ApiTaskV1TaskTypeGrant) GetIdentityUserIdOk() (*string, bool)`

GetIdentityUserIdOk returns a tuple with the IdentityUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityUserId

`func (o *C1ApiTaskV1TaskTypeGrant) SetIdentityUserId(v string)`

SetIdentityUserId sets IdentityUserId field to given value.

### HasIdentityUserId

`func (o *C1ApiTaskV1TaskTypeGrant) HasIdentityUserId() bool`

HasIdentityUserId returns a boolean if a field has been set.

### GetOutcome

`func (o *C1ApiTaskV1TaskTypeGrant) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *C1ApiTaskV1TaskTypeGrant) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *C1ApiTaskV1TaskTypeGrant) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *C1ApiTaskV1TaskTypeGrant) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetOutcomeTime

`func (o *C1ApiTaskV1TaskTypeGrant) GetOutcomeTime() time.Time`

GetOutcomeTime returns the OutcomeTime field if non-nil, zero value otherwise.

### GetOutcomeTimeOk

`func (o *C1ApiTaskV1TaskTypeGrant) GetOutcomeTimeOk() (*time.Time, bool)`

GetOutcomeTimeOk returns a tuple with the OutcomeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeTime

`func (o *C1ApiTaskV1TaskTypeGrant) SetOutcomeTime(v time.Time)`

SetOutcomeTime sets OutcomeTime field to given value.

### HasOutcomeTime

`func (o *C1ApiTaskV1TaskTypeGrant) HasOutcomeTime() bool`

HasOutcomeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


