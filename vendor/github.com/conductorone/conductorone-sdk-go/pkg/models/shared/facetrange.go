// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// The FacetRange message.
type FacetRange struct {
	// The count of items in the range.
	Count *string `json:"count,omitempty"`
	// The display name of the range.
	DisplayName *string `json:"displayName,omitempty"`
	// The starting value of the range.
	From *string `json:"from,omitempty"`
	// The icon of the range.
	IconURL *string `json:"iconUrl,omitempty"`
	// The ending value of the range.
	To *string `json:"to,omitempty"`
}

func (o *FacetRange) GetCount() *string {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *FacetRange) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *FacetRange) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *FacetRange) GetIconURL() *string {
	if o == nil {
		return nil
	}
	return o.IconURL
}

func (o *FacetRange) GetTo() *string {
	if o == nil {
		return nil
	}
	return o.To
}
