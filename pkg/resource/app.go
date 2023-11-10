package resource

import "github.com/conductorone/conductorone-sdk-go/pkg/models/shared"

const TerraformAppType = "conductorone_app"

type AppTemplate struct {
	App shared.App
}

func (a AppTemplate) GetRequired() map[string]string {
	ids := make(map[string]string)
	if a.App.DisplayName != nil {
		ids["display_name"] = *a.App.DisplayName
	}
	return ids
}

func (a AppTemplate) GetType() string {
	return TerraformAppType // Assuming the type is "App"
}

func (a AppTemplate) GetId() string {
	return *a.App.ID
}

func (a AppTemplate) GetDatasourceId() string {
	return "id_" + *a.App.ID
}

func (a AppTemplate) GetOutputId() string {
	return a.GetType() + "_" + a.GetDatasourceId()
}
