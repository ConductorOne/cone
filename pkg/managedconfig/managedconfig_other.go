//go:build !linux && !darwin && !windows

package managedconfig

// readManagedConfig returns no configuration on platforms without a managed
// device configuration store.
func readManagedConfig() map[string]string {
	return nil
}
