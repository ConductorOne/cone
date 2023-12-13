package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/terraform"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

const (
	terraformProviderExample = "https://github.com/ConductorOne/terraform-provider-conductorone/blob/main/examples/provider/provider.tf"
	tempTfFile               = "cone_temp.tf"
)

var objects = []string{"policy", "app_entitlement"}

func terraformGenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen <object-name> <terraform-directory-path>",
		Short: fmt.Sprintf("Import all terraform resources for the specified object type (%s, or * for all). Terraform v1.5 or later is required", strings.Join(objects, ", ")),
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

func getResourceMap(ctx context.Context, c client.C1Client, v *viper.Viper, object string) (map[string]terraform.TemplateData, error) {
	resources := make(map[string]terraform.TemplateData)
	if err := populateResourcesWithApps(ctx, c, object, resources); err != nil {
		return nil, err
	}
	if err := populateResourcesWithPolicies(ctx, c, object, resources); err != nil {
		return nil, err
	}
	if err := populateResourcesWithEntitlements(ctx, c, v, object, resources); err != nil {
		return nil, err
	}
	return resources, nil
}

func populateResourcesWithApps(ctx context.Context, c client.C1Client, object string, resources map[string]terraform.TemplateData) error {
	if object == "app" || object == "*" {
		apps, err := c.ListApps(ctx)
		if err != nil {
			return err
		}
		for _, app := range apps {
			tmplData := terraform.AppTemplate{App: app}
			resources[tmplData.GetOutputId()] = tmplData
		}
	}
	return nil
}

func populateResourcesWithPolicies(ctx context.Context, c client.C1Client, object string, resources map[string]terraform.TemplateData) error {
	if object == "policy" || object == "*" {
		policies, err := c.ListPolicies(ctx)
		if err != nil {
			return err
		}
		for _, policy := range policies {
			tmplData := terraform.PolicyTemplate{Policy: policy}
			resources[tmplData.GetOutputId()] = tmplData
		}
	}
	return nil
}

func populateResourcesWithEntitlements(ctx context.Context, c client.C1Client, v *viper.Viper, object string, resources map[string]terraform.TemplateData) error {
	if object == "app_entitlement" || object == "*" {
		appId := v.GetString(tfAppIdFlag)
		if appId == "" {
			return errors.New("app-id flag is required for app_entitlement object")
		}

		entitlements, err := c.ListEntitlements(ctx, appId)
		if err != nil {
			return err
		}
		for _, entitlement := range entitlements {
			tmplData := terraform.AppEntitlementTemplate{AppEntitlement: entitlement}
			resources[tmplData.GetOutputId()] = tmplData
		}
	}
	return nil
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
		return fmt.Errorf("invalid object name, the following are supported: %s, or * for all)", strings.Join(objects, ", "))
	}

	inputDir := args[1]
	terraformDir, err := filepath.Abs(inputDir)
	if err != nil {
		return fmt.Errorf("terraform directory %s does not exist", terraformDir)
	}

	tempFilePath := path.Join(terraformDir, tempTfFile)

	// Turns objects into dataTemplates
	resources, err := getResourceMap(ctx, c, v, object)
	if err != nil {
		return err
	}

	outputTemplate, err := terraform.ApplyTemplates(maps.Values(resources), terraform.ImportTemplateString)
	if err != nil {
		return err
	}

	err = writeToFile(tempFilePath, outputTemplate)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	cmdTf := exec.Command("terraform", "plan", "-generate-config-output=generated_resources.tf")
	cmdTf.Dir = terraformDir
	cmdTf.Stdout = &buffer
	err = cmdTf.Run()
	if err != nil {
		return fmt.Errorf("terraform plan failed: %w.\nThis relies on an experimental feature from Hashicorp, make sure you are running Terraform v1.5 or later", err)
	}

	err = os.Remove(tempFilePath)
	if err != nil {
		return err
	}
	return nil
}
