package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func installMCPCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install-mcp",
		Short: "Connect Claude Code to ConductorOne's hosted MCP gateway.",
		Long: `Registers ConductorOne's hosted MCP endpoint with Claude Code.

Claude Code handles OAuth (including DCR) on first connection.
Cone just provides the endpoint URL based on your existing login.

Requires 'cone login <tenant>' to have been run first.`,
		RunE: installMCPRun,
	}

	cmd.Flags().String("scope", "user", "Claude Code scope: user or project")
	cmd.Flags().Bool("dry-run", false, "Print what would happen without doing it")
	cmd.Flags().Bool("manual", false, "Print config snippet instead of running claude CLI")

	return cmd
}

func installMCPRun(cmd *cobra.Command, args []string) error {
	v, err := getSubViperForProfile(cmd)
	if err != nil {
		return err
	}

	clientID := v.GetString("client-id")
	if clientID == "" {
		return fmt.Errorf("not authenticated. Run 'cone login <tenant>' first")
	}

	// Parse tenant host from client-id.
	// Format: {name}@{host}/{path}
	tenantHost, err := parseTenantHost(clientID)
	if err != nil {
		return fmt.Errorf("could not determine tenant from client-id %q: %w", clientID, err)
	}

	mcpURL := fmt.Sprintf("https://%s/api/v1alpha/mcp", tenantHost)

	profile := v.GetString("profile")
	if profile == "" {
		profile = "default"
	}
	serverName := "conductorone"
	if profile != "default" {
		serverName = fmt.Sprintf("conductorone-%s", profile)
	}

	scope, _ := cmd.Flags().GetString("scope")
	dryRun, _ := cmd.Flags().GetBool("dry-run")
	manual, _ := cmd.Flags().GetBool("manual")

	if dryRun {
		fmt.Printf("Would run:\n  claude mcp add --transport http --scope %s %s %s\n", scope, serverName, mcpURL)
		return nil
	}

	if manual {
		printManualMCPConfig(serverName, mcpURL)
		return nil
	}

	claudePath, err := exec.LookPath("claude")
	if err != nil {
		fmt.Printf("claude CLI not found on PATH. Falling back to manual config.\n\n")
		printManualMCPConfig(serverName, mcpURL)
		return nil
	}

	spinner, err := pterm.DefaultSpinner.Start(fmt.Sprintf("Registering MCP server %q...", serverName))
	if err != nil {
		return err
	}

	out, err := exec.CommandContext(cmd.Context(), claudePath, "mcp", "add",
		"--transport", "http",
		"--scope", scope,
		serverName, mcpURL,
	).CombinedOutput()
	if err != nil {
		spinner.Fail(fmt.Sprintf("Failed to register: %s", strings.TrimSpace(string(out))))
		return fmt.Errorf("claude mcp add failed: %w", err)
	}

	spinner.Success(fmt.Sprintf("Registered %q (scope: %s)", serverName, scope))
	fmt.Printf("\nEndpoint: %s\n", mcpURL)
	fmt.Printf("Claude Code will handle OAuth on first connection.\n")
	fmt.Printf("Restart Claude Code or run /mcp to connect.\n")

	return nil
}

// parseTenantHost extracts the host from a C1 client-id.
// Client-id format: {name}@{host}/{path}
func parseTenantHost(clientID string) (string, error) {
	parts := strings.SplitN(clientID, "@", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("expected format {name}@{host}/{path}")
	}
	hostPath := parts[1]
	hostParts := strings.SplitN(hostPath, "/", 2)
	if len(hostParts) != 2 {
		return "", fmt.Errorf("expected format {name}@{host}/{path}")
	}
	return hostParts[0], nil
}

func printManualMCPConfig(serverName, mcpURL string) {
	config := map[string]any{
		"type": "http",
		"url":  mcpURL,
	}
	configJSON, _ := json.MarshalIndent(map[string]any{
		serverName: config,
	}, "", "  ")

	fmt.Printf("Add the following to your Claude Code MCP config:\n\n")
	fmt.Printf("%s\n", configJSON)
}
