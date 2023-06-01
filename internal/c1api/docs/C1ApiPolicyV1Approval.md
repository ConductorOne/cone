# C1ApiPolicyV1Approval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowReassignment** | Pointer to **bool** | The allowReassignment field. | [optional] 
**AppOwners** | Pointer to [**NullableC1ApiPolicyV1AppOwnerApproval**](C1ApiPolicyV1AppOwnerApproval.md) |  | [optional] 
**Assigned** | Pointer to **bool** | The assigned field. | [optional] 
**EntitlementOwners** | Pointer to [**NullableC1ApiPolicyV1EntitlementOwnerApproval**](C1ApiPolicyV1EntitlementOwnerApproval.md) |  | [optional] 
**Group** | Pointer to [**NullableC1ApiPolicyV1AppGroupApproval**](C1ApiPolicyV1AppGroupApproval.md) |  | [optional] 
**Manager** | Pointer to [**NullableC1ApiPolicyV1ManagerApproval**](C1ApiPolicyV1ManagerApproval.md) |  | [optional] 
**RequireApprovalReason** | Pointer to **bool** | The requireApprovalReason field. | [optional] 
**RequireReassignmentReason** | Pointer to **bool** | The requireReassignmentReason field. | [optional] 
**Self** | Pointer to [**NullableC1ApiPolicyV1SelfApproval**](C1ApiPolicyV1SelfApproval.md) |  | [optional] 
**Users** | Pointer to [**NullableC1ApiPolicyV1UserApproval**](C1ApiPolicyV1UserApproval.md) |  | [optional] 

## Methods

### NewC1ApiPolicyV1Approval

`func NewC1ApiPolicyV1Approval() *C1ApiPolicyV1Approval`

NewC1ApiPolicyV1Approval instantiates a new C1ApiPolicyV1Approval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1ApprovalWithDefaults

`func NewC1ApiPolicyV1ApprovalWithDefaults() *C1ApiPolicyV1Approval`

NewC1ApiPolicyV1ApprovalWithDefaults instantiates a new C1ApiPolicyV1Approval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowReassignment

`func (o *C1ApiPolicyV1Approval) GetAllowReassignment() bool`

GetAllowReassignment returns the AllowReassignment field if non-nil, zero value otherwise.

### GetAllowReassignmentOk

`func (o *C1ApiPolicyV1Approval) GetAllowReassignmentOk() (*bool, bool)`

GetAllowReassignmentOk returns a tuple with the AllowReassignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReassignment

`func (o *C1ApiPolicyV1Approval) SetAllowReassignment(v bool)`

SetAllowReassignment sets AllowReassignment field to given value.

### HasAllowReassignment

`func (o *C1ApiPolicyV1Approval) HasAllowReassignment() bool`

HasAllowReassignment returns a boolean if a field has been set.

### GetAppOwners

`func (o *C1ApiPolicyV1Approval) GetAppOwners() C1ApiPolicyV1AppOwnerApproval`

GetAppOwners returns the AppOwners field if non-nil, zero value otherwise.

### GetAppOwnersOk

`func (o *C1ApiPolicyV1Approval) GetAppOwnersOk() (*C1ApiPolicyV1AppOwnerApproval, bool)`

GetAppOwnersOk returns a tuple with the AppOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppOwners

`func (o *C1ApiPolicyV1Approval) SetAppOwners(v C1ApiPolicyV1AppOwnerApproval)`

SetAppOwners sets AppOwners field to given value.

### HasAppOwners

`func (o *C1ApiPolicyV1Approval) HasAppOwners() bool`

HasAppOwners returns a boolean if a field has been set.

### SetAppOwnersNil

`func (o *C1ApiPolicyV1Approval) SetAppOwnersNil(b bool)`

 SetAppOwnersNil sets the value for AppOwners to be an explicit nil

### UnsetAppOwners
`func (o *C1ApiPolicyV1Approval) UnsetAppOwners()`

UnsetAppOwners ensures that no value is present for AppOwners, not even an explicit nil
### GetAssigned

`func (o *C1ApiPolicyV1Approval) GetAssigned() bool`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *C1ApiPolicyV1Approval) GetAssignedOk() (*bool, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *C1ApiPolicyV1Approval) SetAssigned(v bool)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *C1ApiPolicyV1Approval) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### GetEntitlementOwners

`func (o *C1ApiPolicyV1Approval) GetEntitlementOwners() C1ApiPolicyV1EntitlementOwnerApproval`

GetEntitlementOwners returns the EntitlementOwners field if non-nil, zero value otherwise.

### GetEntitlementOwnersOk

`func (o *C1ApiPolicyV1Approval) GetEntitlementOwnersOk() (*C1ApiPolicyV1EntitlementOwnerApproval, bool)`

GetEntitlementOwnersOk returns a tuple with the EntitlementOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementOwners

`func (o *C1ApiPolicyV1Approval) SetEntitlementOwners(v C1ApiPolicyV1EntitlementOwnerApproval)`

SetEntitlementOwners sets EntitlementOwners field to given value.

### HasEntitlementOwners

`func (o *C1ApiPolicyV1Approval) HasEntitlementOwners() bool`

HasEntitlementOwners returns a boolean if a field has been set.

### SetEntitlementOwnersNil

`func (o *C1ApiPolicyV1Approval) SetEntitlementOwnersNil(b bool)`

 SetEntitlementOwnersNil sets the value for EntitlementOwners to be an explicit nil

### UnsetEntitlementOwners
`func (o *C1ApiPolicyV1Approval) UnsetEntitlementOwners()`

UnsetEntitlementOwners ensures that no value is present for EntitlementOwners, not even an explicit nil
### GetGroup

`func (o *C1ApiPolicyV1Approval) GetGroup() C1ApiPolicyV1AppGroupApproval`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *C1ApiPolicyV1Approval) GetGroupOk() (*C1ApiPolicyV1AppGroupApproval, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *C1ApiPolicyV1Approval) SetGroup(v C1ApiPolicyV1AppGroupApproval)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *C1ApiPolicyV1Approval) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *C1ApiPolicyV1Approval) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *C1ApiPolicyV1Approval) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetManager

`func (o *C1ApiPolicyV1Approval) GetManager() C1ApiPolicyV1ManagerApproval`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *C1ApiPolicyV1Approval) GetManagerOk() (*C1ApiPolicyV1ManagerApproval, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *C1ApiPolicyV1Approval) SetManager(v C1ApiPolicyV1ManagerApproval)`

SetManager sets Manager field to given value.

### HasManager

`func (o *C1ApiPolicyV1Approval) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *C1ApiPolicyV1Approval) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *C1ApiPolicyV1Approval) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetRequireApprovalReason

`func (o *C1ApiPolicyV1Approval) GetRequireApprovalReason() bool`

GetRequireApprovalReason returns the RequireApprovalReason field if non-nil, zero value otherwise.

### GetRequireApprovalReasonOk

`func (o *C1ApiPolicyV1Approval) GetRequireApprovalReasonOk() (*bool, bool)`

GetRequireApprovalReasonOk returns a tuple with the RequireApprovalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireApprovalReason

`func (o *C1ApiPolicyV1Approval) SetRequireApprovalReason(v bool)`

SetRequireApprovalReason sets RequireApprovalReason field to given value.

### HasRequireApprovalReason

`func (o *C1ApiPolicyV1Approval) HasRequireApprovalReason() bool`

HasRequireApprovalReason returns a boolean if a field has been set.

### GetRequireReassignmentReason

`func (o *C1ApiPolicyV1Approval) GetRequireReassignmentReason() bool`

GetRequireReassignmentReason returns the RequireReassignmentReason field if non-nil, zero value otherwise.

### GetRequireReassignmentReasonOk

`func (o *C1ApiPolicyV1Approval) GetRequireReassignmentReasonOk() (*bool, bool)`

GetRequireReassignmentReasonOk returns a tuple with the RequireReassignmentReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireReassignmentReason

`func (o *C1ApiPolicyV1Approval) SetRequireReassignmentReason(v bool)`

SetRequireReassignmentReason sets RequireReassignmentReason field to given value.

### HasRequireReassignmentReason

`func (o *C1ApiPolicyV1Approval) HasRequireReassignmentReason() bool`

HasRequireReassignmentReason returns a boolean if a field has been set.

### GetSelf

`func (o *C1ApiPolicyV1Approval) GetSelf() C1ApiPolicyV1SelfApproval`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *C1ApiPolicyV1Approval) GetSelfOk() (*C1ApiPolicyV1SelfApproval, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *C1ApiPolicyV1Approval) SetSelf(v C1ApiPolicyV1SelfApproval)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *C1ApiPolicyV1Approval) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### SetSelfNil

`func (o *C1ApiPolicyV1Approval) SetSelfNil(b bool)`

 SetSelfNil sets the value for Self to be an explicit nil

### UnsetSelf
`func (o *C1ApiPolicyV1Approval) UnsetSelf()`

UnsetSelf ensures that no value is present for Self, not even an explicit nil
### GetUsers

`func (o *C1ApiPolicyV1Approval) GetUsers() C1ApiPolicyV1UserApproval`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *C1ApiPolicyV1Approval) GetUsersOk() (*C1ApiPolicyV1UserApproval, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *C1ApiPolicyV1Approval) SetUsers(v C1ApiPolicyV1UserApproval)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *C1ApiPolicyV1Approval) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *C1ApiPolicyV1Approval) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *C1ApiPolicyV1Approval) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


