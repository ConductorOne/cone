# C1ApiPolicyV1PolicyPostActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertifyRemediateImmediately** | Pointer to **NullableBool** |  ONLY valid when used in a CERTIFY Ticket Type:  Causes any deprovision or change in a grant to be applied when Certify Ticket is closed.  This field is part of the &#x60;action&#x60; oneof. See the documentation for &#x60;c1.api.policy.v1.PolicyPostActions&#x60; for more details. | [optional] 

## Methods

### NewC1ApiPolicyV1PolicyPostActions

`func NewC1ApiPolicyV1PolicyPostActions() *C1ApiPolicyV1PolicyPostActions`

NewC1ApiPolicyV1PolicyPostActions instantiates a new C1ApiPolicyV1PolicyPostActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC1ApiPolicyV1PolicyPostActionsWithDefaults

`func NewC1ApiPolicyV1PolicyPostActionsWithDefaults() *C1ApiPolicyV1PolicyPostActions`

NewC1ApiPolicyV1PolicyPostActionsWithDefaults instantiates a new C1ApiPolicyV1PolicyPostActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertifyRemediateImmediately

`func (o *C1ApiPolicyV1PolicyPostActions) GetCertifyRemediateImmediately() bool`

GetCertifyRemediateImmediately returns the CertifyRemediateImmediately field if non-nil, zero value otherwise.

### GetCertifyRemediateImmediatelyOk

`func (o *C1ApiPolicyV1PolicyPostActions) GetCertifyRemediateImmediatelyOk() (*bool, bool)`

GetCertifyRemediateImmediatelyOk returns a tuple with the CertifyRemediateImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifyRemediateImmediately

`func (o *C1ApiPolicyV1PolicyPostActions) SetCertifyRemediateImmediately(v bool)`

SetCertifyRemediateImmediately sets CertifyRemediateImmediately field to given value.

### HasCertifyRemediateImmediately

`func (o *C1ApiPolicyV1PolicyPostActions) HasCertifyRemediateImmediately() bool`

HasCertifyRemediateImmediately returns a boolean if a field has been set.

### SetCertifyRemediateImmediatelyNil

`func (o *C1ApiPolicyV1PolicyPostActions) SetCertifyRemediateImmediatelyNil(b bool)`

 SetCertifyRemediateImmediatelyNil sets the value for CertifyRemediateImmediately to be an explicit nil

### UnsetCertifyRemediateImmediately
`func (o *C1ApiPolicyV1PolicyPostActions) UnsetCertifyRemediateImmediately()`

UnsetCertifyRemediateImmediately ensures that no value is present for CertifyRemediateImmediately, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


