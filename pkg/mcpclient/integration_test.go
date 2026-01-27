package mcpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/conductorone/cone/pkg/mcpclient/mock"
)

// TestIntegration_FullAnalysisFlow runs a complete end-to-end integration test.
func TestIntegration_FullAnalysisFlow(t *testing.T) {
	// Flow:
	// 1. Creates temp connector directory with files
	// 2. Starts mock MCP server
	// 3. Connects cone client
	// 4. Runs analysis with tool callbacks
	// 5. Verifies completion
	// Create a temporary connector directory with some files
	tmpDir, err := os.MkdirTemp("", "connector-integration-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create minimal connector files that the mock server will request
	files := map[string]string{
		"go.mod":       "module test/connector\n\ngo 1.21\n",
		"connector.go": "package connector\n\ntype Connector struct{}\n",
		"README.md":    "# Test Connector\n",
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(tmpDir, name), []byte(content), 0600); err != nil {
			t.Fatalf("failed to create %s: %v", name, err)
		}
	}

	// Create a test scenario that only uses read_files (no user interaction)
	readOnlyScenario := &mock.Scenario{
		Name:        "read_only",
		Description: "Only reads files, no user interaction",
		ToolCalls: []mock.ToolCall{
			{
				Name: "read_files",
				Arguments: map[string]interface{}{
					"patterns": []string{"*.go", "*.md"},
				},
			},
		},
	}

	mockServer := mock.NewServer(readOnlyScenario)
	ts := httptest.NewServer(http.HandlerFunc(mockServer.HandleMCP))
	defer ts.Close()

	// Create tool handler that works non-interactively
	handler := &ToolHandler{
		ConnectorDir: tmpDir,
		DryRun:       true,
	}

	client := &Client{
		ServerURL:   ts.URL,
		HTTPClient:  http.DefaultClient,
		ToolHandler: handler,
	}

	ctx := context.Background()

	// Step 1: Connect (initialize)
	if err := client.Connect(ctx); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	// Step 2: Start analysis - this triggers tool callbacks
	result, err := client.Analyze(ctx, tmpDir, "full")
	if err != nil {
		t.Fatalf("Analyze failed: %v", err)
	}

	// Step 3: Verify we got a completion result
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	t.Logf("Analysis completed: %+v", result)

	// Step 4: Verify the mock server received the tool results
	toolResults := mockServer.ToolResults()
	if len(toolResults) != 1 {
		t.Errorf("expected 1 tool result (read_files), got %d", len(toolResults))
	}
}

// TestIntegration_ManualProtocolFlow tests the raw JSON-RPC protocol
// without going through the client abstraction.
func TestIntegration_ManualProtocolFlow(t *testing.T) {
	scenario := mock.HappyPathScenario()
	mockServer := mock.NewServer(scenario)

	ts := httptest.NewServer(http.HandlerFunc(mockServer.HandleMCP))
	defer ts.Close()

	httpClient := &http.Client{}

	// 1. Initialize
	resp := sendJSONRPC(t, httpClient, ts.URL, "initialize", map[string]interface{}{
		"protocolVersion": "2024-11-05",
		"clientInfo":      map[string]string{"name": "integration-test", "version": "1.0"},
	})
	assertNoRPCError(t, resp)
	t.Log("Initialize: OK")

	// 2. List tools
	resp = sendJSONRPC(t, httpClient, ts.URL, "tools/list", nil)
	assertNoRPCError(t, resp)
	t.Log("tools/list: OK")

	// 3. Call connector_analyze
	resp = sendJSONRPC(t, httpClient, ts.URL, "tools/call", map[string]interface{}{
		"name": "connector_analyze",
		"arguments": map[string]interface{}{
			"connector_path": "/test/connector",
		},
	})
	assertNoRPCError(t, resp)

	// Verify we got a tool_call response
	result := resp["result"].(map[string]interface{})
	if result["status"] != "tool_call" {
		t.Fatalf("expected status 'tool_call', got %v", result["status"])
	}
	toolCall := result["tool_call"].(map[string]interface{})
	if toolCall["name"] != "read_files" {
		t.Fatalf("expected first tool to be 'read_files', got %v", toolCall["name"])
	}
	t.Log("tools/call connector_analyze: OK, got read_files callback")

	// 4. Send tool results for each expected callback
	for i, tc := range scenario.ToolCalls {
		resp = sendJSONRPC(t, httpClient, ts.URL, "tool_result", map[string]interface{}{
			"tool": tc.Name,
			"result": map[string]interface{}{
				"success": true,
				"data":    map[string]interface{}{"mock": "data"},
			},
		})
		assertNoRPCError(t, resp)

		result := resp["result"].(map[string]interface{})
		status := result["status"].(string)
		t.Logf("tool_result %d (%s): status=%s", i+1, tc.Name, status)

		if i == len(scenario.ToolCalls)-1 {
			// Last one should be complete
			if status != "complete" {
				t.Errorf("expected final status 'complete', got %s", status)
			}
		} else {
			// Others should be tool_call
			if status != "tool_call" {
				t.Errorf("expected status 'tool_call', got %s", status)
			}
		}
	}

	// 5. Verify all results were recorded
	results := mockServer.ToolResults()
	if len(results) != len(scenario.ToolCalls) {
		t.Errorf("expected %d tool results, got %d", len(scenario.ToolCalls), len(results))
	}

	t.Log("Full protocol flow completed successfully")
}

func sendJSONRPC(t *testing.T, client *http.Client, url, method string, params interface{}) map[string]interface{} {
	t.Helper()

	req := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  method,
	}
	if params != nil {
		req["params"] = params
	}

	body, _ := json.Marshal(req)
	httpReq, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(httpReq)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	return result
}

func assertNoRPCError(t *testing.T, resp map[string]interface{}) {
	t.Helper()
	if errObj, ok := resp["error"]; ok && errObj != nil {
		t.Fatalf("unexpected JSON-RPC error: %v", errObj)
	}
}
