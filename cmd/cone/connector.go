package main

import (
	"github.com/spf13/cobra"
)

// connectorCmd returns the root command for connector operations
// (subcommands: init, dev, build).
func connectorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connector",
		Short: "Manage ConductorOne connectors",
		Long: `Commands for developing, building, and managing ConductorOne connectors.

The connector subcommands help you:
  - Initialize new connector projects
  - Run a local development server with hot reload
  - Build connector binaries for deployment
  - Publish connectors to the ConductorOne registry`,
	}

	cmd.AddCommand(connectorBuildCmd())
	cmd.AddCommand(connectorInitCmd())
	cmd.AddCommand(connectorDevCmd())
	cmd.AddCommand(connectorPublishCmd())
	cmd.AddCommand(connectorValidateConfigCmd())

	return cmd
}
