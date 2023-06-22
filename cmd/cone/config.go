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

var defaultConfigPath string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "$HOME"
	}

	defaultConfigPath = filepath.Join(homeDir, ".conductorone")
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
		viper.AddConfigPath(defaultConfigPath)
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
