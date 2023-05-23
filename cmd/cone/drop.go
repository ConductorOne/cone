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
	profileName, err := cmd.Flags().GetString("profile")
	if err != nil {
		return err
	}
	_, err = getProfile(cmd, profileName)
	if err != nil {
		return err
	}

	// TODO: Implement this
	return nil
}
