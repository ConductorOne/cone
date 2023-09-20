// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// The UserServiceGetResponse returns a user view which has a user including JSONPATHs to the expanded items in the expanded array.
type UserServiceGetResponse struct {
	// The UserView object provides a user response object, as well as JSONPATHs to related objects provided by expanders.
	UserView *UserView `json:"userView,omitempty"`
	// List of serialized related objects.
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
}

func (o *UserServiceGetResponse) GetUserView() *UserView {
	if o == nil {
		return nil
	}
	return o.UserView
}

func (o *UserServiceGetResponse) GetExpanded() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Expanded
}
