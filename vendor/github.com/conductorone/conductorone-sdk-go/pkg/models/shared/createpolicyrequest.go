// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreatePolicyRequestPolicyType - The policyType field.
type CreatePolicyRequestPolicyType string

const (
	CreatePolicyRequestPolicyTypePolicyTypeUnspecified   CreatePolicyRequestPolicyType = "POLICY_TYPE_UNSPECIFIED"
	CreatePolicyRequestPolicyTypePolicyTypeGrant         CreatePolicyRequestPolicyType = "POLICY_TYPE_GRANT"
	CreatePolicyRequestPolicyTypePolicyTypeRevoke        CreatePolicyRequestPolicyType = "POLICY_TYPE_REVOKE"
	CreatePolicyRequestPolicyTypePolicyTypeCertify       CreatePolicyRequestPolicyType = "POLICY_TYPE_CERTIFY"
	CreatePolicyRequestPolicyTypePolicyTypeAccessRequest CreatePolicyRequestPolicyType = "POLICY_TYPE_ACCESS_REQUEST"
	CreatePolicyRequestPolicyTypePolicyTypeProvision     CreatePolicyRequestPolicyType = "POLICY_TYPE_PROVISION"
)

func (e CreatePolicyRequestPolicyType) ToPointer() *CreatePolicyRequestPolicyType {
	return &e
}

func (e *CreatePolicyRequestPolicyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "POLICY_TYPE_UNSPECIFIED":
		fallthrough
	case "POLICY_TYPE_GRANT":
		fallthrough
	case "POLICY_TYPE_REVOKE":
		fallthrough
	case "POLICY_TYPE_CERTIFY":
		fallthrough
	case "POLICY_TYPE_ACCESS_REQUEST":
		fallthrough
	case "POLICY_TYPE_PROVISION":
		*e = CreatePolicyRequestPolicyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePolicyRequestPolicyType: %v", v)
	}
}

// CreatePolicyRequest - The CreatePolicyRequest message.
type CreatePolicyRequest struct {
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The policySteps field.
	PolicySteps map[string]PolicySteps `json:"policySteps,omitempty"`
	// The policyType field.
	PolicyType *CreatePolicyRequestPolicyType `json:"policyType,omitempty"`
	// The postActions field.
	PostActions []PolicyPostActions `json:"postActions,omitempty"`
	// The reassignTasksToDelegates field.
	ReassignTasksToDelegates *bool `json:"reassignTasksToDelegates,omitempty"`
}
