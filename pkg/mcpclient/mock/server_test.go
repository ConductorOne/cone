package mock

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMockServer_HappyPath(t *testing.T) {
	scenario := HappyPathScenario()
	server := NewServer(scenario)

	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(server.HandleMCP))
	defer ts.Close()

	// Test initialize
	resp := doRequest(t, ts.URL, "initialize", map[string]interface{}{
		"protocolVersion": "2024-11-05",
		"clientInfo":      map[string]string{"name": "test", "version": "1.0"},
	})
	assertNoError(t, resp)

	// Test tools/list
	resp = doRequest(t, ts.URL, "tools/list", nil)
	assertNoError(t, resp)

	var listResult struct {
		Tools []map[string]interface{} `json:"tools"`
	}
	listBytes, _ := json.Marshal(resp["result"])
	if err := json.Unmarshal(listBytes, &listResult); err != nil {
		t.Fatalf("failed to parse tools/list result: %v", err)
	}
	if len(listResult.Tools) == 0 {
		t.Error("expected at least one tool")
	}

	// Test tools/call for connector_analyze
	resp = doRequest(t, ts.URL, "tools/call", map[string]interface{}{
		"name": "connector_analyze",
		"arguments": map[string]interface{}{
			"connector_path": "/test/connector",
		},
	})
	assertNoError(t, resp)

	// Should get first tool callback (read_files)
	var callResult struct {
		Status   string `json:"status"`
		ToolCall struct {
			Name      string                 `json:"name"`
			Arguments map[string]interface{} `json:"arguments"`
		} `json:"tool_call"`
	}
	resultBytes, _ := json.Marshal(resp["result"])
	if err := json.Unmarshal(resultBytes, &callResult); err != nil {
		t.Fatalf("failed to parse tool call result: %v", err)
	}

	if callResult.Status != "tool_call" {
		t.Errorf("expected status 'tool_call', got '%s'", callResult.Status)
	}
	if callResult.ToolCall.Name != "read_files" {
		t.Errorf("expected first tool call to be 'read_files', got '%s'", callResult.ToolCall.Name)
	}
}

func TestMockServer_SessionTracking(t *testing.T) {
	scenario := HappyPathScenario()
	server := NewServer(scenario)

	ts := httptest.NewServer(http.HandlerFunc(server.HandleMCP))
	defer ts.Close()

	// Initialize
	doRequest(t, ts.URL, "initialize", map[string]interface{}{
		"protocolVersion": "2024-11-05",
		"clientInfo":      map[string]string{"name": "test", "version": "1.0"},
	})

	// Start analysis
	doRequest(t, ts.URL, "tools/call", map[string]interface{}{
		"name": "connector_analyze",
		"arguments": map[string]interface{}{
			"connector_path": "/test",
		},
	})

	// Send tool results
	for i := 0; i < len(scenario.ToolCalls); i++ {
		doRequest(t, ts.URL, "tool_result", map[string]interface{}{
			"tool": scenario.ToolCalls[i].Name,
			"result": map[string]interface{}{
				"success": true,
			},
		})
	}

	// Verify all results were recorded
	results := server.ToolResults()
	if len(results) != len(scenario.ToolCalls) {
		t.Errorf("expected %d tool results, got %d", len(scenario.ToolCalls), len(results))
	}
}

func doRequest(t *testing.T, url, method string, params interface{}) map[string]interface{} {
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
	resp, err := http.DefaultClient.Do(httpReq)
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

func assertNoError(t *testing.T, resp map[string]interface{}) {
	t.Helper()
	if errObj, ok := resp["error"]; ok && errObj != nil {
		t.Fatalf("unexpected error: %v", errObj)
	}
}
