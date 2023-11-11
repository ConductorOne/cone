package resource

import (
	"bytes"
	"os"
	"os/exec"
	"text/template"
)

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

// ExecuteTerraform takes the Terraform configuration content, writes it to a file,
// and executes Terraform commands in the given directory.
func ExecuteTerraform(tfConfig string, outputDir string) error {
	// Ensure the output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}
	exec.Command("cd " + outputDir).Run()

	if err := os.WriteFile(outputDir+"/output.tf", []byte(tfConfig), 0644); err != nil {
		return err
	}

	// Apply the Terraform configuration
	return runTerraformCommand(outputDir)
}

// runTerraformCommand runs a Terraform command with the given arguments.
func runTerraformCommand(outputDir string) error {

	cmd := exec.Command("/bin/sh", "-c", "terraform plan "+`| sed 's/\x1b\[[0-9;]*m//g'>`+outputDir+"/plan.txt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
