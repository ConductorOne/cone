package terraform

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

const TerraformCatalogType = "conductorone_catalog"

type CatalogTemplate struct {
	RequestCatalog shared.RequestCatalog
}

func (r CatalogTemplate) GetRequired() map[string]string {
	ids := make(map[string]string)
	if r.RequestCatalog.ID != nil {
		ids["id"] = *r.RequestCatalog.ID
	}
	return ids
}

func (r CatalogTemplate) GetType() string {
	return TerraformCatalogType
}

func (r CatalogTemplate) GetId() string {
	ids := r.GetRequired()
	return ids["id"]
}

func (r CatalogTemplate) GetResourceId() string {
	return resourcePrefix + r.GetId()
}

func (r CatalogTemplate) GetOutputId() string {
	return r.GetType() + "_" + r.GetResourceId()
}
