//go:build darwin

package managedconfig

import "os/exec"

// defaultsRead reads a single managed-preferences value through the `defaults`
// tool. Reading via the managed-preferences layer (rather than the on-disk
// plist) ensures only administrator-pushed policy is honored. It is a variable
// so tests can stub the lookup.
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
