//go:build linux

package managedconfig

import (
	"os"
	"path/filepath"
	"testing"
)

// withLinuxConfigPath points the package at path for the duration of the test.
func withLinuxConfigPath(t *testing.T, path string) {
	t.Helper()
	prev := linuxConfigPath
	linuxConfigPath = path
	t.Cleanup(func() { linuxConfigPath = prev })
}

func writeTempConfig(t *testing.T, contents string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "managed.toml")
	if err := os.WriteFile(path, []byte(contents), 0600); err != nil {
		t.Fatalf("writing temp config: %v", err)
	}
	return path
}

func TestReadManagedConfigLinux_Parse(t *testing.T) {
	withLinuxConfigPath(t, writeTempConfig(t, "TenantDomain = \"acme.conductor.one\"\n"))

	if got := Read().TenantDomain; got != "acme.conductor.one" {
		t.Errorf("TenantDomain = %q, want %q", got, "acme.conductor.one")
	}
	if got := Read().ControlPlaneURL(); got != "https://acme.conductor.one" {
		t.Errorf("ControlPlaneURL() = %q, want %q", got, "https://acme.conductor.one")
	}
}

func TestReadManagedConfigLinux_UnknownKeyIgnored(t *testing.T) {
	withLinuxConfigPath(t, writeTempConfig(t, "Unknown = \"junk\"\nTenantDomain = \"acme.eu.c1.ai\"\nExtra = 42\n"))

	if got := Read().TenantDomain; got != "acme.eu.c1.ai" {
		t.Errorf("TenantDomain = %q, want %q", got, "acme.eu.c1.ai")
	}
}

func TestReadManagedConfigLinux_Absent(t *testing.T) {
	withLinuxConfigPath(t, filepath.Join(t.TempDir(), "does-not-exist.toml"))

	if got := Read(); got != (Config{}) {
		t.Errorf("Read() = %+v, want zero Config", got)
	}
}

func TestReadManagedConfigLinux_Malformed(t *testing.T) {
	withLinuxConfigPath(t, writeTempConfig(t, "this is not = = valid toml"))

	if got := Read(); got != (Config{}) {
		t.Errorf("Read() = %+v, want zero Config", got)
	}
}

func TestReadManagedConfigLinux_BareSlugRejected(t *testing.T) {
	withLinuxConfigPath(t, writeTempConfig(t, "TenantDomain = \"acme\"\n"))

	if got := Read().TenantDomain; got != "" {
		t.Errorf("TenantDomain = %q, want empty (bare slug should be rejected)", got)
	}
}
