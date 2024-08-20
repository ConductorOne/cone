// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// AppResourceServiceListResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type AppResourceServiceListResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (a AppResourceServiceListResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppResourceServiceListResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppResourceServiceListResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *AppResourceServiceListResponseExpanded) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The AppResourceServiceListResponse message contains a list of results and a nextPageToken if applicable.
type AppResourceServiceListResponse struct {
	// List of serialized related objects.
	Expanded []AppResourceServiceListResponseExpanded `json:"expanded,omitempty"`
	// The list of results containing up to X results, where X is the page size defined in the request.
	List []AppResourceView `json:"list,omitempty"`
	// The nextPageToken is shown for the next page if the number of results is larger than the max page size.
	//  The server returns one page of results and the nextPageToken until all results are retreived.
	//  To retrieve the next page, use the same request and append a pageToken field with the value of nextPageToken shown on the previous page.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *AppResourceServiceListResponse) GetExpanded() []AppResourceServiceListResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *AppResourceServiceListResponse) GetList() []AppResourceView {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *AppResourceServiceListResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
