// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The AddManuallyManagedUsersRequest message.
type AddManuallyManagedUsersRequest struct {
	// The userIds field.
	UserIds []string `json:"userIds,omitempty"`
}

func (o *AddManuallyManagedUsersRequest) GetUserIds() []string {
	if o == nil {
		return nil
	}
	return o.UserIds
}
