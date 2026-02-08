package main

import (
	"github.com/spf13/cobra"
)

func functionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "functions",
		Short: "Commands for developing and deploying C1 Functions",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(functionsDevCmd())
	cmd.AddCommand(functionsNewCmd())
	cmd.AddCommand(functionsLogsCmd())
	cmd.AddCommand(functionsListCmd())

	// Hidden until SDK exposes FunctionsService.CreateCommit
	deployCmd := functionsDeployCmd()
	deployCmd.Hidden = true
	cmd.AddCommand(deployCmd)

	return cmd
}
