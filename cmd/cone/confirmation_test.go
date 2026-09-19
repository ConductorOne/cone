package main

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestConfirmMutationRequiresInteractiveMode(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool(confirmFlag, false, "")
	cmd.Flags().Bool(nonInteractiveFlag, false, "")
	if err := cmd.ParseFlags([]string{"--confirm", "--non-interactive"}); err != nil {
		t.Fatalf("ParseFlags: %v", err)
	}

	err := confirmMutation(cmd, "creating an access request")
	if err == nil || !strings.Contains(err.Error(), "requires interactive mode") {
		t.Fatalf("confirmMutation error = %v, want interactive-mode error", err)
	}
}

func TestConfirmMutationIsOptIn(t *testing.T) {
	if err := confirmMutation(&cobra.Command{}, "creating an access request"); err != nil {
		t.Fatalf("confirmMutation without --confirm: %v", err)
	}
}
