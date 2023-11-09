package main

import (
	"os"

	"github.com/conductorone/cone/pkg/resource"
	"github.com/spf13/cobra"
)

const terraformDir = "terraform"

func tfCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "tf",
		RunE: tfRun,
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

	apps, err := c.ListApps(ctx)
	if err != nil {
		return err
	}
	res := ""
	for _, app := range apps {
		// Create an instance of AppTemplate with the app
		appTmpl := resource.AppTemplate{App: app} // Now using the exported field 'App'

		// Apply the template using the appTmpl
		tmpl, err := resource.ApplyTemplate(appTmpl)
		if err != nil {
			return err
		}

		res = res + tmpl
	}

	resource.ExecuteTerraform(res, terraformDir)
	x, err := resource.ParseFieldAttributes("app")
	if err != nil {
		return err
	}

	mappings := make(map[string]map[string]resource.FieldAttribute)
	mappings["conductorone_app"] = x
	result, err := resource.ParseHCLBlocks(terraformDir, "plan", mappings)
	if err != nil {
		return err
	}
	writeToFile("terraform/imports.tf", result)
	return nil
}
