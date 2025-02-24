// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// GetAppEntitlementProxyResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type GetAppEntitlementProxyResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (g GetAppEntitlementProxyResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAppEntitlementProxyResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAppEntitlementProxyResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *GetAppEntitlementProxyResponseExpanded) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The GetAppEntitlementProxyResponse message.
type GetAppEntitlementProxyResponse struct {
	// The AppEntitlementProxyView message.
	AppEntitlementProxyView *AppEntitlementProxyView `json:"appProxyEntitlementView,omitempty"`
	// The expanded field.
	Expanded []GetAppEntitlementProxyResponseExpanded `json:"expanded,omitempty"`
}

func (o *GetAppEntitlementProxyResponse) GetAppEntitlementProxyView() *AppEntitlementProxyView {
	if o == nil {
		return nil
	}
	return o.AppEntitlementProxyView
}

func (o *GetAppEntitlementProxyResponse) GetExpanded() []GetAppEntitlementProxyResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
