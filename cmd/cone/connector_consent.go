package main

import (
	"fmt"

	"github.com/conductorone/cone/pkg/consent"
	"github.com/conductorone/cone/pkg/prompt"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func connectorConsentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "consent",
		Short: "Manage consent for AI-assisted connector analysis",
		Long: `Manage your consent for AI-assisted connector analysis features.

AI-assisted analysis sends your connector source code to ConductorOne's
AI copilot for review and suggestions. This requires explicit consent.

Without any flags, displays the current consent status.

Examples:
  cone connector consent           # Check consent status
  cone connector consent --agree   # Grant consent (interactive)
  cone connector consent --revoke  # Revoke consent
  cone connector consent --status  # Explicit status check`,
		RunE: runConnectorConsent,
	}

	cmd.Flags().Bool("agree", false, "Grant consent for AI-assisted analysis (requires interactive terminal)")
	cmd.Flags().Bool("revoke", false, "Revoke consent for AI-assisted analysis")
	cmd.Flags().Bool("status", false, "Display consent status")

	return cmd
}

func runConnectorConsent(cmd *cobra.Command, args []string) error {
	agree, _ := cmd.Flags().GetBool("agree")
	revoke, _ := cmd.Flags().GetBool("revoke")
	status, _ := cmd.Flags().GetBool("status")

	// Validate mutually exclusive flags
	flagCount := 0
	if agree {
		flagCount++
	}
	if revoke {
		flagCount++
	}
	if status {
		flagCount++
	}
	if flagCount > 1 {
		return fmt.Errorf("only one of --agree, --revoke, or --status can be specified")
	}

	// Handle revoke
	if revoke {
		if err := consent.Revoke(); err != nil {
			return fmt.Errorf("failed to revoke consent: %w", err)
		}
		pterm.Success.Println("Consent revoked. AI-assisted analysis is now disabled.")
		return nil
	}

	// Handle status (explicit or default)
	if status || (!agree && !revoke) {
		fmt.Printf("Consent status: %s\n", consent.Status())
		return nil
	}

	// Handle agree
	if agree {
		return grantConsent()
	}

	return nil
}

func grantConsent() error {
	// Require interactive terminal for consent
	if !prompt.IsInteractive() {
		return fmt.Errorf("--agree requires an interactive terminal; cannot grant consent in non-interactive mode")
	}

	// Check if already consented
	if consent.HasValidConsent() {
		pterm.Info.Println("You have already consented to AI-assisted analysis.")
		fmt.Printf("Current status: %s\n", consent.Status())
		return nil
	}

	// Display consent dialog
	fmt.Println()
	prompt.DisplayBox("AI-Assisted Analysis Consent", consent.ConsentText())
	fmt.Println()

	// Prompt for confirmation
	confirmed, err := prompt.Confirm("Do you consent to AI-assisted analysis?")
	if err != nil {
		return fmt.Errorf("failed to get confirmation: %w", err)
	}

	if !confirmed {
		pterm.Warning.Println("Consent not granted. AI-assisted analysis remains disabled.")
		return nil
	}

	// Save consent
	if err := consent.Save(); err != nil {
		return fmt.Errorf("failed to save consent: %w", err)
	}

	pterm.Success.Println("Consent granted. AI-assisted analysis is now enabled.")
	fmt.Printf("You can revoke consent at any time with: cone connector consent --revoke\n")

	return nil
}
