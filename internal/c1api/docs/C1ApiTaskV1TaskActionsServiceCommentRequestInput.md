# C1ApiTaskV1TaskActionsServiceCommentRequestInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The comment field. | [optional] 
**ExpandMask** | Pointer to [**C1ApiTaskV1TaskExpandMask**](C1ApiTaskV1TaskExpandMask.md) |  | [optional] 

## Methods

### NewC1ApiTaskV1TaskActionsServiceCommentRequestInput

`func NewC1ApiTaskV1TaskActionsServiceCommentRequestInput() *C1ApiTaskV1TaskActionsServiceCommentRequestInput`

NewC1ApiTaskV1TaskActionsServiceCommentRequestInput instantiates a new C1ApiTaskV1TaskActionsServiceCommentRequestInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiTaskV1TaskActionsServiceCommentRequestInputWithDefaults

`func NewC1ApiTaskV1TaskActionsServiceCommentRequestInputWithDefaults() *C1ApiTaskV1TaskActionsServiceCommentRequestInput`

NewC1ApiTaskV1TaskActionsServiceCommentRequestInputWithDefaults instantiates a new C1ApiTaskV1TaskActionsServiceCommentRequestInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *C1ApiTaskV1TaskActionsServiceCommentRequestInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *C1ApiTaskV1TaskActionsServiceCommentRequestInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *C1ApiTaskV1TaskActionsServiceCommentRequestInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *C1ApiTaskV1TaskActionsServiceCommentRequestInput) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpandMask

`func (o *C1ApiTaskV1TaskActionsServiceCommentRequestInput) GetExpandMask() C1ApiTaskV1TaskExpandMask`

GetExpandMask returns the ExpandMask field if non-nil, zero value otherwise.

### GetExpandMaskOk

`func (o *C1ApiTaskV1TaskActionsServiceCommentRequestInput) GetExpandMaskOk() (*C1ApiTaskV1TaskExpandMask, bool)`

GetExpandMaskOk returns a tuple with the ExpandMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandMask

`func (o *C1ApiTaskV1TaskActionsServiceCommentRequestInput) SetExpandMask(v C1ApiTaskV1TaskExpandMask)`

SetExpandMask sets ExpandMask field to given value.

### HasExpandMask

`func (o *C1ApiTaskV1TaskActionsServiceCommentRequestInput) HasExpandMask() bool`

HasExpandMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

