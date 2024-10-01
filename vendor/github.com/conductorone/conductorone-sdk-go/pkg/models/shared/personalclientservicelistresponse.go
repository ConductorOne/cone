// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The PersonalClientServiceListResponse message.
type PersonalClientServiceListResponse struct {
	// The list field.
	List []PersonalClient `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *PersonalClientServiceListResponse) GetList() []PersonalClient {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *PersonalClientServiceListResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}