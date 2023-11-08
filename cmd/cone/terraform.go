package main

import (
	"fmt"

	"github.com/conductorone/cone/pkg/resource"
	"github.com/spf13/cobra"
)

func tfCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "tf",
		RunE: tfRun,
	}

	return cmd
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

	for _, app := range apps {
		// Create an instance of AppTemplate with the app
		appTmpl := resource.AppTemplate{App: app} // Now using the exported field 'App'

		// Apply the template using the appTmpl
		tmpl, err := resource.ApplyTemplate(appTmpl)
		if err != nil {
			return err
		}
		fmt.Print(tmpl)
		if err = resource.ExecuteTerraform(tmpl, "terraform"); err != nil {
			return err
		}
		break
	}

	return nil
}
