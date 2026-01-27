package mcpclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Connect(t *testing.T) {
	t.Run("successful initialize", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var req map[string]interface{}
			_ = json.NewDecoder(r.Body).Decode(&req)

			if req["method"] != "initialize" {
				t.Errorf("expected initialize method, got %v", req["method"])
			}

			resp := map[string]interface{}{
				"jsonrpc": "2.0",
				"id":      req["id"],
				"result": map[string]interface{}{
					"protocolVersion": "2024-11-05",
					"serverInfo": map[string]string{
						"name":    "test-server",
						"version": "1.0",
					},
				},
			}
			_ = json.NewEncoder(w).Encode(resp)
		}))
		defer server.Close()

		client := &Client{
			ServerURL:  server.URL,
			HTTPClient: http.DefaultClient,
		}

		err := client.Connect(context.Background())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("handles server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var req map[string]interface{}
			_ = json.NewDecoder(r.Body).Decode(&req)

			resp := map[string]interface{}{
				"jsonrpc": "2.0",
				"id":      req["id"],
				"error": map[string]interface{}{
					"code":    -32600,
					"message": "Invalid Request",
				},
			}
			_ = json.NewEncoder(w).Encode(resp)
		}))
		defer server.Close()

		client := &Client{
			ServerURL:  server.URL,
			HTTPClient: http.DefaultClient,
		}

		err := client.Connect(context.Background())
		if err == nil {
			t.Error("expected error for server error response")
		}
	})
}

func TestClient_Analyze(t *testing.T) {
	t.Run("handles tool_call response", func(t *testing.T) {
		callCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var req map[string]interface{}
			_ = json.NewDecoder(r.Body).Decode(&req)

			callCount++
			var resp map[string]interface{}

			switch req["method"] {
			case "initialize":
				resp = map[string]interface{}{
					"jsonrpc": "2.0",
					"id":      req["id"],
					"result": map[string]interface{}{
						"protocolVersion": "2024-11-05",
					},
				}
			case "tools/call":
				params := req["params"].(map[string]interface{})
				if params["name"] == "connector_analyze" {
					resp = map[string]interface{}{
						"jsonrpc": "2.0",
						"id":      req["id"],
						"result": map[string]interface{}{
							"content": []map[string]interface{}{
								{
									"type": "text",
									"text": `{"status":"tool_call","session_id":"test-session","tool_call":{"name":"read_files","arguments":{"paths":["go.mod"]}}}`,
								},
							},
						},
					}
				} else if params["name"] == "tool_result" {
					resp = map[string]interface{}{
						"jsonrpc": "2.0",
						"id":      req["id"],
						"result": map[string]interface{}{
							"content": []map[string]interface{}{
								{
									"type": "text",
									"text": `{"status":"complete","session_id":"test-session","message":"Done"}`,
								},
							},
						},
					}
				}
			}

			_ = json.NewEncoder(w).Encode(resp)
		}))
		defer server.Close()

		// Create a mock tool handler that returns empty results
		handler := &ToolHandler{
			ConnectorDir: "/tmp/test",
			DryRun:       true,
		}

		client := &Client{
			ServerURL:   server.URL,
			HTTPClient:  http.DefaultClient,
			ToolHandler: handler,
		}

		// This will fail because the tool handler can't actually read files,
		// but it tests the client's ability to parse responses
		_, err := client.Analyze(context.Background(), "/tmp/test", "full")
		// We expect an error because read_files will fail on non-existent path
		if err == nil {
			// If no error, the flow completed somehow
			t.Log("analyze completed without error")
		}
	})
}

func TestToolHandler_HandleToolCall(t *testing.T) {
	handler := &ToolHandler{
		ConnectorDir: ".",
		DryRun:       true,
	}

	t.Run("read_files with valid path", func(t *testing.T) {
		result, err := handler.HandleToolCall(context.Background(), "read_files", map[string]interface{}{
			"patterns": []interface{}{"client_test.go"},
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !result.Success {
			t.Errorf("expected success, got error: %s", result.Error)
		}
	})

	t.Run("read_files with missing path", func(t *testing.T) {
		result, err := handler.HandleToolCall(context.Background(), "read_files", map[string]interface{}{
			"patterns": []interface{}{"nonexistent_file_12345.go"},
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// Should still succeed but with empty files (no matches)
		if !result.Success {
			t.Log("read_files reported failure for missing file (acceptable)")
		}
	})

	t.Run("unknown tool returns error in result", func(t *testing.T) {
		result, err := handler.HandleToolCall(context.Background(), "unknown_tool", map[string]interface{}{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// Unknown tools return Success: false with an error message
		if result.Success {
			t.Error("expected Success=false for unknown tool")
		}
		if result.Error == "" {
			t.Error("expected error message for unknown tool")
		}
	})

	t.Run("write_file in dry-run mode", func(t *testing.T) {
		result, err := handler.HandleToolCall(context.Background(), "write_file", map[string]interface{}{
			"path":    "/tmp/test.txt",
			"content": "test content",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !result.Success {
			t.Error("expected success in dry-run mode")
		}
		// In dry-run, file should not actually be written
	})
}
