// Package mock provides a mock MCP server for testing cone's MCP client
// before the C1 MCP server is ready.
//
// The mock server simulates the C1 connector analysis workflow:
// 1. Accept connector_analyze call
// 2. Return tool callbacks (read_files, ask_user, edit_file)
// 3. Process tool results
// 4. Complete the session
package mock

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

// Scenario defines a test scenario with canned tool calls and expected responses.
type Scenario struct {
	Name        string
	Description string
	ToolCalls   []ToolCall
}

// ToolCall represents a tool call the mock server will make to the client.
type ToolCall struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
	// ExpectedResult is used for validation in tests
	ExpectedResult map[string]interface{} `json:"-"`
}

// Server is a mock MCP server for testing.
type Server struct {
	scenario    *Scenario
	currentStep int
	mu          sync.Mutex
	addr        string
	server      *http.Server
	toolResults []map[string]interface{}
	sessionID   string
	initialized bool
}

// NewServer creates a new mock MCP server with the given scenario.
func NewServer(scenario *Scenario) *Server {
	return &Server{
		scenario:    scenario,
		toolResults: make([]map[string]interface{}, 0),
	}
}

// Start starts the mock server on the given address.
func (s *Server) Start(addr string) error {
	s.addr = addr
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HandleMCP)

	s.server = &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("mock server error: %v\n", err)
		}
	}()

	return nil
}

// Stop stops the mock server.
func (s *Server) Stop(ctx context.Context) error {
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

// Addr returns the server address.
func (s *Server) Addr() string {
	return s.addr
}

// ToolResults returns the results received from tool calls.
func (s *Server) ToolResults() []map[string]interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.toolResults
}

// HandleMCP handles MCP protocol messages. Exported for use in tests.
func (s *Server) HandleMCP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		JSONRPC string                 `json:"jsonrpc"`
		ID      interface{}            `json:"id"`
		Method  string                 `json:"method"`
		Params  map[string]interface{} `json:"params"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.sendError(w, nil, -32700, "parse error")
		return
	}

	switch req.Method {
	case "initialize":
		s.handleInitialize(w, req.ID, req.Params)
	case "tools/list":
		s.handleToolsList(w, req.ID)
	case "tools/call":
		s.handleToolsCall(w, req.ID, req.Params)
	case "tool_result":
		s.handleToolResult(w, req.ID, req.Params)
	default:
		s.sendError(w, req.ID, -32601, fmt.Sprintf("method not found: %s", req.Method))
	}
}

func (s *Server) handleInitialize(w http.ResponseWriter, id interface{}, params map[string]interface{}) {
	s.mu.Lock()
	s.initialized = true
	s.sessionID = fmt.Sprintf("mock-session-%d", s.currentStep)
	s.mu.Unlock()

	s.sendResult(w, id, map[string]interface{}{
		"protocolVersion": "2024-11-05",
		"serverInfo": map[string]string{
			"name":    "c1-mock-mcp",
			"version": "0.1.0-test",
		},
		"capabilities": map[string]interface{}{
			"tools": map[string]bool{"supported": true},
		},
	})
}

func (s *Server) handleToolsList(w http.ResponseWriter, id interface{}) {
	tools := []map[string]interface{}{
		{
			"name":        "connector_analyze",
			"description": "Analyze a connector for issues and improvements",
			"inputSchema": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"connector_path": map[string]string{"type": "string"},
					"mode":           map[string]string{"type": "string"},
				},
				"required": []string{"connector_path"},
			},
		},
	}

	s.sendResult(w, id, map[string]interface{}{
		"tools": tools,
	})
}

func (s *Server) handleToolsCall(w http.ResponseWriter, id interface{}, params map[string]interface{}) {
	name, _ := params["name"].(string)

	if name != "connector_analyze" {
		s.sendError(w, id, -32602, fmt.Sprintf("unknown tool: %s", name))
		return
	}

	s.mu.Lock()
	s.currentStep = 0
	s.mu.Unlock()

	// Return the first tool callback
	s.sendNextToolCallback(w, id)
}

func (s *Server) handleToolResult(w http.ResponseWriter, id interface{}, params map[string]interface{}) {
	s.mu.Lock()
	s.toolResults = append(s.toolResults, params)
	s.currentStep++
	step := s.currentStep
	s.mu.Unlock()

	// Check if we have more tool calls
	if step < len(s.scenario.ToolCalls) {
		s.sendNextToolCallback(w, id)
		return
	}

	// Analysis complete
	s.sendResult(w, id, map[string]interface{}{
		"status":  "complete",
		"message": "Analysis finished",
		"summary": map[string]interface{}{
			"issues_found":   len(s.scenario.ToolCalls),
			"files_analyzed": len(s.toolResults),
		},
	})
}

func (s *Server) sendNextToolCallback(w http.ResponseWriter, id interface{}) {
	s.mu.Lock()
	if s.currentStep >= len(s.scenario.ToolCalls) {
		s.mu.Unlock()
		s.sendResult(w, id, map[string]interface{}{
			"status":  "complete",
			"message": "No more tool calls",
		})
		return
	}
	toolCall := s.scenario.ToolCalls[s.currentStep]
	s.mu.Unlock()

	s.sendResult(w, id, map[string]interface{}{
		"status":    "tool_call",
		"tool_call": toolCall,
	})
}

func (s *Server) sendResult(w http.ResponseWriter, id interface{}, result interface{}) {
	resp := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      id,
		"result":  result,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Fprintf(os.Stderr, "mock server: failed to encode result response (id=%v): %v\n", id, err)
	}
}

func (s *Server) sendError(w http.ResponseWriter, id interface{}, code int, message string) {
	resp := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      id,
		"error": map[string]interface{}{
			"code":    code,
			"message": message,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Fprintf(os.Stderr, "mock server: failed to encode error response (id=%v, code=%d): %v\n", id, code, err)
	}
}

// Predefined test scenarios

// HappyPathScenario returns a scenario that reads files and suggests an edit.
func HappyPathScenario() *Scenario {
	return &Scenario{
		Name:        "happy_path",
		Description: "Read connector files, suggest an edit",
		ToolCalls: []ToolCall{
			{
				Name: "read_files",
				Arguments: map[string]interface{}{
					"patterns": []string{"*.go", "*.yaml"},
				},
			},
			{
				Name: "ask_user",
				Arguments: map[string]interface{}{
					"question": "Found a missing error check. Should I fix it?",
					"type":     "confirm",
				},
			},
			{
				Name: "edit_file",
				Arguments: map[string]interface{}{
					"path": "connector.go",
					"old":  "result, _ := client.Get()",
					"new":  "result, err := client.Get()\nif err != nil {\n    return nil, err\n}",
				},
			},
		},
	}
}

// UserDeclinesScenario returns a scenario where the user declines an edit.
func UserDeclinesScenario() *Scenario {
	return &Scenario{
		Name:        "user_declines",
		Description: "User declines a suggested edit",
		ToolCalls: []ToolCall{
			{
				Name: "read_files",
				Arguments: map[string]interface{}{
					"patterns": []string{"*.go"},
				},
			},
			{
				Name: "ask_user",
				Arguments: map[string]interface{}{
					"question": "Should I refactor this function?",
					"type":     "confirm",
				},
			},
		},
	}
}

// InvalidToolCallScenario returns a scenario with an invalid tool call.
func InvalidToolCallScenario() *Scenario {
	return &Scenario{
		Name:        "invalid_tool",
		Description: "Server sends an unknown tool call",
		ToolCalls: []ToolCall{
			{
				Name: "nonexistent_tool",
				Arguments: map[string]interface{}{
					"foo": "bar",
				},
			},
		},
	}
}
