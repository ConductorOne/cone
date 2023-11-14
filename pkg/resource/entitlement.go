package resource

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

const TerraformAppEntilementType = "conductorone_app_entitlement"

type AppEntitlementTemplate struct {
	AppEntitlement shared.AppEntitlement
}

func (ae AppEntitlementTemplate) GetRequired() map[string]string {
	ids := make(map[string]string)
	// Should probably be an error if any are nil
	if ae.AppEntitlement.AppID != nil {
		ids["app_id"] = *ae.AppEntitlement.AppID
	}
	if ae.AppEntitlement.AppResourceID != nil {
		ids["id"] = *ae.AppEntitlement.ID
	}
	return ids
}

func (ae AppEntitlementTemplate) GetType() string {
	return TerraformAppEntilementType
}

func (ae AppEntitlementTemplate) GetId() string {
	ids := ae.GetRequired()
	return ids["id"] + "." + ids["app_id"]
}

func (ae AppEntitlementTemplate) GetDatasourceId() string {
	return datasourcePrefix + ae.GetId()
}

func (ae AppEntitlementTemplate) GetOutputId() string {
	return ae.GetType() + "_" + ae.GetDatasourceId()
}
