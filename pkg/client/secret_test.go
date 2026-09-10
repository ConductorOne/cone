package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestSearchSecretsSharedWithMePaginatesWithFiltersPreserved(t *testing.T) {
	type capturedRequest struct {
		body map[string]any
	}
	var requests []capturedRequest
	page := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/search/secrets/shared_with_me" {
			t.Errorf("path = %q, want /api/v1/search/secrets/shared_with_me", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("Decode() unexpected error: %v", err)
		}
		requests = append(requests, capturedRequest{body: body})
		page++
		w.Header().Set("Content-Type", "application/json")
		resp := shared.PaperSecretServiceSearchResponse{}
		if page == 1 {
			resp.NextPageToken = new("page-two")
			resp.List = []shared.PaperSecret{{VaultID: new("vault-1")}}
		} else {
			resp.List = []shared.PaperSecret{{VaultID: new("vault-2")}}
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("Encode() unexpected error: %v", err)
		}
	}))
	defer server.Close()

	c := newPaperSecretTestClient(server.URL, server.Client())
	req := &shared.PaperSecretServiceSearchSecretsSharedWithMeRequest{
		Query:      new("deploy"),
		Statuses:   []shared.PaperSecretServiceSearchSecretsSharedWithMeRequestStatuses{shared.PaperSecretServiceSearchSecretsSharedWithMeRequestStatusesSecretStatusActive},
		SecretType: shared.PaperSecretServiceSearchSecretsSharedWithMeRequestSecretTypeSecretTypeText.ToPointer(),
	}
	secrets, err := c.SearchSecretsSharedWithMe(context.Background(), req)
	if err != nil {
		t.Fatalf("SearchSecretsSharedWithMe() unexpected error: %v", err)
	}
	if len(secrets) != 2 {
		t.Fatalf("secrets returned = %d, want 2 across two pages", len(secrets))
	}
	if len(requests) != 2 {
		t.Fatalf("requests sent = %d, want 2", len(requests))
	}
	for i, cr := range requests {
		if _, hasUserID := cr.body["userId"]; hasUserID {
			t.Errorf("request %d carried userId; the endpoint is caller-bound", i+1)
		}
		if _, hasSortBy := cr.body["sortBy"]; hasSortBy {
			t.Errorf("request %d carried sortBy; the endpoint does not accept it", i+1)
		}
		if _, hasSharing := cr.body["sharingMode"]; hasSharing {
			t.Errorf("request %d carried sharingMode; the endpoint does not accept it", i+1)
		}
		if got, _ := cr.body["query"].(string); got != "deploy" {
			t.Errorf("request %d query = %v, want deploy preserved across pages", i+1, got)
		}
	}
	if _, hasToken := requests[1].body["pageToken"]; !hasToken {
		t.Error("second request must carry pageToken from first response")
	}
	if req.PageToken != nil {
		t.Errorf("caller request was mutated: PageToken = %q, want unchanged", *req.PageToken)
	}
}

func TestSearchSecretsSharedWithMeStopsOnRepeatedPageToken(t *testing.T) {
	var requests int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		w.Header().Set("Content-Type", "application/json")
		resp := shared.PaperSecretServiceSearchResponse{
			NextPageToken: new("same-token"),
			List:          []shared.PaperSecret{{VaultID: new("vault-1")}},
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("Encode() unexpected error: %v", err)
		}
	}))
	defer server.Close()

	c := newPaperSecretTestClient(server.URL, server.Client())
	_, err := c.SearchSecretsSharedWithMe(context.Background(), &shared.PaperSecretServiceSearchSecretsSharedWithMeRequest{})
	if err == nil {
		t.Fatal("a constant next_page_token must abort the listing instead of looping")
	}
	if !strings.Contains(err.Error(), "next_page_token") {
		t.Fatalf("error = %v, want repeated-token message", err)
	}
	if requests != 2 {
		t.Fatalf("requests sent = %d, want 2 (the repeated token must stop the loop)", requests)
	}
}

func TestSearchSecretsSharedWithMeStopsOnCyclingPageToken(t *testing.T) {
	// A server that alternates tokens (A → B → A → …) defeats a guard that only
	// compares against the immediately preceding token; every token the server
	// has handed out must be tracked.
	var requests int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		w.Header().Set("Content-Type", "application/json")
		tokens := []string{"a", "b", "a"}
		resp := shared.PaperSecretServiceSearchResponse{
			NextPageToken: &tokens[(requests-1)%len(tokens)],
			List:          []shared.PaperSecret{{VaultID: new("vault-1")}},
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("Encode() unexpected error: %v", err)
		}
	}))
	defer server.Close()

	c := newPaperSecretTestClient(server.URL, server.Client())
	_, err := c.SearchSecretsSharedWithMe(context.Background(), &shared.PaperSecretServiceSearchSecretsSharedWithMeRequest{})
	if err == nil {
		t.Fatal("a cycling next_page_token must abort the listing instead of looping")
	}
	if !strings.Contains(err.Error(), "next_page_token") {
		t.Fatalf("error = %v, want repeated-token message", err)
	}
	if requests != 3 {
		t.Fatalf("requests sent = %d, want 3 (the token cycle must stop the loop)", requests)
	}
}

func TestSearchSecretsSharedWithMeStopsOnUnboundedFreshTokens(t *testing.T) {
	// A server that mints a fresh unique token on every page defeats the
	// repeat checks; the page cap is what finally terminates the listing.
	var requests int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		w.Header().Set("Content-Type", "application/json")
		token := fmt.Sprintf("fresh-%d", requests)
		resp := shared.PaperSecretServiceSearchResponse{
			NextPageToken: &token,
			List:          []shared.PaperSecret{{VaultID: new("vault-1")}},
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("Encode() unexpected error: %v", err)
		}
	}))
	defer server.Close()

	c := newPaperSecretTestClient(server.URL, server.Client())
	_, err := c.SearchSecretsSharedWithMe(context.Background(), &shared.PaperSecretServiceSearchSecretsSharedWithMeRequest{})
	if err == nil {
		t.Fatal("an endless stream of fresh tokens must abort at the page cap instead of looping")
	}
	if !strings.Contains(err.Error(), "pages") {
		t.Fatalf("error = %v, want page-cap message", err)
	}
	if requests != paginationPageCap {
		t.Fatalf("requests sent = %d, want %d (the cap must stop the loop)", requests, paginationPageCap)
	}
}

func TestSearchMySecretsDoesNotMutateCallerRequest(t *testing.T) {
	var requests int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		w.Header().Set("Content-Type", "application/json")
		resp := shared.PaperSecretServiceSearchResponse{}
		if requests == 1 {
			resp.NextPageToken = new("page-two")
			resp.List = []shared.PaperSecret{{VaultID: new("vault-1")}}
		} else {
			resp.List = []shared.PaperSecret{{VaultID: new("vault-2")}}
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("Encode() unexpected error: %v", err)
		}
	}))
	defer server.Close()

	c := newPaperSecretTestClient(server.URL, server.Client())
	req := &shared.PaperSecretServiceSearchMySecretsRequest{}
	if _, err := c.SearchMySecrets(context.Background(), req); err != nil {
		t.Fatalf("SearchMySecrets() unexpected error: %v", err)
	}
	if req.PageToken != nil {
		t.Fatalf("caller request was mutated: PageToken = %q, want unchanged", *req.PageToken)
	}
}

func TestSearchSecretsSharedWithMeNeverSendsUserID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("Decode() unexpected error: %v", err)
		}
		if _, hasUserID := body["userId"]; hasUserID {
			t.Error("request carried userId; the endpoint is caller-bound")
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(shared.PaperSecretServiceSearchResponse{}); err != nil {
			t.Errorf("Encode() unexpected error: %v", err)
		}
	}))
	defer server.Close()

	c := newPaperSecretTestClient(server.URL, server.Client())
	if _, err := c.SearchSecretsSharedWithMe(context.Background(), &shared.PaperSecretServiceSearchSecretsSharedWithMeRequest{}); err != nil {
		t.Fatalf("SearchSecretsSharedWithMe() unexpected error: %v", err)
	}
}

func TestSearchSecretsSharedWithMePreservesHTTPErrorMapping(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"code":"permission_denied"}`))
	}))
	defer server.Close()

	c := newPaperSecretTestClient(server.URL, server.Client())
	_, err := c.SearchSecretsSharedWithMe(context.Background(), &shared.PaperSecretServiceSearchSecretsSharedWithMeRequest{})
	var httpErr *HTTPError
	if !errors.As(err, &httpErr) {
		t.Fatalf("error = %T, want *HTTPError", err)
	}
	if httpErr.StatusCode != http.StatusForbidden {
		t.Fatalf("status code = %d, want %d", httpErr.StatusCode, http.StatusForbidden)
	}
}

func TestSearchSecretsSharedWithMeHonorsContextDeadline(t *testing.T) {
	httpClient := &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		<-req.Context().Done()
		return nil, req.Context().Err()
	})}
	c := newPaperSecretTestClient("https://example.invalid", httpClient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	_, err := c.SearchSecretsSharedWithMe(ctx, &shared.PaperSecretServiceSearchSecretsSharedWithMeRequest{})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error = %v, want context deadline exceeded", err)
	}
}
