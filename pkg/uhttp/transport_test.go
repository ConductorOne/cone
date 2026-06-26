package uhttp

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"golang.org/x/oauth2"
)

func TestTokenSourceTripperKeepsAuthOnTrustedHostRedirect(t *testing.T) {
	const bearerToken = "test-token"

	var startAuth string
	var redirectedAuth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/start":
			startAuth = r.Header.Get("Authorization")
			http.Redirect(w, r, "/next", http.StatusFound)
		case "/next":
			redirectedAuth = r.Header.Get("Authorization")
			w.WriteHeader(http.StatusNoContent)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	client, err := NewClient(
		context.Background(),
		WithTokenSource(staticBearerTokenSource(bearerToken), serverURL.Host),
	)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL+"/start", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()

	expectedAuth := "Bearer " + bearerToken
	if startAuth != expectedAuth {
		t.Fatalf("start Authorization header = %q, want %q", startAuth, expectedAuth)
	}
	if redirectedAuth != expectedAuth {
		t.Fatalf("redirected Authorization header = %q, want %q", redirectedAuth, expectedAuth)
	}
}

func TestTokenSourceTripperSkipsAuthOnCrossHostRedirect(t *testing.T) {
	const bearerToken = "test-token"

	var trustedAuth string
	var redirectedAuth string
	redirectedServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redirectedAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer redirectedServer.Close()

	trustedServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		trustedAuth = r.Header.Get("Authorization")
		http.Redirect(w, r, redirectedServer.URL, http.StatusFound)
	}))
	defer trustedServer.Close()

	trustedURL, err := url.Parse(trustedServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	client, err := NewClient(
		context.Background(),
		WithTokenSource(staticBearerTokenSource(bearerToken), trustedURL.Host),
	)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, trustedServer.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()

	expectedAuth := "Bearer " + bearerToken
	if trustedAuth != expectedAuth {
		t.Fatalf("trusted Authorization header = %q, want %q", trustedAuth, expectedAuth)
	}
	if redirectedAuth != "" {
		t.Fatalf("cross-host redirected Authorization header = %q, want empty", redirectedAuth)
	}
}

func TestTokenSourceTripperRemovesAuthOnUntrustedHost(t *testing.T) {
	var auth string
	tripper := &tokenSourceTripper{
		next: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			auth = req.Header.Get("Authorization")
			return &http.Response{
				StatusCode: http.StatusNoContent,
				Body:       http.NoBody,
				Header:     make(http.Header),
				Request:    req,
			}, nil
		}),
		tokenSource: staticBearerTokenSource("fresh-token"),
		tokenHost:   "trusted.example",
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://other.example/path", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer copied-token")

	resp, err := tripper.RoundTrip(req)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()

	if auth != "" {
		t.Fatalf("untrusted host Authorization header = %q, want empty", auth)
	}
}

func staticBearerTokenSource(accessToken string) oauth2.TokenSource {
	return oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: accessToken,
	})
}

type roundTripFunc func(req *http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
