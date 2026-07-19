package client

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func TestRequirePaperSecretAgeSuite(t *testing.T) {
	if err := requirePaperSecretAgeSuite("create", requiredAgeSuite); err != nil {
		t.Fatalf("requirePaperSecretAgeSuite() rejected hybrid suite: %v", err)
	}

	for _, returned := range []string{"", "AGE_SUITE_UNSPECIFIED", "AGE_SUITE_X25519", "AGE_SUITE_FUTURE"} {
		t.Run(returned, func(t *testing.T) {
			err := requirePaperSecretAgeSuite("create", returned)
			if err == nil {
				t.Fatal("requirePaperSecretAgeSuite() accepted mismatched suite")
			}
			var mismatch *AgeSuiteMismatchError
			if !errors.As(err, &mismatch) {
				t.Fatalf("error type = %T, want *AgeSuiteMismatchError", err)
			}
			if !mismatch.Temporary() {
				t.Fatal("AgeSuiteMismatchError must be retryable")
			}
		})
	}
}

func TestCreateSecretRequiresAndVerifiesHybridAgeSuite(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		external bool
	}{
		{name: "internal", path: "/api/v1/secrets/internal"},
		{name: "external", path: "/api/v1/secrets/external", external: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != tt.path {
					t.Errorf("path = %q, want %q", r.URL.Path, tt.path)
				}
				var body map[string]any
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					t.Fatalf("Decode() unexpected error: %v", err)
				}
				if body["requiredAgeSuite"] != requiredAgeSuite {
					t.Errorf("requiredAgeSuite = %v, want %s", body["requiredAgeSuite"], requiredAgeSuite)
				}
				if r.Header.Get("Authorization") != "Bearer test-token" {
					t.Errorf("Authorization = %q, want configured client token", r.Header.Get("Authorization"))
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"vaultId":"vault-id","ageRecipient":"age1pq1recipient","ageSuite":"AGE_SUITE_MLKEM768X25519"}`))
			}))
			defer server.Close()
			baseURL, err := url.Parse(server.URL)
			if err != nil {
				t.Fatal(err)
			}
			httpClient := server.Client()
			httpClient.Transport = authorizationTransport{base: httpClient.Transport}
			c := &client{httpClient: httpClient, baseURL: baseURL}

			var response *shared.PaperSecretServiceCreateResponse
			if tt.external {
				response, err = c.CreateExternalSecret(context.Background(), &shared.PaperSecretServiceCreateExternalRequest{})
			} else {
				response, err = c.CreateInternalSecret(context.Background(), &shared.PaperSecretServiceCreateInternalRequest{})
			}
			if err != nil {
				t.Fatalf("create unexpected error: %v", err)
			}
			if response.GetVaultID() == nil || *response.GetVaultID() != "vault-id" {
				t.Fatalf("VaultID = %v, want vault-id", response.GetVaultID())
			}
		})
	}
}

type authorizationTransport struct {
	base http.RoundTripper
}

func (t authorizationTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer test-token")
	return t.base.RoundTrip(req)
}

func TestCreateSecretRejectsDowngradedAgeSuite(t *testing.T) {
	for _, suite := range []string{"", "AGE_SUITE_UNSPECIFIED", "AGE_SUITE_X25519"} {
		t.Run(suite, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(map[string]string{"ageSuite": suite})
			}))
			defer server.Close()
			baseURL, err := url.Parse(server.URL)
			if err != nil {
				t.Fatal(err)
			}
			c := &client{httpClient: server.Client(), baseURL: baseURL}
			_, err = c.CreateInternalSecret(context.Background(), &shared.PaperSecretServiceCreateInternalRequest{})
			var mismatch *AgeSuiteMismatchError
			if !errors.As(err, &mismatch) {
				t.Fatalf("error = %v, want *AgeSuiteMismatchError", err)
			}
		})
	}
}
