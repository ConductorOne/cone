package terraform

import (
	"bytes"
	"text/template"
)

const resourcePrefix = "id_"
const ImportTemplateString = `import   {{"{"}}
	to = {{.GetType}}.{{.GetResourceId}}
	id = "{{.GetId}}"
{{"}\n"}}`

type TemplateData interface {
	GetRequired() map[string]string
	GetType() string
	GetId() string
	GetResourceId() string
	GetOutputId() string
}

func ApplyTemplate(data TemplateData, tmpl string) (string, error) {
	// Prepare a buffer to hold the combined output
	var combinedOutput bytes.Buffer

	// Process the datasource template
	templateString := template.New("tmpl")

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
