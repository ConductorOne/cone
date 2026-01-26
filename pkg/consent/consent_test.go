package consent

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// setupTestDir creates a temp directory and sets HOME to point to it.
// Returns a cleanup function that restores the original HOME.
func setupTestDir(t *testing.T) func() {
	t.Helper()

	originalHome := os.Getenv("HOME")
	tmpDir := t.TempDir()
	if err := os.Setenv("HOME", tmpDir); err != nil {
		t.Fatalf("failed to set HOME: %v", err)
	}

	return func() {
		os.Setenv("HOME", originalHome)
	}
}

func TestConsentDir(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	dir, err := consentDir()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !contains(dir, ".cone") {
		t.Errorf("expected dir to contain .cone, got %s", dir)
	}
}

func TestConsentFilePath(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	path, err := consentFilePath()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !contains(path, "consent.json") {
		t.Errorf("expected path to contain consent.json, got %s", path)
	}
	if !contains(path, ".cone") {
		t.Errorf("expected path to contain .cone, got %s", path)
	}
}

func TestEnsureConsentDir(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	if err := ensureConsentDir(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	dir, err := consentDir()
	if err != nil {
		t.Fatalf("failed to get consent dir: %v", err)
	}

	info, err := os.Stat(dir)
	if err != nil {
		t.Fatalf("failed to stat directory: %v", err)
	}
	if !info.IsDir() {
		t.Error("expected directory, got file")
	}
	if info.Mode().Perm() != 0700 {
		t.Errorf("expected permissions 0700, got %o", info.Mode().Perm())
	}
}

func TestLoad_NoConsent(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	record, err := Load()
	if record != nil {
		t.Error("expected nil record")
	}
	if err != ErrNoConsent {
		t.Errorf("expected ErrNoConsent, got %v", err)
	}
}

func TestLoad_ValidConsent(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	// Create consent file manually
	if err := ensureConsentDir(); err != nil {
		t.Fatalf("failed to ensure dir: %v", err)
	}

	path, err := consentFilePath()
	if err != nil {
		t.Fatalf("failed to get path: %v", err)
	}

	record := ConsentRecord{
		ConsentedAt: time.Now().UTC(),
		Version:     CurrentConsentVersion,
	}
	data, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	// Load and verify
	loaded, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if loaded.Version != CurrentConsentVersion {
		t.Errorf("expected version %s, got %s", CurrentConsentVersion, loaded.Version)
	}
	if loaded.ConsentedAt.IsZero() {
		t.Error("expected non-zero consented_at")
	}
}

func TestLoad_VersionMismatch(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	// Create consent file with old version
	if err := ensureConsentDir(); err != nil {
		t.Fatalf("failed to ensure dir: %v", err)
	}

	path, err := consentFilePath()
	if err != nil {
		t.Fatalf("failed to get path: %v", err)
	}

	record := ConsentRecord{
		ConsentedAt: time.Now().UTC(),
		Version:     "0.9", // Old version
	}
	data, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	// Load should return version mismatch error
	loaded, err := Load()
	if loaded == nil {
		t.Error("expected record to be returned even on version mismatch")
	}
	if err != ErrConsentVersionMismatch {
		t.Errorf("expected ErrConsentVersionMismatch, got %v", err)
	}
	if loaded != nil && loaded.Version != "0.9" {
		t.Errorf("expected version 0.9, got %s", loaded.Version)
	}
}

func TestLoad_InvalidJSON(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	if err := ensureConsentDir(); err != nil {
		t.Fatalf("failed to ensure dir: %v", err)
	}

	path, err := consentFilePath()
	if err != nil {
		t.Fatalf("failed to get path: %v", err)
	}

	if err := os.WriteFile(path, []byte("not valid json"), 0600); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	record, err := Load()
	if record != nil {
		t.Error("expected nil record for invalid JSON")
	}
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
	if !contains(err.Error(), "failed to parse consent file") {
		t.Errorf("expected parse error, got: %v", err)
	}
}

func TestHasValidConsent_NoConsent(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	if HasValidConsent() {
		t.Error("expected no valid consent")
	}
}

func TestHasValidConsent_WithConsent(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	if err := Save(); err != nil {
		t.Fatalf("failed to save: %v", err)
	}

	if !HasValidConsent() {
		t.Error("expected valid consent")
	}
}

func TestSave(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	if err := Save(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify file exists with correct permissions
	path, err := consentFilePath()
	if err != nil {
		t.Fatalf("failed to get path: %v", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("failed to stat file: %v", err)
	}
	if info.Mode().Perm() != 0600 {
		t.Errorf("expected permissions 0600, got %o", info.Mode().Perm())
	}

	// Verify content
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	var record ConsentRecord
	if err := json.Unmarshal(data, &record); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	if record.Version != CurrentConsentVersion {
		t.Errorf("expected version %s, got %s", CurrentConsentVersion, record.Version)
	}
	if record.ConsentedAt.IsZero() {
		t.Error("expected non-zero consented_at")
	}
}

func TestRevoke(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	// Save then revoke
	if err := Save(); err != nil {
		t.Fatalf("failed to save: %v", err)
	}
	if !HasValidConsent() {
		t.Error("expected valid consent after save")
	}

	if err := Revoke(); err != nil {
		t.Fatalf("failed to revoke: %v", err)
	}
	if HasValidConsent() {
		t.Error("expected no valid consent after revoke")
	}
}

func TestRevoke_NoConsent(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	// Revoking when no consent exists should not error
	if err := Revoke(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestStatus_NotConsented(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	status := Status()
	if status != "Not consented" {
		t.Errorf("expected 'Not consented', got %s", status)
	}
}

func TestStatus_Consented(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	if err := Save(); err != nil {
		t.Fatalf("failed to save: %v", err)
	}

	status := Status()
	if !contains(status, "Consented on") {
		t.Errorf("expected status to contain 'Consented on', got %s", status)
	}
	if !contains(status, CurrentConsentVersion) {
		t.Errorf("expected status to contain version %s, got %s", CurrentConsentVersion, status)
	}
}

func TestStatus_VersionMismatch(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	// Create old version consent
	if err := ensureConsentDir(); err != nil {
		t.Fatalf("failed to ensure dir: %v", err)
	}

	path, err := consentFilePath()
	if err != nil {
		t.Fatalf("failed to get path: %v", err)
	}

	record := ConsentRecord{
		ConsentedAt: time.Now().UTC(),
		Version:     "0.9",
	}
	data, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	status := Status()
	if !contains(status, "outdated") {
		t.Errorf("expected status to contain 'outdated', got %s", status)
	}
	if !contains(status, "0.9") {
		t.Errorf("expected status to contain '0.9', got %s", status)
	}
}

func TestConsentText(t *testing.T) {
	text := ConsentText()
	if !contains(text, "AI-Assisted Connector Analysis") {
		t.Error("expected consent text to contain 'AI-Assisted Connector Analysis'")
	}
	if !contains(text, "ConductorOne") {
		t.Error("expected consent text to contain 'ConductorOne'")
	}
	if !contains(text, "privacy") {
		t.Error("expected consent text to contain 'privacy'")
	}
}

func TestCurrentConsentVersion(t *testing.T) {
	if CurrentConsentVersion == "" {
		t.Error("expected non-empty consent version")
	}
	if CurrentConsentVersion != "1.0" {
		t.Errorf("expected version 1.0, got %s", CurrentConsentVersion)
	}
}

func TestDirectoryPermissions(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	// Create directory with wrong permissions
	dir, err := consentDir()
	if err != nil {
		t.Fatalf("failed to get dir: %v", err)
	}

	if err := os.MkdirAll(dir, 0755); err != nil { // Wrong permissions
		t.Fatalf("failed to create dir: %v", err)
	}

	// ensureConsentDir should fix them
	if err := ensureConsentDir(); err != nil {
		t.Fatalf("failed to ensure dir: %v", err)
	}

	info, err := os.Stat(dir)
	if err != nil {
		t.Fatalf("failed to stat dir: %v", err)
	}
	if info.Mode().Perm() != 0700 {
		t.Errorf("expected permissions 0700, got %o", info.Mode().Perm())
	}
}

func TestConsentRecordSerialization(t *testing.T) {
	now := time.Date(2026, 1, 25, 12, 0, 0, 0, time.UTC)
	record := ConsentRecord{
		ConsentedAt: now,
		Version:     "1.0",
	}

	data, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded ConsentRecord
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if decoded.Version != record.Version {
		t.Errorf("expected version %s, got %s", record.Version, decoded.Version)
	}
	if !record.ConsentedAt.Equal(decoded.ConsentedAt) {
		t.Errorf("timestamps don't match: %v vs %v", record.ConsentedAt, decoded.ConsentedAt)
	}
}

func TestConsentDirCreatesParentDirectories(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	// Verify the directory doesn't exist yet
	dir, err := consentDir()
	if err != nil {
		t.Fatalf("failed to get dir: %v", err)
	}
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		t.Error("expected directory to not exist initially")
	}

	// Save should create the directory
	if err := Save(); err != nil {
		t.Fatalf("failed to save: %v", err)
	}

	// Verify directory now exists
	info, err := os.Stat(dir)
	if err != nil {
		t.Fatalf("failed to stat dir: %v", err)
	}
	if !info.IsDir() {
		t.Error("expected directory")
	}
}

func TestMultipleSaveOverwrites(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	// Save twice
	if err := Save(); err != nil {
		t.Fatalf("first save failed: %v", err)
	}

	time.Sleep(10 * time.Millisecond) // Ensure different timestamp

	if err := Save(); err != nil {
		t.Fatalf("second save failed: %v", err)
	}

	// Should still be valid
	if !HasValidConsent() {
		t.Error("expected valid consent")
	}

	// Only one file should exist
	dir, err := consentDir()
	if err != nil {
		t.Fatalf("failed to get dir: %v", err)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("failed to read dir: %v", err)
	}
	if len(entries) != 1 {
		t.Errorf("expected 1 entry, got %d", len(entries))
	}
	if entries[0].Name() != "consent.json" {
		t.Errorf("expected consent.json, got %s", entries[0].Name())
	}
}

func TestConsentFileInSubdirectory(t *testing.T) {
	cleanup := setupTestDir(t)
	defer cleanup()

	path, err := consentFilePath()
	if err != nil {
		t.Fatalf("failed to get path: %v", err)
	}

	// Should be in .cone subdirectory
	dir := filepath.Dir(path)
	if filepath.Base(dir) != ".cone" {
		t.Errorf("expected parent dir to be .cone, got %s", filepath.Base(dir))
	}
}

// contains is a helper for string containment checks
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
