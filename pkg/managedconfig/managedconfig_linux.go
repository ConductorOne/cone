//go:build linux

package managedconfig

import "os"

// linuxConfigPath is the location of the Linux managed device configuration
// file. It is a variable so tests can point it at a temporary file.
var linuxConfigPath = "/etc/c1/managed.toml"

// readManagedConfig reads the managed device configuration from the Linux
// managed-config file. A missing or unreadable file yields a nil map.
func readManagedConfig() map[string]string {
	data, err := os.ReadFile(linuxConfigPath)
	if err != nil {
		return nil
	}
	return parseManagedTOML(data)
}
