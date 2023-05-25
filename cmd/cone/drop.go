package main

import (
	"github.com/spf13/cobra"
)

func dropCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "drop",
		Short: "Remove a specific entitlement grant",
		RunE:  runDrop,
	}

	addWaitFlag(cmd)
	addAppIdFlag(cmd)
	addEntitlementIdFlag(cmd)
	addEntitlementAliasFlag(cmd)

	return cmd
}

func runDrop(cmd *cobra.Command, args []string) error {
	// TODO: Implement this
	return nil
}
