package main

import "github.com/spf13/cobra"

const (
	tfAppIdFlag = "app-id"
)

func addTfAppIdFlag(cmd *cobra.Command) {
	cmd.Flags().String(tfAppIdFlag, "", "App ID to get entitlements for.")
}
