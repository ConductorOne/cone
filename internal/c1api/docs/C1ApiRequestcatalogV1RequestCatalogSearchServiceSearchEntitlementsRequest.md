# C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementAlias** | Pointer to **string** | The entitlementAlias field. | [optional] 
**ExpandMask** | Pointer to [**C1ApiAppV1AppEntitlementExpandMask**](C1ApiAppV1AppEntitlementExpandMask.md) |  | [optional] 
**PageSize** | Pointer to **float32** | The pageSize field. | [optional] 
**PageToken** | Pointer to **string** | The pageToken field. | [optional] 
**Query** | Pointer to **string** | The query field. | [optional] 

## Methods

### NewC1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest

`func NewC1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest() *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest`

NewC1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest instantiates a new C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequestWithDefaults

`func NewC1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequestWithDefaults() *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest`

NewC1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequestWithDefaults instantiates a new C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementAlias

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) GetEntitlementAlias() string`

GetEntitlementAlias returns the EntitlementAlias field if non-nil, zero value otherwise.

### GetEntitlementAliasOk

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) GetEntitlementAliasOk() (*string, bool)`

GetEntitlementAliasOk returns a tuple with the EntitlementAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementAlias

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) SetEntitlementAlias(v string)`

SetEntitlementAlias sets EntitlementAlias field to given value.

### HasEntitlementAlias

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) HasEntitlementAlias() bool`

HasEntitlementAlias returns a boolean if a field has been set.

### GetExpandMask

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) GetExpandMask() C1ApiAppV1AppEntitlementExpandMask`

GetExpandMask returns the ExpandMask field if non-nil, zero value otherwise.

### GetExpandMaskOk

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) GetExpandMaskOk() (*C1ApiAppV1AppEntitlementExpandMask, bool)`

GetExpandMaskOk returns a tuple with the ExpandMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandMask

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) SetExpandMask(v C1ApiAppV1AppEntitlementExpandMask)`

SetExpandMask sets ExpandMask field to given value.

### HasExpandMask

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) HasExpandMask() bool`

HasExpandMask returns a boolean if a field has been set.

### GetPageSize

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### GetQuery

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *C1ApiRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


