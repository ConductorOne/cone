package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

// searchPage builds a search response page whose single entitlement points at
// its own expanded[0] app, mirroring the server's per-response indexing.
func searchPage(entID, appID, nextToken string) map[string]any {
	page := map[string]any{
		"list": []any{
			map[string]any{
				"entitlement": map[string]any{
					"appEntitlement": map[string]any{"id": entID, "appId": appID},
					"appPath":        "$.expanded[0]",
				},
			},
		},
		"expanded": []any{
			map[string]any{"@type": atTypeApp, "id": appID, "displayName": "App " + appID},
		},
	}
	if nextToken != "" {
		page["nextPageToken"] = nextToken
	}
	return page
}

func TestSearchEntitlementsFollowsPageToken(t *testing.T) {
	var tokens []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req shared.RequestCatalogSearchServiceSearchEntitlementsRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("Decode() unexpected error: %v", err)
		}
		if req.PageSize == nil || *req.PageSize != 100 {
			t.Errorf("page size = %v, want 100", req.PageSize)
		}
		tokens = append(tokens, StringFromPtr(req.PageToken))

		var page map[string]any
		switch StringFromPtr(req.PageToken) {
		case "":
			page = searchPage("e1", "a1", "t1")
		case "t1":
			page = searchPage("e2", "a2", "t2")
		case "t2":
			page = searchPage("e3", "a3", "")
		default:
			t.Errorf("unexpected page token %q", StringFromPtr(req.PageToken))
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(page)
	}))
	defer server.Close()

	c := newPaperSecretTestClient(server.URL, server.Client())
	got, err := c.SearchEntitlements(context.Background(), &SearchEntitlementsFilter{})
	if err != nil {
		t.Fatalf("SearchEntitlements() unexpected error: %v", err)
	}

	if want := []string{"", "t1", "t2"}; fmt.Sprint(tokens) != fmt.Sprint(want) {
		t.Fatalf("page tokens = %v, want %v", tokens, want)
	}
	if len(got) != 3 {
		t.Fatalf("got %d entitlements, want 3", len(got))
	}
	for i, id := range []string{"1", "2", "3"} {
		if gotID := StringFromPtr(got[i].Entitlement.ID); gotID != "e"+id {
			t.Errorf("entitlement[%d] id = %q, want e%s", i, gotID, id)
		}
		// Every page uses expanded[0]; each entitlement must resolve to its own page's app.
		app := GetExpanded[shared.App](got[i], ExpandedApp)
		if app == nil || StringFromPtr(app.ID) != "a"+id {
			t.Errorf("entitlement[%d] expanded app = %v, want a%s", i, app, id)
		}
	}
}

func TestSearchEntitlementsHonorsContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// Always advertise another page; cancel after the first response.
		cancel()
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(searchPage("e", "a", "next"))
	}))
	defer server.Close()

	c := newPaperSecretTestClient(server.URL, server.Client())
	if _, err := c.SearchEntitlements(ctx, &SearchEntitlementsFilter{}); err == nil {
		t.Fatal("SearchEntitlements() error = nil, want context error")
	}
}
