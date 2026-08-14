//go:build reprocredrecursion

// End-to-end reproduction harness for IGA-3789, run against a real AWS CLI.
//
// The claim: `cone aws credentials <profile>` is installed as a
// credential_process, cone shells out to `aws sso get-role-credentials`, and
// that child `aws` resolves the same profile's credential_process — re-invoking
// cone, unbounded.
//
// This file is excluded from normal builds and CI by the `reprocredrecursion`
// build tag, because it needs a real AWS CLI and makes live calls to the public
// AWS SSO OIDC endpoint. Run it explicitly:
//
//	go test -tags reprocredrecursion -run TestIGA3789 -v ./cmd/cone/...
//
// It fails loudly (t.Fatal, never t.Skip) if `aws` is missing, so a
// misconfigured run cannot be mistaken for a pass.
//
// A NOTE ON THE FIRST VERSION OF THIS HARNESS, which reported the claim refuted:
// resolution order in botocore puts the web-identity provider AHEAD of the
// custom-process (credential_process) provider. Any environment that exports
// AWS_ROLE_ARN + AWS_WEB_IDENTITY_TOKEN_FILE — Kubernetes IRSA, and many CI
// runners — therefore satisfies the credential lookup before credential_process
// is ever consulted, and the harness records zero invocations no matter how
// recursive the configuration is. Clearing only AWS_ACCESS_KEY_ID /
// AWS_SECRET_ACCESS_KEY / AWS_SESSION_TOKEN is not enough. scrubAWSEnv below
// clears the whole set; do not weaken it.
//
// Safety: the credential_process target here does re-invoke `aws`, so it can
// genuinely recurse — that is the point. Growth is bounded by a depth counter
// carried in the child environment (the script exits once it is exceeded) and
// by a context timeout on every call, so the process table is never at risk.
// HOME is redirected to a per-test temp dir, so the real ~/.aws is never read
// or written.
package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const maxReproDepth = 4

// recursiveCredentialProcess stands in for `cone aws credentials <profile>` as it
// behaved before the fix: it logs its invocation and then shells out to
// `aws sso get-role-credentials` with a fully inherited environment.
// __MAXDEPTH__ is substituted by setupReproHome.
const recursiveCredentialProcess = `#!/bin/sh
DEPTH=$((${CONE_REPRO_DEPTH:-0} + 1))
export CONE_REPRO_DEPTH="$DEPTH"
echo "INVOKED depth=$DEPTH pid=$$ ppid=$PPID profile=${AWS_PROFILE:-<unset>}" >> "$CONE_REPRO_LOG"
if [ "$DEPTH" -gt __MAXDEPTH__ ]; then
  echo "DEPTH-CAP-HIT depth=$DEPTH" >> "$CONE_REPRO_LOG"
  exit 97
fi
aws sso get-role-credentials --access-token bogus-token-value \
  --account-id 123456789012 --role-name TestRole --region us-east-1 \
  --output json >/dev/null 2>&1
echo '{"Version":1,"AccessKeyId":"ASIAFAKEFAKEFAKEFAKE","SecretAccessKey":"fake","SessionToken":"fake","Expiration":"2099-01-01T00:00:00Z"}'
`

const reproSSOStartURL = "https://d-9067example.awsapps.com/start"

// coneProfileConfig mirrors what createAWSProfile writes, shape for shape.
const coneProfileConfig = `[profile %s]
credential_process = %s
cone_app_id = app-123
cone_entitlement_id = ent-123
cone_sso_account_id = 123456789012
cone_sso_role_name = TestRole
cone_sso_region = us-east-1
cone_sso_start_url = %s
cone_sso_registration_scopes = sso:account:access
sso_session = cone-sso
region = us-east-1
output = json

[sso-session cone-sso]
sso_start_url = %s
sso_region = us-east-1
sso_registration_scopes = sso:account:access
`

// scrubAWSEnv clears every ambient credential source that outranks
// credential_process in botocore's resolution chain. Without this the harness
// silently reports a false negative. See the file comment.
func scrubAWSEnv(t *testing.T) {
	t.Helper()
	for _, k := range []string{
		"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_SESSION_TOKEN",
		"AWS_ROLE_ARN", "AWS_WEB_IDENTITY_TOKEN_FILE", "AWS_ROLE_SESSION_NAME",
		"AWS_CONTAINER_CREDENTIALS_FULL_URI", "AWS_CONTAINER_CREDENTIALS_RELATIVE_URI",
		"AWS_PROFILE", "AWS_DEFAULT_PROFILE",
	} {
		t.Setenv(k, "") // registers the restore
		os.Unsetenv(k)
	}
	t.Setenv("AWS_EC2_METADATA_DISABLED", "true")
}

// setupReproHome builds an isolated HOME holding a cone-shaped profile whose
// credential_process really does shell back out to `aws`, and returns the path
// of the invocation log.
func setupReproHome(t *testing.T, profile string) string {
	t.Helper()
	if _, err := exec.LookPath("aws"); err != nil {
		t.Fatalf("this test requires a real AWS CLI on PATH: %v", err)
	}
	home := t.TempDir()
	if err := os.MkdirAll(filepath.Join(home, ".aws"), 0o700); err != nil {
		t.Fatalf("mkdir .aws: %v", err)
	}
	logPath := filepath.Join(home, "invocations.log")
	script := filepath.Join(home, "fake-cone.sh")
	body := strings.ReplaceAll(recursiveCredentialProcess, "__MAXDEPTH__", fmt.Sprintf("%d", maxReproDepth))
	if err := os.WriteFile(script, []byte(body), 0o700); err != nil {
		t.Fatalf("write credential_process script: %v", err)
	}
	cfg := fmt.Sprintf(coneProfileConfig, profile, script, reproSSOStartURL, reproSSOStartURL)
	if err := os.WriteFile(filepath.Join(home, ".aws", "config"), []byte(cfg), 0o600); err != nil {
		t.Fatalf("write aws config: %v", err)
	}

	scrubAWSEnv(t)
	t.Setenv("HOME", home)
	t.Setenv("CONE_REPRO_LOG", logPath)
	t.Setenv("CONE_REPRO_DEPTH", "0")
	return logPath
}

func invocationCount(t *testing.T, logPath string) int {
	t.Helper()
	b, err := os.ReadFile(logPath) //nolint:gosec // test-owned temp path
	if os.IsNotExist(err) {
		return 0
	}
	if err != nil {
		t.Fatalf("read invocation log: %v", err)
	}
	return strings.Count(string(b), "INVOKED ")
}

// TestIGA3789_RecursionChannelExists reproduces the reported mechanism against the
// real AWS CLI: an `aws sso get-role-credentials` invocation that inherits a
// cone-managed AWS_PROFILE — exactly what cone spawned before the fix — resolves
// that profile's credential_process and recurses.
func TestIGA3789_RecursionChannelExists(t *testing.T) {
	logPath := setupReproHome(t, "aws-testrole")
	t.Setenv("AWS_PROFILE", "aws-testrole")

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	// The pre-fix spawn shape: no cmd.Env, so the child inherits AWS_PROFILE.
	cmd := exec.CommandContext(ctx, "aws", "sso", "get-role-credentials",
		"--access-token", "bogus-token-value",
		"--account-id", "123456789012",
		"--role-name", "TestRole",
		"--region", "us-east-1",
		"--output", "json")
	out, err := cmd.CombinedOutput()
	t.Logf("inherited-environment child: err=%v out=%s", err, strings.TrimSpace(string(out)))

	n := invocationCount(t, logPath)
	t.Logf("credential_process invocations with an inherited environment: %d", n)
	if n < 2 {
		t.Fatalf("expected the credential_process to re-enter itself (>=2 invocations), got %d — "+
			"if this is 0, check that scrubAWSEnv still clears every provider that outranks credential_process", n)
	}
}

// TestIGA3789_FixedChildEnvDoesNotRecurse exercises cone's real getRoleCredentials
// against the real AWS CLI, in the same recursive configuration, and asserts the
// credential_process is never entered.
func TestIGA3789_FixedChildEnvDoesNotRecurse(t *testing.T) {
	logPath := setupReproHome(t, "aws-testrole")
	t.Setenv("AWS_PROFILE", "aws-testrole")

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	// A bogus token: the call is expected to fail at AWS with UnauthorizedException.
	// What matters is the environment the child was handed, not the API result.
	out, err := getRoleCredentials(ctx, "bogus-token-value", "123456789012", "TestRole", "us-east-1")
	t.Logf("getRoleCredentials: err=%v out=%q", err, string(out))
	if err == nil {
		t.Fatal("expected an authentication failure from AWS with a bogus token; got success, so this run did not reach AWS")
	}
	if !strings.Contains(err.Error(), "UnauthorizedException") && !strings.Contains(err.Error(), "Session token not found") {
		t.Fatalf("expected a live AWS UnauthorizedException, got %v — the call may not have reached AWS at all", err)
	}

	if n := invocationCount(t, logPath); n != 0 {
		t.Fatalf("credential_process was entered %d time(s); the fixed child environment must make that impossible", n)
	}
}

// TestIGA3789_SSOLoginDoesNotRecurse covers cone's other AWS CLI spawn site.
func TestIGA3789_SSOLoginDoesNotRecurse(t *testing.T) {
	logPath := setupReproHome(t, "aws-testrole")
	t.Setenv("AWS_PROFILE", "aws-testrole")

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	// Fails at RegisterClient against the placeholder start URL; the invocation
	// count is what is under test.
	err := ssoLogin(ctx)
	t.Logf("ssoLogin: err=%v", err)

	if n := invocationCount(t, logPath); n != 0 {
		t.Fatalf("credential_process was entered %d time(s) via ssoLogin", n)
	}
}
