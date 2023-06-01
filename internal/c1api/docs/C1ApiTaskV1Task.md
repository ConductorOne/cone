# C1ApiTaskV1Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | The actions field. | [optional] 
**AnalysisId** | Pointer to **string** | The analysisId field. | [optional] 
**Annotations** | Pointer to [**[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner**](C1ApiAppV1AppResourceServiceGetResponseExpandedInner.md) | The annotations field. | [optional] 
**CommentCount** | Pointer to **float32** | The commentCount field. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedByUserId** | Pointer to **string** | The createdByUserId field. | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** | The description field. | [optional] 
**DisplayName** | Pointer to **string** | The displayName field. | [optional] 
**ExternalRefs** | Pointer to [**[]C1ApiTaskV1ExternalRef**](C1ApiTaskV1ExternalRef.md) | The externalRefs field. | [optional] 
**Id** | Pointer to **string** |  General Metadata  | [optional] 
**NumericId** | Pointer to **string** | The numericId field. | [optional] 
**Policy** | Pointer to [**C1ApiPolicyV1PolicyInstance**](C1ApiPolicyV1PolicyInstance.md) |  | [optional] 
**Processing** | Pointer to **string** | The processing field. | [optional] 
**State** | Pointer to **string** |  State  | [optional] 
**StepApproverIds** | Pointer to **[]string** | The stepApproverIds field. | [optional] 
**Type** | Pointer to [**C1ApiTaskV1TaskType**](C1ApiTaskV1TaskType.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **string** |  External IDS  | [optional] 

## Methods

### NewC1ApiTaskV1Task

`func NewC1ApiTaskV1Task() *C1ApiTaskV1Task`

NewC1ApiTaskV1Task instantiates a new C1ApiTaskV1Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiTaskV1TaskWithDefaults

`func NewC1ApiTaskV1TaskWithDefaults() *C1ApiTaskV1Task`

NewC1ApiTaskV1TaskWithDefaults instantiates a new C1ApiTaskV1Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *C1ApiTaskV1Task) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *C1ApiTaskV1Task) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *C1ApiTaskV1Task) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *C1ApiTaskV1Task) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *C1ApiTaskV1Task) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *C1ApiTaskV1Task) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetAnalysisId

`func (o *C1ApiTaskV1Task) GetAnalysisId() string`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *C1ApiTaskV1Task) GetAnalysisIdOk() (*string, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *C1ApiTaskV1Task) SetAnalysisId(v string)`

SetAnalysisId sets AnalysisId field to given value.

### HasAnalysisId

`func (o *C1ApiTaskV1Task) HasAnalysisId() bool`

HasAnalysisId returns a boolean if a field has been set.

### GetAnnotations

`func (o *C1ApiTaskV1Task) GetAnnotations() []C1ApiAppV1AppResourceServiceGetResponseExpandedInner`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *C1ApiTaskV1Task) GetAnnotationsOk() (*[]C1ApiAppV1AppResourceServiceGetResponseExpandedInner, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *C1ApiTaskV1Task) SetAnnotations(v []C1ApiAppV1AppResourceServiceGetResponseExpandedInner)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *C1ApiTaskV1Task) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *C1ApiTaskV1Task) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *C1ApiTaskV1Task) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetCommentCount

`func (o *C1ApiTaskV1Task) GetCommentCount() float32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *C1ApiTaskV1Task) GetCommentCountOk() (*float32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *C1ApiTaskV1Task) SetCommentCount(v float32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *C1ApiTaskV1Task) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *C1ApiTaskV1Task) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *C1ApiTaskV1Task) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *C1ApiTaskV1Task) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *C1ApiTaskV1Task) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *C1ApiTaskV1Task) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *C1ApiTaskV1Task) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *C1ApiTaskV1Task) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *C1ApiTaskV1Task) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetDeletedAt

`func (o *C1ApiTaskV1Task) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *C1ApiTaskV1Task) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *C1ApiTaskV1Task) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *C1ApiTaskV1Task) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *C1ApiTaskV1Task) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *C1ApiTaskV1Task) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *C1ApiTaskV1Task) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *C1ApiTaskV1Task) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *C1ApiTaskV1Task) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *C1ApiTaskV1Task) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *C1ApiTaskV1Task) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *C1ApiTaskV1Task) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalRefs

`func (o *C1ApiTaskV1Task) GetExternalRefs() []C1ApiTaskV1ExternalRef`

GetExternalRefs returns the ExternalRefs field if non-nil, zero value otherwise.

### GetExternalRefsOk

`func (o *C1ApiTaskV1Task) GetExternalRefsOk() (*[]C1ApiTaskV1ExternalRef, bool)`

GetExternalRefsOk returns a tuple with the ExternalRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefs

`func (o *C1ApiTaskV1Task) SetExternalRefs(v []C1ApiTaskV1ExternalRef)`

SetExternalRefs sets ExternalRefs field to given value.

### HasExternalRefs

`func (o *C1ApiTaskV1Task) HasExternalRefs() bool`

HasExternalRefs returns a boolean if a field has been set.

### SetExternalRefsNil

`func (o *C1ApiTaskV1Task) SetExternalRefsNil(b bool)`

 SetExternalRefsNil sets the value for ExternalRefs to be an explicit nil

### UnsetExternalRefs
`func (o *C1ApiTaskV1Task) UnsetExternalRefs()`

UnsetExternalRefs ensures that no value is present for ExternalRefs, not even an explicit nil
### GetId

`func (o *C1ApiTaskV1Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *C1ApiTaskV1Task) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *C1ApiTaskV1Task) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *C1ApiTaskV1Task) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumericId

`func (o *C1ApiTaskV1Task) GetNumericId() string`

GetNumericId returns the NumericId field if non-nil, zero value otherwise.

### GetNumericIdOk

`func (o *C1ApiTaskV1Task) GetNumericIdOk() (*string, bool)`

GetNumericIdOk returns a tuple with the NumericId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericId

`func (o *C1ApiTaskV1Task) SetNumericId(v string)`

SetNumericId sets NumericId field to given value.

### HasNumericId

`func (o *C1ApiTaskV1Task) HasNumericId() bool`

HasNumericId returns a boolean if a field has been set.

### GetPolicy

`func (o *C1ApiTaskV1Task) GetPolicy() C1ApiPolicyV1PolicyInstance`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *C1ApiTaskV1Task) GetPolicyOk() (*C1ApiPolicyV1PolicyInstance, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *C1ApiTaskV1Task) SetPolicy(v C1ApiPolicyV1PolicyInstance)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *C1ApiTaskV1Task) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetProcessing

`func (o *C1ApiTaskV1Task) GetProcessing() string`

GetProcessing returns the Processing field if non-nil, zero value otherwise.

### GetProcessingOk

`func (o *C1ApiTaskV1Task) GetProcessingOk() (*string, bool)`

GetProcessingOk returns a tuple with the Processing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessing

`func (o *C1ApiTaskV1Task) SetProcessing(v string)`

SetProcessing sets Processing field to given value.

### HasProcessing

`func (o *C1ApiTaskV1Task) HasProcessing() bool`

HasProcessing returns a boolean if a field has been set.

### GetState

`func (o *C1ApiTaskV1Task) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *C1ApiTaskV1Task) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *C1ApiTaskV1Task) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *C1ApiTaskV1Task) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStepApproverIds

`func (o *C1ApiTaskV1Task) GetStepApproverIds() []string`

GetStepApproverIds returns the StepApproverIds field if non-nil, zero value otherwise.

### GetStepApproverIdsOk

`func (o *C1ApiTaskV1Task) GetStepApproverIdsOk() (*[]string, bool)`

GetStepApproverIdsOk returns a tuple with the StepApproverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepApproverIds

`func (o *C1ApiTaskV1Task) SetStepApproverIds(v []string)`

SetStepApproverIds sets StepApproverIds field to given value.

### HasStepApproverIds

`func (o *C1ApiTaskV1Task) HasStepApproverIds() bool`

HasStepApproverIds returns a boolean if a field has been set.

### SetStepApproverIdsNil

`func (o *C1ApiTaskV1Task) SetStepApproverIdsNil(b bool)`

 SetStepApproverIdsNil sets the value for StepApproverIds to be an explicit nil

### UnsetStepApproverIds
`func (o *C1ApiTaskV1Task) UnsetStepApproverIds()`

UnsetStepApproverIds ensures that no value is present for StepApproverIds, not even an explicit nil
### GetType

`func (o *C1ApiTaskV1Task) GetType() C1ApiTaskV1TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *C1ApiTaskV1Task) GetTypeOk() (*C1ApiTaskV1TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *C1ApiTaskV1Task) SetType(v C1ApiTaskV1TaskType)`

SetType sets Type field to given value.

### HasType

`func (o *C1ApiTaskV1Task) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *C1ApiTaskV1Task) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *C1ApiTaskV1Task) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *C1ApiTaskV1Task) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *C1ApiTaskV1Task) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *C1ApiTaskV1Task) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *C1ApiTaskV1Task) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *C1ApiTaskV1Task) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *C1ApiTaskV1Task) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


