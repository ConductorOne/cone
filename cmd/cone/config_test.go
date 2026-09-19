package main

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func TestInitConfigBindsFlagsWithoutConfigFile(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	t.Setenv("CONE_CONFIG_PATH", t.TempDir())

	cmd := &cobra.Command{}

	cmd.PersistentFlags().Bool("debug", false, "")
	if err := initConfig(cmd); err != nil {
		t.Fatalf("initConfig: %v", err)
	}
	if err := cmd.ParseFlags([]string{"--debug"}); err != nil {
		t.Fatalf("ParseFlags: %v", err)
	}
	if !viper.GetBool("debug") {
		t.Fatal("--debug was not bound when no config file exists")
	}
}
