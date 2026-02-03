package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

// connectorDevCmd returns the command for running a local development server.
// It watches for file changes and automatically rebuilds/restarts the connector.
func connectorDevCmd() *cobra.Command {
	var port int
	var noWatch bool

	cmd := &cobra.Command{
		Use:   "dev [path]",
		Short: "Run a connector in development mode with hot reload",
		Long: `Run a connector in development mode with automatic rebuilding on file changes.

This command:
1. Builds the connector
2. Runs it with the specified flags
3. Watches for .go file changes
4. Automatically rebuilds and restarts on changes

Press Ctrl+C to stop.

Examples:
  cone connector dev
  cone connector dev ./my-connector
  cone connector dev --no-watch`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			devPath := "."
			if len(args) > 0 {
				devPath = args[0]
			}

			absPath, err := filepath.Abs(devPath)
			if err != nil {
				return fmt.Errorf("failed to resolve path: %w", err)
			}

			// Check if directory contains go.mod
			goModPath := filepath.Join(absPath, "go.mod")
			if _, err := os.Stat(goModPath); os.IsNotExist(err) {
				return fmt.Errorf("no go.mod found in %s - is this a Go project?", absPath)
			}

			ctx, cancel := context.WithCancel(cmd.Context())
			defer cancel()

			// Handle shutdown signals
			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
			go func() {
				<-sigCh
				fmt.Println("\nShutting down...")
				cancel()
			}()

			if noWatch {
				// Just build and run once
				return buildAndRun(ctx, absPath, port)
			}

			// Watch mode: rebuild on file changes
			return watchAndRun(ctx, absPath, port)
		},
	}

	cmd.Flags().IntVarP(&port, "port", "P", 8080, "Port for the connector to listen on (if applicable)")
	cmd.Flags().BoolVar(&noWatch, "no-watch", false, "Disable file watching (run once)")

	return cmd
}

// buildAndRun builds the connector and runs it.
func buildAndRun(ctx context.Context, path string, port int) error {
	binaryPath := filepath.Join(path, "connector-dev")

	// Build
	fmt.Println("Building connector...")
	buildCmd := exec.CommandContext(ctx, "go", "build", "-o", binaryPath, ".")
	buildCmd.Dir = path
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr

	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	// Run
	fmt.Printf("Starting connector (port %d)...\n", port)
	runCmd := exec.CommandContext(ctx, binaryPath)
	runCmd.Dir = path
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	runCmd.Env = append(os.Environ(), fmt.Sprintf("PORT=%d", port))

	if err := runCmd.Run(); err != nil {
		// Context cancelled is expected on shutdown
		if ctx.Err() != nil {
			return nil
		}
		return fmt.Errorf("connector exited with error: %w", err)
	}

	return nil
}

// watchAndRun watches for file changes and rebuilds/restarts the connector.
func watchAndRun(ctx context.Context, path string, port int) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create file watcher: %w", err)
	}
	defer watcher.Close()

	// Add directories to watch
	if err := addWatchDirs(watcher, path); err != nil {
		return fmt.Errorf("failed to watch directories: %w", err)
	}

	var runCmd *exec.Cmd
	var runCancel context.CancelFunc
	binaryPath := filepath.Join(path, "connector-dev")

	// Initial build and run
	build := func() error {
		fmt.Println("\n[dev] Building connector...")
		buildCmd := exec.CommandContext(ctx, "go", "build", "-o", binaryPath, ".")
		buildCmd.Dir = path
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr

		if err := buildCmd.Run(); err != nil {
			fmt.Printf("[dev] Build failed: %v\n", err)
			return err
		}
		fmt.Println("[dev] Build successful")
		return nil
	}

	start := func() {
		// Stop previous run if any
		if runCancel != nil {
			runCancel()
		}
		if runCmd != nil && runCmd.Process != nil {
			_ = runCmd.Process.Kill()
			_ = runCmd.Wait()
		}

		fmt.Printf("[dev] Starting connector (port %d)...\n", port)
		var runCtx context.Context
		runCtx, runCancel = context.WithCancel(ctx)
		runCmd = exec.CommandContext(runCtx, binaryPath)
		runCmd.Dir = path
		runCmd.Stdout = os.Stdout
		runCmd.Stderr = os.Stderr
		runCmd.Env = append(os.Environ(), fmt.Sprintf("PORT=%d", port))

		go func() {
			if err := runCmd.Run(); err != nil && runCtx.Err() == nil {
				fmt.Printf("[dev] Connector exited: %v\n", err)
			}
		}()
	}

	// Initial build and run
	if err := build(); err != nil {
		fmt.Println("[dev] Initial build failed, waiting for file changes...")
	} else {
		start()
	}

	// Debounce timer for file changes
	var debounceTimer *time.Timer
	debounce := func() {
		if debounceTimer != nil {
			debounceTimer.Stop()
		}
		debounceTimer = time.AfterFunc(500*time.Millisecond, func() {
			if err := build(); err == nil {
				start()
			}
		})
	}

	fmt.Println("[dev] Watching for file changes... (Ctrl+C to stop)")

	for {
		select {
		case <-ctx.Done():
			if runCancel != nil {
				runCancel()
			}
			// Clean up binary
			_ = os.Remove(binaryPath)
			return nil

		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			// Only watch .go files
			if filepath.Ext(event.Name) == ".go" {
				if event.Op&(fsnotify.Write|fsnotify.Create) != 0 {
					fmt.Printf("[dev] Change detected: %s\n", filepath.Base(event.Name))
					debounce()
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			fmt.Printf("[dev] Watch error: %v\n", err)
		}
	}
}

// addWatchDirs adds all directories containing .go files to the watcher.
func addWatchDirs(watcher *fsnotify.Watcher, root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip vendor, .git, and other hidden directories
		if info.IsDir() {
			name := info.Name()
			if name == "vendor" || name == ".git" || (name[0] == '.' && len(name) > 1) {
				return filepath.SkipDir
			}
			return watcher.Add(path)
		}

		return nil
	})
}
