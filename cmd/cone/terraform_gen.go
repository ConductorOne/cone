package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/resource"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var tempFile = "cone_temp.txt"
var tempTfFile = "cone_temp.tf"

var objects = []string{"app", "policy", "app_entitlement"}

func terraformGenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen <object-name> <terraform-directory-path>",
		Short: "Import all terraform resources for the specified object type",
		RunE:  terraformGen,
	}
	addTfAppIdFlag(cmd)

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

func getResourceMap(ctx context.Context, c client.C1Client, v *viper.Viper, object string) (map[string]resource.TemplateData, error) {
	resources := make(map[string]resource.TemplateData)
	if object == "app" || object == "*" {
		apps, err := c.ListApps(ctx)
		if err != nil {
			return nil, err
		}
		for _, app := range apps {
			tmplData := resource.AppTemplate{App: app}
			resources[tmplData.GetOutputId()] = tmplData
		}
	}
	if object == "policy" || object == "*" {
		policies, err := c.ListPolicies(ctx)
		if err != nil {
			return nil, err
		}
		for _, policy := range policies {
			tmplData := resource.PolicyTemplate{Policy: policy}
			resources[tmplData.GetOutputId()] = tmplData
		}
	}

	if object == "app_entitlement" || object == "*" {
		appId := v.GetString(tfAppIdFlag)
		if appId == "" {
			return nil, errors.New("app-id flag is required for app_entitlement object")
		}

		entitlements, err := c.ListEntitlements(ctx, appId)
		if err != nil {
			return nil, err
		}
		for _, entitlement := range entitlements {
			tmplData := resource.AppEntitlementTemplate{AppEntitlement: entitlement}
			resources[tmplData.GetOutputId()] = tmplData
		}
	}
	return resources, nil
}

func terraformGen(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
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

	resources, err := getResourceMap(ctx, c, v, object)
	if err != nil {
		return err
	}

	outputTemplate, err := resource.ApplyTemplates(maps.Values(resources), resource.DataTemplateString, resource.OutputTemplateString)
	if err != nil {
		return err
	}
	err = writeToFile(terraformDir+"/cone_temp.tf", outputTemplate)
	if err != nil {
		return err
	}
	pterm.Info.Println("Please run this command in the terraform directory:")
	pterm.Info.Printfln(`touch %s; terraform plan | sed 's/\x1b\[[0-9;]*m//g'> %s`, tempFile, tempFile)

	ok, err := pterm.DefaultInteractiveConfirm.WithDefaultText("Have you run the command?").Show()
	if err != nil {
		return err
	}
	if !ok {
		pterm.Error.Println("You must run the command to continue")
		return nil
	}

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
	result, err := resource.ParseHCLBlocks(terraformDir+"/"+tempFile, mappings, resources)
	if err != nil {
		return err
	}

	// Writes the final imports and deletes the temp files
	err = writeToFile(terraformDir+"/cone_output.tf", importTemplate)
	if err != nil {
		return err
	}
	err = writeToFile(terraformDir+"/cone_imports.tf", result)
	if err != nil {
		return err
	}
	err = os.Remove(fmt.Sprintf("%s/%s", terraformDir, tempFile))
	if err != nil {
		return err
	}
	err = os.Remove(fmt.Sprintf("%s/%s", terraformDir, tempTfFile))
	if err != nil {
		return err
	}
	pterm.Info.Println("You can now run terraform refresh to import the resources. After that you can delete the cone_output.tf file")
	return nil
}
