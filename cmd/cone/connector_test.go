package main

import (
	"testing"
)

func TestConnectorCmd(t *testing.T) {
	cmd := connectorCmd()

	if cmd.Use != "connector" {
		t.Errorf("expected Use to be 'connector', got %s", cmd.Use)
	}
	if cmd.Short != "Manage ConductorOne connectors" {
		t.Errorf("expected Short to be 'Manage ConductorOne connectors', got %s", cmd.Short)
	}
	if !cmd.HasSubCommands() {
		t.Error("connector command should have subcommands")
	}
}

func TestConnectorBuildCmd(t *testing.T) {
	cmd := connectorBuildCmd()

	if cmd.Use != "build [path]" {
		t.Errorf("expected Use to be 'build [path]', got %s", cmd.Use)
	}
	if cmd.Short != "Build a connector binary" {
		t.Errorf("expected Short to be 'Build a connector binary', got %s", cmd.Short)
	}

	// Verify flags exist
	outputFlag := cmd.Flag("output")
	if outputFlag == nil {
		t.Error("should have --output flag")
	} else if outputFlag.Shorthand != "o" {
		t.Errorf("expected output shorthand to be 'o', got %s", outputFlag.Shorthand)
	}

	osFlag := cmd.Flag("os")
	if osFlag == nil {
		t.Error("should have --os flag")
	}

	archFlag := cmd.Flag("arch")
	if archFlag == nil {
		t.Error("should have --arch flag")
	}
}
