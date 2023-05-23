package main

import (
	"github.com/spf13/cobra"
)

const (
	waitFlag          = "wait"
	entitlementIdFlag = "entitlement-id"
	entitlementAlias  = "alias"
	AppIdFlag         = "app-id"
)

func addWaitFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(waitFlag, "w", "", "Wait for the task to be approved and provisioned.")
}

func addEntitlementIdFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(entitlementIdFlag, "e", "", "The entitlement id to filter by")
}

func addEntitlementAliasFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(entitlementAlias, "x", "", "The entitlement alias to filter by")
}

func addAppIdFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(AppIdFlag, "a", "", "The app id to filter by")
}
