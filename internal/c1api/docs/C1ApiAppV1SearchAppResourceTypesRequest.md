# C1ApiAppV1SearchAppResourceTypesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppIds** | Pointer to **[]string** | The appIds field. | [optional] 
**ExcludeResourceTypeIds** | Pointer to **[]string** | The excludeResourceTypeIds field. | [optional] 
**ExcludeResourceTypeTraitIds** | Pointer to **[]string** | The excludeResourceTypeTraitIds field. | [optional] 
**PageSize** | Pointer to **float32** | The pageSize field. | [optional] 
**PageToken** | Pointer to **string** | The pageToken field. | [optional] 
**Query** | Pointer to **string** | The query field. | [optional] 
**ResourceTypeIds** | Pointer to **[]string** | The resourceTypeIds field. | [optional] 
**ResourceTypeTraitIds** | Pointer to **[]string** | The resourceTypeTraitIds field. | [optional] 

## Methods

### NewC1ApiAppV1SearchAppResourceTypesRequest

`func NewC1ApiAppV1SearchAppResourceTypesRequest() *C1ApiAppV1SearchAppResourceTypesRequest`

NewC1ApiAppV1SearchAppResourceTypesRequest instantiates a new C1ApiAppV1SearchAppResourceTypesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiAppV1SearchAppResourceTypesRequestWithDefaults

`func NewC1ApiAppV1SearchAppResourceTypesRequestWithDefaults() *C1ApiAppV1SearchAppResourceTypesRequest`

NewC1ApiAppV1SearchAppResourceTypesRequestWithDefaults instantiates a new C1ApiAppV1SearchAppResourceTypesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### SetAppIdsNil

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) SetAppIdsNil(b bool)`

 SetAppIdsNil sets the value for AppIds to be an explicit nil

### UnsetAppIds
`func (o *C1ApiAppV1SearchAppResourceTypesRequest) UnsetAppIds()`

UnsetAppIds ensures that no value is present for AppIds, not even an explicit nil
### GetExcludeResourceTypeIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetExcludeResourceTypeIds() []string`

GetExcludeResourceTypeIds returns the ExcludeResourceTypeIds field if non-nil, zero value otherwise.

### GetExcludeResourceTypeIdsOk

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetExcludeResourceTypeIdsOk() (*[]string, bool)`

GetExcludeResourceTypeIdsOk returns a tuple with the ExcludeResourceTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeResourceTypeIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) SetExcludeResourceTypeIds(v []string)`

SetExcludeResourceTypeIds sets ExcludeResourceTypeIds field to given value.

### HasExcludeResourceTypeIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) HasExcludeResourceTypeIds() bool`

HasExcludeResourceTypeIds returns a boolean if a field has been set.

### SetExcludeResourceTypeIdsNil

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) SetExcludeResourceTypeIdsNil(b bool)`

 SetExcludeResourceTypeIdsNil sets the value for ExcludeResourceTypeIds to be an explicit nil

### UnsetExcludeResourceTypeIds
`func (o *C1ApiAppV1SearchAppResourceTypesRequest) UnsetExcludeResourceTypeIds()`

UnsetExcludeResourceTypeIds ensures that no value is present for ExcludeResourceTypeIds, not even an explicit nil
### GetExcludeResourceTypeTraitIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetExcludeResourceTypeTraitIds() []string`

GetExcludeResourceTypeTraitIds returns the ExcludeResourceTypeTraitIds field if non-nil, zero value otherwise.

### GetExcludeResourceTypeTraitIdsOk

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetExcludeResourceTypeTraitIdsOk() (*[]string, bool)`

GetExcludeResourceTypeTraitIdsOk returns a tuple with the ExcludeResourceTypeTraitIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeResourceTypeTraitIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) SetExcludeResourceTypeTraitIds(v []string)`

SetExcludeResourceTypeTraitIds sets ExcludeResourceTypeTraitIds field to given value.

### HasExcludeResourceTypeTraitIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) HasExcludeResourceTypeTraitIds() bool`

HasExcludeResourceTypeTraitIds returns a boolean if a field has been set.

### SetExcludeResourceTypeTraitIdsNil

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) SetExcludeResourceTypeTraitIdsNil(b bool)`

 SetExcludeResourceTypeTraitIdsNil sets the value for ExcludeResourceTypeTraitIds to be an explicit nil

### UnsetExcludeResourceTypeTraitIds
`func (o *C1ApiAppV1SearchAppResourceTypesRequest) UnsetExcludeResourceTypeTraitIds()`

UnsetExcludeResourceTypeTraitIds ensures that no value is present for ExcludeResourceTypeTraitIds, not even an explicit nil
### GetPageSize

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### GetQuery

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetResourceTypeIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetResourceTypeIds() []string`

GetResourceTypeIds returns the ResourceTypeIds field if non-nil, zero value otherwise.

### GetResourceTypeIdsOk

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetResourceTypeIdsOk() (*[]string, bool)`

GetResourceTypeIdsOk returns a tuple with the ResourceTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypeIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) SetResourceTypeIds(v []string)`

SetResourceTypeIds sets ResourceTypeIds field to given value.

### HasResourceTypeIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) HasResourceTypeIds() bool`

HasResourceTypeIds returns a boolean if a field has been set.

### SetResourceTypeIdsNil

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) SetResourceTypeIdsNil(b bool)`

 SetResourceTypeIdsNil sets the value for ResourceTypeIds to be an explicit nil

### UnsetResourceTypeIds
`func (o *C1ApiAppV1SearchAppResourceTypesRequest) UnsetResourceTypeIds()`

UnsetResourceTypeIds ensures that no value is present for ResourceTypeIds, not even an explicit nil
### GetResourceTypeTraitIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetResourceTypeTraitIds() []string`

GetResourceTypeTraitIds returns the ResourceTypeTraitIds field if non-nil, zero value otherwise.

### GetResourceTypeTraitIdsOk

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) GetResourceTypeTraitIdsOk() (*[]string, bool)`

GetResourceTypeTraitIdsOk returns a tuple with the ResourceTypeTraitIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypeTraitIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) SetResourceTypeTraitIds(v []string)`

SetResourceTypeTraitIds sets ResourceTypeTraitIds field to given value.

### HasResourceTypeTraitIds

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) HasResourceTypeTraitIds() bool`

HasResourceTypeTraitIds returns a boolean if a field has been set.

### SetResourceTypeTraitIdsNil

`func (o *C1ApiAppV1SearchAppResourceTypesRequest) SetResourceTypeTraitIdsNil(b bool)`

 SetResourceTypeTraitIdsNil sets the value for ResourceTypeTraitIds to be an explicit nil

### UnsetResourceTypeTraitIds
`func (o *C1ApiAppV1SearchAppResourceTypesRequest) UnsetResourceTypeTraitIds()`

UnsetResourceTypeTraitIds ensures that no value is present for ResourceTypeTraitIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


