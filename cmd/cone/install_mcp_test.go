package main

import (
	"testing"
)

func TestParseTenantHost(t *testing.T) {
	tests := []struct {
		name     string
		clientID string
		want     string
		wantErr  bool
	}{
		{
			name:     "standard client-id",
			clientID: "myapp@mycompany.conductor.one/api",
			want:     "mycompany.conductor.one",
		},
		{
			name:     "client-id with longer path",
			clientID: "svc@staging.conductor.one/api/v1/something",
			want:     "staging.conductor.one",
		},
		{
			name:     "missing @ separator",
			clientID: "nohostpart",
			wantErr:  true,
		},
		{
			name:     "missing / after host",
			clientID: "name@hostonly",
			wantErr:  true,
		},
		{
			name:     "empty string",
			clientID: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTenantHost(tt.clientID)
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseTenantHost(%q) expected error, got %q", tt.clientID, got)
				}
				return
			}
			if err != nil {
				t.Errorf("parseTenantHost(%q) unexpected error: %v", tt.clientID, err)
				return
			}
			if got != tt.want {
				t.Errorf("parseTenantHost(%q) = %q, want %q", tt.clientID, got, tt.want)
			}
		})
	}
}
