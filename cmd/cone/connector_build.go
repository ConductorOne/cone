package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

// connectorBuildCmd returns the command for building connector binaries.
func connectorBuildCmd() *cobra.Command {
	var outputPath string
	var targetOS string
	var targetArch string

	cmd := &cobra.Command{
		Use:   "build [path]",
		Short: "Build a connector binary",
		Long: `Build a connector binary from the specified path.

If no path is provided, builds from the current directory.

Examples:
  cone connector build
  cone connector build ./my-connector
  cone connector build -o ./dist/connector
  cone connector build --os linux --arch amd64`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			buildPath := "."
			if len(args) > 0 {
				buildPath = args[0]
			}

			// Resolve absolute path
			absPath, err := filepath.Abs(buildPath)
			if err != nil {
				return fmt.Errorf("failed to resolve path: %w", err)
			}

			// Check if directory exists and contains go.mod
			goModPath := filepath.Join(absPath, "go.mod")
			if _, err := os.Stat(goModPath); os.IsNotExist(err) {
				return fmt.Errorf("no go.mod found in %s - is this a Go project?", absPath)
			}

			// Determine output path
			if outputPath == "" {
				outputPath = filepath.Join(absPath, "connector")
				if targetOS == "windows" || runtime.GOOS == "windows" {
					outputPath += ".exe"
				}
			}

			// Set up build environment
			buildEnv := os.Environ()
			if targetOS != "" {
				buildEnv = append(buildEnv, "GOOS="+targetOS)
			}
			if targetArch != "" {
				buildEnv = append(buildEnv, "GOARCH="+targetArch)
			}

			// Build the connector
			// Template creates main.go at root, so build from "."
			buildCmd := exec.CommandContext(cmd.Context(), "go", "build", "-o", outputPath, ".")
			buildCmd.Dir = absPath
			buildCmd.Env = buildEnv
			buildCmd.Stdout = os.Stdout
			buildCmd.Stderr = os.Stderr

			fmt.Printf("Building connector in %s...\n", absPath)
			if err := buildCmd.Run(); err != nil {
				return fmt.Errorf("build failed: %w", err)
			}

			fmt.Printf("Built: %s\n", outputPath)
			return nil
		},
	}

	cmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output path for the binary")
	cmd.Flags().StringVar(&targetOS, "os", "", "Target operating system (e.g., linux, darwin, windows)")
	cmd.Flags().StringVar(&targetArch, "arch", "", "Target architecture (e.g., amd64, arm64)")

	return cmd
}
