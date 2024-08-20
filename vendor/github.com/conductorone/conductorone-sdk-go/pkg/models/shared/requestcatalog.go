// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// The RequestCatalog is used for managing which entitlements are requestable, and who can request them.
type RequestCatalog struct {
	// An array of app entitlements that, if the user has, can view the contents of this catalog.
	AccessEntitlements []AppEntitlement `json:"accessEntitlements,omitempty"`
	// The Apps contained in this request catalog.
	AppIds    []string   `json:"appIds,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The id of the user this request catalog was created by.
	CreatedByUserID *string    `json:"createdByUserId,omitempty"`
	DeletedAt       *time.Time `json:"deletedAt,omitempty"`
	// The description of the request catalog.
	Description *string `json:"description,omitempty"`
	// The display name of the request catalog.
	DisplayName *string `json:"displayName,omitempty"`
	// The id of the request catalog.
	ID *string `json:"id,omitempty"`
	// Whether or not this catalog is published.
	Published *bool `json:"published,omitempty"`
	// Whether all the entitlements in the catalog can be requests at once. Your tenant must have the bundles feature to use this.
	RequestBundle *bool      `json:"requestBundle,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	// If this is true, the access entitlement requirement is ignored.
	VisibleToEveryone *bool `json:"visibleToEveryone,omitempty"`
}

func (r RequestCatalog) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestCatalog) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestCatalog) GetAccessEntitlements() []AppEntitlement {
	if o == nil {
		return nil
	}
	return o.AccessEntitlements
}

func (o *RequestCatalog) GetAppIds() []string {
	if o == nil {
		return nil
	}
	return o.AppIds
}

func (o *RequestCatalog) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RequestCatalog) GetCreatedByUserID() *string {
	if o == nil {
		return nil
	}
	return o.CreatedByUserID
}

func (o *RequestCatalog) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *RequestCatalog) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *RequestCatalog) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *RequestCatalog) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RequestCatalog) GetPublished() *bool {
	if o == nil {
		return nil
	}
	return o.Published
}

func (o *RequestCatalog) GetRequestBundle() *bool {
	if o == nil {
		return nil
	}
	return o.RequestBundle
}

func (o *RequestCatalog) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RequestCatalog) GetVisibleToEveryone() *bool {
	if o == nil {
		return nil
	}
	return o.VisibleToEveryone
}
