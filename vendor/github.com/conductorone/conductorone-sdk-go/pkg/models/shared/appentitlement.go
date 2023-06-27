// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type AppEntitlementMaxGrantDurationDurationUnset struct {
}

type AppEntitlementMaxGrantDuration struct {
	DurationGrant *string                                      `json:"durationGrant,omitempty"`
	DurationUnset *AppEntitlementMaxGrantDurationDurationUnset `json:"durationUnset,omitempty"`
}

// AppEntitlement - The AppEntitlement message.
//
// This message contains a oneof named max_grant_duration. Only a single field of the following list may be set at a time:
//   - durationUnset
//   - durationGrant
type AppEntitlement struct {
	// The alias field.
	Alias *string `json:"alias,omitempty"`
	// The appId field.
	AppID *string `json:"appId,omitempty"`
	// The appResourceId field.
	AppResourceID *string `json:"appResourceId,omitempty"`
	// The appResourceTypeId field.
	AppResourceTypeID *string `json:"appResourceTypeId,omitempty"`
	// The certifyPolicyId field.
	CertifyPolicyID *string `json:"certifyPolicyId,omitempty"`
	// The complianceFrameworkValueIds field.
	ComplianceFrameworkValueIds []string   `json:"complianceFrameworkValueIds,omitempty"`
	CreatedAt                   *time.Time `json:"createdAt,omitempty"`
	DeletedAt                   *time.Time `json:"deletedAt,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The emergencyGrantEnabled field.
	EmergencyGrantEnabled *bool `json:"emergencyGrantEnabled,omitempty"`
	// The emergencyGrantPolicyId field.
	EmergencyGrantPolicyID *string `json:"emergencyGrantPolicyId,omitempty"`
	// The grantCount field.
	GrantCount *string `json:"grantCount,omitempty"`
	// The grantPolicyId field.
	GrantPolicyID *string `json:"grantPolicyId,omitempty"`
	// The id field.
	ID               *string                         `json:"id,omitempty"`
	MaxGrantDuration *AppEntitlementMaxGrantDuration `json:"max_grant_duration,omitempty"`
	// The ProvisionPolicy message.
	//
	// This message contains a oneof named typ. Only a single field of the following list may be set at a time:
	//   - connector
	//   - manual
	//   - delegated
	//
	ProvisionerPolicy *ProvisionPolicy `json:"provisionerPolicy,omitempty"`
	// The revokePolicyId field.
	RevokePolicyID *string `json:"revokePolicyId,omitempty"`
	// The riskLevelValueId field.
	RiskLevelValueID *string `json:"riskLevelValueId,omitempty"`
	// The slug field.
	Slug *string `json:"slug,omitempty"`
	// The systemBuiltin field.
	SystemBuiltin *bool      `json:"systemBuiltin,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
}
