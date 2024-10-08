// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The PersonalClientSearchServiceSearchResponse message.
type PersonalClientSearchServiceSearchResponse struct {
	// The list field.
	List []PersonalClient `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *PersonalClientSearchServiceSearchResponse) GetList() []PersonalClient {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *PersonalClientSearchServiceSearchResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
