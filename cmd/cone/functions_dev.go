package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/conductorone/dpop/pkg/dpop"
	"github.com/go-jose/go-jose/v4"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/conductorone/cone/pkg/client"
	"github.com/conductorone/cone/pkg/proxy"
)

func functionsDevCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dev [directory]",
		Short: "Run a local development server for C1 Functions",
		Long: `Starts a local development server with production parity.

The server runs your TypeScript function through a MITM proxy that:
- Injects OAuth2 + DPoP authentication for C1 API calls
- Enforces the egress allowlist from deno.json or C1 API
- Injects secrets from C1 API as environment variables
- Uses restricted Deno permissions matching production

If no directory is specified, uses the current directory.

Requires 'cone login' for C1 API authentication.
Requires 'deno' to be installed (https://deno.land).`,
		RunE: functionsDevRun,
	}

	cmd.Flags().IntP("port", "P", 8000, "Port to listen on")
	cmd.Flags().BoolP("watch", "w", true, "Enable hot reload on file changes")
	cmd.Flags().Bool("no-proxy", false, "Disable MITM proxy (no auth injection or allowlist)")
	cmd.Flags().String("function-id", "", "Function ID for fetching secrets and config from C1")

	return cmd
}

// denoConfig represents relevant fields from deno.json
type denoConfig struct {
	Permissions struct {
		Net []string `json:"net"`
	} `json:"permissions"`
}

func functionsDevRun(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// Check if deno is installed
	denoPath, err := exec.LookPath("deno")
	if err != nil {
		return fmt.Errorf("deno not found in PATH - install from https://deno.land")
	}

	// Determine source directory
	sourceDir := "."
	if len(args) > 0 {
		sourceDir = args[0]
	}

	// Convert to absolute path
	absDir, err := filepath.Abs(sourceDir)
	if err != nil {
		return fmt.Errorf("failed to resolve directory: %w", err)
	}

	// Check directory exists
	info, err := os.Stat(absDir)
	if err != nil {
		return fmt.Errorf("directory not found: %s", absDir)
	}
	if !info.IsDir() {
		return fmt.Errorf("not a directory: %s", absDir)
	}

	// Find main file
	mainFile := findMainFile(absDir)
	if mainFile == "" {
		return fmt.Errorf("no main.ts or index.ts found in %s", absDir)
	}

	port, _ := cmd.Flags().GetInt("port")
	watch, _ := cmd.Flags().GetBool("watch")
	noProxy, _ := cmd.Flags().GetBool("no-proxy")
	functionID, _ := cmd.Flags().GetString("function-id")

	// Set up environment
	env := os.Environ()
	env = append(env, fmt.Sprintf("PORT=%d", port))

	var proxyInstance *proxy.Proxy

	if !noProxy {
		// Load allowlist from deno.json
		allowlist := loadAllowlist(absDir)

		// Get C1 credentials
		v, err := getSubViperForProfile(cmd)
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		var proofer *dpop.Proofer
		var c1Host string
		var secrets map[string]string
		var proxyConfig proxy.Config

		clientID, clientSecret, err := getCredentials(v)
		if err != nil {
			pterm.Warning.Println("No credentials found - C1 API calls will not be authenticated")
			pterm.Warning.Println("Run 'cone login' to authenticate")
		} else {
			ts, _, tokenHost, err := client.NewC1TokenSource(ctx,
				clientID, clientSecret,
				v.GetString("api-endpoint"),
				v.GetBool("debug"),
			)
			if err != nil {
				pterm.Warning.Printf("Failed to create token source: %v\n", err)
			} else {
				proxyConfig.TokenSource = ts
				c1Host = tokenHost

				// Create DPoP proofer
				proofer, err = createDPoPProofer()
				if err != nil {
					pterm.Warning.Printf("Failed to create DPoP proofer: %v\n", err)
				}

				// Fetch secrets if function ID provided
				if functionID != "" {
					c1Client, err := client.New(ctx, clientID, clientSecret, v, "functions:dev")
					if err == nil {
						fetchedSecrets, fetchedAllowlist, err := proxy.FetchSecrets(ctx, c1Client.SDK(), functionID)
						if err != nil {
							pterm.Warning.Printf("Failed to fetch secrets: %v\n", err)
						} else {
							secrets = fetchedSecrets
							if len(fetchedAllowlist) > 0 {
								// Merge with local allowlist, preferring remote
								allowlist = append(fetchedAllowlist, allowlist...)
							}
							if len(secrets) > 0 {
								pterm.Info.Printf("Loaded %d secrets from C1\n", len(secrets))
							}
						}
					}
				}
			}
		}

		// Get config directory for CA storage
		configDir := defaultConfigPath()

		// Load or create CA
		ca, err := proxy.LoadOrCreateCA(configDir)
		if err != nil {
			return fmt.Errorf("failed to initialize CA: %w", err)
		}

		// Complete proxy config
		proxyConfig.Allowlist = proxy.ParseAllowlist(allowlist)
		proxyConfig.C1APIHost = c1Host
		proxyConfig.Secrets = secrets
		proxyConfig.SourceDir = absDir
		proxyConfig.Proofer = proofer

		// Create and start proxy
		proxyInstance, err = proxy.New(ca, proxyConfig)
		if err != nil {
			return fmt.Errorf("failed to create proxy: %w", err)
		}

		proxyAddr, err := proxyInstance.Start(ctx)
		if err != nil {
			return fmt.Errorf("failed to start proxy: %w", err)
		}
		defer proxyInstance.Stop()

		pterm.Info.Printf("MITM proxy started on %s\n", proxyAddr)
		if len(proxyConfig.Allowlist) > 0 {
			pterm.Info.Printf("Egress allowlist: %v\n", proxyConfig.Allowlist)
		} else {
			pterm.Info.Println("Egress allowlist: all hosts allowed")
		}
		if proxyConfig.TokenSource != nil {
			pterm.Info.Printf("Auth injection enabled for %s\n", c1Host)
			if proofer != nil {
				pterm.Info.Println("DPoP proofs enabled")
			}
		}

		// Set proxy environment variables
		env = append(env, fmt.Sprintf("HTTPS_PROXY=http://%s", proxyAddr))
		env = append(env, fmt.Sprintf("HTTP_PROXY=http://%s", proxyAddr))
		env = append(env, fmt.Sprintf("DENO_CERT=%s", ca.CertPath()))

		// Inject secrets as environment variables
		for key, value := range secrets {
			env = append(env, fmt.Sprintf("%s=%s", strings.ToUpper(key), value))
		}
	}

	// Build deno command with restricted permissions (matching production)
	denoArgs := []string{"run"}

	// Production-like permissions
	if !noProxy {
		denoArgs = append(denoArgs,
			"--allow-net", // Proxy handles network restrictions
			"--allow-env",
			fmt.Sprintf("--allow-read=%s", absDir),
			"--deny-write",
			"--deny-sys",
			"--deny-run",
			"--deny-ffi",
		)
	} else {
		// Permissive mode for --no-proxy
		denoArgs = append(denoArgs,
			"--allow-net",
			"--allow-read",
			"--allow-env",
		)
	}

	if watch {
		denoArgs = append(denoArgs, "--watch")
	}

	denoArgs = append(denoArgs, mainFile)

	// Create the command
	execCmd := exec.CommandContext(ctx, denoPath, denoArgs...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	execCmd.Stdin = os.Stdin
	execCmd.Dir = absDir
	execCmd.Env = env

	// Set up signal handling for graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		if proxyInstance != nil {
			proxyInstance.Stop()
		}
		if execCmd.Process != nil {
			_ = execCmd.Process.Signal(syscall.SIGTERM)
		}
	}()

	pterm.Info.Printf("Starting dev server on port %d\n", port)
	pterm.Info.Printf("Source: %s\n", mainFile)
	if watch {
		pterm.Info.Println("Hot reload enabled")
	}
	pterm.Info.Println("Press Ctrl+C to stop")
	fmt.Println()

	err = execCmd.Run()
	if err != nil {
		// Check if it was killed by signal (expected on Ctrl+C)
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == -1 {
				return nil
			}
		}
		return fmt.Errorf("deno exited with error: %w", err)
	}

	return nil
}

func loadAllowlist(dir string) []string {
	denoJSON := filepath.Join(dir, "deno.json")
	data, err := os.ReadFile(denoJSON)
	if err != nil {
		return nil
	}

	var config denoConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil
	}

	return config.Permissions.Net
}

// createDPoPProofer creates a DPoP proofer with a new EC key.
func createDPoPProofer() (*dpop.Proofer, error) {
	// Generate EC P-256 key for DPoP
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate EC key: %w", err)
	}

	jwk := jose.JSONWebKey{
		Key:       privateKey,
		Algorithm: string(jose.ES256),
		Use:       "sig",
	}

	return dpop.NewProofer(&jwk)
}

