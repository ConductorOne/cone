// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The AppUserExpandMask message contains a list of paths to expand in the response.
type AppUserExpandMask struct {
	// The paths to expand in the response. May be any combination of "*", "identity_user_id", "app_id", and "last_usage".
	Paths []string `json:"paths,omitempty"`
}

func (o *AppUserExpandMask) GetPaths() []string {
	if o == nil {
		return nil
	}
	return o.Paths
}
