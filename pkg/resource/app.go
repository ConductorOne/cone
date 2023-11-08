package resource

import "github.com/conductorone/conductorone-sdk-go/pkg/models/shared"

const TerraformAppType = "conductorone_app"

type AppTemplate struct {
	App shared.App
}

func (a AppTemplate) GetName() string {
	return a.GetType() + "_" + a.GetPk()
}

func (a AppTemplate) GetIds() map[string]string {
	ids := make(map[string]string)
	if a.App.ID != nil {
		ids["id"] = *a.App.ID
	}
	return ids
}

func (a AppTemplate) GetResourceType() string {
	return "app"
}

func (a AppTemplate) GetType() string {
	return TerraformAppType // Assuming the type is "App"
}

func (a AppTemplate) GetPk() string {
	return GeneratePK(a)
}
