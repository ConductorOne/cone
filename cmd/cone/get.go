package main

import (
	"github.com/spf13/cobra"
)

func getCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Create an access request for an entitlement by slug",
		RunE:  runGet,
	}

	addWaitFlag(cmd)
	addAppIdFlag(cmd)
	addEntitlementIdFlag(cmd)
	addEntitlementAliasFlag(cmd)

	return cmd
}

func runGet(cmd *cobra.Command, args []string) error {
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
