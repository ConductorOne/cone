# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Cone is a CLI tool for the ConductorOne platform, written in Go. It manages access to entitlements (request, approve, deny, revoke access) via the ConductorOne API.

## Build & Development Commands

```bash
make build          # Build binary → dist/<OS>_<ARCH>/cone
make lint           # Run golangci-lint (strict config in .golangci.yml, ~30 linters)
go test ./...       # Run all tests
go test -v -run TestName ./pkg/client/  # Run a single test
make update-deps    # Update all Go dependencies and re-vendor
```

Uses vendored dependencies (`vendor/`). After modifying `go.mod`, run `go mod tidy -v && go mod vendor`.

## Architecture

### Entry Point & Commands (`cmd/cone/`)

- `main.go` — Cobra root command setup, signal handling, registers all subcommands
- `cmd.go` — `cmdContext()` creates authenticated `C1Client` + viper config for each command. Auth priority: access token env var → OIDC token exchange → client credentials
- Each command file (e.g., `task.go`, `search_entitlements.go`) returns a `*cobra.Command` and calls `cmdContext()` to get the client

### Core Packages (`pkg/`)

- **`client/`** — Wraps `conductorone-sdk-go`. The `C1Client` interface defines all API operations. Auth uses JWT bearer assertion (Ed25519 signed) via `token_source.go`, with RFC 8693 token exchange in `token_exchange.go`. Client ID format: `name@host/suffix` (host is parsed to determine API endpoint).
- **`output/`** — Pluggable output formatting. `Manager` interface with table/JSON implementations. Data types implement `TablePrint` (`Header() []string`, `Rows() [][]string`) for table output, and optionally `WideTablePrint` (`WideHeader()`, `WideRows()`) for wide mode. JSON output serializes the struct directly.
- **`uhttp/`** — HTTP client factory with OAuth2 token source, debug logging, custom transport.
- **`logging/`** — Singleton zap logger initialized once at startup.

### Adding a New Command

1. Create file in `cmd/cone/`
2. Return a `*cobra.Command` that calls `cmdContext(cmd)` to get `(ctx, client, viper, error)`
3. Use `output.NewManager(ctx, v)` to format output
4. Register in `main.go` via `cliCmd.AddCommand()`

### Adding a New Client Method

1. Add the method signature to the `C1Client` interface in `pkg/client/client.go`
2. Implement on `*client` in the appropriate file under `pkg/client/`
3. Use `c.sdk.<Service>.<Method>(ctx, operationsRequest)` to call the SDK
4. Check `NewHTTPError(resp.RawResponse)` for HTTP-level errors

## Linting Rules

Key non-obvious rules from `.golangci.yml`:

- Line length limit: 200 characters
- No naked returns (any function length)
- No named returns
- Comments must end in a period (except TODOs)
- Variable naming: use `ID`, `URL`, `HTTP`, `API` (not `Id`, `Url`, etc.)
- No `init()` functions
- All errors must be checked (exceptions: `fmt.Printf/Println`, `fmt.Fprintf/Fprintln`)
- `goimports` for import formatting
