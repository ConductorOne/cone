package main

import (
	"fmt"
	"path/filepath"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/toqueteos/webbrowser"

	conductoroneapi "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/cone/pkg/client"
)

func loginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login <tenant-name>",
		Short: fmt.Sprintf("Authenticate to ConductorOne, creating config.yaml in %s if it doesn't exist.", defaultConfigPath),
		RunE:  loginRun,
		Args:  cobra.ExactArgs(1),
	}

	cmd.Flags().String("profile", "default", "Config profile to create or update.")

	return cmd
}

func loginRun(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	tenant := args[0]

	spinner, err := pterm.DefaultSpinner.Start("Logging in...")
	if err != nil {
		return err
	}

	v, err := getSubViperForProfile(cmd)
	if err != nil {
		return err
	}

	profile := "default"
	configFileUsed := filepath.Join(defaultConfigPath, "config.yaml")
	if v.GetString("profile") != "" {
		profile = v.GetString("profile")
	}

	creds, err := conductoroneapi.LoginFlow(
		ctx,
		tenant,
		client.ConeClientID,
		"Created by Cone",
		webbrowser.Open,
	)
	if err != nil {
		spinner.Fail(err)
		return err
	}

	profiles := viper.GetStringMap("profiles")

	v = viper.New()
	v.Set(fmt.Sprintf("profiles.%s", profile), map[string]string{
		"client-id":     creds.ClientID,
		"client-secret": creds.ClientSecret,
	})

	if err := v.MergeConfigMap(map[string]interface{}{"profiles": profiles}); err != nil {
		return err
	}

	if err := v.WriteConfigAs(configFileUsed); err != nil {
		spinner.Fail(err)
		return err
	}

	spinner.Success("Config written to " + configFileUsed)

	return nil
}
