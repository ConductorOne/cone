// Package consent manages user consent for AI-assisted features that send code to C1.
//
// Security rationale for design decisions:
// - Consent stored in ~/.cone/consent.json (separate from ~/.conductorone/ credentials)
// - File permissions: 0600 (user-only read/write) to prevent other users from modifying
// - Version tracking enables re-prompting when consent terms change
// - Requires interactive terminal for --agree to prevent scripted consent bypass
package consent

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// CurrentConsentVersion should be incremented when consent text changes materially.
// This triggers re-prompting users who consented to a previous version.
const CurrentConsentVersion = "1.0"

// ConsentRecord stores the user's consent decision.
type ConsentRecord struct {
	ConsentedAt time.Time `json:"consented_at"`
	Version     string    `json:"version"`
}

// ErrNoConsent is returned when the user has not given consent.
var ErrNoConsent = errors.New("consent: user has not consented to AI-assisted analysis")

// ErrConsentVersionMismatch is returned when consent version is outdated.
var ErrConsentVersionMismatch = errors.New("consent: consent version has changed, re-consent required")

// consentDir returns the path to the cone config directory.
func consentDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(home, ".cone"), nil
}

// consentFilePath returns the path to the consent file.
func consentFilePath() (string, error) {
	dir, err := consentDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "consent.json"), nil
}

// ensureConsentDir ensures ~/.cone directory exists with correct permissions.
// Security: 0700 permissions (rwx------) prevent other users from listing contents.
func ensureConsentDir() error {
	dir, err := consentDir()
	if err != nil {
		return err
	}

	// Create with restrictive permissions (0700 = rwx------)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("failed to create .cone directory: %w", err)
	}

	// Verify permissions in case directory already existed with wrong perms
	info, err := os.Stat(dir)
	if err != nil {
		return err
	}
	if info.Mode().Perm() != 0700 {
		if err := os.Chmod(dir, 0700); err != nil {
			return fmt.Errorf("failed to set directory permissions: %w", err)
		}
	}

	return nil
}

// Load reads the consent record from disk.
// Returns ErrNoConsent if no consent file exists.
// Returns ErrConsentVersionMismatch if consent version doesn't match current.
func Load() (*ConsentRecord, error) {
	path, err := consentFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoConsent
		}
		return nil, fmt.Errorf("failed to read consent file: %w", err)
	}

	var record ConsentRecord
	if err := json.Unmarshal(data, &record); err != nil {
		return nil, fmt.Errorf("failed to parse consent file: %w", err)
	}

	if record.Version != CurrentConsentVersion {
		return &record, ErrConsentVersionMismatch
	}

	return &record, nil
}

// HasValidConsent returns true if the user has valid, current consent.
func HasValidConsent() bool {
	_, err := Load()
	return err == nil
}

// Save writes a new consent record to disk.
// Security: File written with 0600 permissions (rw-------).
func Save() error {
	if err := ensureConsentDir(); err != nil {
		return err
	}

	path, err := consentFilePath()
	if err != nil {
		return err
	}

	record := ConsentRecord{
		ConsentedAt: time.Now().UTC(),
		Version:     CurrentConsentVersion,
	}

	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal consent record: %w", err)
	}

	// Write with restrictive permissions (0600 = rw-------)
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("failed to write consent file: %w", err)
	}

	return nil
}

// Revoke removes the consent record from disk.
func Revoke() error {
	path, err := consentFilePath()
	if err != nil {
		return err
	}

	if err := os.Remove(path); err != nil {
		if os.IsNotExist(err) {
			return nil // Already revoked
		}
		return fmt.Errorf("failed to remove consent file: %w", err)
	}

	return nil
}

// Status returns a human-readable string describing consent status.
func Status() string {
	record, err := Load()
	if err != nil {
		if errors.Is(err, ErrNoConsent) {
			return "Not consented"
		}
		if errors.Is(err, ErrConsentVersionMismatch) {
			return fmt.Sprintf("Consent outdated (v%s, current is v%s)", record.Version, CurrentConsentVersion)
		}
		return fmt.Sprintf("Error checking consent: %v", err)
	}
	return fmt.Sprintf("Consented on %s (v%s)", record.ConsentedAt.Format(time.RFC3339), record.Version)
}

// ConsentText returns the full consent text to display to users.
func ConsentText() string {
	return `AI-Assisted Connector Analysis Consent

This command sends your connector source code to ConductorOne for AI analysis.

What happens:
  - Your connector code is sent to ConductorOne's AI copilot
  - The AI analyzes your code and suggests improvements
  - Your code is processed in memory and is NOT stored permanently
  - Analysis results are returned to your local machine

Your code is:
  - Processed only for the duration of the analysis
  - Not used for AI training
  - Not shared with third parties
  - Subject to ConductorOne's privacy policy

For more information, see: https://www.conductorone.com/privacy

Do you consent to AI-assisted analysis of your connector code?`
}
