// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UpdateRolesResponse is the response message containing the updated role.
type UpdateRolesResponse struct {
	// Role is a role that can be assigned to a user in ConductorOne.
	Role *Role `json:"role,omitempty"`
}

func (o *UpdateRolesResponse) GetRole() *Role {
	if o == nil {
		return nil
	}
	return o.Role
}
