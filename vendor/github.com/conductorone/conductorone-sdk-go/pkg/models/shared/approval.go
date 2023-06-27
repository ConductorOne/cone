// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ApprovalTyp struct {
	// The AppOwnerApproval message.
	AppOwners *AppOwnerApproval `json:"appOwners,omitempty"`
	// The EntitlementOwnerApproval message.
	EntitlementOwners *EntitlementOwnerApproval `json:"entitlementOwners,omitempty"`
	// The AppGroupApproval message.
	Group *AppGroupApproval `json:"group,omitempty"`
	// The ManagerApproval message.
	Manager *ManagerApproval `json:"manager,omitempty"`
	// The SelfApproval message.
	Self *SelfApproval `json:"self,omitempty"`
	// The UserApproval message.
	Users *UserApproval `json:"users,omitempty"`
}

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
	// The allowReassignment field.
	AllowReassignment *bool `json:"allowReassignment,omitempty"`
	// The assigned field.
	Assigned *bool `json:"assigned,omitempty"`
	// The requireApprovalReason field.
	RequireApprovalReason *bool `json:"requireApprovalReason,omitempty"`
	// The requireReassignmentReason field.
	RequireReassignmentReason *bool        `json:"requireReassignmentReason,omitempty"`
	Typ                       *ApprovalTyp `json:"typ,omitempty"`
}
