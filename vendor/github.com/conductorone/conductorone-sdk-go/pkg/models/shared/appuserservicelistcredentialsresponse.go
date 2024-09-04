// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// The AppUserServiceListCredentialsResponse message.
type AppUserServiceListCredentialsResponse struct {
	// The list field.
	List []AppUserCredential `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *AppUserServiceListCredentialsResponse) GetList() []AppUserCredential {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *AppUserServiceListCredentialsResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}