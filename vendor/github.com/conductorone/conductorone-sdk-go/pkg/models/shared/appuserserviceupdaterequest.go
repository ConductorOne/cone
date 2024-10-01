// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The AppUserServiceUpdateRequest message contains the app user and the fields to be updated.
type AppUserServiceUpdateRequest struct {
	// Application User that represents an account in the application.
	AppUser *AppUserInput `json:"appUser,omitempty"`
	// The AppUserExpandMask message contains a list of paths to expand in the response.
	AppUserExpandMask *AppUserExpandMask `json:"expandMask,omitempty"`
	UpdateMask        *string            `json:"updateMask,omitempty"`
}

func (o *AppUserServiceUpdateRequest) GetAppUser() *AppUserInput {
	if o == nil {
		return nil
	}
	return o.AppUser
}

func (o *AppUserServiceUpdateRequest) GetAppUserExpandMask() *AppUserExpandMask {
	if o == nil {
		return nil
	}
	return o.AppUserExpandMask
}

func (o *AppUserServiceUpdateRequest) GetUpdateMask() *string {
	if o == nil {
		return nil
	}
	return o.UpdateMask
}
