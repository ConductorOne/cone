package main

import (
	"github.com/spf13/cobra"
)

func loginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login <tenant-name>",
		Short: "Authenticate to ConductorOne",
		RunE:  loginRun,
		Args:  cobra.ExactArgs(1),
	}

	return cmd
}

func loginRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	tenantName := args[0]

	return nil
}
