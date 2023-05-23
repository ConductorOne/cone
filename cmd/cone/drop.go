package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func dropCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
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
	profile, err := getProfile(cmd, profileName)
	if err != nil {
		return err
	}
	fmt.Println(profile.ClientID, profile.ClientSecret)

	// TODO: Implement this
	return nil
}
