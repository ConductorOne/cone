package managedconfig

import "testing"

func TestConfigControlPlaneURL(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		want   string
	}{
		{name: "empty", config: Config{}, want: ""},
		{name: "commercial", config: Config{TenantDomain: "acme.conductor.one"}, want: "https://acme.conductor.one"},
		{name: "eu", config: Config{TenantDomain: "acme.eu.c1.ai"}, want: "https://acme.eu.c1.ai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.config.ControlPlaneURL(); got != tt.want {
				t.Errorf("ControlPlaneURL() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestConfigFromMap(t *testing.T) {
	tests := []struct {
		name string
		in   map[string]string
		want string
	}{
		{name: "nil map", in: nil, want: ""},
		{name: "empty map", in: map[string]string{}, want: ""},
		{name: "valid commercial", in: map[string]string{KeyTenantDomain: "acme.conductor.one"}, want: "acme.conductor.one"},
		{name: "valid eu", in: map[string]string{KeyTenantDomain: "acme.eu.c1.ai"}, want: "acme.eu.c1.ai"},
		{name: "trims whitespace", in: map[string]string{KeyTenantDomain: "  acme.conductor.one\n"}, want: "acme.conductor.one"},
		{name: "unknown keys ignored", in: map[string]string{"SomethingElse": "value", KeyTenantDomain: "acme.conductor.one"}, want: "acme.conductor.one"},
		{name: "bare slug rejected", in: map[string]string{KeyTenantDomain: "acme"}, want: ""},
		{name: "two-label rejected", in: map[string]string{KeyTenantDomain: "conductor.one"}, want: ""},
		{name: "scheme rejected", in: map[string]string{KeyTenantDomain: "https://acme.conductor.one"}, want: ""},
		{name: "path rejected", in: map[string]string{KeyTenantDomain: "acme.conductor.one/foo"}, want: ""},
		{name: "empty value", in: map[string]string{KeyTenantDomain: ""}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configFromMap(tt.in).TenantDomain; got != tt.want {
				t.Errorf("configFromMap(%v).TenantDomain = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestIsValidTenantDomain(t *testing.T) {
	valid := []string{"acme.conductor.one", "acme.eu.c1.ai", "a.b.c"}
	for _, s := range valid {
		if !isValidTenantDomain(s) {
			t.Errorf("isValidTenantDomain(%q) = false, want true", s)
		}
	}
	invalid := []string{"", "acme", "conductor.one", "acme.conductor.one/x", "https://acme.conductor.one", "acme.conductor.one:8080", "acme..one", "acme conductor one"}
	for _, s := range invalid {
		if isValidTenantDomain(s) {
			t.Errorf("isValidTenantDomain(%q) = true, want false", s)
		}
	}
}

func TestParseManagedTOML(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string // expected TenantDomain after configFromMap; "" means absent/ignored
	}{
		{name: "valid", in: "TenantDomain = \"acme.conductor.one\"\n", want: "acme.conductor.one"},
		{name: "unknown keys ignored", in: "TenantDomain = \"acme.conductor.one\"\nUnknown = \"x\"\nCount = 3\n", want: "acme.conductor.one"},
		{name: "absent key", in: "SomethingElse = \"x\"\n", want: ""},
		{name: "empty file", in: "", want: ""},
		{name: "malformed", in: "TenantDomain = = = broken", want: ""},
		{name: "non-string value ignored", in: "TenantDomain = 123\n", want: ""},
		{name: "nested table ignored", in: "[section]\nTenantDomain = \"acme.conductor.one\"\n", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := configFromMap(parseManagedTOML([]byte(tt.in))).TenantDomain
			if got != tt.want {
				t.Errorf("parseManagedTOML(%q) TenantDomain = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestParseDefaultsValue(t *testing.T) {
	if m := parseDefaultsValue(KeyTenantDomain, []byte("acme.conductor.one\n")); m[KeyTenantDomain] != "acme.conductor.one" {
		t.Errorf("parseDefaultsValue got %v", m)
	}
	if m := parseDefaultsValue(KeyTenantDomain, []byte("   \n")); m != nil {
		t.Errorf("parseDefaultsValue(empty) = %v, want nil", m)
	}
}
