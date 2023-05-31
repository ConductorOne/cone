# C1ApiUserV1User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DelegatedUserId** | Pointer to **string** | The delegatedUserId field. | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Department** | Pointer to **string** | The department field. | [optional] 
**DepartmentSources** | Pointer to [**[]C1ApiUserV1UserAttributeMappingSource**](C1ApiUserV1UserAttributeMappingSource.md) | The departmentSources field. | [optional] 
**DirectoryIds** | Pointer to **[]string** | The directoryIds field. | [optional] 
**DirectoryStatus** | Pointer to **string** | The directoryStatus field. | [optional] 
**DirectoryStatusSources** | Pointer to [**[]C1ApiUserV1UserAttributeMappingSource**](C1ApiUserV1UserAttributeMappingSource.md) | The directoryStatusSources field. | [optional] 
**DisplayName** | Pointer to **string** | The displayName field. | [optional] 
**Email** | Pointer to **string** | The email field. | [optional] 
**EmploymentStatus** | Pointer to **string** | The employmentStatus field. | [optional] 
**EmploymentStatusSources** | Pointer to [**[]C1ApiUserV1UserAttributeMappingSource**](C1ApiUserV1UserAttributeMappingSource.md) | The employmentStatusSources field. | [optional] 
**EmploymentType** | Pointer to **string** | The employmentType field. | [optional] 
**EmploymentTypeSources** | Pointer to [**[]C1ApiUserV1UserAttributeMappingSource**](C1ApiUserV1UserAttributeMappingSource.md) | The employmentTypeSources field. | [optional] 
**Id** | Pointer to **string** | The id field. | [optional] 
**JobTitle** | Pointer to **string** | The jobTitle field. | [optional] 
**JobTitleSources** | Pointer to [**[]C1ApiUserV1UserAttributeMappingSource**](C1ApiUserV1UserAttributeMappingSource.md) | The jobTitleSources field. | [optional] 
**ManagerIds** | Pointer to **[]string** | The managerIds field. | [optional] 
**ManagerSources** | Pointer to [**[]C1ApiUserV1UserAttributeMappingSource**](C1ApiUserV1UserAttributeMappingSource.md) | The managerSources field. | [optional] 
**RoleIds** | Pointer to **[]string** | The roleIds field. | [optional] 
**Status** | Pointer to **string** | The status field. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewC1ApiUserV1User

`func NewC1ApiUserV1User() *C1ApiUserV1User`

NewC1ApiUserV1User instantiates a new C1ApiUserV1User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiUserV1UserWithDefaults

`func NewC1ApiUserV1UserWithDefaults() *C1ApiUserV1User`

NewC1ApiUserV1UserWithDefaults instantiates a new C1ApiUserV1User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *C1ApiUserV1User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *C1ApiUserV1User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *C1ApiUserV1User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *C1ApiUserV1User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDelegatedUserId

`func (o *C1ApiUserV1User) GetDelegatedUserId() string`

GetDelegatedUserId returns the DelegatedUserId field if non-nil, zero value otherwise.

### GetDelegatedUserIdOk

`func (o *C1ApiUserV1User) GetDelegatedUserIdOk() (*string, bool)`

GetDelegatedUserIdOk returns a tuple with the DelegatedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedUserId

`func (o *C1ApiUserV1User) SetDelegatedUserId(v string)`

SetDelegatedUserId sets DelegatedUserId field to given value.

### HasDelegatedUserId

`func (o *C1ApiUserV1User) HasDelegatedUserId() bool`

HasDelegatedUserId returns a boolean if a field has been set.

### GetDeletedAt

`func (o *C1ApiUserV1User) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *C1ApiUserV1User) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *C1ApiUserV1User) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *C1ApiUserV1User) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDepartment

`func (o *C1ApiUserV1User) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *C1ApiUserV1User) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *C1ApiUserV1User) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *C1ApiUserV1User) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDepartmentSources

`func (o *C1ApiUserV1User) GetDepartmentSources() []C1ApiUserV1UserAttributeMappingSource`

GetDepartmentSources returns the DepartmentSources field if non-nil, zero value otherwise.

### GetDepartmentSourcesOk

`func (o *C1ApiUserV1User) GetDepartmentSourcesOk() (*[]C1ApiUserV1UserAttributeMappingSource, bool)`

GetDepartmentSourcesOk returns a tuple with the DepartmentSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentSources

`func (o *C1ApiUserV1User) SetDepartmentSources(v []C1ApiUserV1UserAttributeMappingSource)`

SetDepartmentSources sets DepartmentSources field to given value.

### HasDepartmentSources

`func (o *C1ApiUserV1User) HasDepartmentSources() bool`

HasDepartmentSources returns a boolean if a field has been set.

### SetDepartmentSourcesNil

`func (o *C1ApiUserV1User) SetDepartmentSourcesNil(b bool)`

 SetDepartmentSourcesNil sets the value for DepartmentSources to be an explicit nil

### UnsetDepartmentSources
`func (o *C1ApiUserV1User) UnsetDepartmentSources()`

UnsetDepartmentSources ensures that no value is present for DepartmentSources, not even an explicit nil
### GetDirectoryIds

`func (o *C1ApiUserV1User) GetDirectoryIds() []string`

GetDirectoryIds returns the DirectoryIds field if non-nil, zero value otherwise.

### GetDirectoryIdsOk

`func (o *C1ApiUserV1User) GetDirectoryIdsOk() (*[]string, bool)`

GetDirectoryIdsOk returns a tuple with the DirectoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryIds

`func (o *C1ApiUserV1User) SetDirectoryIds(v []string)`

SetDirectoryIds sets DirectoryIds field to given value.

### HasDirectoryIds

`func (o *C1ApiUserV1User) HasDirectoryIds() bool`

HasDirectoryIds returns a boolean if a field has been set.

### SetDirectoryIdsNil

`func (o *C1ApiUserV1User) SetDirectoryIdsNil(b bool)`

 SetDirectoryIdsNil sets the value for DirectoryIds to be an explicit nil

### UnsetDirectoryIds
`func (o *C1ApiUserV1User) UnsetDirectoryIds()`

UnsetDirectoryIds ensures that no value is present for DirectoryIds, not even an explicit nil
### GetDirectoryStatus

`func (o *C1ApiUserV1User) GetDirectoryStatus() string`

GetDirectoryStatus returns the DirectoryStatus field if non-nil, zero value otherwise.

### GetDirectoryStatusOk

`func (o *C1ApiUserV1User) GetDirectoryStatusOk() (*string, bool)`

GetDirectoryStatusOk returns a tuple with the DirectoryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryStatus

`func (o *C1ApiUserV1User) SetDirectoryStatus(v string)`

SetDirectoryStatus sets DirectoryStatus field to given value.

### HasDirectoryStatus

`func (o *C1ApiUserV1User) HasDirectoryStatus() bool`

HasDirectoryStatus returns a boolean if a field has been set.

### GetDirectoryStatusSources

`func (o *C1ApiUserV1User) GetDirectoryStatusSources() []C1ApiUserV1UserAttributeMappingSource`

GetDirectoryStatusSources returns the DirectoryStatusSources field if non-nil, zero value otherwise.

### GetDirectoryStatusSourcesOk

`func (o *C1ApiUserV1User) GetDirectoryStatusSourcesOk() (*[]C1ApiUserV1UserAttributeMappingSource, bool)`

GetDirectoryStatusSourcesOk returns a tuple with the DirectoryStatusSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryStatusSources

`func (o *C1ApiUserV1User) SetDirectoryStatusSources(v []C1ApiUserV1UserAttributeMappingSource)`

SetDirectoryStatusSources sets DirectoryStatusSources field to given value.

### HasDirectoryStatusSources

`func (o *C1ApiUserV1User) HasDirectoryStatusSources() bool`

HasDirectoryStatusSources returns a boolean if a field has been set.

### SetDirectoryStatusSourcesNil

`func (o *C1ApiUserV1User) SetDirectoryStatusSourcesNil(b bool)`

 SetDirectoryStatusSourcesNil sets the value for DirectoryStatusSources to be an explicit nil

### UnsetDirectoryStatusSources
`func (o *C1ApiUserV1User) UnsetDirectoryStatusSources()`

UnsetDirectoryStatusSources ensures that no value is present for DirectoryStatusSources, not even an explicit nil
### GetDisplayName

`func (o *C1ApiUserV1User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *C1ApiUserV1User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *C1ApiUserV1User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *C1ApiUserV1User) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *C1ApiUserV1User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *C1ApiUserV1User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *C1ApiUserV1User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *C1ApiUserV1User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmploymentStatus

`func (o *C1ApiUserV1User) GetEmploymentStatus() string`

GetEmploymentStatus returns the EmploymentStatus field if non-nil, zero value otherwise.

### GetEmploymentStatusOk

`func (o *C1ApiUserV1User) GetEmploymentStatusOk() (*string, bool)`

GetEmploymentStatusOk returns a tuple with the EmploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentStatus

`func (o *C1ApiUserV1User) SetEmploymentStatus(v string)`

SetEmploymentStatus sets EmploymentStatus field to given value.

### HasEmploymentStatus

`func (o *C1ApiUserV1User) HasEmploymentStatus() bool`

HasEmploymentStatus returns a boolean if a field has been set.

### GetEmploymentStatusSources

`func (o *C1ApiUserV1User) GetEmploymentStatusSources() []C1ApiUserV1UserAttributeMappingSource`

GetEmploymentStatusSources returns the EmploymentStatusSources field if non-nil, zero value otherwise.

### GetEmploymentStatusSourcesOk

`func (o *C1ApiUserV1User) GetEmploymentStatusSourcesOk() (*[]C1ApiUserV1UserAttributeMappingSource, bool)`

GetEmploymentStatusSourcesOk returns a tuple with the EmploymentStatusSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentStatusSources

`func (o *C1ApiUserV1User) SetEmploymentStatusSources(v []C1ApiUserV1UserAttributeMappingSource)`

SetEmploymentStatusSources sets EmploymentStatusSources field to given value.

### HasEmploymentStatusSources

`func (o *C1ApiUserV1User) HasEmploymentStatusSources() bool`

HasEmploymentStatusSources returns a boolean if a field has been set.

### SetEmploymentStatusSourcesNil

`func (o *C1ApiUserV1User) SetEmploymentStatusSourcesNil(b bool)`

 SetEmploymentStatusSourcesNil sets the value for EmploymentStatusSources to be an explicit nil

### UnsetEmploymentStatusSources
`func (o *C1ApiUserV1User) UnsetEmploymentStatusSources()`

UnsetEmploymentStatusSources ensures that no value is present for EmploymentStatusSources, not even an explicit nil
### GetEmploymentType

`func (o *C1ApiUserV1User) GetEmploymentType() string`

GetEmploymentType returns the EmploymentType field if non-nil, zero value otherwise.

### GetEmploymentTypeOk

`func (o *C1ApiUserV1User) GetEmploymentTypeOk() (*string, bool)`

GetEmploymentTypeOk returns a tuple with the EmploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentType

`func (o *C1ApiUserV1User) SetEmploymentType(v string)`

SetEmploymentType sets EmploymentType field to given value.

### HasEmploymentType

`func (o *C1ApiUserV1User) HasEmploymentType() bool`

HasEmploymentType returns a boolean if a field has been set.

### GetEmploymentTypeSources

`func (o *C1ApiUserV1User) GetEmploymentTypeSources() []C1ApiUserV1UserAttributeMappingSource`

GetEmploymentTypeSources returns the EmploymentTypeSources field if non-nil, zero value otherwise.

### GetEmploymentTypeSourcesOk

`func (o *C1ApiUserV1User) GetEmploymentTypeSourcesOk() (*[]C1ApiUserV1UserAttributeMappingSource, bool)`

GetEmploymentTypeSourcesOk returns a tuple with the EmploymentTypeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentTypeSources

`func (o *C1ApiUserV1User) SetEmploymentTypeSources(v []C1ApiUserV1UserAttributeMappingSource)`

SetEmploymentTypeSources sets EmploymentTypeSources field to given value.

### HasEmploymentTypeSources

`func (o *C1ApiUserV1User) HasEmploymentTypeSources() bool`

HasEmploymentTypeSources returns a boolean if a field has been set.

### SetEmploymentTypeSourcesNil

`func (o *C1ApiUserV1User) SetEmploymentTypeSourcesNil(b bool)`

 SetEmploymentTypeSourcesNil sets the value for EmploymentTypeSources to be an explicit nil

### UnsetEmploymentTypeSources
`func (o *C1ApiUserV1User) UnsetEmploymentTypeSources()`

UnsetEmploymentTypeSources ensures that no value is present for EmploymentTypeSources, not even an explicit nil
### GetId

`func (o *C1ApiUserV1User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *C1ApiUserV1User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *C1ApiUserV1User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *C1ApiUserV1User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobTitle

`func (o *C1ApiUserV1User) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *C1ApiUserV1User) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *C1ApiUserV1User) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *C1ApiUserV1User) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetJobTitleSources

`func (o *C1ApiUserV1User) GetJobTitleSources() []C1ApiUserV1UserAttributeMappingSource`

GetJobTitleSources returns the JobTitleSources field if non-nil, zero value otherwise.

### GetJobTitleSourcesOk

`func (o *C1ApiUserV1User) GetJobTitleSourcesOk() (*[]C1ApiUserV1UserAttributeMappingSource, bool)`

GetJobTitleSourcesOk returns a tuple with the JobTitleSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitleSources

`func (o *C1ApiUserV1User) SetJobTitleSources(v []C1ApiUserV1UserAttributeMappingSource)`

SetJobTitleSources sets JobTitleSources field to given value.

### HasJobTitleSources

`func (o *C1ApiUserV1User) HasJobTitleSources() bool`

HasJobTitleSources returns a boolean if a field has been set.

### SetJobTitleSourcesNil

`func (o *C1ApiUserV1User) SetJobTitleSourcesNil(b bool)`

 SetJobTitleSourcesNil sets the value for JobTitleSources to be an explicit nil

### UnsetJobTitleSources
`func (o *C1ApiUserV1User) UnsetJobTitleSources()`

UnsetJobTitleSources ensures that no value is present for JobTitleSources, not even an explicit nil
### GetManagerIds

`func (o *C1ApiUserV1User) GetManagerIds() []string`

GetManagerIds returns the ManagerIds field if non-nil, zero value otherwise.

### GetManagerIdsOk

`func (o *C1ApiUserV1User) GetManagerIdsOk() (*[]string, bool)`

GetManagerIdsOk returns a tuple with the ManagerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerIds

`func (o *C1ApiUserV1User) SetManagerIds(v []string)`

SetManagerIds sets ManagerIds field to given value.

### HasManagerIds

`func (o *C1ApiUserV1User) HasManagerIds() bool`

HasManagerIds returns a boolean if a field has been set.

### SetManagerIdsNil

`func (o *C1ApiUserV1User) SetManagerIdsNil(b bool)`

 SetManagerIdsNil sets the value for ManagerIds to be an explicit nil

### UnsetManagerIds
`func (o *C1ApiUserV1User) UnsetManagerIds()`

UnsetManagerIds ensures that no value is present for ManagerIds, not even an explicit nil
### GetManagerSources

`func (o *C1ApiUserV1User) GetManagerSources() []C1ApiUserV1UserAttributeMappingSource`

GetManagerSources returns the ManagerSources field if non-nil, zero value otherwise.

### GetManagerSourcesOk

`func (o *C1ApiUserV1User) GetManagerSourcesOk() (*[]C1ApiUserV1UserAttributeMappingSource, bool)`

GetManagerSourcesOk returns a tuple with the ManagerSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerSources

`func (o *C1ApiUserV1User) SetManagerSources(v []C1ApiUserV1UserAttributeMappingSource)`

SetManagerSources sets ManagerSources field to given value.

### HasManagerSources

`func (o *C1ApiUserV1User) HasManagerSources() bool`

HasManagerSources returns a boolean if a field has been set.

### SetManagerSourcesNil

`func (o *C1ApiUserV1User) SetManagerSourcesNil(b bool)`

 SetManagerSourcesNil sets the value for ManagerSources to be an explicit nil

### UnsetManagerSources
`func (o *C1ApiUserV1User) UnsetManagerSources()`

UnsetManagerSources ensures that no value is present for ManagerSources, not even an explicit nil
### GetRoleIds

`func (o *C1ApiUserV1User) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *C1ApiUserV1User) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *C1ApiUserV1User) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *C1ApiUserV1User) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### SetRoleIdsNil

`func (o *C1ApiUserV1User) SetRoleIdsNil(b bool)`

 SetRoleIdsNil sets the value for RoleIds to be an explicit nil

### UnsetRoleIds
`func (o *C1ApiUserV1User) UnsetRoleIds()`

UnsetRoleIds ensures that no value is present for RoleIds, not even an explicit nil
### GetStatus

`func (o *C1ApiUserV1User) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *C1ApiUserV1User) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *C1ApiUserV1User) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *C1ApiUserV1User) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *C1ApiUserV1User) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *C1ApiUserV1User) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *C1ApiUserV1User) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *C1ApiUserV1User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


