package main

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type previewClient struct{ baseURL string }

func (c previewClient) BaseURL() string { return c.baseURL }

func TestPreviewMutations(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.PersistentFlags().Bool(dryRunFlag, false, "")
	if err := cmd.ParseFlags([]string{"--dry-run"}); err != nil {
		t.Fatalf("ParseFlags: %v", err)
	}
	out := &bytes.Buffer{}
	cmd.SetOut(out)
	v := viper.New()
	v.Set("profile", "admin")

	if !previewMutations(cmd, previewClient{baseURL: "https://tenant.example"}, v, mutation{
		Action:  "Approve task",
		Target:  "task T-1",
		Details: []string{"Comment: reviewed"},
	}) {
		t.Fatal("previewMutations = false, want true when --dry-run is set")
	}
	want := "Dry run — no changes will be made.\n\nTenant: https://tenant.example\nProfile: admin\n\nPlan:\n  Approve task: task T-1\n    Comment: reviewed\n"
	if got := out.String(); got != want {
		t.Fatalf("preview output = %q, want %q", got, want)
	}
}

func TestPreviewMutationsDisabled(t *testing.T) {
	cmd := &cobra.Command{}
	out := &bytes.Buffer{}
	cmd.SetOut(out)
	v := viper.New()
	if previewMutations(cmd, previewClient{baseURL: "https://tenant.example"}, v, mutation{}) {
		t.Fatal("previewMutations = true, want false without --dry-run")
	}
	if got := out.String(); got != "" {
		t.Fatalf("preview wrote %q without --dry-run", got)
	}
}

func TestDryRunSupportIsScopedToPreviewCommands(t *testing.T) {
	for _, cmd := range []*cobra.Command{
		getCmd(),
		dropCmd(),
		approveTasksCmd(),
		denyTasksCmd(),
		tasksCommentCmd(),
		escalateTasksCmd(),
		generateAliasCmd(),
		installMCPCmd(),
	} {
		if !dryRunSupported(cmd) {
			t.Fatalf("%s does not support --dry-run", cmd.CommandPath())
		}
	}

	for _, cmd := range []*cobra.Command{tasksCmd(), getTasksCmd(), secretCmd()} {
		if dryRunSupported(cmd) {
			t.Fatalf("%s unexpectedly supports --dry-run", cmd.CommandPath())
		}
	}
}
