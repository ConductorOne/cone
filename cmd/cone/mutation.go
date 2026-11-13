package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const dryRunFlag = "dry-run"

const dryRunAnnotation = "cone.conductorone.com/dry-run"

// mutation describes a change that a command would make. It stays deliberately
// human-readable: cone's preview is a safety check, not a transport trace.
type mutation struct {
	Action  string
	Target  string
	Details []string
}

type tenantTarget interface {
	BaseURL() string
}

func supportsDryRun(cmd *cobra.Command) *cobra.Command {
	if cmd.Annotations == nil {
		cmd.Annotations = make(map[string]string)
	}
	cmd.Annotations[dryRunAnnotation] = "true"
	return cmd
}

func dryRunSupported(cmd *cobra.Command) bool {
	return cmd.Annotations[dryRunAnnotation] == "true"
}

func dryRunEnabled(cmd *cobra.Command) bool {
	enabled, _ := cmd.Flags().GetBool(dryRunFlag)
	return enabled
}

// previewMutations prints the resolved tenant, profile, and planned mutations,
// then reports true when the caller must return without making a change.
func previewMutations(cmd *cobra.Command, c tenantTarget, v *viper.Viper, changes ...mutation) bool {
	if !dryRunEnabled(cmd) {
		return false
	}

	out := cmd.OutOrStdout()
	_, _ = fmt.Fprintln(out, "Dry run — no changes will be made.")
	_, _ = fmt.Fprintf(out, "\nTenant: %s\nProfile: %s\n\nPlan:\n", c.BaseURL(), v.GetString("profile"))
	for _, change := range changes {
		_, _ = fmt.Fprintf(out, "  %s: %s\n", change.Action, change.Target)
		for _, detail := range change.Details {
			if detail != "" {
				_, _ = fmt.Fprintf(out, "    %s\n", detail)
			}
		}
	}
	return true
}

func nonEmptyDetail(label, value string) string {
	if strings.TrimSpace(value) == "" {
		return ""
	}
	return label + ": " + value
}
