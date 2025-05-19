package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envPrefix = "cone"
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

// configAwsCmd creates the main AWS configuration command
// This command was renamed from 'config' to 'config-aws' to be more specific about its AWS functionality
// It provides subcommands for managing AWS SSO settings, particularly the SSO start URL
// which is required for AWS credential management and permission set operations
func configAwsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config-aws",
		Short: "Manage AWS SSO configuration for ConductorOne",
		Long: `Manage AWS SSO configuration for ConductorOne.
This command helps you configure AWS SSO settings needed for AWS credential management.
It allows you to set and get the AWS SSO start URL, which is required for:
- Getting AWS credentials via 'cone aws-credentials'
- Managing AWS permission sets
- Accessing AWS resources through ConductorOne`,
	}

	cmd.AddCommand(setAWSSSOStartURLCmd())
	cmd.AddCommand(getAWSSSOStartURLCmd())

	return cmd
}

// setAWSSSOStartURLCmd creates the command for setting the AWS SSO start URL
// This command was renamed from 'set-aws-sso-url' to 'set-sso-url' for simplicity
// The URL is stored in the Viper configuration and is used by other AWS-related commands
// to authenticate with AWS SSO and manage permissions
func setAWSSSOStartURLCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-sso-url <url>",
		Short: "Set the AWS SSO start URL for your organization",
		Long: `Set the AWS SSO start URL for your organization.
This URL is required for AWS SSO authentication and is used when:
- Getting AWS credentials via 'cone aws-credentials'
- Managing AWS permission sets
- Accessing AWS resources through ConductorOne

Example: cone config-aws set-sso-url https://your-org.awsapps.com/start`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			url := args[0]
			viper.Set("aws_sso_start_url", url)
			if err := viper.WriteConfig(); err != nil {
				return fmt.Errorf("failed to write config: %w", err)
			}
			fmt.Printf("AWS SSO start URL set to: %s\n", url)
			return nil
		},
	}
	return cmd
}

// getAWSSSOStartURLCmd creates the command for retrieving the current AWS SSO start URL
// This command was renamed from 'get-aws-sso-url' to 'get-sso-url' for consistency
// It reads the URL from the Viper configuration and displays it to the user
// If no URL is set, it informs the user that the URL needs to be configured
func getAWSSSOStartURLCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-sso-url",
		Short: "Get the currently configured AWS SSO start URL",
		Long: `Get the currently configured AWS SSO start URL.
This URL is used for AWS SSO authentication and is required for:
- Getting AWS credentials via 'cone aws-credentials'
- Managing AWS permission sets
- Accessing AWS resources through ConductorOne`,
		RunE: func(cmd *cobra.Command, args []string) error {
			url := viper.GetString("aws_sso_start_url")
			if url == "" {
				fmt.Println("AWS SSO start URL is not set")
				return nil
			}
			fmt.Printf("AWS SSO start URL: %s\n", url)
			return nil
		},
	}
	return cmd
}
