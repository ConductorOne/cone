# C1ApiIamV1Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**DisplayName** | Pointer to **string** | The displayName field. | [optional] 
**Id** | Pointer to **string** | The id field. | [optional] 
**Name** | Pointer to **string** | The name field. | [optional] 
**Permissions** | Pointer to **[]string** | The permissions field. | [optional] 
**ServiceRoles** | Pointer to **[]string** | The serviceRoles field. | [optional] 
**SystemBuiltin** | Pointer to **bool** | The systemBuiltin field. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewC1ApiIamV1Role

`func NewC1ApiIamV1Role() *C1ApiIamV1Role`

NewC1ApiIamV1Role instantiates a new C1ApiIamV1Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiIamV1RoleWithDefaults

`func NewC1ApiIamV1RoleWithDefaults() *C1ApiIamV1Role`

NewC1ApiIamV1RoleWithDefaults instantiates a new C1ApiIamV1Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *C1ApiIamV1Role) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *C1ApiIamV1Role) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *C1ApiIamV1Role) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *C1ApiIamV1Role) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *C1ApiIamV1Role) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *C1ApiIamV1Role) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *C1ApiIamV1Role) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *C1ApiIamV1Role) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *C1ApiIamV1Role) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *C1ApiIamV1Role) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *C1ApiIamV1Role) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *C1ApiIamV1Role) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *C1ApiIamV1Role) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *C1ApiIamV1Role) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *C1ApiIamV1Role) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *C1ApiIamV1Role) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *C1ApiIamV1Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *C1ApiIamV1Role) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *C1ApiIamV1Role) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *C1ApiIamV1Role) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *C1ApiIamV1Role) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *C1ApiIamV1Role) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *C1ApiIamV1Role) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *C1ApiIamV1Role) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *C1ApiIamV1Role) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *C1ApiIamV1Role) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetServiceRoles

`func (o *C1ApiIamV1Role) GetServiceRoles() []string`

GetServiceRoles returns the ServiceRoles field if non-nil, zero value otherwise.

### GetServiceRolesOk

`func (o *C1ApiIamV1Role) GetServiceRolesOk() (*[]string, bool)`

GetServiceRolesOk returns a tuple with the ServiceRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRoles

`func (o *C1ApiIamV1Role) SetServiceRoles(v []string)`

SetServiceRoles sets ServiceRoles field to given value.

### HasServiceRoles

`func (o *C1ApiIamV1Role) HasServiceRoles() bool`

HasServiceRoles returns a boolean if a field has been set.

### SetServiceRolesNil

`func (o *C1ApiIamV1Role) SetServiceRolesNil(b bool)`

 SetServiceRolesNil sets the value for ServiceRoles to be an explicit nil

### UnsetServiceRoles
`func (o *C1ApiIamV1Role) UnsetServiceRoles()`

UnsetServiceRoles ensures that no value is present for ServiceRoles, not even an explicit nil
### GetSystemBuiltin

`func (o *C1ApiIamV1Role) GetSystemBuiltin() bool`

GetSystemBuiltin returns the SystemBuiltin field if non-nil, zero value otherwise.

### GetSystemBuiltinOk

`func (o *C1ApiIamV1Role) GetSystemBuiltinOk() (*bool, bool)`

GetSystemBuiltinOk returns a tuple with the SystemBuiltin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemBuiltin

`func (o *C1ApiIamV1Role) SetSystemBuiltin(v bool)`

SetSystemBuiltin sets SystemBuiltin field to given value.

### HasSystemBuiltin

`func (o *C1ApiIamV1Role) HasSystemBuiltin() bool`

HasSystemBuiltin returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *C1ApiIamV1Role) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *C1ApiIamV1Role) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *C1ApiIamV1Role) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *C1ApiIamV1Role) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


