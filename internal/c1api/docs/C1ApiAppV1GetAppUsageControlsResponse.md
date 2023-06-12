# C1ApiAppV1GetAppUsageControlsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUsageControls** | Pointer to [**C1ApiAppV1AppUsageControls**](C1ApiAppV1AppUsageControls.md) |  | [optional] 
**HasUsageData** | Pointer to **bool** | The hasUsageData field. | [optional] 

## Methods

### NewC1ApiAppV1GetAppUsageControlsResponse

`func NewC1ApiAppV1GetAppUsageControlsResponse() *C1ApiAppV1GetAppUsageControlsResponse`

NewC1ApiAppV1GetAppUsageControlsResponse instantiates a new C1ApiAppV1GetAppUsageControlsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiAppV1GetAppUsageControlsResponseWithDefaults

`func NewC1ApiAppV1GetAppUsageControlsResponseWithDefaults() *C1ApiAppV1GetAppUsageControlsResponse`

NewC1ApiAppV1GetAppUsageControlsResponseWithDefaults instantiates a new C1ApiAppV1GetAppUsageControlsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUsageControls

`func (o *C1ApiAppV1GetAppUsageControlsResponse) GetAppUsageControls() C1ApiAppV1AppUsageControls`

GetAppUsageControls returns the AppUsageControls field if non-nil, zero value otherwise.

### GetAppUsageControlsOk

`func (o *C1ApiAppV1GetAppUsageControlsResponse) GetAppUsageControlsOk() (*C1ApiAppV1AppUsageControls, bool)`

GetAppUsageControlsOk returns a tuple with the AppUsageControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUsageControls

`func (o *C1ApiAppV1GetAppUsageControlsResponse) SetAppUsageControls(v C1ApiAppV1AppUsageControls)`

SetAppUsageControls sets AppUsageControls field to given value.

### HasAppUsageControls

`func (o *C1ApiAppV1GetAppUsageControlsResponse) HasAppUsageControls() bool`

HasAppUsageControls returns a boolean if a field has been set.

### GetHasUsageData

`func (o *C1ApiAppV1GetAppUsageControlsResponse) GetHasUsageData() bool`

GetHasUsageData returns the HasUsageData field if non-nil, zero value otherwise.

### GetHasUsageDataOk

`func (o *C1ApiAppV1GetAppUsageControlsResponse) GetHasUsageDataOk() (*bool, bool)`

GetHasUsageDataOk returns a tuple with the HasUsageData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUsageData

`func (o *C1ApiAppV1GetAppUsageControlsResponse) SetHasUsageData(v bool)`

SetHasUsageData sets HasUsageData field to given value.

### HasHasUsageData

`func (o *C1ApiAppV1GetAppUsageControlsResponse) HasHasUsageData() bool`

HasHasUsageData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


