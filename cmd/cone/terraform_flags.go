package main

import "github.com/spf13/cobra"

const (
	tfAppIdFlag  = "app-id"
	tfOutputFlag = "out"
)

func addTfAppIdFlag(cmd *cobra.Command) {
	cmd.Flags().String(tfAppIdFlag, "", "App ID to get entitlements for.")
}

func addTfOutputFlag(cmd *cobra.Command) {
	cmd.Flags().String(tfOutputFlag, "generated_resources.tf", "File name for generated resources, default name is generated_resources.tf")
}
