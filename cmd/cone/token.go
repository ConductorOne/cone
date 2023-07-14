package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func tokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token",
		Short: "",
		RunE:  tokenRun,
	}

	return cmd
}

func tokenRun(cmd *cobra.Command, args []string) error {
	ctx, c, v, err := cmdContext(cmd)
	if err != nil {
		return err
	}

	if len(args) != 0 {
		usageErrorString := cmd.UsageString()
		return fmt.Errorf("expected 0 arguments, got %d\n"+usageErrorString, len(args))
	}

	return nil
}
