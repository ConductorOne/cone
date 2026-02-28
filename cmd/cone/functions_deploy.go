package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func functionsDeployCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy [directory]",
		Short: "Deploy a C1 Function to ConductorOne",
		Long: `Uploads and deploys a function to ConductorOne.

Uses your existing cone credentials (from 'cone login').
If no directory is specified, uses the current directory.`,
		RunE: functionsDeployRun,
	}

	cmd.Flags().String("function-id", "", "Function ID to deploy to (required for updates)")
	cmd.Flags().String("name", "", "Function name (required for new functions)")
	cmd.Flags().String("description", "", "Function description")
	cmd.Flags().Bool("dry-run", false, "Show what would be deployed without deploying")

	return cmd
}

func functionsDeployRun(cmd *cobra.Command, args []string) error {
	// Determine source directory
	sourceDir := "."
	if len(args) > 0 {
		sourceDir = args[0]
	}

	absDir, err := filepath.Abs(sourceDir)
	if err != nil {
		return fmt.Errorf("failed to resolve directory: %w", err)
	}

	// Verify main.ts or index.ts exists
	mainFile := findMainFile(absDir)
	if mainFile == "" {
		return fmt.Errorf("no main.ts or index.ts found in %s", absDir)
	}

	dryRun, _ := cmd.Flags().GetBool("dry-run")
	functionID, _ := cmd.Flags().GetString("function-id")
	functionName, _ := cmd.Flags().GetString("name")

	// Build the deployment bundle
	spinner, _ := pterm.DefaultSpinner.Start("Building deployment bundle...")

	bundle, hash, err := buildDeploymentBundle(absDir)
	if err != nil {
		spinner.Fail(err)
		return err
	}

	spinner.Success(fmt.Sprintf("Bundle built: %d bytes, sha256:%s", len(bundle), hash[:12]))

	if dryRun {
		pterm.Info.Println("Dry run - not deploying")
		pterm.Info.Printf("Would deploy %d bytes to function %s\n", len(bundle), functionID)
		return nil
	}

	// Get client
	_, c1Client, _, err := cmdContext(cmd)
	if err != nil {
		return fmt.Errorf("not logged in - run 'cone login' first: %w", err)
	}

	_ = c1Client // Will be used when we add the SDK Functions API

	spinner, _ = pterm.DefaultSpinner.Start("Deploying function...")

	// TODO: Call FunctionsService.CreateCommit with the bundle
	// This requires the SDK to expose Functions APIs
	// For now, we'll show what would happen

	if functionID == "" && functionName == "" {
		spinner.Fail("either --function-id or --name is required")
		return fmt.Errorf("either --function-id (for updates) or --name (for new functions) is required")
	}

	spinner.Warning("Deploy API integration pending - SDK support needed")
	pterm.Info.Printf("Would deploy bundle (sha256:%s) to function %s\n", hash[:12], functionID)

	return nil
}

func findMainFile(dir string) string {
	for _, name := range []string{"main.ts", "index.ts", "mod.ts"} {
		path := filepath.Join(dir, name)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return ""
}

func buildDeploymentBundle(dir string) ([]byte, string, error) {
	var buf bytes.Buffer
	gzw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gzw)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden files and directories (except .gitignore)
		name := info.Name()
		if strings.HasPrefix(name, ".") && name != ".gitignore" {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip node_modules, .git, etc.
		if info.IsDir() {
			switch name {
			case "node_modules", ".git", "__pycache__", ".vscode":
				return filepath.SkipDir
			}
			return nil
		}

		// Only include TypeScript and JSON files
		ext := filepath.Ext(name)
		if ext != ".ts" && ext != ".json" && name != ".gitignore" {
			return nil
		}

		// Get relative path
		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		// Create tar header
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		header.Name = relPath

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// Write file contents
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(tw, f)
		return err
	})

	if err != nil {
		return nil, "", fmt.Errorf("failed to build bundle: %w", err)
	}

	if err := tw.Close(); err != nil {
		return nil, "", err
	}
	if err := gzw.Close(); err != nil {
		return nil, "", err
	}

	data := buf.Bytes()
	hash := sha256.Sum256(data)

	return data, hex.EncodeToString(hash[:]), nil
}
