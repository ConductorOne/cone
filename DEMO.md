# Cone Registry CLI Demo

Demonstrate the `cone registry` commands for browsing, downloading, and publishing connectors.

---

## Connector Run Modes

ConductorOne connectors can run in different deployment modes:

| Mode | Description | Use Case |
|------|-------------|----------|
| **Managed (Lambda)** | Hosted by ConductorOne in AWS Lambda | Default for cloud-hosted connectors; zero infrastructure to manage |
| **Self-Hosted (Local)** | Downloaded binary running in your infrastructure | Air-gapped environments, custom networks, on-prem systems |
| **Vendored** | Built into ConductorOne platform | Legacy connectors, special integrations |

### When You Download a Connector

Using `cone registry download` implies **self-hosted mode**:

- You download the connector binary to your infrastructure
- You run it locally (or in your cloud, Kubernetes, etc.)
- You configure it to sync back to ConductorOne
- You manage updates, scaling, and availability

This is different from **managed connectors** which:
- Run automatically in ConductorOne's infrastructure
- Are configured entirely through the UI/API
- Update automatically when new versions are published
- Require no infrastructure management

### Choosing a Run Mode

| Scenario | Recommended Mode |
|----------|------------------|
| Standard SaaS integrations (Okta, Google, etc.) | Managed |
| On-premises systems (Active Directory, databases) | Self-Hosted |
| Air-gapped or regulated environments | Self-Hosted |
| Custom/internal applications | Self-Hosted |
| Testing connector changes | Self-Hosted |

---

## Assumptions & Prerequisites

### For Production Use

| Requirement | Description |
|-------------|-------------|
| ConductorOne Tenant | Active tenant at `your-tenant.conductor.one` |
| Okta SSO (Optional) | If using Okta for SSO, configure OIDC application |
| Publisher Role | User must have publisher scope assigned in ConductorOne |
| Admin Role | Admin commands require admin scope in ConductorOne |

### Authentication Flow

The `cone login` command uses OAuth 2.0 with PKCE:

1. Opens browser to ConductorOne login page
2. User authenticates (directly or via SSO like Okta)
3. ConductorOne issues JWT with user's scopes
4. Token stored in `~/.conductorone/config.yaml`
5. Registry validates JWT against ConductorOne's JWKS endpoint

**JWT Claims Used by Registry:**

| Claim | Purpose |
|-------|---------|
| `iss` | Must match ConductorOne tenant URL |
| `aud` | Must match registry's configured audience |
| `sub` | User identifier for audit logging |
| `org` | Organization for resource scoping |
| `c1scp` | ConductorOne scope IDs (matched against publisher/admin roles) |

### ConductorOne Role Configuration

To grant registry access, assign these scopes in ConductorOne:

| Permission Level | Required Scope |
|------------------|----------------|
| Publisher | Scope ID configured in registry's `roles.publisher` |
| Admin | Scope ID configured in registry's `roles.admin` |

### For Local Development

| Requirement | Description |
|-------------|-------------|
| Docker | For LocalStack (DynamoDB + S3) |
| Go 1.22+ | For building cone and registry |
| cosign (Optional) | For signature verification |

---

## Part 1: Browse & Download (No Auth Required)

These commands work without authentication against the public registry.

### List All Connectors

```bash
# List all available connectors (currently 51 from dist.conductorone.com)
cone registry list

# Output as JSON for scripting
cone registry list --output json
```

> **Note**: The description overlay contains 167 entries (for future connectors), but dist.conductorone.com currently publishes 51 connectors. 46 of these have matching descriptions.

### Show Connector Details

```bash
# Show connector metadata (includes description from overlay)
cone registry show ConductorOne/baton-okta

# Show a specific version
cone registry show ConductorOne/baton-okta v0.1.0

# Show available platforms for a version
cone registry show ConductorOne/baton-okta v0.1.0 --platforms
```

### List Versions

```bash
# List all versions of a connector
cone registry versions ConductorOne/baton-okta
```

### Download a Connector

```bash
# Download stable version for your platform
cone registry download ConductorOne/baton-okta

# Download specific version
cone registry download ConductorOne/baton-okta v0.1.0

# Download for a different platform
cone registry download ConductorOne/baton-okta --platform linux-amd64

# Download to specific directory
cone registry download ConductorOne/baton-okta --output ./bin/

# Skip verification (not recommended)
cone registry download ConductorOne/baton-okta --skip-verify
```

---

## Part 2: Authentication

Publisher and admin commands require ConductorOne authentication.

### Login to ConductorOne

```bash
# Login (opens browser for OAuth)
cone login your-tenant.conductor.one

# Verify you're logged in
cone whoami
```

### Using Multiple Profiles

```bash
# Login with named profiles
cone login prod-tenant.conductor.one --profile prod
cone login dev-tenant.conductor.one --profile dev

# Use a specific profile
cone registry status --profile prod
```

---

## Part 3: Signing Keys (Publisher)

Manage cryptographic keys for signing connector releases.

### Generate a Key Pair

```bash
# Generate ECDSA P-256 key pair (cosign-compatible)
cone registry keys generate release-2024

# Generates:
#   release-2024.key (private - keep secret!)
#   release-2024.pub (public - register with registry)
```

### Register Your Public Key

```bash
# Register the public key with your organization
cone registry keys add ConductorOne \
  --name "Release Signing Key 2024" \
  --type cosign \
  --key-file release-2024.pub
```

### List Your Organization's Keys

```bash
cone registry keys list ConductorOne
```

### Show Key Details

```bash
cone registry keys show ConductorOne <key-id>
```

---

## Part 4: Publishing (Publisher)

Publish new connector versions to the registry.

### Build Your Binaries

```bash
# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o dist/baton-myapp-linux-amd64 .
GOOS=linux GOARCH=arm64 go build -o dist/baton-myapp-linux-arm64 .
GOOS=darwin GOARCH=amd64 go build -o dist/baton-myapp-darwin-amd64 .
GOOS=darwin GOARCH=arm64 go build -o dist/baton-myapp-darwin-arm64 .
GOOS=windows GOARCH=amd64 go build -o dist/baton-myapp-windows-amd64.exe .
```

### Sign Your Binaries (Optional but Recommended)

```bash
# Sign each binary with cosign
for binary in dist/baton-myapp-*; do
  cosign sign-blob --key release-2024.key \
    --output-signature "${binary}.sig" \
    "$binary"
done
```

### Preview What Will Be Published

```bash
# Dry run - shows what would be uploaded
cone registry publish MyOrg/baton-myapp v1.0.0 \
  --binary-dir ./dist/ \
  --dry-run
```

### Publish

```bash
# Publish with metadata
cone registry publish MyOrg/baton-myapp v1.0.0 \
  --binary-dir ./dist/ \
  --description "Initial release" \
  --changelog "- User sync\n- Group sync\n- Entitlement sync" \
  --license "Apache-2.0"
```

### Check Publication Status

```bash
# View all your published versions
cone registry status

# Filter by connector
cone registry status --connector MyOrg/baton-myapp

# Filter by state
cone registry status --state VALIDATING
cone registry status --state PUBLISHED
cone registry status --state FAILED
```

---

## Part 5: Admin Commands

These require admin role permissions.

### Set Stable Version

```bash
# Mark a version as the stable/recommended version
cone registry set-stable MyOrg/baton-myapp v1.0.0
```

### Yank a Version

```bash
# Withdraw a version (e.g., security issue)
cone registry yank MyOrg/baton-myapp v0.9.0 --reason "Security vulnerability CVE-2024-XXXX"
```

### Deprecate a Connector

```bash
# Mark connector as deprecated
cone registry deprecate MyOrg/baton-legacy --reason "Replaced by baton-myapp"
```

---

## Part 6: Data Sync (Admin)

Compare registry against dist.conductorone.com and sync changes.

### Check Sync Status (Diff)

```bash
# Compare registry (DynamoDB) against dist.conductorone.com
cone registry diff

# Example output:
# Registry vs https://dist.conductorone.com
# Generated: 2026-01-07 15:30:05
#
# Summary:
#   Connectors: +1 added, -2 removed, ~49 modified
#   Versions:   +15 added, -0 removed
#
# Added Connectors (in registry, not in dist):
#   + TestOrg/baton-test
#
# Removed Connectors (in dist, not in registry):
#   - ConductorOne/baton-old
#
# Modified Connectors:
#   ~ ConductorOne/baton-okta
#       stable: v0.4.3 → v0.4.4
#       versions added: v0.4.4
```

```bash
# Output as JSON for scripting
cone registry diff --output json

# Compare against staging dist
cone registry diff --base-url https://staging-dist.conductorone.com
```

### Export to Dist (Push)

```bash
# Preview what would be exported (dry run)
cone registry export --dry-run

# Example output:
# Would export all connectors:
#   Connectors: 50
#   Releases:   136
#   Files:      187
#
# (dry run - no files written)
```

```bash
# Export all connectors to S3 (dist layout)
cone registry export

# Export specific connectors only
cone registry export ConductorOne/baton-okta ConductorOne/baton-aws

# Output as JSON
cone registry export --output json
```

### Sync Workflow

The `sync` command provides convenient aliases:

```bash
# Check what would change (alias for 'registry diff')
cone registry sync status

# Push changes to dist (alias for 'registry export')
cone registry sync push --dry-run
cone registry sync push

# Push specific connectors
cone registry sync push ConductorOne/baton-okta
```

### Typical Sync Workflow

```bash
# 1. Check current differences
cone registry diff

# 2. Review what will be exported
cone registry export --dry-run

# 3. Export to dist
cone registry export

# 4. Verify sync completed
cone registry diff
# Should show: "No differences found - registry matches dist."
```

---

## Part 7: JSON Output & Scripting

All commands support JSON output for automation.

```bash
# List as JSON
cone registry list --output json

# Pretty-printed JSON
cone registry show ConductorOne/baton-okta --output json-pretty

# Use with jq for filtering
cone registry list --output json | jq '.[] | select(.stableVersion != "")'

# Get download URL programmatically
cone registry show ConductorOne/baton-okta v0.1.0 --output json | jq -r '.downloadUrl'
```

---

## Local Development Setup

For testing against a local registry server with real connector data.

### Start Local Infrastructure

```bash
cd /path/to/connector-registry

# Start LocalStack (DynamoDB + S3)
docker compose up -d localstack

# Wait for LocalStack to be ready
until curl -s http://localhost:4566/_localstack/health | grep -q '"dynamodb": "running"'; do
  sleep 1
done
```

### Import Connector Data

There are two ways to populate the registry with connector data:

#### Option A: Live Import (requires network)

```bash
# Fetch directly from dist.conductorone.com
make import

# This fetches:
# - ~51 connectors from the live catalog
# - ~147 release manifests with platform/checksum info
# - Applies description overlay (167 connector descriptions)
```

#### Option B: Snapshot Import (offline capable)

```bash
# First, download a snapshot (do this once while online)
make snapshot-download
# Saves to data/catalog_snapshot.json (~900KB)

# Later, load from snapshot (works offline)
make snapshot-load
```

The snapshot approach is useful for:
- Air-gapped environments
- Reproducible demo data
- Faster local development (no network fetches)

### Description Overlay

The import applies rich descriptions from `pkg/importer/data/description_overlay.json`:

- **baton-okta**: "Syncs users, groups, roles, applications, and custom roles from Okta..."
- **baton-aws**: "Syncs IAM users, groups, roles, and accounts with optional AWS Identity Center..."
- **baton-azure**: "Syncs Entra ID users, groups, roles, resource groups, tenants..."

The overlay contains descriptions for 167 connectors (more than currently published).

### Start the Registry Server

```bash
# Basic local server
make run-local

# Or with dev config (for testing auth)
AWS_ACCESS_KEY_ID=test \
AWS_SECRET_ACCESS_KEY=test \
AWS_REGION=us-east-1 \
AWS_ENDPOINT=http://localhost:4566 \
DYNAMODB_TABLE=connector-registry \
S3_BUCKET=connector-registry-binaries \
./registry-local serve --config=config.dev.json --port=8080
```

### Test Against Local Server

```bash
# All commands accept --registry-url flag
cone registry list --registry-url http://localhost:8080

# Should show 51 connectors with descriptions
cone registry list --registry-url http://localhost:8080 | wc -l
# Output: 52 (51 connectors + header)
```

### Test Authenticated Endpoints Locally

The `config.dev.json` enables `skip_signature_verification` for testing without real JWTs:

```bash
# Publisher test token (c1scp: ["publisher-scope"])
export PUBLISHER_TOKEN="eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJzdWIiOiJ0ZXN0LXVzZXIiLCJpc3MiOiJ0ZXN0LWlzc3VlciIsImF1ZCI6WyJjb25uZWN0b3ItcmVnaXN0cnkiXSwiZXhwIjo0MTAyNDQ0ODAwLCJpYXQiOjE3MDQwNjcyMDAsIm9yZyI6IlRlc3RPcmciLCJyb2xlcyI6WyJwdWJsaXNoZXIiXSwiYzFzY3AiOlsicHVibGlzaGVyLXNjb3BlIl0sImVtYWlsIjoidGVzdEBleGFtcGxlLmNvbSJ9."

# Admin test token (c1scp: ["admin-scope"])
export ADMIN_TOKEN="eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJzdWIiOiJ0ZXN0LWFkbWluIiwiaXNzIjoidGVzdC1pc3N1ZXIiLCJhdWQiOlsiY29ubmVjdG9yLXJlZ2lzdHJ5Il0sImV4cCI6NDEwMjQ0NDgwMCwiaWF0IjoxNzA0MDY3MjAwLCJjMXNjcCI6WyJhZG1pbi1zY29wZSJdLCJlbWFpbCI6ImFkbWluQHRlc3QuY29tIn0."

# Test publisher commands
cone registry keys list TestOrg \
  --registry-url http://localhost:8080 \
  --registry-token "$PUBLISHER_TOKEN"

cone registry status \
  --registry-url http://localhost:8080 \
  --registry-token "$PUBLISHER_TOKEN"

# Test admin commands (diff/export)
cone registry diff \
  --registry-url http://localhost:8080 \
  --registry-token "$ADMIN_TOKEN"

cone registry export --dry-run \
  --registry-url http://localhost:8080 \
  --registry-token "$ADMIN_TOKEN"

# Export specific connector
cone registry export ConductorOne/baton-okta \
  --registry-url http://localhost:8080 \
  --registry-token "$ADMIN_TOKEN" \
  --dry-run
```

---

## Command Reference

| Command | Auth | Description |
|---------|------|-------------|
| `registry list` | No | List all connectors |
| `registry show <connector>` | No | Show connector details |
| `registry show <connector> <version>` | No | Show version details |
| `registry versions <connector>` | No | List all versions |
| `registry download <connector>` | No | Download binary |
| `registry keys generate <name>` | No | Generate key pair locally |
| `registry keys list <org>` | Publisher | List org's signing keys |
| `registry keys add <org>` | Publisher | Register a signing key |
| `registry keys show <org> <id>` | Publisher | Show key details |
| `registry status` | Publisher | Show your published versions |
| `registry publish <connector> <version>` | Publisher | Publish a new version |
| `registry set-stable <connector> <version>` | Admin | Set stable version |
| `registry yank <connector> <version>` | Admin | Yank a version |
| `registry deprecate <connector>` | Admin | Deprecate connector |
| `registry diff` | Admin | Compare registry vs dist |
| `registry export [connectors...]` | Admin | Export to dist layout (S3) |
| `registry sync status` | Admin | Alias for `diff` |
| `registry sync push [connectors...]` | Admin | Alias for `export` |

---

## Troubleshooting

### "authentication required"

```bash
cone login your-tenant.conductor.one
cone whoami  # verify
```

### "no stable version available"

Specify version explicitly:
```bash
cone registry download ConductorOne/baton-okta v0.1.0
```

### "checksum verification failed"

Try again or skip verification:
```bash
cone registry download ConductorOne/baton-okta --skip-verify
```

### "connector not found"

Check spelling and org/name format:
```bash
cone registry list | grep -i okta
```

### "permission denied" 

Verify your ConductorOne account has the required scope:
```bash
# Check your token claims
cone whoami --output json | jq '.scopes'
```
