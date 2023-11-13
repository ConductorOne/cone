package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/conductorone/cone/pkg/resource"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var objects = []string{"app", "policy"}

func tfCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tf <resource-type> <terraform-directory-path>",
		Short: "Import all terraform resources for the specified resource type",
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

func executeTerraformPlan(outputDir string) error {
	err := exec.Command("cd " + outputDir).Run()
	if err != nil {
		return err
	}
	cmd := exec.Command("/bin/sh", "-c", "terraform plan "+`| sed 's/\x1b\[[0-9;]*m//g'>`+outputDir+"/cone_temp.txt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func tfRun(cmd *cobra.Command, args []string) error {
	ctx, c, _, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(2, args, cmd); err != nil {
		return err
	}

	object := args[0]
	if !slices.Contains(objects, object) && object != "*" {
		return fmt.Errorf("invalid object name, only support %v and * for all", objects)
	}

	terraformDir := args[1]
	if _, err := os.Stat(terraformDir); err != nil {
		return fmt.Errorf("terraform directory %s does not exist", terraformDir)
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

	outputTemplate, err := resource.ApplyTemplates(maps.Values(resources), resource.DataTemplateString, resource.OutputTemplateString)
	writeToFile(terraformDir+"/cone_temp.tf", outputTemplate)
	executeTerraformPlan(terraformDir)

	// Creates the import template
	importTemplate, err := resource.ApplyTemplates(maps.Values(resources), resource.ImportTemplateString)
	if err != nil {
		return err
	}

	// Creates the mappings to parse the terraform plan
	mappings := make(map[string](map[string]map[string]resource.FieldAttribute))
	for _, v := range objects {
		if object == v || object == "*" {
			x, err := resource.ParseFieldAttributes(object)
			if err != nil {
				return err
			}
			mappings[resource.ObjectNameToTerraformType(object)] = x
		}
	}

	// Parses the terraform plan and generates the imports.tf file
	result, err := resource.ParseHCLBlocks(terraformDir, mappings, resources)
	if err != nil {
		return err
	}

	// Writes the final imports and deletes the temp files
	writeToFile(terraformDir+"/cone_output.tf", importTemplate+result)
	writeToFile(terraformDir+"/cone_imports.tf", result)
	os.Remove(terraformDir + "/cone_temp.txt")
	os.Remove(terraformDir + "/cone_temp.tf")

	return nil
}
