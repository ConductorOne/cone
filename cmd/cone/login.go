package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/toqueteos/webbrowser"

	conductoroneapi "github.com/conductorone/conductorone-sdk-go"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/managedconfig"
)

func loginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login [tenant-name or tenant-url]",
		Short: fmt.Sprintf("Authenticate to ConductorOne, creating config.yaml in %s if it doesn't exist.", defaultConfigPath()),
		Long: fmt.Sprintf("Authenticate to ConductorOne, creating config.yaml in %s if it doesn't exist.\n\n"+
			"If a managed device configuration is present, the tenant is discovered from it automatically "+
			"and the tenant argument may be omitted.", defaultConfigPath()),
		RunE: loginRun,
	}

	cmd.Flags().String("profile", "default", "Config profile to create or update.")
	return cmd
}

// resolveLoginTenant determines the tenant (name or URL) to authenticate
// against. Managed device configuration pushed by an administrator takes
// precedence over an argument supplied on the command line, allowing a bare
// "cone login" to discover its tenant automatically. When no managed
// configuration is present the behavior is unchanged: the tenant must be passed
// as an argument. The returned bool reports whether the tenant was sourced from
// managed configuration.
func resolveLoginTenant(cmd *cobra.Command, args []string) (string, bool, error) {
	if serverURL := managedconfig.Read().ControlPlaneURL(); serverURL != "" {
		return serverURL, true, nil
	}
	if err := validateArgLenth(1, args, cmd); err != nil {
		return "", false, err
	}
	return args[0], false, nil
}

func loginRun(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	tenant, fromManaged, err := resolveLoginTenant(cmd, args)
	if err != nil {
		return err
	}

	if fromManaged {
		pterm.Info.Printfln("Using tenant %q from managed device configuration.", tenant)
	}

	spinner, err := pterm.DefaultSpinner.Start("Logging in...")
	if err != nil {
		return err
	}

	v, err := getSubViperForProfile(cmd)
	if err != nil {
		return err
	}

	profile := "default"
	configFileUsed := filepath.Join(defaultConfigPath(), "config.yaml")
	if v.GetString("profile") != "" {
		profile = v.GetString("profile")
	}

	if _, err := os.Stat(defaultConfigPath()); os.IsNotExist(err) {
		if err := os.MkdirAll(defaultConfigPath(), 0700); err != nil {
			return err
		}
	}

	creds, err := conductoroneapi.LoginFlow(
		ctx,
		tenant,
		client.ConeClientID,
		"Created by Cone",
		func(validateDetails *conductoroneapi.DeviceCodeResponse) error {
			pterm.Printf("Attempting to open the device authorization page in your browser.\n"+
				"If your browser does not open or you wish to use a different device to authorize this request, visit the following:"+
				"\n\n    %s\n\nAnd verify the code in the browser matches the code below:\n\n    %s\n\n", validateDetails.VerificationURI, validateDetails.UserCode)
			// Ignore errors here, as we'll print the URL anyway
			_ = webbrowser.Open(validateDetails.VerificationURI)
			return nil
		},
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
