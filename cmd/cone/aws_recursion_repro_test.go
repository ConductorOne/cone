//go:build reprocredrecursion

// Reproduction harness for IGA-3789.
//
// Hypothesis under test: getRoleCredentials (aws.go:618-634) shells out to
// `aws sso get-role-credentials` without setting cmd.Env, so the child `aws`
// process inherits cone's environment, resolves the same ~/.aws/config
// profile, and re-triggers that profile's own `credential_process = cone ...`
// entry — recursing until the process table is exhausted.
//
// This file is excluded from normal builds/CI by the `reprocredrecursion`
// build tag. Run explicitly:
//
//	go test -tags reprocredrecursion -run TestIGA3789 -v ./cmd/cone/...
//
// Requires a real `aws` CLI (v2 preferred — v1 lacks `aws sso login`) on
// PATH. The test fails loudly (t.Fatal, not t.Skip) if `aws` is missing,
// so a misconfigured run cannot be mistaken for a pass.
//
// Safety: the `credential_process` target used here is a throwaway marker
// script that only appends a line to a counter file and exits — it never
// re-invokes `cone`, `aws`, or itself. One recorded invocation is already
// conclusive evidence of the recursion channel, so there is no need (and no
// mechanism) for the harness itself to recurse or exhaust the process table.
// HOME is redirected to a per-test temp dir so ~/.aws/config on the real
// machine is never read or written.
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

const markerScript = `#!/bin/sh
echo "FIRED $(date +%s%N) pid=$$ ppid=$PPID" >> "$MARKER_LOG"
exit 1
`

// writeMarker installs a non-recursive stand-in credential_process target
// and returns the path to the invocation-count log it appends to.
func writeMarker(t *testing.T, dir string) (scriptPath, logPath string) {
	t.Helper()
	logPath = filepath.Join(dir, "marker.log")
	scriptPath = filepath.Join(dir, "marker.sh")
	if err := os.WriteFile(scriptPath, []byte(markerScript), 0o700); err != nil {
		t.Fatalf("write marker script: %v", err)
	}
	return scriptPath, logPath
}

func markerFireCount(t *testing.T, logPath string) int {
	t.Helper()
	b, err := os.ReadFile(logPath) //nolint:gosec // test-owned temp path
	if os.IsNotExist(err) {
		return 0
	}
	if err != nil {
		t.Fatalf("read marker log: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return 0
	}
	return len(lines)
}

func requireRealAWSCLI(t *testing.T) {
	t.Helper()
	if _, err := exec.LookPath("aws"); err != nil {
		t.Fatalf("real `aws` CLI not found on PATH — this repro requires it (install AWS CLI v2); cannot silently skip, see IGA-1658: %v", err)
	}
}

// setIsolatedAWSEnv points HOME/AWS_PROFILE at a scratch dir for the
// duration of the test and restores the prior values on cleanup. It never
// touches the real ~/.aws/config.
func setIsolatedAWSEnv(t *testing.T, home, profile string) {
	t.Helper()
	origHome, hadHome := os.LookupEnv("HOME")
	origProfile, hadProfile := os.LookupEnv("AWS_PROFILE")
	origKey, hadKey := os.LookupEnv("AWS_ACCESS_KEY_ID")
	origSecret, hadSecret := os.LookupEnv("AWS_SECRET_ACCESS_KEY")
	origSession, hadSession := os.LookupEnv("AWS_SESSION_TOKEN")

	os.Setenv("HOME", home)
	if profile != "" {
		os.Setenv("AWS_PROFILE", profile)
	} else {
		os.Unsetenv("AWS_PROFILE")
	}
	os.Unsetenv("AWS_ACCESS_KEY_ID")
	os.Unsetenv("AWS_SECRET_ACCESS_KEY")
	os.Unsetenv("AWS_SESSION_TOKEN")

	t.Cleanup(func() {
		restore(hadHome, "HOME", origHome)
		restore(hadProfile, "AWS_PROFILE", origProfile)
		restore(hadKey, "AWS_ACCESS_KEY_ID", origKey)
		restore(hadSecret, "AWS_SECRET_ACCESS_KEY", origSecret)
		restore(hadSession, "AWS_SESSION_TOKEN", origSession)
	})
}

func restore(had bool, key, val string) {
	if had {
		os.Setenv(key, val)
	} else {
		os.Unsetenv(key)
	}
}

// TestIGA3789_GetRoleCredentialsDoesNotTriggerCredentialProcess is the actual
// reproduction attempt for the ticket's claim. It calls cone's real
// getRoleCredentials (aws.go:618) and ssoLogin (aws.go:609) — the two exec.
// CommandContext call sites in this file, neither of which sets cmd.Env —
// against a profile whose credential_process points at our marker, and
// checks whether the child `aws` process ever invokes it.
func TestIGA3789_GetRoleCredentialsDoesNotTriggerCredentialProcess(t *testing.T) {
	requireRealAWSCLI(t)

	dir := t.TempDir()
	home := filepath.Join(dir, "home")
	if err := os.MkdirAll(filepath.Join(home, ".aws"), 0o755); err != nil {
		t.Fatalf("mkdir HOME/.aws: %v", err)
	}
	markerPath, logPath := writeMarker(t, dir)

	const profileName = "iga3789-test-profile"
	// Mirrors the exact shape cone writes at aws.go:244-257, with
	// credential_process pointed at our marker instead of `cone`.
	config := fmt.Sprintf(`
[profile %s]
credential_process = %s
cone_sso_account_id = 123456789012
cone_sso_role_name = TestRole
cone_sso_region = us-east-1
sso_session = cone-sso
region = us-east-1
output = json

[sso-session cone-sso]
sso_start_url = https://example.awsapps.com/start
sso_region = us-east-1
sso_registration_scopes = sso:account:access
`, profileName, markerPath)

	if err := os.WriteFile(filepath.Join(home, ".aws", "config"), []byte(config), 0o600); err != nil {
		t.Fatalf("write ~/.aws/config: %v", err)
	}

	setIsolatedAWSEnv(t, home, profileName)
	os.Setenv("MARKER_LOG", logPath)
	t.Cleanup(func() { os.Unsetenv("MARKER_LOG") })

	// Bound #1: hard context timeout. Bound #2: the marker script itself
	// never recurses (see file header) — belt and suspenders, not strictly
	// needed given #2, but keeps any future edit to this test honest.
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	t.Run("getRoleCredentials", func(t *testing.T) {
		out, err := getRoleCredentials(ctx, "bogus-access-token", "123456789012", "TestRole", "us-east-1")
		t.Logf("getRoleCredentials returned err=%v out=%q", err, out)
		if err == nil {
			t.Fatalf("expected an error for a bogus SSO access token, got nil (out=%q)", out)
		}

		n := markerFireCount(t, logPath)
		if n > 0 {
			t.Fatalf("RECURSION CONFIRMED: credential_process (marker) invoked %d time(s) by the aws sso get-role-credentials child process — IGA-3789 reproduces at cmd/cone/aws.go:618-634", n)
		}
		t.Logf("credential_process invocation count after getRoleCredentials: %d (0 = hypothesis refuted for this call)", n)
	})

	t.Run("ssoLogin", func(t *testing.T) {
		err := ssoLogin(ctx)
		t.Logf("ssoLogin returned err=%v", err)

		n := markerFireCount(t, logPath)
		if n > 0 {
			t.Fatalf("RECURSION CONFIRMED: credential_process (marker) invoked %d time(s) by the aws sso login child process — IGA-3789 reproduces at cmd/cone/aws.go:609-615", n)
		}
		t.Logf("credential_process invocation count after ssoLogin: %d (0 = hypothesis refuted for this call)", n)
	})
}

// TestIGA3789_MarkerFiresOnOrdinarySigV4Call is a positive control. It does
// NOT go through cone's code — it proves the marker/harness setup used above
// is capable of detecting a credential_process invocation at all, so a "0
// fires" result in the test above is evidence of no recursion rather than a
// broken harness.
func TestIGA3789_MarkerFiresOnOrdinarySigV4Call(t *testing.T) {
	requireRealAWSCLI(t)

	dir := t.TempDir()
	home := filepath.Join(dir, "home")
	if err := os.MkdirAll(filepath.Join(home, ".aws"), 0o755); err != nil {
		t.Fatalf("mkdir HOME/.aws: %v", err)
	}
	markerPath, logPath := writeMarker(t, dir)

	const profileName = "iga3789-control-profile"
	config := fmt.Sprintf(`
[profile %s]
region = us-east-1
credential_process = %s
`, profileName, markerPath)
	if err := os.WriteFile(filepath.Join(home, ".aws", "config"), []byte(config), 0o600); err != nil {
		t.Fatalf("write ~/.aws/config: %v", err)
	}

	setIsolatedAWSEnv(t, home, "")
	os.Setenv("MARKER_LOG", logPath)
	t.Cleanup(func() { os.Unsetenv("MARKER_LOG") })

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// A plain SigV4 operation — unlike sso get-role-credentials / sso login,
	// this one genuinely needs resolved IAM credentials, so the CLI must
	// consult credential_process.
	cmd := exec.CommandContext(ctx, "aws", "sts", "get-caller-identity", "--profile", profileName) //nolint:gosec // test-controlled args
	out, err := cmd.CombinedOutput()
	t.Logf("aws sts get-caller-identity: err=%v output=%s", err, strings.TrimSpace(string(out)))

	n := markerFireCount(t, logPath)
	if n == 0 {
		t.Fatalf("control failed: credential_process was never invoked for an ordinary SigV4 call — this test harness (or the installed aws CLI) cannot detect credential_process invocations, so the recursion test's 0-fire result is not meaningful")
	}
	t.Logf("control OK: credential_process fired %d time(s) for an ordinary SigV4 call", n)
}
