// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// PolicyPolicyType - The policyType field.
type PolicyPolicyType string

const (
	PolicyPolicyTypePolicyTypeUnspecified   PolicyPolicyType = "POLICY_TYPE_UNSPECIFIED"
	PolicyPolicyTypePolicyTypeGrant         PolicyPolicyType = "POLICY_TYPE_GRANT"
	PolicyPolicyTypePolicyTypeRevoke        PolicyPolicyType = "POLICY_TYPE_REVOKE"
	PolicyPolicyTypePolicyTypeCertify       PolicyPolicyType = "POLICY_TYPE_CERTIFY"
	PolicyPolicyTypePolicyTypeAccessRequest PolicyPolicyType = "POLICY_TYPE_ACCESS_REQUEST"
	PolicyPolicyTypePolicyTypeProvision     PolicyPolicyType = "POLICY_TYPE_PROVISION"
)

func (e PolicyPolicyType) ToPointer() *PolicyPolicyType {
	return &e
}

func (e *PolicyPolicyType) UnmarshalJSON(data []byte) error {
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
		*e = PolicyPolicyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PolicyPolicyType: %v", v)
	}
}

// Policy - The Policy message.
type Policy struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	//  Properties
	//
	ID *string `json:"id,omitempty"`
	// The policySteps field.
	PolicySteps map[string]PolicySteps `json:"policySteps,omitempty"`
	// The policyType field.
	PolicyType *PolicyPolicyType `json:"policyType,omitempty"`
	// The postActions field.
	PostActions []PolicyPostActions `json:"postActions,omitempty"`
	// The reassignTasksToDelegates field.
	ReassignTasksToDelegates *bool `json:"reassignTasksToDelegates,omitempty"`
	// The systemBuiltin field.
	SystemBuiltin *bool      `json:"systemBuiltin,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
}
