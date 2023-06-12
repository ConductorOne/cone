# C1ApiDirectoryV1DirectoryServiceCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | The appId field. | [optional] 
**ExpandMask** | Pointer to [**C1ApiDirectoryV1DirectoryExpandMask**](C1ApiDirectoryV1DirectoryExpandMask.md) |  | [optional] 

## Methods

### NewC1ApiDirectoryV1DirectoryServiceCreateRequest

`func NewC1ApiDirectoryV1DirectoryServiceCreateRequest() *C1ApiDirectoryV1DirectoryServiceCreateRequest`

NewC1ApiDirectoryV1DirectoryServiceCreateRequest instantiates a new C1ApiDirectoryV1DirectoryServiceCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiDirectoryV1DirectoryServiceCreateRequestWithDefaults

`func NewC1ApiDirectoryV1DirectoryServiceCreateRequestWithDefaults() *C1ApiDirectoryV1DirectoryServiceCreateRequest`

NewC1ApiDirectoryV1DirectoryServiceCreateRequestWithDefaults instantiates a new C1ApiDirectoryV1DirectoryServiceCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *C1ApiDirectoryV1DirectoryServiceCreateRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *C1ApiDirectoryV1DirectoryServiceCreateRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *C1ApiDirectoryV1DirectoryServiceCreateRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *C1ApiDirectoryV1DirectoryServiceCreateRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetExpandMask

`func (o *C1ApiDirectoryV1DirectoryServiceCreateRequest) GetExpandMask() C1ApiDirectoryV1DirectoryExpandMask`

GetExpandMask returns the ExpandMask field if non-nil, zero value otherwise.

### GetExpandMaskOk

`func (o *C1ApiDirectoryV1DirectoryServiceCreateRequest) GetExpandMaskOk() (*C1ApiDirectoryV1DirectoryExpandMask, bool)`

GetExpandMaskOk returns a tuple with the ExpandMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandMask

`func (o *C1ApiDirectoryV1DirectoryServiceCreateRequest) SetExpandMask(v C1ApiDirectoryV1DirectoryExpandMask)`

SetExpandMask sets ExpandMask field to given value.

### HasExpandMask

`func (o *C1ApiDirectoryV1DirectoryServiceCreateRequest) HasExpandMask() bool`

HasExpandMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


