// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The UpdateOrgDomainRequest message.
type UpdateOrgDomainRequest struct {
	// The newDomains field.
	NewDomains []string `json:"newDomains,omitempty"`
}

func (o *UpdateOrgDomainRequest) GetNewDomains() []string {
	if o == nil {
		return nil
	}
	return o.NewDomains
}
