package client

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	sdk "github.com/conductorone/conductorone-sdk-go"
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func TestRequirePaperSecretAgeSuite(t *testing.T) {
	hybrid := shared.PaperSecretServiceCreateResponseAgeSuiteAgeSuiteMlkem768X25519
	if err := requirePaperSecretAgeSuite("create", &hybrid); err != nil {
		t.Fatalf("requirePaperSecretAgeSuite() rejected hybrid suite: %v", err)
	}

	tests := []struct {
		name     string
		returned *shared.PaperSecretServiceCreateResponseAgeSuite
	}{
		{name: "missing"},
		{
			name:     "unspecified",
			returned: shared.PaperSecretServiceCreateResponseAgeSuiteAgeSuiteUnspecified.ToPointer(),
		},
		{
			name:     "x25519",
			returned: shared.PaperSecretServiceCreateResponseAgeSuiteAgeSuiteX25519.ToPointer(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := requirePaperSecretAgeSuite("create", tt.returned)
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
				if tt.external {
					var request shared.PaperSecretServiceCreateExternalRequest
					if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
						t.Errorf("Decode() unexpected error: %v", err)
						return
					}
					if request.GetRequiredAgeSuite() == nil || *request.GetRequiredAgeSuite() != shared.RequiredAgeSuiteAgeSuiteMlkem768X25519 {
						t.Errorf("required age suite = %v, want %s", request.GetRequiredAgeSuite(), shared.RequiredAgeSuiteAgeSuiteMlkem768X25519)
					}
				} else {
					var request shared.PaperSecretServiceCreateInternalRequest
					if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
						t.Errorf("Decode() unexpected error: %v", err)
						return
					}
					if request.GetRequiredAgeSuite() == nil || *request.GetRequiredAgeSuite() != shared.PaperSecretServiceCreateInternalRequestRequiredAgeSuiteAgeSuiteMlkem768X25519 {
						t.Errorf("required age suite = %v, want %s", request.GetRequiredAgeSuite(), shared.PaperSecretServiceCreateInternalRequestRequiredAgeSuiteAgeSuiteMlkem768X25519)
					}
				}
				if r.Header.Get("Authorization") != "Bearer test-token" {
					t.Errorf("Authorization = %q, want configured client token", r.Header.Get("Authorization"))
				}

				vaultID := "vault-id"
				ageRecipient := "age1pq1recipient"
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(shared.PaperSecretServiceCreateResponse{
					VaultID:      &vaultID,
					AgeRecipient: &ageRecipient,
					AgeSuite:     shared.PaperSecretServiceCreateResponseAgeSuiteAgeSuiteMlkem768X25519.ToPointer(),
				}); err != nil {
					t.Errorf("Encode() unexpected error: %v", err)
				}
			}))
			defer server.Close()

			httpClient := server.Client()
			httpClient.Transport = authorizationTransport{base: httpClient.Transport}
			c := newPaperSecretTestClient(server.URL, httpClient)

			var response *shared.PaperSecretServiceCreateResponse
			var err error
			if tt.external {
				request := &shared.PaperSecretServiceCreateExternalRequest{
					RequiredAgeSuite: shared.RequiredAgeSuiteAgeSuiteX25519.ToPointer(),
				}
				response, err = c.CreateExternalSecret(context.Background(), request)
				if request.GetRequiredAgeSuite() == nil || *request.GetRequiredAgeSuite() != shared.RequiredAgeSuiteAgeSuiteX25519 {
					t.Fatalf("caller request was mutated: %v", request.GetRequiredAgeSuite())
				}
			} else {
				request := &shared.PaperSecretServiceCreateInternalRequest{
					RequiredAgeSuite: shared.PaperSecretServiceCreateInternalRequestRequiredAgeSuiteAgeSuiteX25519.ToPointer(),
				}
				response, err = c.CreateInternalSecret(context.Background(), request)
				if request.GetRequiredAgeSuite() == nil || *request.GetRequiredAgeSuite() != shared.PaperSecretServiceCreateInternalRequestRequiredAgeSuiteAgeSuiteX25519 {
					t.Fatalf("caller request was mutated: %v", request.GetRequiredAgeSuite())
				}
			}
			if err != nil {
				t.Fatalf("create unexpected error: %v", err)
			}
			if response.GetVaultID() == nil || *response.GetVaultID() != "vault-id" {
				t.Fatalf("VaultID = %v, want vault-id", response.GetVaultID())
			}
			if response.GetAgeSuite() == nil || *response.GetAgeSuite() != shared.PaperSecretServiceCreateResponseAgeSuiteAgeSuiteMlkem768X25519 {
				t.Fatalf("AgeSuite = %v, want %s", response.GetAgeSuite(), shared.PaperSecretServiceCreateResponseAgeSuiteAgeSuiteMlkem768X25519)
			}
		})
	}
}

func TestCreateSecretRejectsDowngradedAgeSuite(t *testing.T) {
	tests := []struct {
		name  string
		suite *shared.PaperSecretServiceCreateResponseAgeSuite
	}{
		{name: "missing"},
		{
			name:  "unspecified",
			suite: shared.PaperSecretServiceCreateResponseAgeSuiteAgeSuiteUnspecified.ToPointer(),
		},
		{
			name:  "x25519",
			suite: shared.PaperSecretServiceCreateResponseAgeSuiteAgeSuiteX25519.ToPointer(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(shared.PaperSecretServiceCreateResponse{AgeSuite: tt.suite}); err != nil {
					t.Errorf("Encode() unexpected error: %v", err)
				}
			}))
			defer server.Close()

			c := newPaperSecretTestClient(server.URL, server.Client())
			_, err := c.CreateInternalSecret(context.Background(), &shared.PaperSecretServiceCreateInternalRequest{})
			var mismatch *AgeSuiteMismatchError
			if !errors.As(err, &mismatch) {
				t.Fatalf("error = %v, want *AgeSuiteMismatchError", err)
			}
		})
	}
}

func TestCreateSecretPreservesHTTPErrorMapping(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"code":"invalid_argument"}`))
	}))
	defer server.Close()

	c := newPaperSecretTestClient(server.URL, server.Client())
	_, err := c.CreateExternalSecret(context.Background(), &shared.PaperSecretServiceCreateExternalRequest{})
	var httpErr *HTTPError
	if !errors.As(err, &httpErr) {
		t.Fatalf("error = %T, want *HTTPError", err)
	}
	if httpErr.StatusCode != http.StatusBadRequest {
		t.Fatalf("status code = %d, want %d", httpErr.StatusCode, http.StatusBadRequest)
	}
	if httpErr.Body != `{"code":"invalid_argument"}` {
		t.Fatalf("body = %q", httpErr.Body)
	}
}

func TestCreateSecretHonorsContextDeadline(t *testing.T) {
	httpClient := &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		<-req.Context().Done()
		return nil, req.Context().Err()
	})}
	c := newPaperSecretTestClient("https://example.invalid", httpClient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	_, err := c.CreateInternalSecret(ctx, &shared.PaperSecretServiceCreateInternalRequest{})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error = %v, want context deadline exceeded", err)
	}
}

func newPaperSecretTestClient(serverURL string, httpClient *http.Client) *client {
	return &client{
		httpClient: httpClient,
		sdk: sdk.New(
			sdk.WithClient(httpClient),
			sdk.WithServerURL(serverURL),
		),
	}
}

type authorizationTransport struct {
	base http.RoundTripper
}

func (t authorizationTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer test-token")
	return t.base.RoundTrip(req)
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
