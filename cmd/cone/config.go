package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envPrefix = "cone"
)

type Config struct {
	Profiles map[string]ConfigProfile `yaml:"profiles"`
}

type ConfigProfile struct {
	ClientID     string `yaml:"client-id"`
	ClientSecret string `yaml:"client-secret"`
}

func initConfig(cmd *cobra.Command) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	configPath := viper.GetString("config-path")
	if configPath != "" {
		viper.AddConfigPath(configPath)
	}
	viper.AddConfigPath("$HOME/.cone")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		notFoundErr := &viper.ConfigFileNotFoundError{}
		if ok := errors.As(err, notFoundErr); ok {
			err = createDefaultConfig()
			if err != nil {
				return err
			}
			return nil
		}
		return fmt.Errorf("fatal error config file: %w", err)
	}

	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	if err := viper.BindPFlag("profile", cmd.PersistentFlags().Lookup("profile")); err != nil {
		return err
	}

	return nil
}

func createDefaultConfig() error {
	defaultConfig := Config{
		Profiles: map[string]ConfigProfile{
			"default": {
				ClientID:     "<client-id>",
				ClientSecret: "<client-secret>",
			},
		},
	}

	viper.SetDefault("profiles", defaultConfig.Profiles)
	err := viper.SafeWriteConfigAs("config")
	if err != nil {
		return fmt.Errorf("failed creating config file: %w", err)
	}
	return nil
}

func getSubViperForProfile(cmd *cobra.Command) (*viper.Viper, error) {
	profile := viper.GetString("profile")
	if profile == "" {
		profile = "default"
	}

	v := viper.Sub(fmt.Sprintf("profiles.%s", profile))
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
