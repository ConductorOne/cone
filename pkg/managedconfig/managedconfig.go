// Package managedconfig reads the ConductorOne managed device configuration:
// administrator-defined policy delivered to a device through an MDM. It lets any
// ConductorOne client auto-discover the tenant it belongs to without per-machine
// manual setup.
//
// The configuration lives in a company-level, read-only managed store addressed
// by the namespace "ai.c1". Reads are best-effort by design: an absent,
// unreadable, or malformed store yields a zero Config. Nothing in this package
// returns an error or panics, so callers can consult it unconditionally at the
// top of their configuration-resolution chain.
package managedconfig

import (
	"strings"

	"github.com/pelletier/go-toml/v2"
)

const (
	// Namespace is the managed-config store the client reads from.
	Namespace = "ai.c1"

	// KeyTenantDomain is the key holding the full DNS host of the tenant's
	// control plane, for example "acme.conductor.one".
	KeyTenantDomain = "TenantDomain"
)

// Config holds the managed device configuration values the client understands.
// Keys the client does not recognize are ignored.
type Config struct {
	// TenantDomain is the full DNS host of the tenant's control plane, for
	// example "acme.conductor.one". It is empty when unset or invalid.
	TenantDomain string
}

// ControlPlaneURL returns the tenant control-plane URL ("https://" + TenantDomain),
// or an empty string when no valid TenantDomain is configured.
func (c Config) ControlPlaneURL() string {
	if c.TenantDomain == "" {
		return ""
	}
	return "https://" + c.TenantDomain
}

// Read returns the managed device configuration for the current operating
// system. It never returns an error: an absent, unreadable, or malformed store
// yields a zero Config.
func Read() Config {
	return configFromMap(readManagedConfig())
}

// configFromMap extracts the recognized keys from a raw key/value view of the
// store, ignoring unknown keys and dropping values that do not satisfy the
// contract.
func configFromMap(m map[string]string) Config {
	var c Config
	if domain := strings.TrimSpace(m[KeyTenantDomain]); isValidTenantDomain(domain) {
		c.TenantDomain = domain
	}
	return c
}

// isValidTenantDomain reports whether s is a full DNS host suitable for use as a
// control-plane locator: at least three dot-separated labels (for example
// "acme.conductor.one" or "acme.eu.c1.ai") with no scheme, path, port, or
// whitespace. A bare tenant slug is intentionally rejected.
func isValidTenantDomain(s string) bool {
	if strings.ContainsAny(s, "/:@ \t\r\n") {
		return false
	}
	labels := strings.Split(s, ".")
	if len(labels) < 3 {
		return false
	}
	for _, l := range labels {
		if l == "" {
			return false
		}
	}
	return true
}

// parseManagedTOML parses the Linux managed-config TOML into a flat map of
// top-level string values. Nested tables and non-string values are ignored.
// Malformed input yields a nil map.
func parseManagedTOML(data []byte) map[string]string {
	raw := map[string]any{}
	if err := toml.Unmarshal(data, &raw); err != nil {
		return nil
	}
	return stringValues(raw)
}

// stringValues returns only the top-level string entries of raw.
func stringValues(raw map[string]any) map[string]string {
	out := make(map[string]string, len(raw))
	for k, v := range raw {
		if s, ok := v.(string); ok {
			out[k] = s
		}
	}
	return out
}

// parseDefaultsValue converts the raw output of `defaults read <domain> <key>`
// into a single-key map. Empty output yields a nil map.
func parseDefaultsValue(key string, out []byte) map[string]string {
	v := strings.TrimSpace(string(out))
	if v == "" {
		return nil
	}
	return map[string]string{key: v}
}
