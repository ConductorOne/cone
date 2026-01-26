package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	c1client "github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/consent"
	"github.com/conductorone/cone/pkg/mcpclient"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func connectorAnalyzeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "analyze [path]",
		Short: "Analyze a connector with AI assistance",
		Long: `Analyze a connector using ConductorOne's AI copilot.

The AI will review your connector code and suggest improvements for:
  - Resource model completeness (users, groups, entitlements, grants)
  - SDK usage patterns and best practices
  - Error handling and edge cases
  - Performance and efficiency

This command requires consent for AI-assisted analysis.
Grant consent with: cone connector consent --agree

Examples:
  cone connector analyze                 # Analyze current directory
  cone connector analyze ./my-connector  # Analyze specific path
  cone connector analyze --offline       # Run offline checks only
  cone connector analyze --dry-run       # Preview without applying changes`,
		RunE: runConnectorAnalyze,
	}

	cmd.Flags().Bool("offline", false, "Run offline analysis only (no AI)")
	cmd.Flags().Bool("dry-run", false, "Preview changes without applying them")
	cmd.Flags().String("mode", "interactive", "Analysis mode: interactive or batch")
	cmd.Flags().String("server", "", "Override MCP server URL (for testing)")

	return cmd
}

func runConnectorAnalyze(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	offline, _ := cmd.Flags().GetBool("offline")
	dryRun, _ := cmd.Flags().GetBool("dry-run")
	mode, _ := cmd.Flags().GetString("mode")
	serverOverride, _ := cmd.Flags().GetString("server")

	// Determine connector path
	connectorPath := "."
	if len(args) > 0 {
		connectorPath = args[0]
	}

	// Resolve to absolute path
	absPath, err := filepath.Abs(connectorPath)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}

	// Verify path exists and is a directory
	info, err := os.Stat(absPath)
	if err != nil {
		return fmt.Errorf("path not found: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("path must be a directory: %s", absPath)
	}

	// Check consent
	if !offline && !consent.HasValidConsent() {
		pterm.Warning.Println("AI-assisted analysis requires consent.")
		fmt.Println()
		fmt.Println("To enable AI analysis, run:")
		fmt.Println("  cone connector consent --agree")
		fmt.Println()
		fmt.Println("Running offline analysis instead...")
		fmt.Println()
		offline = true
	}

	if offline {
		return runOfflineAnalysis(ctx, absPath)
	}

	return runOnlineAnalysis(ctx, absPath, mode, dryRun, serverOverride)
}

// runOfflineAnalysis runs basic checks without connecting to C1.
func runOfflineAnalysis(ctx context.Context, connectorPath string) error {
	spinner, _ := pterm.DefaultSpinner.Start("Running offline analysis...")

	// Check for connector configuration files
	checks := []struct {
		name   string
		files  []string
		passed bool
	}{
		{"Configuration file", []string{"connector.yaml", ".baton.yaml", "config.yaml"}, false},
		{"Go module", []string{"go.mod"}, false},
		{"Main package", []string{"main.go", "cmd/baton-*/main.go"}, false},
	}

	for i, check := range checks {
		for _, file := range check.files {
			matches, _ := filepath.Glob(filepath.Join(connectorPath, file))
			if len(matches) > 0 {
				checks[i].passed = true
				break
			}
		}
	}

	spinner.Success("Offline analysis complete")
	fmt.Println()

	// Display results
	fmt.Println("Checks:")
	allPassed := true
	for _, check := range checks {
		if check.passed {
			fmt.Printf("  [PASS] %s\n", check.name)
		} else {
			fmt.Printf("  [FAIL] %s\n", check.name)
			allPassed = false
		}
	}

	fmt.Println()
	if !allPassed {
		fmt.Println("Some checks failed. For full AI analysis, run:")
		fmt.Println("  cone connector consent --agree")
		fmt.Println("  cone connector analyze")
	} else {
		fmt.Println("Basic checks passed. For deeper AI analysis, run:")
		fmt.Println("  cone connector consent --agree")
		fmt.Println("  cone connector analyze")
	}

	return nil
}

// runOnlineAnalysis connects to C1 for AI-assisted analysis.
func runOnlineAnalysis(ctx context.Context, connectorPath, mode string, dryRun bool, serverOverride string) error {
	// Get server URL
	serverURL := serverOverride
	if serverURL == "" {
		// Use configured tenant
		v, err := getSubViperForProfile(nil)
		if err == nil {
			tenant := v.GetString("tenant")
			if tenant != "" {
				serverURL = fmt.Sprintf("https://%s.conductorone.com/api/v1alpha/mcp/cone", tenant)
			}
		}
		if serverURL == "" {
			serverURL = viper.GetString("mcp-server")
		}
		if serverURL == "" {
			return fmt.Errorf("no MCP server configured. Use --server or run 'cone login' first")
		}
	}

	// Get auth token using cone's OAuth credential flow (same as other commands)
	v, err := getSubViperForProfile(nil)
	if err != nil {
		return fmt.Errorf("failed to get config: %w", err)
	}

	clientId, clientSecret, err := getCredentials(v)
	if err != nil {
		return fmt.Errorf("no credentials available. Run 'cone login' first: %w", err)
	}

	tokenSrc, _, _, err := c1client.NewC1TokenSource(ctx, clientId, clientSecret, v.GetString("api-endpoint"), v.GetBool("debug"))
	if err != nil {
		return fmt.Errorf("failed to create token source: %w", err)
	}

	token, err := tokenSrc.Token()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}
	authToken := token.AccessToken

	// Create tool handler
	toolHandler := mcpclient.NewToolHandler(connectorPath)
	toolHandler.DryRun = dryRun

	// Create client
	client := mcpclient.NewClient(serverURL, authToken, toolHandler)

	// Run analysis with timeout
	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	spinner, _ := pterm.DefaultSpinner.Start("Connecting to C1...")

	if err := client.Connect(ctx); err != nil {
		spinner.Fail("Connection failed")
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer client.Close()

	spinner.Success("Connected")

	// Run analysis
	fmt.Println()
	pterm.Info.Printf("Analyzing connector at: %s\n", connectorPath)
	if dryRun {
		pterm.Warning.Println("Dry run mode - no changes will be applied")
	}
	fmt.Println()

	result, err := client.Analyze(ctx, connectorPath, mode)
	if err != nil {
		return fmt.Errorf("analysis failed: %w", err)
	}

	// Display results
	fmt.Println()
	fmt.Println("Analysis Complete")
	fmt.Println("=================")
	fmt.Printf("Status: %s\n", result.Status)
	if result.Message != "" {
		fmt.Printf("Message: %s\n", result.Message)
	}
	if result.FilesScanned > 0 {
		fmt.Printf("Files scanned: %d\n", result.FilesScanned)
	}
	if result.IssuesFound > 0 {
		fmt.Printf("Issues found: %d\n", result.IssuesFound)
	}

	return nil
}
