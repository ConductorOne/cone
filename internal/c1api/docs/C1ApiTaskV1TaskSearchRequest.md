# C1ApiTaskV1TaskSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessReviewIds** | Pointer to **[]string** | The accessReviewIds field. | [optional] 
**AccountOwnerIds** | Pointer to **[]string** | The accountOwnerIds field. | [optional] 
**ActorId** | Pointer to **string** | The actorId field. | [optional] 
**AppEntitlementIds** | Pointer to **[]string** | The appEntitlementIds field. | [optional] 
**AppResourceIds** | Pointer to **[]string** | The appResourceIds field. | [optional] 
**AppResourceTypeIds** | Pointer to **[]string** | The appResourceTypeIds field. | [optional] 
**AppUserSubjectIds** | Pointer to **[]string** |  Find Tasks which are referncing a Set of AppUserIDs  | [optional] 
**ApplicationIds** | Pointer to **[]string** | The applicationIds field. | [optional] 
**AssigneesInIds** | Pointer to **[]string** |  Search tasks by  List of UserIDs which are currently assigned these Tasks  | [optional] 
**CreatedAfter** | Pointer to **time.Time** |  | [optional] 
**CreatedBefore** | Pointer to **time.Time** |  | [optional] 
**CurrentStep** | Pointer to **string** | The currentStep field. | [optional] 
**ExcludeAppEntitlementIds** | Pointer to **[]string** | The excludeAppEntitlementIds field. | [optional] 
**ExcludeIds** | Pointer to **[]string** |  Exclude Specific TaskIDs from this serach result.  | [optional] 
**ExpandMask** | Pointer to [**C1ApiTaskV1TaskExpandMask**](C1ApiTaskV1TaskExpandMask.md) |  | [optional] 
**IncludeDeleted** | Pointer to **bool** | The includeDeleted field. | [optional] 
**MyWorkUserIds** | Pointer to **[]string** |  Search tasks by a List of UserIDs which are currently assigned to OR have previously acted upon this Task  | [optional] 
**OpenerIds** | Pointer to **[]string** |  Find a Task which was opened by UserIDs  | [optional] 
**PageSize** | Pointer to **float32** | The pageSize field. | [optional] 
**PageToken** | Pointer to **string** | The pageToken field. | [optional] 
**PreviouslyActedOnIds** | Pointer to **[]string** |  Search tasks by a  List of UserIDs which have previously approved or otherwise acted upon this Task  | [optional] 
**Query** | Pointer to **string** | The query field. | [optional] 
**Refs** | Pointer to [**[]C1ApiTaskV1TaskRef**](C1ApiTaskV1TaskRef.md) | The refs field. | [optional] 
**SortBy** | Pointer to **string** | The sortBy field. | [optional] 
**SubjectIds** | Pointer to **[]string** |  Find Tasks which are referncing this C1 UserID  | [optional] 
**TaskStates** | Pointer to **[]string** | The taskStates field. | [optional] 
**TaskTypes** | Pointer to [**[]C1ApiTaskV1TaskType**](C1ApiTaskV1TaskType.md) |  TODO(pquerna): why is this a MESSAGE that only CONTAINS AN ENUM?  | [optional] 

## Methods

### NewC1ApiTaskV1TaskSearchRequest

`func NewC1ApiTaskV1TaskSearchRequest() *C1ApiTaskV1TaskSearchRequest`

NewC1ApiTaskV1TaskSearchRequest instantiates a new C1ApiTaskV1TaskSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiTaskV1TaskSearchRequestWithDefaults

`func NewC1ApiTaskV1TaskSearchRequestWithDefaults() *C1ApiTaskV1TaskSearchRequest`

NewC1ApiTaskV1TaskSearchRequestWithDefaults instantiates a new C1ApiTaskV1TaskSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessReviewIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetAccessReviewIds() []string`

GetAccessReviewIds returns the AccessReviewIds field if non-nil, zero value otherwise.

### GetAccessReviewIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetAccessReviewIdsOk() (*[]string, bool)`

GetAccessReviewIdsOk returns a tuple with the AccessReviewIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessReviewIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetAccessReviewIds(v []string)`

SetAccessReviewIds sets AccessReviewIds field to given value.

### HasAccessReviewIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasAccessReviewIds() bool`

HasAccessReviewIds returns a boolean if a field has been set.

### SetAccessReviewIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetAccessReviewIdsNil(b bool)`

 SetAccessReviewIdsNil sets the value for AccessReviewIds to be an explicit nil

### UnsetAccessReviewIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetAccessReviewIds()`

UnsetAccessReviewIds ensures that no value is present for AccessReviewIds, not even an explicit nil
### GetAccountOwnerIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetAccountOwnerIds() []string`

GetAccountOwnerIds returns the AccountOwnerIds field if non-nil, zero value otherwise.

### GetAccountOwnerIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetAccountOwnerIdsOk() (*[]string, bool)`

GetAccountOwnerIdsOk returns a tuple with the AccountOwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOwnerIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetAccountOwnerIds(v []string)`

SetAccountOwnerIds sets AccountOwnerIds field to given value.

### HasAccountOwnerIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasAccountOwnerIds() bool`

HasAccountOwnerIds returns a boolean if a field has been set.

### SetAccountOwnerIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetAccountOwnerIdsNil(b bool)`

 SetAccountOwnerIdsNil sets the value for AccountOwnerIds to be an explicit nil

### UnsetAccountOwnerIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetAccountOwnerIds()`

UnsetAccountOwnerIds ensures that no value is present for AccountOwnerIds, not even an explicit nil
### GetActorId

`func (o *C1ApiTaskV1TaskSearchRequest) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *C1ApiTaskV1TaskSearchRequest) SetActorId(v string)`

SetActorId sets ActorId field to given value.

### HasActorId

`func (o *C1ApiTaskV1TaskSearchRequest) HasActorId() bool`

HasActorId returns a boolean if a field has been set.

### GetAppEntitlementIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetAppEntitlementIds() []string`

GetAppEntitlementIds returns the AppEntitlementIds field if non-nil, zero value otherwise.

### GetAppEntitlementIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetAppEntitlementIdsOk() (*[]string, bool)`

GetAppEntitlementIdsOk returns a tuple with the AppEntitlementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEntitlementIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetAppEntitlementIds(v []string)`

SetAppEntitlementIds sets AppEntitlementIds field to given value.

### HasAppEntitlementIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasAppEntitlementIds() bool`

HasAppEntitlementIds returns a boolean if a field has been set.

### SetAppEntitlementIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetAppEntitlementIdsNil(b bool)`

 SetAppEntitlementIdsNil sets the value for AppEntitlementIds to be an explicit nil

### UnsetAppEntitlementIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetAppEntitlementIds()`

UnsetAppEntitlementIds ensures that no value is present for AppEntitlementIds, not even an explicit nil
### GetAppResourceIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetAppResourceIds() []string`

GetAppResourceIds returns the AppResourceIds field if non-nil, zero value otherwise.

### GetAppResourceIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetAppResourceIdsOk() (*[]string, bool)`

GetAppResourceIdsOk returns a tuple with the AppResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppResourceIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetAppResourceIds(v []string)`

SetAppResourceIds sets AppResourceIds field to given value.

### HasAppResourceIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasAppResourceIds() bool`

HasAppResourceIds returns a boolean if a field has been set.

### SetAppResourceIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetAppResourceIdsNil(b bool)`

 SetAppResourceIdsNil sets the value for AppResourceIds to be an explicit nil

### UnsetAppResourceIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetAppResourceIds()`

UnsetAppResourceIds ensures that no value is present for AppResourceIds, not even an explicit nil
### GetAppResourceTypeIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetAppResourceTypeIds() []string`

GetAppResourceTypeIds returns the AppResourceTypeIds field if non-nil, zero value otherwise.

### GetAppResourceTypeIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetAppResourceTypeIdsOk() (*[]string, bool)`

GetAppResourceTypeIdsOk returns a tuple with the AppResourceTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppResourceTypeIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetAppResourceTypeIds(v []string)`

SetAppResourceTypeIds sets AppResourceTypeIds field to given value.

### HasAppResourceTypeIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasAppResourceTypeIds() bool`

HasAppResourceTypeIds returns a boolean if a field has been set.

### SetAppResourceTypeIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetAppResourceTypeIdsNil(b bool)`

 SetAppResourceTypeIdsNil sets the value for AppResourceTypeIds to be an explicit nil

### UnsetAppResourceTypeIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetAppResourceTypeIds()`

UnsetAppResourceTypeIds ensures that no value is present for AppResourceTypeIds, not even an explicit nil
### GetAppUserSubjectIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetAppUserSubjectIds() []string`

GetAppUserSubjectIds returns the AppUserSubjectIds field if non-nil, zero value otherwise.

### GetAppUserSubjectIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetAppUserSubjectIdsOk() (*[]string, bool)`

GetAppUserSubjectIdsOk returns a tuple with the AppUserSubjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUserSubjectIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetAppUserSubjectIds(v []string)`

SetAppUserSubjectIds sets AppUserSubjectIds field to given value.

### HasAppUserSubjectIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasAppUserSubjectIds() bool`

HasAppUserSubjectIds returns a boolean if a field has been set.

### SetAppUserSubjectIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetAppUserSubjectIdsNil(b bool)`

 SetAppUserSubjectIdsNil sets the value for AppUserSubjectIds to be an explicit nil

### UnsetAppUserSubjectIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetAppUserSubjectIds()`

UnsetAppUserSubjectIds ensures that no value is present for AppUserSubjectIds, not even an explicit nil
### GetApplicationIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetApplicationIds() []string`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetApplicationIdsOk() (*[]string, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetApplicationIds(v []string)`

SetApplicationIds sets ApplicationIds field to given value.

### HasApplicationIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasApplicationIds() bool`

HasApplicationIds returns a boolean if a field has been set.

### SetApplicationIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetApplicationIdsNil(b bool)`

 SetApplicationIdsNil sets the value for ApplicationIds to be an explicit nil

### UnsetApplicationIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetApplicationIds()`

UnsetApplicationIds ensures that no value is present for ApplicationIds, not even an explicit nil
### GetAssigneesInIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetAssigneesInIds() []string`

GetAssigneesInIds returns the AssigneesInIds field if non-nil, zero value otherwise.

### GetAssigneesInIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetAssigneesInIdsOk() (*[]string, bool)`

GetAssigneesInIdsOk returns a tuple with the AssigneesInIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneesInIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetAssigneesInIds(v []string)`

SetAssigneesInIds sets AssigneesInIds field to given value.

### HasAssigneesInIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasAssigneesInIds() bool`

HasAssigneesInIds returns a boolean if a field has been set.

### SetAssigneesInIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetAssigneesInIdsNil(b bool)`

 SetAssigneesInIdsNil sets the value for AssigneesInIds to be an explicit nil

### UnsetAssigneesInIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetAssigneesInIds()`

UnsetAssigneesInIds ensures that no value is present for AssigneesInIds, not even an explicit nil
### GetCreatedAfter

`func (o *C1ApiTaskV1TaskSearchRequest) GetCreatedAfter() time.Time`

GetCreatedAfter returns the CreatedAfter field if non-nil, zero value otherwise.

### GetCreatedAfterOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetCreatedAfterOk() (*time.Time, bool)`

GetCreatedAfterOk returns a tuple with the CreatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAfter

`func (o *C1ApiTaskV1TaskSearchRequest) SetCreatedAfter(v time.Time)`

SetCreatedAfter sets CreatedAfter field to given value.

### HasCreatedAfter

`func (o *C1ApiTaskV1TaskSearchRequest) HasCreatedAfter() bool`

HasCreatedAfter returns a boolean if a field has been set.

### GetCreatedBefore

`func (o *C1ApiTaskV1TaskSearchRequest) GetCreatedBefore() time.Time`

GetCreatedBefore returns the CreatedBefore field if non-nil, zero value otherwise.

### GetCreatedBeforeOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetCreatedBeforeOk() (*time.Time, bool)`

GetCreatedBeforeOk returns a tuple with the CreatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBefore

`func (o *C1ApiTaskV1TaskSearchRequest) SetCreatedBefore(v time.Time)`

SetCreatedBefore sets CreatedBefore field to given value.

### HasCreatedBefore

`func (o *C1ApiTaskV1TaskSearchRequest) HasCreatedBefore() bool`

HasCreatedBefore returns a boolean if a field has been set.

### GetCurrentStep

`func (o *C1ApiTaskV1TaskSearchRequest) GetCurrentStep() string`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetCurrentStepOk() (*string, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *C1ApiTaskV1TaskSearchRequest) SetCurrentStep(v string)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *C1ApiTaskV1TaskSearchRequest) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### GetExcludeAppEntitlementIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetExcludeAppEntitlementIds() []string`

GetExcludeAppEntitlementIds returns the ExcludeAppEntitlementIds field if non-nil, zero value otherwise.

### GetExcludeAppEntitlementIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetExcludeAppEntitlementIdsOk() (*[]string, bool)`

GetExcludeAppEntitlementIdsOk returns a tuple with the ExcludeAppEntitlementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAppEntitlementIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetExcludeAppEntitlementIds(v []string)`

SetExcludeAppEntitlementIds sets ExcludeAppEntitlementIds field to given value.

### HasExcludeAppEntitlementIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasExcludeAppEntitlementIds() bool`

HasExcludeAppEntitlementIds returns a boolean if a field has been set.

### SetExcludeAppEntitlementIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetExcludeAppEntitlementIdsNil(b bool)`

 SetExcludeAppEntitlementIdsNil sets the value for ExcludeAppEntitlementIds to be an explicit nil

### UnsetExcludeAppEntitlementIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetExcludeAppEntitlementIds()`

UnsetExcludeAppEntitlementIds ensures that no value is present for ExcludeAppEntitlementIds, not even an explicit nil
### GetExcludeIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetExcludeIds() []string`

GetExcludeIds returns the ExcludeIds field if non-nil, zero value otherwise.

### GetExcludeIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetExcludeIdsOk() (*[]string, bool)`

GetExcludeIdsOk returns a tuple with the ExcludeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetExcludeIds(v []string)`

SetExcludeIds sets ExcludeIds field to given value.

### HasExcludeIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasExcludeIds() bool`

HasExcludeIds returns a boolean if a field has been set.

### SetExcludeIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetExcludeIdsNil(b bool)`

 SetExcludeIdsNil sets the value for ExcludeIds to be an explicit nil

### UnsetExcludeIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetExcludeIds()`

UnsetExcludeIds ensures that no value is present for ExcludeIds, not even an explicit nil
### GetExpandMask

`func (o *C1ApiTaskV1TaskSearchRequest) GetExpandMask() C1ApiTaskV1TaskExpandMask`

GetExpandMask returns the ExpandMask field if non-nil, zero value otherwise.

### GetExpandMaskOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetExpandMaskOk() (*C1ApiTaskV1TaskExpandMask, bool)`

GetExpandMaskOk returns a tuple with the ExpandMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandMask

`func (o *C1ApiTaskV1TaskSearchRequest) SetExpandMask(v C1ApiTaskV1TaskExpandMask)`

SetExpandMask sets ExpandMask field to given value.

### HasExpandMask

`func (o *C1ApiTaskV1TaskSearchRequest) HasExpandMask() bool`

HasExpandMask returns a boolean if a field has been set.

### GetIncludeDeleted

`func (o *C1ApiTaskV1TaskSearchRequest) GetIncludeDeleted() bool`

GetIncludeDeleted returns the IncludeDeleted field if non-nil, zero value otherwise.

### GetIncludeDeletedOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetIncludeDeletedOk() (*bool, bool)`

GetIncludeDeletedOk returns a tuple with the IncludeDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDeleted

`func (o *C1ApiTaskV1TaskSearchRequest) SetIncludeDeleted(v bool)`

SetIncludeDeleted sets IncludeDeleted field to given value.

### HasIncludeDeleted

`func (o *C1ApiTaskV1TaskSearchRequest) HasIncludeDeleted() bool`

HasIncludeDeleted returns a boolean if a field has been set.

### GetMyWorkUserIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetMyWorkUserIds() []string`

GetMyWorkUserIds returns the MyWorkUserIds field if non-nil, zero value otherwise.

### GetMyWorkUserIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetMyWorkUserIdsOk() (*[]string, bool)`

GetMyWorkUserIdsOk returns a tuple with the MyWorkUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyWorkUserIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetMyWorkUserIds(v []string)`

SetMyWorkUserIds sets MyWorkUserIds field to given value.

### HasMyWorkUserIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasMyWorkUserIds() bool`

HasMyWorkUserIds returns a boolean if a field has been set.

### SetMyWorkUserIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetMyWorkUserIdsNil(b bool)`

 SetMyWorkUserIdsNil sets the value for MyWorkUserIds to be an explicit nil

### UnsetMyWorkUserIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetMyWorkUserIds()`

UnsetMyWorkUserIds ensures that no value is present for MyWorkUserIds, not even an explicit nil
### GetOpenerIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetOpenerIds() []string`

GetOpenerIds returns the OpenerIds field if non-nil, zero value otherwise.

### GetOpenerIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetOpenerIdsOk() (*[]string, bool)`

GetOpenerIdsOk returns a tuple with the OpenerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenerIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetOpenerIds(v []string)`

SetOpenerIds sets OpenerIds field to given value.

### HasOpenerIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasOpenerIds() bool`

HasOpenerIds returns a boolean if a field has been set.

### SetOpenerIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetOpenerIdsNil(b bool)`

 SetOpenerIdsNil sets the value for OpenerIds to be an explicit nil

### UnsetOpenerIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetOpenerIds()`

UnsetOpenerIds ensures that no value is present for OpenerIds, not even an explicit nil
### GetPageSize

`func (o *C1ApiTaskV1TaskSearchRequest) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *C1ApiTaskV1TaskSearchRequest) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *C1ApiTaskV1TaskSearchRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *C1ApiTaskV1TaskSearchRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *C1ApiTaskV1TaskSearchRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *C1ApiTaskV1TaskSearchRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### GetPreviouslyActedOnIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetPreviouslyActedOnIds() []string`

GetPreviouslyActedOnIds returns the PreviouslyActedOnIds field if non-nil, zero value otherwise.

### GetPreviouslyActedOnIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetPreviouslyActedOnIdsOk() (*[]string, bool)`

GetPreviouslyActedOnIdsOk returns a tuple with the PreviouslyActedOnIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviouslyActedOnIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetPreviouslyActedOnIds(v []string)`

SetPreviouslyActedOnIds sets PreviouslyActedOnIds field to given value.

### HasPreviouslyActedOnIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasPreviouslyActedOnIds() bool`

HasPreviouslyActedOnIds returns a boolean if a field has been set.

### SetPreviouslyActedOnIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetPreviouslyActedOnIdsNil(b bool)`

 SetPreviouslyActedOnIdsNil sets the value for PreviouslyActedOnIds to be an explicit nil

### UnsetPreviouslyActedOnIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetPreviouslyActedOnIds()`

UnsetPreviouslyActedOnIds ensures that no value is present for PreviouslyActedOnIds, not even an explicit nil
### GetQuery

`func (o *C1ApiTaskV1TaskSearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *C1ApiTaskV1TaskSearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *C1ApiTaskV1TaskSearchRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRefs

`func (o *C1ApiTaskV1TaskSearchRequest) GetRefs() []C1ApiTaskV1TaskRef`

GetRefs returns the Refs field if non-nil, zero value otherwise.

### GetRefsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetRefsOk() (*[]C1ApiTaskV1TaskRef, bool)`

GetRefsOk returns a tuple with the Refs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefs

`func (o *C1ApiTaskV1TaskSearchRequest) SetRefs(v []C1ApiTaskV1TaskRef)`

SetRefs sets Refs field to given value.

### HasRefs

`func (o *C1ApiTaskV1TaskSearchRequest) HasRefs() bool`

HasRefs returns a boolean if a field has been set.

### SetRefsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetRefsNil(b bool)`

 SetRefsNil sets the value for Refs to be an explicit nil

### UnsetRefs
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetRefs()`

UnsetRefs ensures that no value is present for Refs, not even an explicit nil
### GetSortBy

`func (o *C1ApiTaskV1TaskSearchRequest) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *C1ApiTaskV1TaskSearchRequest) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *C1ApiTaskV1TaskSearchRequest) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetSubjectIds

`func (o *C1ApiTaskV1TaskSearchRequest) GetSubjectIds() []string`

GetSubjectIds returns the SubjectIds field if non-nil, zero value otherwise.

### GetSubjectIdsOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetSubjectIdsOk() (*[]string, bool)`

GetSubjectIdsOk returns a tuple with the SubjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectIds

`func (o *C1ApiTaskV1TaskSearchRequest) SetSubjectIds(v []string)`

SetSubjectIds sets SubjectIds field to given value.

### HasSubjectIds

`func (o *C1ApiTaskV1TaskSearchRequest) HasSubjectIds() bool`

HasSubjectIds returns a boolean if a field has been set.

### SetSubjectIdsNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetSubjectIdsNil(b bool)`

 SetSubjectIdsNil sets the value for SubjectIds to be an explicit nil

### UnsetSubjectIds
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetSubjectIds()`

UnsetSubjectIds ensures that no value is present for SubjectIds, not even an explicit nil
### GetTaskStates

`func (o *C1ApiTaskV1TaskSearchRequest) GetTaskStates() []string`

GetTaskStates returns the TaskStates field if non-nil, zero value otherwise.

### GetTaskStatesOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetTaskStatesOk() (*[]string, bool)`

GetTaskStatesOk returns a tuple with the TaskStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStates

`func (o *C1ApiTaskV1TaskSearchRequest) SetTaskStates(v []string)`

SetTaskStates sets TaskStates field to given value.

### HasTaskStates

`func (o *C1ApiTaskV1TaskSearchRequest) HasTaskStates() bool`

HasTaskStates returns a boolean if a field has been set.

### SetTaskStatesNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetTaskStatesNil(b bool)`

 SetTaskStatesNil sets the value for TaskStates to be an explicit nil

### UnsetTaskStates
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetTaskStates()`

UnsetTaskStates ensures that no value is present for TaskStates, not even an explicit nil
### GetTaskTypes

`func (o *C1ApiTaskV1TaskSearchRequest) GetTaskTypes() []C1ApiTaskV1TaskType`

GetTaskTypes returns the TaskTypes field if non-nil, zero value otherwise.

### GetTaskTypesOk

`func (o *C1ApiTaskV1TaskSearchRequest) GetTaskTypesOk() (*[]C1ApiTaskV1TaskType, bool)`

GetTaskTypesOk returns a tuple with the TaskTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTypes

`func (o *C1ApiTaskV1TaskSearchRequest) SetTaskTypes(v []C1ApiTaskV1TaskType)`

SetTaskTypes sets TaskTypes field to given value.

### HasTaskTypes

`func (o *C1ApiTaskV1TaskSearchRequest) HasTaskTypes() bool`

HasTaskTypes returns a boolean if a field has been set.

### SetTaskTypesNil

`func (o *C1ApiTaskV1TaskSearchRequest) SetTaskTypesNil(b bool)`

 SetTaskTypesNil sets the value for TaskTypes to be an explicit nil

### UnsetTaskTypes
`func (o *C1ApiTaskV1TaskSearchRequest) UnsetTaskTypes()`

UnsetTaskTypes ensures that no value is present for TaskTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


