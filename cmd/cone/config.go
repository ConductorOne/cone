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

// configAwsCmd creates the main AWS configuration command.
func configAwsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config-aws",
		Short: "Configure AWS settings",
	}
	cmd.AddCommand(setAWSSSOStartURLCmd())
	cmd.AddCommand(getAWSSSOStartURLCmd())
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
			if err := viper.WriteConfig(); err != nil {
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
