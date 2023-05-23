package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Profiles map[string]ConfigProfile `yaml:"profiles"`
}

type ConfigProfile struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.cone")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = createDefaultConfig()
			if err != nil {
				return err
			}
			return nil
		}
		return fmt.Errorf("fatal error config file: %w", err)
	}
	return nil
}

func getProfile(cmd *cobra.Command, key string) (*ConfigProfile, error) {
	var config Config

	if err := viper.Unmarshal(&config, viper.DecoderConfigOption(func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "yaml"
	})); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	if config.Profiles == nil {
		return nil, fmt.Errorf("no profiles found in config file")
	}
	configProfile, ok := config.Profiles[key]
	if !ok {
		return nil, fmt.Errorf("'%s' profile is not in the config file", key)
	}
	return &configProfile, nil
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
