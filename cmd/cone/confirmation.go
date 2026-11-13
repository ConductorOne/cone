package main

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const (
	confirmFlag       = "confirm"
	confirmAnnotation = "cone.conductorone.com/confirm"
)

func supportsConfirmation(cmd *cobra.Command) *cobra.Command {
	if cmd.Annotations == nil {
		cmd.Annotations = make(map[string]string)
	}
	cmd.Annotations[confirmAnnotation] = "true"
	return cmd
}

func confirmationSupported(cmd *cobra.Command) bool {
	return cmd.Annotations[confirmAnnotation] == "true"
}

func confirmMutation(cmd *cobra.Command, action string) error {
	enabled, _ := cmd.Flags().GetBool(confirmFlag)
	if !enabled {
		return nil
	}

	nonInteractive, _ := cmd.Flags().GetBool(nonInteractiveFlag)
	if nonInteractive {
		return fmt.Errorf("--confirm requires interactive mode")
	}

	confirmed, err := pterm.DefaultInteractiveConfirm.Show("Proceed with " + action + "?")
	if err != nil {
		return err
	}
	if !confirmed {
		return fmt.Errorf("mutation cancelled")
	}
	return nil
}
