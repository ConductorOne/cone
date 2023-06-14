package main

import (
	"github.com/spf13/cobra"
)

const (
	waitFlag             = "wait"
	entitlementIdFlag    = "entitlement-id"
	appIdFlag            = "app-id"
	entitlementAliasFlag = "alias"
	queryFlag            = "query"
	justificationFlag    = "justification"
	durationFlag         = "duration"
	forceFlag            = "force"
	nonInteractiveFlag   = "non-interactive"
)

func addWaitFlag(cmd *cobra.Command) {
	cmd.Flags().BoolP(waitFlag, "w", false, "Wait for the task to be approved and provisioned.")
}

func addEntitlementIdFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(entitlementIdFlag, "e", "", "The entitlement id to filter by")
}

func addAppIdFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(appIdFlag, "a", "", "The app id to filter by")
}

func addQueryFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(queryFlag, "q", "", "The query to filter by")
}

func addEntitlementAliasFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(entitlementAliasFlag, "", "", "The entitlement alias to filter by")
}

func addJustificationFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(justificationFlag, "j", "Made with cone", "The justification for the request")
}

func addGrantDurationFlag(cmd *cobra.Command) {
	usageStr := "A sequence of decimal numbers, each with optional fraction and a unit suffix, such as \"12h\", \"1w2d\" or \"2h45m\". Valid units are (m)inutes, (h)ours, (d)ays, (w)eeks."
	cmd.Flags().StringP(durationFlag, "d", "", usageStr)
}

func addForceTaskCreateFlag(cmd *cobra.Command) {
	cmd.Flags().Bool(forceFlag, false, "Force the creation of a task even if the user already has (or doesn't have) the entitlement.")
}
