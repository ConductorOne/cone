//go:build darwin && !cgo

package managedconfig

import "os/exec"

// defaultsRead reads a single managed-preferences value through the `defaults`
// tool. Reading via the managed-preferences layer (rather than the on-disk
// plist) ensures only administrator-pushed policy is honored. It is a variable
// so tests can stub the lookup.
//
// This CGO-off fallback exists so cone still cross-compiles for darwin from a
// pure-Go toolchain (for example darwin/arm64 built on Linux with
// CGO_ENABLED=0). The darwin && cgo build uses the native CoreFoundation reader
// in managedconfig_darwin_cgo.go instead.
var defaultsRead = func(domain, key string) ([]byte, error) {
	return exec.Command("defaults", "read", domain, key).Output()
}

// readManagedConfig reads the managed device configuration from the macOS
// managed-preferences domain. A missing key or lookup error yields a nil map.
func readManagedConfig() map[string]string {
	out, err := defaultsRead(Namespace, KeyTenantDomain)
	if err != nil {
		return nil
	}
	return parseDefaultsValue(KeyTenantDomain, out)
}
