// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// UserServiceListResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type UserServiceListResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (u UserServiceListResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UserServiceListResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UserServiceListResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *UserServiceListResponseExpanded) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The UserServiceListResponse message contains a list of results and a nextPageToken if applicable.
type UserServiceListResponse struct {
	// List of serialized related objects.
	Expanded []UserServiceListResponseExpanded `json:"expanded,omitempty"`
	// The list of results containing up to X results, where X is the page size defined in the request
	List []UserView `json:"list,omitempty"`
	// The nextPageToken is shown for the next page if the number of results is larger than the max page size.
	//  The server returns one page of results and the nextPageToken until all results are retreived.
	//  To retrieve the next page, use the same request and append a pageToken field with the value of nextPageToken shown on the previous page.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *UserServiceListResponse) GetExpanded() []UserServiceListResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *UserServiceListResponse) GetList() []UserView {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *UserServiceListResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
