// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Approval - The Approval message.
//
// This message contains a oneof named typ. Only a single field of the following list may be set at a time:
//   - users
//   - manager
//   - appOwners
//   - group
//   - self
//   - entitlementOwners
type Approval struct {
	// The AppGroupApproval message.
	AppGroupApproval *AppGroupApproval `json:"group,omitempty"`
	// The AppOwnerApproval message.
	AppOwnerApproval *AppOwnerApproval `json:"appOwners,omitempty"`
	// The EntitlementOwnerApproval message.
	EntitlementOwnerApproval *EntitlementOwnerApproval `json:"entitlementOwners,omitempty"`
	// The ManagerApproval message.
	ManagerApproval *ManagerApproval `json:"manager,omitempty"`
	// The SelfApproval message.
	SelfApproval *SelfApproval `json:"self,omitempty"`
	// The UserApproval message.
	UserApproval *UserApproval `json:"users,omitempty"`
	// The allowReassignment field.
	AllowReassignment *bool `json:"allowReassignment,omitempty"`
	// The assigned field.
	Assigned *bool `json:"assigned,omitempty"`
	// The requireApprovalReason field.
	RequireApprovalReason *bool `json:"requireApprovalReason,omitempty"`
	// The requireReassignmentReason field.
	RequireReassignmentReason *bool `json:"requireReassignmentReason,omitempty"`
}
