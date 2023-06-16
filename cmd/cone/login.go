package main

import (
	"fmt"

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
	ctx, c, _, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	tenant := args[0]

	clientID, clientSecret, err := c.Login(ctx, tenant)
	if err != nil {
		return err
	}

	fmt.Printf("Client ID: %s\nClient Secret: %s\n", clientID, clientSecret)

	return nil
}
