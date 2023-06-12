# C1ApiAppV1AppReportServiceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]C1ApiAppV1AppPopulationReport**](C1ApiAppV1AppPopulationReport.md) | The list field. | [optional] 
**NextPageToken** | Pointer to **string** | The nextPageToken field. | [optional] 

## Methods

### NewC1ApiAppV1AppReportServiceListResponse

`func NewC1ApiAppV1AppReportServiceListResponse() *C1ApiAppV1AppReportServiceListResponse`

NewC1ApiAppV1AppReportServiceListResponse instantiates a new C1ApiAppV1AppReportServiceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiAppV1AppReportServiceListResponseWithDefaults

`func NewC1ApiAppV1AppReportServiceListResponseWithDefaults() *C1ApiAppV1AppReportServiceListResponse`

NewC1ApiAppV1AppReportServiceListResponseWithDefaults instantiates a new C1ApiAppV1AppReportServiceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *C1ApiAppV1AppReportServiceListResponse) GetList() []C1ApiAppV1AppPopulationReport`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *C1ApiAppV1AppReportServiceListResponse) GetListOk() (*[]C1ApiAppV1AppPopulationReport, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *C1ApiAppV1AppReportServiceListResponse) SetList(v []C1ApiAppV1AppPopulationReport)`

SetList sets List field to given value.

### HasList

`func (o *C1ApiAppV1AppReportServiceListResponse) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *C1ApiAppV1AppReportServiceListResponse) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *C1ApiAppV1AppReportServiceListResponse) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetNextPageToken

`func (o *C1ApiAppV1AppReportServiceListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *C1ApiAppV1AppReportServiceListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *C1ApiAppV1AppReportServiceListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *C1ApiAppV1AppReportServiceListResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


