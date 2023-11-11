package main

import (
	"fmt"
	"os"

	"github.com/conductorone/cone/pkg/resource"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

const terraformDir = "terraform"

var objects = []string{"app", "policy"}

func tfCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tf <object-name>",
		Short: "Import unmanaged terraform resources for the specified object",
		RunE:  tfRun,
	}

	return cmd
}

func writeToFile(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func tfRun(cmd *cobra.Command, args []string) error {
	ctx, c, _, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}

	object := args[0]
	if !slices.Contains[string](objects, object) && object != "*" {
		return fmt.Errorf("invalid object name, only support %v and * for all", objects)
	}

	resources := make(map[string]resource.TemplateData)
	if object == "app" || object == "*" {
		apps, err := c.ListApps(ctx)
		if err != nil {
			return err
		}
		for _, app := range apps {
			tmplData := resource.AppTemplate{App: app}
			resources[tmplData.GetOutputId()] = tmplData
		}
	}
	if object == "policy" || object == "*" {
		policies, err := c.ListPolicies(ctx)
		if err != nil {
			return err
		}
		for _, policy := range policies {
			tmplData := resource.PolicyTemplate{Policy: policy}
			resources[tmplData.GetOutputId()] = tmplData
		}
	}

	outputTemplate := ""
	for _, r := range resources {
		tmpl1, err := resource.ApplyTemplate(r, resource.DataTemplateString)
		if err != nil {
			return err
		}
		tmpl2, err := resource.ApplyTemplate(r, resource.OutputTemplateString)
		if err != nil {
			return err
		}
		outputTemplate = outputTemplate + tmpl1 + tmpl2
	}
	resource.ExecuteTerraform(outputTemplate, terraformDir)

	// Creates the import template
	importTemplate := ""
	for _, v := range resources {
		tmpl, err := resource.ApplyTemplate(v, resource.ImportTemplateString)
		if err != nil {
			return err
		}
		importTemplate = importTemplate + tmpl
	}
	writeToFile("terraform/output.tf", importTemplate)

	// Creates the mappings to parse the terraform plan
	mappings := make(map[string](map[string]map[string]resource.FieldAttribute))
	if object == "*" {
		for _, object := range objects {
			x, err := resource.ParseFieldAttributes(object)
			if err != nil {
				return err
			}
			mappings[resource.ObjectNameToTerraformType(object)] = x
		}
	} else {
		x, err := resource.ParseFieldAttributes(object)
		if err != nil {
			return err
		}
		mappings[resource.ObjectNameToTerraformType(object)] = x
	}
	// Parses the terraform plan and generates the imports.tf file
	result, err := resource.ParseHCLBlocks(terraformDir, mappings, resources)
	if err != nil {
		return err
	}
	writeToFile("terraform/imports.tf", result)

	return nil
}
