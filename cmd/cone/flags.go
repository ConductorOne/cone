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
	cmd.Flags().StringP(durationFlag, "d", "", "The duration of the grant in seconds")
}

func addTaskCommentFlag(cmd *cobra.Command) {
	cmd.Flags().StringP(durationFlag, "d", "", "The comment to add to the task")
}
