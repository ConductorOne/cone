package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// Regression coverage for IGA-3789: `cone aws credentials <profile>` is wired into
// ~/.aws/config as a credential_process, and it shells out to the AWS CLI. If the
// child `aws` process can resolve a cone-managed profile, its credential_process
// re-invokes cone and the pair recurses until the process table is exhausted.
//
// These tests install a fake `aws` on PATH that records the environment it was
// handed, so they assert the actual child environment without needing the real
// AWS CLI, network, or credentials.

const envRecorderScript = `#!/bin/sh
env > "$AWS_ENV_DUMP"
echo '{"roleCredentials":{"accessKeyId":"AKIA","secretAccessKey":"s","sessionToken":"t","expiration":0}}'
exit 0
`

// installFakeAWS puts an `aws` on PATH that dumps its environment to the returned
// path and exits successfully.
func installFakeAWS(t *testing.T) string {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("fake aws stub is a /bin/sh script; cone's CI runs these on linux")
	}
	dir := t.TempDir()
	script := filepath.Join(dir, "aws")
	if err := os.WriteFile(script, []byte(envRecorderScript), 0o700); err != nil { //nolint:gosec // must be executable; test-owned temp dir
		t.Fatalf("write fake aws: %v", err)
	}
	dump := filepath.Join(dir, "env.dump")
	t.Setenv("AWS_ENV_DUMP", dump)
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
	if got, err := exec.LookPath("aws"); err != nil || got != script {
		t.Fatalf("fake aws not first on PATH: got %q err %v", got, err)
	}
	return dump
}

func readChildEnv(t *testing.T, dump string) map[string]string {
	t.Helper()
	b, err := os.ReadFile(dump) //nolint:gosec // test-owned temp path
	if err != nil {
		t.Fatalf("child aws never ran (no env dump at %s): %v", dump, err)
	}
	env := map[string]string{}
	for _, line := range strings.Split(string(b), "\n") {
		if k, v, ok := strings.Cut(line, "="); ok {
			env[k] = v
		}
	}
	if len(env) == 0 {
		t.Fatal("child aws env dump was empty")
	}
	return env
}

func TestGetRoleCredentialsChildCannotReenterCredentialProcess(t *testing.T) {
	dump := installFakeAWS(t)
	// The parent environment a user hits this through: a cone-managed profile
	// selected by env var, which is what the child would otherwise inherit.
	t.Setenv("AWS_PROFILE", "aws-testrole")
	t.Setenv("AWS_DEFAULT_PROFILE", "aws-testrole")

	if _, err := getRoleCredentials(context.Background(), "token", "123456789012", "TestRole", "us-east-1"); err != nil {
		t.Fatalf("getRoleCredentials: %v", err)
	}

	env := readChildEnv(t, dump)
	for _, k := range []string{"AWS_PROFILE", "AWS_DEFAULT_PROFILE"} {
		if v, ok := env[k]; ok {
			t.Errorf("child aws inherited %s=%q; it can resolve a cone-managed profile and recurse", k, v)
		}
	}
	// A [default] profile carrying the same credential_process is just as recursive
	// as a named one, so the config files are pinned away from ~/.aws entirely.
	for _, k := range []string{"AWS_CONFIG_FILE", "AWS_SHARED_CREDENTIALS_FILE"} {
		if env[k] != os.DevNull {
			t.Errorf("child aws %s = %q, want %q", k, env[k], os.DevNull)
		}
	}
	// The rest of the environment must survive — proxies, CA bundles, HOME.
	if env["HOME"] != os.Getenv("HOME") {
		t.Errorf("child aws HOME = %q, want %q", env["HOME"], os.Getenv("HOME"))
	}
}

func TestSSOLoginChildCannotReenterCredentialProcess(t *testing.T) {
	dump := installFakeAWS(t)
	t.Setenv("AWS_PROFILE", "aws-testrole")
	t.Setenv("AWS_DEFAULT_PROFILE", "aws-testrole")

	if err := ssoLogin(context.Background()); err != nil {
		t.Fatalf("ssoLogin: %v", err)
	}

	env := readChildEnv(t, dump)
	for _, k := range []string{"AWS_PROFILE", "AWS_DEFAULT_PROFILE"} {
		if v, ok := env[k]; ok {
			t.Errorf("child aws inherited %s=%q; it can resolve a cone-managed profile and recurse", k, v)
		}
	}
	// `aws sso login --sso-session cone-sso` reads the [sso-session] block out of
	// ~/.aws/config, so this call must NOT be pinned at os.DevNull.
	if v, ok := env["AWS_CONFIG_FILE"]; ok && v == os.DevNull {
		t.Error("ssoLogin pinned AWS_CONFIG_FILE at os.DevNull; the [sso-session cone-sso] block would be unreadable")
	}
}

func TestAWSChildEnvOverridesReplaceRatherThanDuplicate(t *testing.T) {
	t.Setenv("AWS_PROFILE", "aws-testrole")
	t.Setenv("AWS_CONFIG_FILE", "/home/user/.aws/config")

	env := awsChildEnv("AWS_CONFIG_FILE=" + os.DevNull)

	var configFiles []string
	for _, kv := range env {
		if k, v, ok := strings.Cut(kv, "="); ok && k == "AWS_CONFIG_FILE" {
			configFiles = append(configFiles, v)
		}
		if strings.HasPrefix(kv, "AWS_PROFILE=") {
			t.Errorf("awsChildEnv kept %q", kv)
		}
	}
	// exec applies the last occurrence, so an inherited value ahead of the override
	// is harmless — but a value AFTER it would win and silently defeat the fix.
	if n := len(configFiles); n == 0 {
		t.Fatal("awsChildEnv dropped the AWS_CONFIG_FILE override")
	} else if configFiles[n-1] != os.DevNull {
		t.Errorf("last AWS_CONFIG_FILE = %q, want %q", configFiles[n-1], os.DevNull)
	}
}
