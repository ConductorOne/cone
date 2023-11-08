package resource

import (
	"encoding/json"
	"os"
)

// Define your JSON structure in Go struct
type Schema struct {
	FormatVersion   string                    `json:"format_version"`
	ProviderSchemas map[string]ProviderSchema `json:"provider_schemas"`
}

type ProviderSchema struct {
	Provider        SchemaDetails             `json:"provider"`
	ResourceSchemas map[string]ResourceSchema `json:"resource_schemas"`
}

type SchemaDetails struct {
	Version int         `json:"version"`
	Block   SchemaBlock `json:"block"`
}

type ResourceSchema struct {
	Version int         `json:"version"`
	Block   SchemaBlock `json:"block"`
}

type SchemaBlock struct {
	Attributes      map[string]AttributeDetail `json:"attributes"`
	Description     string                     `json:"description"`
	DescriptionKind string                     `json:"description_kind"`
}

type AttributeDetail struct {
	Type            string `json:"type"`
	Description     string `json:"description"`
	DescriptionKind string `json:"description_kind"`
	Required        bool   `json:"required,omitempty"`
	Optional        bool   `json:"optional,omitempty"`
	Computed        bool   `json:"computed,omitempty"`
}

func GetSchema() (map[string]ResourceSchema, error) {
	// This is where your JSON would be loaded from wherever it's sourced
	fileBytes, err := os.ReadFile("providerschema.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON into the struct
	var schema Schema
	err = json.Unmarshal([]byte(fileBytes), &schema)
	if err != nil {
		return nil, err
	}

	// Now you can access the resource schemas
	resourceSchemas := schema.ProviderSchemas["registry.terraform.io/conductorone/conductorone"].ResourceSchemas

	return resourceSchemas, nil
}
