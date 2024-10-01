// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// CreateAppEntitlementResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type CreateAppEntitlementResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (c CreateAppEntitlementResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAppEntitlementResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAppEntitlementResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *CreateAppEntitlementResponseExpanded) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The CreateAppEntitlementResponse message.
type CreateAppEntitlementResponse struct {
	// The app entitlement view contains the serialized app entitlement and paths to objects referenced by the app entitlement.
	AppEntitlementView *AppEntitlementView `json:"appEntitlementView,omitempty"`
	// The expanded field.
	Expanded []CreateAppEntitlementResponseExpanded `json:"expanded,omitempty"`
}

func (o *CreateAppEntitlementResponse) GetAppEntitlementView() *AppEntitlementView {
	if o == nil {
		return nil
	}
	return o.AppEntitlementView
}

func (o *CreateAppEntitlementResponse) GetExpanded() []CreateAppEntitlementResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
