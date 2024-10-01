// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// AppResourceTypeInput - The AppResourceType is referenced by an app entitlement defining its resource types. Commonly things like Group or Role.
type AppResourceTypeInput struct {
	// The display name of the app resource type.
	DisplayName *string `json:"displayName,omitempty"`
}

func (o *AppResourceTypeInput) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}
