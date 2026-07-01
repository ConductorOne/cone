//go:build windows

package managedconfig

import "golang.org/x/sys/windows/registry"

// registryPath is the HKLM policy key holding the managed device configuration.
const registryPath = `SOFTWARE\Policies\ConductorOne\C1`

// readManagedConfig reads the managed device configuration from the Windows
// registry policy key. A missing key or read error yields a nil map.
func readManagedConfig() map[string]string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, registryPath, registry.QUERY_VALUE)
	if err != nil {
		return nil
	}
	defer k.Close()

	val, _, err := k.GetStringValue(KeyTenantDomain)
	if err != nil {
		return nil
	}
	return map[string]string{KeyTenantDomain: val}
}
