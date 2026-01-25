package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// connectorPublishCmd creates the publish command for uploading connectors to the registry.
func connectorPublishCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "publish",
		Short: "Publish connector to the ConductorOne registry",
		Long: `Publish a connector version to the ConductorOne registry.

This command performs the following steps:
  1. Reads connector metadata from go.mod and connector.yaml
  2. Finds built binaries in the dist/ directory
  3. Creates a new version in the registry
  4. Uploads binaries for each platform
  5. Uploads checksums
  6. Finalizes the version

Prerequisites:
  - Run 'cone login' first to authenticate
  - Build binaries with 'make build' or 'goreleaser'
  - Have a connector.yaml with metadata (optional but recommended)`,
		Example: `  # Publish from current directory
  cone connector publish --version v1.0.0

  # Publish with specific binary directory
  cone connector publish --version v1.0.0 --dist ./dist

  # Dry run to see what would be published
  cone connector publish --version v1.0.0 --dry-run

  # Publish specific platforms only
  cone connector publish --version v1.0.0 --platform linux-amd64 --platform darwin-arm64`,
		RunE: runConnectorPublish,
	}

	cmd.Flags().String("version", "", "Version to publish (e.g., v1.0.0)")
	cmd.Flags().String("dist", "dist", "Directory containing built binaries")
	cmd.Flags().StringSlice("platform", nil, "Platforms to publish (default: auto-detect)")
	cmd.Flags().Bool("dry-run", false, "Show what would be published without publishing")
	cmd.Flags().String("registry-url", "https://registry.conductorone.com", "Registry API URL")
	cmd.Flags().String("signing-key", "", "Signing key ID for this release")

	cmd.MarkFlagRequired("version")

	return cmd
}

func runConnectorPublish(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// Get flags
	version, _ := cmd.Flags().GetString("version")
	distDir, _ := cmd.Flags().GetString("dist")
	platforms, _ := cmd.Flags().GetStringSlice("platform")
	dryRun, _ := cmd.Flags().GetBool("dry-run")
	registryURL, _ := cmd.Flags().GetString("registry-url")
	signingKey, _ := cmd.Flags().GetString("signing-key")

	// Validate version format
	if !isValidVersion(version) {
		return fmt.Errorf("invalid version format %q, expected semver like v1.0.0", version)
	}

	// Read connector metadata
	metadata, err := readConnectorMetadata()
	if err != nil {
		return fmt.Errorf("failed to read connector metadata: %w", err)
	}

	fmt.Printf("Publishing %s/%s@%s\n", metadata.Org, metadata.Name, version)

	// Find binaries
	binaries, err := findPublishBinaries(distDir, metadata.Name, platforms)
	if err != nil {
		return fmt.Errorf("failed to find binaries: %w", err)
	}

	if len(binaries) == 0 {
		return fmt.Errorf("no binaries found in %s", distDir)
	}

	fmt.Printf("Found %d platform(s):\n", len(binaries))
	for _, b := range binaries {
		fmt.Printf("  - %s (%s, %d bytes)\n", b.Platform, b.Filename, b.Size)
	}

	if dryRun {
		fmt.Println("\nDry run - no changes made")
		return nil
	}

	// Get auth token
	token, err := getAuthToken(ctx, cmd)
	if err != nil {
		return fmt.Errorf("not authenticated, run 'cone login' first: %w", err)
	}

	// Create registry client
	client := newRegistryClient(registryURL, token)

	// Step 0: Ensure connector exists
	fmt.Println("\nEnsuring connector exists...")
	if err := client.EnsureConnector(ctx, metadata.Org, metadata.Name); err != nil {
		return fmt.Errorf("failed to ensure connector exists: %w", err)
	}

	// Step 1: Create version
	fmt.Println("Creating version...")
	platformNames := make([]string, len(binaries))
	for i, b := range binaries {
		platformNames[i] = b.Platform
	}

	createResp, err := client.CreateVersion(ctx, &createVersionRequest{
		Org:          metadata.Org,
		Name:         metadata.Name,
		Version:      version,
		Description:  metadata.Description,
		RepositoryURL: metadata.RepositoryURL,
		HomepageURL:  metadata.HomepageURL,
		License:      metadata.License,
		Changelog:    metadata.Changelog,
		CommitSHA:    getGitCommitSHA(),
		Platforms:    platformNames,
		SigningKeyID: signingKey,
	})
	if err != nil {
		return fmt.Errorf("failed to create version: %w", err)
	}

	fmt.Printf("Created version %s (state: PENDING)\n", version)

	// Step 2: Upload binaries
	fmt.Println("\nUploading binaries...")
	var assetMetadata []assetMeta
	for _, binary := range binaries {
		fmt.Printf("  Uploading %s...", binary.Platform)

		// Get upload URL
		uploadKey := fmt.Sprintf("%s/binary", binary.Platform)
		uploadURL, ok := createResp.UploadURLs[uploadKey]
		if !ok {
			fmt.Println(" SKIP (no upload URL)")
			continue
		}

		// Upload binary
		if err := uploadFile(ctx, uploadURL, binary.Path); err != nil {
			return fmt.Errorf("failed to upload %s: %w", binary.Platform, err)
		}

		// Upload checksum
		checksumKey := fmt.Sprintf("%s/checksum", binary.Platform)
		if checksumURL, ok := createResp.UploadURLs[checksumKey]; ok {
			checksumContent := fmt.Sprintf("%s  %s\n", binary.Checksum, binary.Filename)
			if err := uploadContent(ctx, checksumURL, []byte(checksumContent)); err != nil {
				return fmt.Errorf("failed to upload checksum for %s: %w", binary.Platform, err)
			}
		}

		assetMetadata = append(assetMetadata, assetMeta{
			Platform:  binary.Platform,
			Filename:  binary.Filename,
			SHA256:    binary.Checksum,
			SizeBytes: binary.Size,
			MediaType: "application/octet-stream",
		})

		fmt.Println(" OK")
	}

	// Step 3: Finalize version
	fmt.Println("\nFinalizing version...")
	finalResp, err := client.FinalizeVersion(ctx, &finalizeVersionRequest{
		Org:     metadata.Org,
		Name:    metadata.Name,
		Version: version,
		Assets:  assetMetadata,
	})
	if err != nil {
		return fmt.Errorf("failed to finalize version: %w", err)
	}

	if finalResp.Release.State == "FAILED" {
		return fmt.Errorf("version validation failed: %s", finalResp.Release.FailureReason)
	}

	fmt.Printf("\nPublished %s/%s@%s\n", metadata.Org, metadata.Name, version)
	fmt.Printf("View at: %s/connectors/%s/%s\n", registryURL, metadata.Org, metadata.Name)

	return nil
}

// connectorMetadata holds connector information for publishing.
type connectorMetadata struct {
	Org           string
	Name          string
	Description   string
	RepositoryURL string
	HomepageURL   string
	License       string
	Changelog     string
}

// readConnectorMetadata reads metadata from go.mod and connector.yaml.
func readConnectorMetadata() (*connectorMetadata, error) {
	// Read module path from go.mod
	modulePath, err := readModulePath()
	if err != nil {
		return nil, fmt.Errorf("failed to read go.mod: %w", err)
	}

	// Parse org and name from module path
	// Expected: github.com/org/baton-name
	parts := strings.Split(modulePath, "/")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid module path %q, expected github.com/org/name", modulePath)
	}

	org := parts[len(parts)-2]
	name := parts[len(parts)-1]

	// Strip "baton-" prefix if present for registry name
	registryName := strings.TrimPrefix(name, "baton-")

	metadata := &connectorMetadata{
		Org:  org,
		Name: registryName,
	}

	// Try to read connector.yaml for additional metadata
	if data, err := os.ReadFile("connector.yaml"); err == nil {
		parseConnectorYAML(data, metadata)
	}

	// Default repository URL from module path
	if metadata.RepositoryURL == "" {
		metadata.RepositoryURL = "https://" + modulePath
	}

	return metadata, nil
}

// readModulePath reads the module path from go.mod.
func readModulePath() (string, error) {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return strings.TrimPrefix(line, "module "), nil
		}
	}

	return "", fmt.Errorf("module directive not found in go.mod")
}

// parseConnectorYAML parses connector.yaml into metadata.
// Simple YAML parsing without external dependency.
func parseConnectorYAML(data []byte, metadata *connectorMetadata) {
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "description:") {
			metadata.Description = strings.TrimSpace(strings.TrimPrefix(line, "description:"))
		} else if strings.HasPrefix(line, "license:") {
			metadata.License = strings.TrimSpace(strings.TrimPrefix(line, "license:"))
		} else if strings.HasPrefix(line, "homepage_url:") {
			metadata.HomepageURL = strings.TrimSpace(strings.TrimPrefix(line, "homepage_url:"))
		} else if strings.HasPrefix(line, "repository_url:") {
			metadata.RepositoryURL = strings.TrimSpace(strings.TrimPrefix(line, "repository_url:"))
		}
	}
}

// publishBinary represents a binary to publish.
type publishBinary struct {
	Platform string
	Path     string
	Filename string
	Checksum string
	Size     int64
}

// findPublishBinaries finds binaries in the dist directory.
func findPublishBinaries(distDir, connectorName string, platforms []string) ([]publishBinary, error) {
	var binaries []publishBinary

	// If platforms specified, only look for those
	if len(platforms) > 0 {
		for _, platform := range platforms {
			binary, err := findBinaryForPlatform(distDir, connectorName, platform)
			if err != nil {
				return nil, fmt.Errorf("platform %s: %w", platform, err)
			}
			binaries = append(binaries, *binary)
		}
		return binaries, nil
	}

	// Auto-detect platforms by scanning dist directory
	entries, err := os.ReadDir(distDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			// Check for platform directories (e.g., linux_amd64, darwin_arm64)
			platform := normalizePlatform(entry.Name())
			if platform != "" {
				binary, err := findBinaryForPlatform(distDir, connectorName, platform)
				if err == nil {
					binaries = append(binaries, *binary)
				}
			}
		} else {
			// Check for direct binary files with platform suffix
			name := entry.Name()
			if strings.HasPrefix(name, connectorName) || strings.HasPrefix(name, "baton-"+connectorName) {
				platform := extractPlatformFromFilename(name)
				if platform != "" {
					path := filepath.Join(distDir, name)
					checksum, size, err := computeFileChecksum(path)
					if err != nil {
						continue
					}
					binaries = append(binaries, publishBinary{
						Platform: platform,
						Path:     path,
						Filename: name,
						Checksum: checksum,
						Size:     size,
					})
				}
			}
		}
	}

	return binaries, nil
}

// findBinaryForPlatform finds a specific platform binary.
func findBinaryForPlatform(distDir, connectorName, platform string) (*publishBinary, error) {
	// Try various naming conventions
	patterns := []string{
		filepath.Join(distDir, platform, connectorName),
		filepath.Join(distDir, platform, "baton-"+connectorName),
		filepath.Join(distDir, strings.ReplaceAll(platform, "-", "_"), connectorName),
		filepath.Join(distDir, strings.ReplaceAll(platform, "-", "_"), "baton-"+connectorName),
		filepath.Join(distDir, fmt.Sprintf("%s_%s", connectorName, platform)),
		filepath.Join(distDir, fmt.Sprintf("baton-%s_%s", connectorName, platform)),
	}

	// Add .exe suffix for Windows
	if strings.HasPrefix(platform, "windows") {
		for i := range patterns {
			patterns = append(patterns, patterns[i]+".exe")
		}
	}

	for _, path := range patterns {
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			checksum, size, err := computeFileChecksum(path)
			if err != nil {
				return nil, err
			}
			return &publishBinary{
				Platform: platform,
				Path:     path,
				Filename: filepath.Base(path),
				Checksum: checksum,
				Size:     size,
			}, nil
		}
	}

	return nil, fmt.Errorf("binary not found")
}

// normalizePlatform converts directory names to platform strings.
func normalizePlatform(name string) string {
	// Convert goreleaser-style names (linux_amd64) to registry style (linux-amd64)
	name = strings.ReplaceAll(name, "_", "-")

	// Validate it looks like a platform
	parts := strings.Split(name, "-")
	if len(parts) != 2 {
		return ""
	}

	os := parts[0]
	arch := parts[1]

	validOS := map[string]bool{"linux": true, "darwin": true, "windows": true}
	validArch := map[string]bool{"amd64": true, "arm64": true, "386": true}

	if validOS[os] && validArch[arch] {
		return name
	}
	return ""
}

// extractPlatformFromFilename extracts platform from filename.
func extractPlatformFromFilename(name string) string {
	// Handle patterns like: baton-okta_linux_amd64, baton-okta-linux-amd64
	name = strings.TrimSuffix(name, ".exe")

	for _, sep := range []string{"_", "-"} {
		parts := strings.Split(name, sep)
		if len(parts) >= 3 {
			os := parts[len(parts)-2]
			arch := parts[len(parts)-1]
			platform := normalizePlatform(os + "-" + arch)
			if platform != "" {
				return platform
			}
		}
	}

	return ""
}

// computeFileChecksum computes SHA256 checksum of a file.
func computeFileChecksum(path string) (string, int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", 0, err
	}
	defer f.Close()

	h := sha256.New()
	size, err := io.Copy(h, f)
	if err != nil {
		return "", 0, err
	}

	return hex.EncodeToString(h.Sum(nil)), size, nil
}

// isValidVersion validates semantic version format.
func isValidVersion(v string) bool {
	if !strings.HasPrefix(v, "v") {
		return false
	}
	parts := strings.Split(strings.TrimPrefix(v, "v"), ".")
	if len(parts) < 3 {
		return false
	}
	// Basic validation - just check it starts with v and has dots
	return true
}

// getGitCommitSHA returns the current git commit SHA.
func getGitCommitSHA() string {
	// Try to read from .git/HEAD
	data, err := os.ReadFile(".git/HEAD")
	if err != nil {
		return ""
	}

	content := strings.TrimSpace(string(data))
	if strings.HasPrefix(content, "ref: ") {
		// It's a symbolic ref, read the actual ref
		refPath := strings.TrimPrefix(content, "ref: ")
		data, err = os.ReadFile(filepath.Join(".git", refPath))
		if err != nil {
			return ""
		}
		return strings.TrimSpace(string(data))
	}
	return content
}

// getAuthToken retrieves the auth token for API calls.
func getAuthToken(ctx context.Context, cmd *cobra.Command) (string, error) {
	// TODO: Integrate with cone's existing auth system
	// For now, check environment variable
	token := os.Getenv("CONE_REGISTRY_TOKEN")
	if token == "" {
		token = os.Getenv("C1_TOKEN")
	}
	if token == "" {
		return "", fmt.Errorf("no auth token found, set CONE_REGISTRY_TOKEN or run 'cone login'")
	}
	return token, nil
}

// Registry client types and methods

type registryClient struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

func newRegistryClient(baseURL, token string) *registryClient {
	return &registryClient{
		baseURL:    baseURL,
		token:      token,
		httpClient: &http.Client{},
	}
}

type createConnectorRequest struct {
	Org  string `json:"org"`
	Name string `json:"name"`
}

type createConnectorResponse struct {
	Connector connectorInfo `json:"connector"`
}

type connectorInfo struct {
	Org  string `json:"org"`
	Name string `json:"name"`
}

type createVersionRequest struct {
	Org           string   `json:"org"`
	Name          string   `json:"name"`
	Version       string   `json:"version"`
	Description   string   `json:"description,omitempty"`
	RepositoryURL string   `json:"repository_url,omitempty"`
	HomepageURL   string   `json:"homepage_url,omitempty"`
	License       string   `json:"license,omitempty"`
	Changelog     string   `json:"changelog,omitempty"`
	CommitSHA     string   `json:"commit_sha,omitempty"`
	Platforms     []string `json:"platforms"`
	SigningKeyID  string   `json:"signing_key_id,omitempty"`
}

type createVersionResponse struct {
	Release    releaseManifest   `json:"release"`
	UploadURLs map[string]string `json:"upload_urls"`
}

type releaseManifest struct {
	Org           string `json:"org"`
	Name          string `json:"name"`
	Version       string `json:"version"`
	State         string `json:"state"`
	FailureReason string `json:"failure_reason,omitempty"`
}

type finalizeVersionRequest struct {
	Org     string      `json:"org"`
	Name    string      `json:"name"`
	Version string      `json:"version"`
	Assets  []assetMeta `json:"assets"`
}

type assetMeta struct {
	Platform  string `json:"platform"`
	Filename  string `json:"filename"`
	SHA256    string `json:"sha256"`
	SizeBytes int64  `json:"size_bytes"`
	MediaType string `json:"media_type"`
}

type finalizeVersionResponse struct {
	Release releaseManifest `json:"release"`
}

// EnsureConnector creates the connector if it doesn't exist.
// Returns nil if connector already exists or was created successfully.
func (c *registryClient) EnsureConnector(ctx context.Context, org, name string) error {
	url := fmt.Sprintf("%s/api/v1/connectors", c.baseURL)

	reqBody := createConnectorRequest{Org: org, Name: name}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// 201 Created = new connector
	// 409 Conflict = already exists (that's fine)
	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusConflict {
		return nil
	}

	respBody, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("failed to ensure connector exists (status %d): %s", resp.StatusCode, string(respBody))
}

func (c *registryClient) CreateVersion(ctx context.Context, req *createVersionRequest) (*createVersionResponse, error) {
	url := fmt.Sprintf("%s/api/v1/connectors/%s/%s/versions", c.baseURL, req.Org, req.Name)

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var result createVersionResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

func (c *registryClient) FinalizeVersion(ctx context.Context, req *finalizeVersionRequest) (*finalizeVersionResponse, error) {
	url := fmt.Sprintf("%s/api/v1/connectors/%s/%s/versions/%s/finalize", c.baseURL, req.Org, req.Name, req.Version)

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var result finalizeVersionResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

func uploadFile(ctx context.Context, url, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, f)
	if err != nil {
		return err
	}
	req.ContentLength = info.Size()
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("upload failed with status %d", resp.StatusCode)
	}

	return nil
}

func uploadContent(ctx context.Context, url string, content []byte) error {
	req, err := http.NewRequestWithContext(ctx, "PUT", url, strings.NewReader(string(content)))
	if err != nil {
		return err
	}
	req.ContentLength = int64(len(content))
	req.Header.Set("Content-Type", "text/plain")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("upload failed with status %d", resp.StatusCode)
	}

	return nil
}

// getCurrentPlatform returns the current OS/arch.
func getCurrentPlatform() string {
	return fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
}
