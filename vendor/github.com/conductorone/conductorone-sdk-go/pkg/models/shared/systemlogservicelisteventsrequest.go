// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// SortDirection - The sortDirection field.
type SortDirection string

const (
	SortDirectionSortDirectionUnspecified SortDirection = "SORT_DIRECTION_UNSPECIFIED"
	SortDirectionSortDirectionAsc         SortDirection = "SORT_DIRECTION_ASC"
	SortDirectionSortDirectionDesc        SortDirection = "SORT_DIRECTION_DESC"
)

func (e SortDirection) ToPointer() *SortDirection {
	return &e
}

// The SystemLogServiceListEventsRequest message.
type SystemLogServiceListEventsRequest struct {
	// The pageSize field.
	PageSize *int `json:"pageSize,omitempty"`
	// The pageToken field.
	PageToken *string    `json:"pageToken,omitempty"`
	Since     *time.Time `json:"since,omitempty"`
	// The sinceEventUid field.
	SinceEventUID *string `json:"sinceEventUid,omitempty"`
	// The sortDirection field.
	SortDirection *SortDirection `json:"sortDirection,omitempty"`
	Until         *time.Time     `json:"until,omitempty"`
}

func (s SystemLogServiceListEventsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SystemLogServiceListEventsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SystemLogServiceListEventsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *SystemLogServiceListEventsRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *SystemLogServiceListEventsRequest) GetSince() *time.Time {
	if o == nil {
		return nil
	}
	return o.Since
}

func (o *SystemLogServiceListEventsRequest) GetSinceEventUID() *string {
	if o == nil {
		return nil
	}
	return o.SinceEventUID
}

func (o *SystemLogServiceListEventsRequest) GetSortDirection() *SortDirection {
	if o == nil {
		return nil
	}
	return o.SortDirection
}

func (o *SystemLogServiceListEventsRequest) GetUntil() *time.Time {
	if o == nil {
		return nil
	}
	return o.Until
}
