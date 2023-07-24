// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectorView - The ConnectorView message.
type ConnectorView struct {
	// The Connector message.
	Connector *Connector `json:"connector,omitempty"`
	// The appPath field.
	AppPath *string `json:"appPath,omitempty"`
	// The usersPath field.
	UsersPath *string `json:"usersPath,omitempty"`
}

func (o *ConnectorView) GetConnector() *Connector {
	if o == nil {
		return nil
	}
	return o.Connector
}

func (o *ConnectorView) GetAppPath() *string {
	if o == nil {
		return nil
	}
	return o.AppPath
}

func (o *ConnectorView) GetUsersPath() *string {
	if o == nil {
		return nil
	}
	return o.UsersPath
}