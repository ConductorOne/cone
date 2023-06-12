# C1ApiUserV1SearchUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeIds** | Pointer to **[]string** | The excludeIds field. | [optional] 
**ExpandMask** | Pointer to [**C1ApiUserV1UserExpandMask**](C1ApiUserV1UserExpandMask.md) |  | [optional] 
**Ids** | Pointer to **[]string** | The ids field. | [optional] 
**PageSize** | Pointer to **float32** | The pageSize field. | [optional] 
**PageToken** | Pointer to **string** | The pageToken field. | [optional] 
**Query** | Pointer to **string** | The query field. | [optional] 
**Refs** | Pointer to [**[]C1ApiUserV1UserRef**](C1ApiUserV1UserRef.md) | The refs field. | [optional] 
**RoleIds** | Pointer to **[]string** | The roleIds field. | [optional] 
**UserStatuses** | Pointer to **[]string** | The userStatuses field. | [optional] 

## Methods

### NewC1ApiUserV1SearchUsersRequest

`func NewC1ApiUserV1SearchUsersRequest() *C1ApiUserV1SearchUsersRequest`

NewC1ApiUserV1SearchUsersRequest instantiates a new C1ApiUserV1SearchUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiUserV1SearchUsersRequestWithDefaults

`func NewC1ApiUserV1SearchUsersRequestWithDefaults() *C1ApiUserV1SearchUsersRequest`

NewC1ApiUserV1SearchUsersRequestWithDefaults instantiates a new C1ApiUserV1SearchUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeIds

`func (o *C1ApiUserV1SearchUsersRequest) GetExcludeIds() []string`

GetExcludeIds returns the ExcludeIds field if non-nil, zero value otherwise.

### GetExcludeIdsOk

`func (o *C1ApiUserV1SearchUsersRequest) GetExcludeIdsOk() (*[]string, bool)`

GetExcludeIdsOk returns a tuple with the ExcludeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeIds

`func (o *C1ApiUserV1SearchUsersRequest) SetExcludeIds(v []string)`

SetExcludeIds sets ExcludeIds field to given value.

### HasExcludeIds

`func (o *C1ApiUserV1SearchUsersRequest) HasExcludeIds() bool`

HasExcludeIds returns a boolean if a field has been set.

### SetExcludeIdsNil

`func (o *C1ApiUserV1SearchUsersRequest) SetExcludeIdsNil(b bool)`

 SetExcludeIdsNil sets the value for ExcludeIds to be an explicit nil

### UnsetExcludeIds
`func (o *C1ApiUserV1SearchUsersRequest) UnsetExcludeIds()`

UnsetExcludeIds ensures that no value is present for ExcludeIds, not even an explicit nil
### GetExpandMask

`func (o *C1ApiUserV1SearchUsersRequest) GetExpandMask() C1ApiUserV1UserExpandMask`

GetExpandMask returns the ExpandMask field if non-nil, zero value otherwise.

### GetExpandMaskOk

`func (o *C1ApiUserV1SearchUsersRequest) GetExpandMaskOk() (*C1ApiUserV1UserExpandMask, bool)`

GetExpandMaskOk returns a tuple with the ExpandMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandMask

`func (o *C1ApiUserV1SearchUsersRequest) SetExpandMask(v C1ApiUserV1UserExpandMask)`

SetExpandMask sets ExpandMask field to given value.

### HasExpandMask

`func (o *C1ApiUserV1SearchUsersRequest) HasExpandMask() bool`

HasExpandMask returns a boolean if a field has been set.

### GetIds

`func (o *C1ApiUserV1SearchUsersRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *C1ApiUserV1SearchUsersRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *C1ApiUserV1SearchUsersRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *C1ApiUserV1SearchUsersRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *C1ApiUserV1SearchUsersRequest) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *C1ApiUserV1SearchUsersRequest) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetPageSize

`func (o *C1ApiUserV1SearchUsersRequest) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *C1ApiUserV1SearchUsersRequest) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *C1ApiUserV1SearchUsersRequest) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *C1ApiUserV1SearchUsersRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *C1ApiUserV1SearchUsersRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *C1ApiUserV1SearchUsersRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *C1ApiUserV1SearchUsersRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *C1ApiUserV1SearchUsersRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### GetQuery

`func (o *C1ApiUserV1SearchUsersRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *C1ApiUserV1SearchUsersRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *C1ApiUserV1SearchUsersRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *C1ApiUserV1SearchUsersRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRefs

`func (o *C1ApiUserV1SearchUsersRequest) GetRefs() []C1ApiUserV1UserRef`

GetRefs returns the Refs field if non-nil, zero value otherwise.

### GetRefsOk

`func (o *C1ApiUserV1SearchUsersRequest) GetRefsOk() (*[]C1ApiUserV1UserRef, bool)`

GetRefsOk returns a tuple with the Refs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefs

`func (o *C1ApiUserV1SearchUsersRequest) SetRefs(v []C1ApiUserV1UserRef)`

SetRefs sets Refs field to given value.

### HasRefs

`func (o *C1ApiUserV1SearchUsersRequest) HasRefs() bool`

HasRefs returns a boolean if a field has been set.

### SetRefsNil

`func (o *C1ApiUserV1SearchUsersRequest) SetRefsNil(b bool)`

 SetRefsNil sets the value for Refs to be an explicit nil

### UnsetRefs
`func (o *C1ApiUserV1SearchUsersRequest) UnsetRefs()`

UnsetRefs ensures that no value is present for Refs, not even an explicit nil
### GetRoleIds

`func (o *C1ApiUserV1SearchUsersRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *C1ApiUserV1SearchUsersRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *C1ApiUserV1SearchUsersRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *C1ApiUserV1SearchUsersRequest) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### SetRoleIdsNil

`func (o *C1ApiUserV1SearchUsersRequest) SetRoleIdsNil(b bool)`

 SetRoleIdsNil sets the value for RoleIds to be an explicit nil

### UnsetRoleIds
`func (o *C1ApiUserV1SearchUsersRequest) UnsetRoleIds()`

UnsetRoleIds ensures that no value is present for RoleIds, not even an explicit nil
### GetUserStatuses

`func (o *C1ApiUserV1SearchUsersRequest) GetUserStatuses() []string`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *C1ApiUserV1SearchUsersRequest) GetUserStatusesOk() (*[]string, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatuses

`func (o *C1ApiUserV1SearchUsersRequest) SetUserStatuses(v []string)`

SetUserStatuses sets UserStatuses field to given value.

### HasUserStatuses

`func (o *C1ApiUserV1SearchUsersRequest) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatusesNil

`func (o *C1ApiUserV1SearchUsersRequest) SetUserStatusesNil(b bool)`

 SetUserStatusesNil sets the value for UserStatuses to be an explicit nil

### UnsetUserStatuses
`func (o *C1ApiUserV1SearchUsersRequest) UnsetUserStatuses()`

UnsetUserStatuses ensures that no value is present for UserStatuses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


