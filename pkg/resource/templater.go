package resource

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

const datasourceTemplateString = `data "{{.GetType}}" "{{.GetPk}}" {{"{"}}
{{- range $key, $value := .GetIds}}
	{{$key}} = "{{$value}}"
{{- end}}
{{"}\n"}}`
const outputTemplateString = `output {{.GetName}} {"{"}}
	value = data.{{GetType}}.{{GetPk}}
{{"}\n"}}`

type templateData interface {
	GetIds() map[string]string
	GetType() string
	GetResourceType() string
	GetPk() string
	GetName() string
}

func GeneratePK(data templateData) string {
	ids := data.GetIds()
	var pkParts []string
	for key, value := range ids {
		pkParts = append(pkParts, key+"_"+value)
	}
	// Sort pkParts to ensure the order is consistent
	sort.Strings(pkParts)
	return strings.Join(pkParts, "_")
}

func ApplyTemplate(data templateData) (string, error) {
	// Prepare a buffer to hold the combined output
	var combinedOutput bytes.Buffer

	// Create a FuncMap to register functions.
	funcMap := template.FuncMap{
		"GetType": data.GetType, // Pass the method itself
		"GetPk":   data.GetPk,   // Pass the method itself
		"GetIds":  data.GetIds,  // Pass the method itself
	}

	// Process the datasource template
	datasourceTemplate := template.New("datasource").Funcs(funcMap)

	// Parse the template file
	datasourceTemplate, err := datasourceTemplate.Parse(datasourceTemplateString)
	if err != nil {
		return "", err
	}
	err = datasourceTemplate.Execute(&combinedOutput, data)
	if err != nil {
		return "", err
	}

	outputTemplate := template.New("output").Funcs(funcMap)
	outputTemplate, err = outputTemplate.Parse(outputTemplateString)
	if err != nil {
		return "", err
	}
	err = outputTemplate.Execute(&combinedOutput, data)
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

	// Write the Terraform configuration to a file
	tfFilePath := filepath.Join("main.tf")
	if err := os.WriteFile(tfFilePath, []byte(tfConfig), 0644); err != nil {
		return err
	}

	// Apply the Terraform configuration
	return runTerraformCommand()
}

// runTerraformCommand runs a Terraform command with the given arguments.
func runTerraformCommand() error {

	cmd := exec.Command("/bin/sh", "-c", "terraform init; terraform plan")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
