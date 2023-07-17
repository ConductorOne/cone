// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// AppResource - The AppResource message.
type AppResource struct {
	// The appId field.
	AppID *string `json:"appId,omitempty"`
	// The appResourceTypeId field.
	AppResourceTypeID *string    `json:"appResourceTypeId,omitempty"`
	CreatedAt         *time.Time `json:"createdAt,omitempty"`
	// The customDescription field.
	CustomDescription *string    `json:"customDescription,omitempty"`
	DeletedAt         *time.Time `json:"deletedAt,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The grantCount field.
	GrantCount *string `json:"grantCount,omitempty"`
	// The id field.
	ID        *string    `json:"id,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

func (o *AppResource) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppResource) GetAppResourceTypeID() *string {
	if o == nil {
		return nil
	}
	return o.AppResourceTypeID
}

func (o *AppResource) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AppResource) GetCustomDescription() *string {
	if o == nil {
		return nil
	}
	return o.CustomDescription
}

func (o *AppResource) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *AppResource) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AppResource) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AppResource) GetGrantCount() *string {
	if o == nil {
		return nil
	}
	return o.GrantCount
}

func (o *AppResource) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AppResource) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
