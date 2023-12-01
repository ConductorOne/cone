package resource

import (
	"bytes"
	"text/template"
)

const datasourcePrefix = "id_"
const DataTemplateString = `data "{{.GetType}}" "{{.GetDatasourceId}}" {{"{"}}
{{- range $key, $value := .GetRequired}}
	{{$key}} = "{{$value}}"
{{- end}}
{{"}\n"}}`
const OutputTemplateString = `output  {{.GetOutputId}} {{"{"}}
	value = data.{{.GetType}}.{{.GetDatasourceId}}
{{"}\n"}}`
const ImportTemplateString = `import   {{"{"}}
	to = {{.GetType}}.{{.GetDatasourceId}}
	id = "{{.GetId}}"
{{"}\n"}}`

type TemplateData interface {
	GetRequired() map[string]string
	GetType() string
	GetId() string
	GetDatasourceId() string
	GetOutputId() string
}

func ObjectNameToTerraformType(objectName string) string {
	switch objectName {
	case "app":
		return TerraformAppType
	case "policy":
		return TerraformPolicyType
	case "app_entitlement":
		return TerraformAppEntilementType
	default:
		return ""
	}
}

func ApplyTemplate(data TemplateData, tmpl string) (string, error) {
	// Prepare a buffer to hold the combined output
	var combinedOutput bytes.Buffer

	// Create a FuncMap to register functions.
	funcMap := template.FuncMap{
		"GetType":     data.GetType,     // Pass the method itself
		"GetRequired": data.GetRequired, // Pass the method itself
		"GetId":       data.GetId,       // Pass the method itself
	}

	// Process the datasource template
	templateString := template.New("tmpl").Funcs(funcMap)

	// Parse the template file
	templateString, err := templateString.Parse(tmpl)
	if err != nil {
		return "", err
	}
	err = templateString.Execute(&combinedOutput, data)
	if err != nil {
		return "", err
	}

	// Return the combined output as a string
	return combinedOutput.String(), nil
}

func ApplyTemplates(data []TemplateData, templates ...string) (string, error) {
	res := ""
	for _, v := range data {
		for _, template := range templates {
			str, err := ApplyTemplate(v, template)
			if err != nil {
				return "", err
			}
			res += str
		}
	}
	return res, nil
}
