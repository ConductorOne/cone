package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// entitlementPage is one canned SearchEntitlements response.
type entitlementPage struct {
	ids           []string
	nextPageToken string
}

// serveEntitlementPages returns a server that hands back pages in order and
// records the page_token it was sent for each request.
func serveEntitlementPages(t *testing.T, pages []entitlementPage) (*httptest.Server, *[]string) {
	t.Helper()
	seen := make([]string, 0, len(pages))
	call := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req map[string]any
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("decode request: %v", err)
		}
		token, _ := req["pageToken"].(string)
		seen = append(seen, token)

		if call >= len(pages) {
			t.Errorf("server called %d times, only %d pages configured", call+1, len(pages))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		page := pages[call]
		call++

		entries := make([]string, 0, len(page.ids))
		for _, id := range page.ids {
			entries = append(entries, fmt.Sprintf(
				`{"entitlement":{"appEntitlement":{"id":%q,"appId":"app1","displayName":%q}}}`, id, id))
		}
		body := fmt.Sprintf(`{"list":[%s]`, strings.Join(entries, ","))
		if page.nextPageToken != "" {
			body += fmt.Sprintf(`,"nextPageToken":%q`, page.nextPageToken)
		}
		body += "}"

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)

	return server, &seen
}

func entitlementIDs(got []*EntitlementWithBindings) []string {
	ids := make([]string, 0, len(got))
	for _, e := range got {
		ids = append(ids, e.Entitlement.GetAppId()+"/"+StringFromPtr(e.Entitlement.ID))
	}
	return ids
}

// TestSearchEntitlementsFollowsNextPageToken is the regression this whole change
// exists for: before it, the client made exactly one call and silently dropped
// every entitlement past the first page.
func TestSearchEntitlementsFollowsNextPageToken(t *testing.T) {
	server, seen := serveEntitlementPages(t, []entitlementPage{
		{ids: []string{"ent1", "ent2"}, nextPageToken: "tok1"},
		{ids: []string{"ent3"}, nextPageToken: "tok2"},
		{ids: []string{"ent4"}},
	})

	c := newPaperSecretTestClient(server.URL, server.Client())
	got, err := c.SearchEntitlements(context.Background(), &SearchEntitlementsFilter{})
	if err != nil {
		t.Fatalf("SearchEntitlements: %v", err)
	}

	if len(got) != 4 {
		t.Fatalf("got %d entitlements (%v), want 4 across 3 pages", len(got), entitlementIDs(got))
	}
	want := []string{"app1/ent1", "app1/ent2", "app1/ent3", "app1/ent4"}
	for i, w := range want {
		if entitlementIDs(got)[i] != w {
			t.Errorf("entitlement %d = %q, want %q (page order must be preserved)", i, entitlementIDs(got)[i], w)
		}
	}

	wantTokens := []string{"", "tok1", "tok2"}
	if len(*seen) != len(wantTokens) {
		t.Fatalf("server saw %d requests (%v), want %d", len(*seen), *seen, len(wantTokens))
	}
	for i, w := range wantTokens {
		if (*seen)[i] != w {
			t.Errorf("request %d sent pageToken %q, want %q", i, (*seen)[i], w)
		}
	}
}

// TestSearchEntitlementsDoesNotStopOnEmptyPage covers the server's post-filter:
// granted_status is applied after the page is cut, so an intermediate page can be
// empty while later pages still hold results. Stopping on a short page would
// silently truncate.
func TestSearchEntitlementsDoesNotStopOnEmptyPage(t *testing.T) {
	server, _ := serveEntitlementPages(t, []entitlementPage{
		{ids: []string{"ent1"}, nextPageToken: "tok1"},
		{ids: nil, nextPageToken: "tok2"},
		{ids: []string{"ent2"}},
	})

	c := newPaperSecretTestClient(server.URL, server.Client())
	got, err := c.SearchEntitlements(context.Background(), &SearchEntitlementsFilter{})
	if err != nil {
		t.Fatalf("SearchEntitlements: %v", err)
	}

	if len(got) != 2 {
		t.Fatalf("got %d entitlements (%v), want 2: an empty middle page must not end pagination",
			len(got), entitlementIDs(got))
	}
}

// TestSearchEntitlementsEmptyResultIsNotAnError pins that a wholly empty result
// returns an empty slice. The wire format omits an empty repeated field, so the
// list arrives nil.
func TestSearchEntitlementsEmptyResultIsNotAnError(t *testing.T) {
	server, _ := serveEntitlementPages(t, []entitlementPage{{ids: nil}})

	c := newPaperSecretTestClient(server.URL, server.Client())
	got, err := c.SearchEntitlements(context.Background(), &SearchEntitlementsFilter{})
	if err != nil {
		t.Fatalf("SearchEntitlements: %v", err)
	}
	if len(got) != 0 {
		t.Fatalf("got %d entitlements, want 0", len(got))
	}
}

// TestSearchEntitlementsStopsOnStuckPageToken keeps a server that never advances
// the token from spinning the client forever.
func TestSearchEntitlementsStopsOnStuckPageToken(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		calls++
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"list":[{"entitlement":{"appEntitlement":{"id":"ent1","appId":"app1"}}}],"nextPageToken":"stuck"}`))
	}))
	defer server.Close()

	c := newPaperSecretTestClient(server.URL, server.Client())
	_, err := c.SearchEntitlements(context.Background(), &SearchEntitlementsFilter{})
	if err == nil {
		t.Fatal("SearchEntitlements returned nil error on a non-advancing page token")
	}
	if !strings.Contains(err.Error(), "not advancing") {
		t.Errorf("error = %v, want a non-advancing pagination error", err)
	}
	if calls > maxRepeatedSearchEntitlementsPageToken+2 {
		t.Errorf("server called %d times, guard should have stopped it near %d",
			calls, maxRepeatedSearchEntitlementsPageToken)
	}
}
