// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// CreateManuallyManagedResourceTypeResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type CreateManuallyManagedResourceTypeResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (c CreateManuallyManagedResourceTypeResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateManuallyManagedResourceTypeResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateManuallyManagedResourceTypeResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *CreateManuallyManagedResourceTypeResponseExpanded) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The CreateManuallyManagedResourceTypeResponse message.
type CreateManuallyManagedResourceTypeResponse struct {
	// The AppResourceType is referenced by an app entitlement defining its resource types. Commonly things like Group or Role.
	AppResourceType *AppResourceType `json:"appResourceType,omitempty"`
	// The expanded field.
	Expanded []CreateManuallyManagedResourceTypeResponseExpanded `json:"expanded,omitempty"`
}

func (o *CreateManuallyManagedResourceTypeResponse) GetAppResourceType() *AppResourceType {
	if o == nil {
		return nil
	}
	return o.AppResourceType
}

func (o *CreateManuallyManagedResourceTypeResponse) GetExpanded() []CreateManuallyManagedResourceTypeResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}