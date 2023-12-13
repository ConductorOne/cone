package main

import (
	"github.com/spf13/cobra"
)

func terraformCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "terraform",
		Aliases: []string{"tf"},
		Short:   "A group of commands related to interacting with a terraform provider.",
		RunE:    terraformRun,
	}

	cmd.AddCommand(terraformGenCmd())

	return cmd
}

func terraformRun(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}
