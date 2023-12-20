package terraform

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

const TerraformPolicyType = "conductorone_policy"

type PolicyTemplate struct {
	Policy shared.Policy
}

func (p PolicyTemplate) GetRequired() map[string]string {
	ids := make(map[string]string)
	if p.Policy.DisplayName != nil {
		ids["display_name"] = *p.Policy.DisplayName
	}
	return ids
}

func (p PolicyTemplate) GetType() string {
	return TerraformPolicyType // Assuming the type is "App"
}

func (p PolicyTemplate) GetId() string {
	return *p.Policy.ID
}

func (p PolicyTemplate) GetResourceId() string {
	return resourcePrefix + *p.Policy.ID
}

func (p PolicyTemplate) GetOutputId() string {
	return p.GetType() + "_" + p.GetResourceId()
}
