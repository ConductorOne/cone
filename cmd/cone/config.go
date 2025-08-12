package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envPrefix             = "cone"
	nativeIntegrationMode = "native"
)

func defaultConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	return filepath.Join(homeDir, ".conductorone")
}

var (
	ErrNoCredentials = errors.New("client-id and client-secret must be set")
)

func initConfig(cmd *cobra.Command) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	configPath := os.Getenv("CONE_CONFIG_PATH")
	if configPath != "" {
		viper.AddConfigPath(configPath)
	} else {
		viper.AddConfigPath(defaultConfigPath())
	}

	err := viper.ReadInConfig()
	if err != nil {
		notFoundErr := &viper.ConfigFileNotFoundError{}
		// Explicitly ignore the not found error case
		if ok := errors.As(err, notFoundErr); ok {
			return nil
		}
		return fmt.Errorf("fatal error config file: %w", err)
	}

	if err := viper.BindPFlags(cmd.PersistentFlags()); err != nil {
		return err
	}
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	return nil
}

func getSubViperForProfile(cmd *cobra.Command) (*viper.Viper, error) {
	profile := viper.GetString("profile")
	if profile == "" {
		profile = "default"
	}

	v := viper.Sub(fmt.Sprintf("profiles.%s", profile))
	if v == nil {
		// No profile found, so create a new viper instance
		v = viper.New()
	}
	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	if err := v.BindPFlags(cmd.PersistentFlags()); err != nil {
		return nil, err
	}
	if err := v.BindPFlags(cmd.Flags()); err != nil {
		return nil, err
	}

	return v, nil
}

// Validate credentials are set, and return them (client-id, client-secret, error).
func getCredentials(v *viper.Viper) (string, string, error) {
	clientId := v.GetString("client-id")
	clientSecret := v.GetString("client-secret")

	if clientId == "" || clientSecret == "" {
		return "", "", ErrNoCredentials
	}
	return clientId, clientSecret, nil
}

// showAWSConfigCmd creates the command for displaying all AWS configuration settings.
func showAWSConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show all AWS configuration settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			raw, _ := cmd.Flags().GetBool("raw")
			if raw {
				// Get the config file path
				configFile := viper.ConfigFileUsed()
				if configFile == "" {
					return fmt.Errorf("no config file found")
				}

				// Read and print the raw YAML content
				content, err := os.ReadFile(configFile)
				if err != nil {
					return fmt.Errorf("error reading config file: %w", err)
				}
				pterm.Println(string(content))
				return nil
			}

			// Get SSO start URL
			ssoStartURL := viper.GetString("aws_sso_start_url")
			if ssoStartURL == "" {
				pterm.Warning.Println("AWS SSO start URL is not set")
			} else {
				pterm.Info.Printf("AWS SSO start URL: %s\n", ssoStartURL)
			}

			// Get integration mode
			integrationMode := viper.GetString("aws_integration_mode")
			if integrationMode == "" {
				integrationMode = nativeIntegrationMode // Default to native if not set
			}
			pterm.Info.Printf("AWS integration mode: %s\n", integrationMode)

			// Show default behavior
			pterm.Println("\nDefault behavior:")
			if integrationMode == "cone" {
				pterm.Println("- AWS profiles will be created automatically")
				pterm.Println("- Uses Cone's credential process for authentication")
			} else {
				pterm.Println("- No AWS profiles will be created")
				pterm.Println("- Uses AWS CLI's native SSO integration")
			}

			return nil
		},
	}

	// Add the raw flag
	cmd.Flags().Bool("raw", false, "Show raw YAML content of the config file")
	return cmd
}

// configAwsCmd creates the main AWS configuration command.
func configAwsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config-aws",
		Short: "Configure AWS settings",
	}
	cmd.AddCommand(setAWSSSOStartURLCmd())
	cmd.AddCommand(getAWSSSOStartURLCmd())
	cmd.AddCommand(setAWSIntegrationModeCmd())
	cmd.AddCommand(getAWSIntegrationModeCmd())
	cmd.AddCommand(showAWSConfigCmd())
	return cmd
}

// setAWSSSOStartURLCmd creates the command for setting the AWS SSO start URL.
func setAWSSSOStartURLCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-sso-url <url>",
		Short: "Set the AWS SSO start URL",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("must provide a URL")
			}
			url := args[0]
			viper.Set("aws_sso_start_url", url)
			if err := viper.SafeWriteConfig(); err != nil {
				return err
			}
			pterm.Info.Printf("AWS SSO start URL set to: %s\n", url)
			return nil
		},
	}
	return cmd
}

// getAWSSSOStartURLCmd creates the command for retrieving the current AWS SSO start URL.
func getAWSSSOStartURLCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-sso-url",
		Short: "Get the current AWS SSO start URL",
		RunE: func(cmd *cobra.Command, args []string) error {
			url := viper.GetString("aws_sso_start_url")
			if url == "" {
				pterm.Warning.Println("AWS SSO start URL is not set")
				return nil
			}
			pterm.Info.Printf("AWS SSO start URL: %s\n", url)
			return nil
		},
	}
	return cmd
}

// setAWSIntegrationModeCmd creates the command for setting the AWS integration mode.
func setAWSIntegrationModeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-integration-mode <mode>",
		Short: "Set the AWS integration mode (cone|native)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("must provide a mode (cone|native)")
			}
			mode := strings.ToLower(args[0])
			if mode != "cone" && mode != "native" {
				return fmt.Errorf("mode must be either 'cone' or 'native'")
			}
			viper.Set("aws_integration_mode", mode)
			if err := viper.SafeWriteConfig(); err != nil {
				return err
			}
			pterm.Info.Printf("AWS integration mode set to: %s\n", mode)
			return nil
		},
	}
	return cmd
}

// getAWSIntegrationModeCmd creates the command for retrieving the current AWS integration mode.
func getAWSIntegrationModeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-integration-mode",
		Short: "Get the current AWS integration mode",
		RunE: func(cmd *cobra.Command, args []string) error {
			mode := viper.GetString("aws_integration_mode")
			if mode == "" {
				mode = "native" // Default to native if not set
			}
			pterm.Info.Printf("AWS integration mode: %s\n", mode)
			return nil
		},
	}
	return cmd
}
