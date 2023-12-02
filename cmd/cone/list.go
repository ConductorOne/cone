package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listObjects = []string{"apps"}
var listObjectsStr = strings.Join(listObjects, ", ")

func listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <object-name>",
		Short: fmt.Sprintf("list all objects of the specified type: %s are supported", listObjectsStr),
		RunE:  listRun,
	}

	return cmd
}

func listApps(ctx context.Context, c client.C1Client, v *viper.Viper) error {
	apps, err := c.ListApps(ctx)
	if err != nil {
		return err
	}

	var appResources []*App
	for _, app := range apps {
		appCopy := app
		newApp := App{app: &appCopy, client: c}
		appResources = append(appResources, &newApp)
	}
	if len(appResources) == 0 {
		return nil
	}

	outputManager := output.NewManager(ctx, v)
	if output.IsWide(v) {
		err = outputManager.Output(ctx, &output.WideOutputList[*App]{Items: appResources})
		if err != nil {
			return err
		}
	} else {
		err = outputManager.Output(ctx, &output.OutputList[*App]{Items: appResources})
		if err != nil {
			return err
		}
	}

	return nil
}

func listRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if err := validateArgLenth(1, args, cmd); err != nil {
		return err
	}
	object := args[0]

	if v.GetString("output") != "" || v.GetString("output") != "wide" {
		return fmt.Errorf("only default and wide output is supported for list")
	}

	switch object {
	case "apps":
		return listApps(ctx, c, v)
	default:
		return fmt.Errorf("invalid object type, must be one of: %s", listObjectsStr)
	}
}
