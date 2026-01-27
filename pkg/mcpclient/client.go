package mcpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Status values for MCP analysis responses.
const (
	statusToolCall = "tool_call"
)

// Client is an MCP client for connecting to C1's connector analysis service.
type Client struct {
	// ServerURL is the URL of the MCP server (e.g., https://tenant.conductorone.com/api/v1alpha/cone/mcp)
	ServerURL string

	// AuthToken is the authentication token from cone login.
	AuthToken string

	// HTTPClient is the HTTP client to use. If nil, a default client is used.
	HTTPClient *http.Client

	// Timeout is the request timeout.
	Timeout time.Duration

	// ToolHandler handles tool callbacks from the server.
	ToolHandler *ToolHandler

	// initialized tracks whether we've completed the MCP handshake.
	initialized bool

	// requestID is a counter for JSON-RPC request IDs.
	requestID int
}

// NewClient creates a new MCP client.
func NewClient(serverURL, authToken string, toolHandler *ToolHandler) *Client {
	return &Client{
		ServerURL:   serverURL,
		AuthToken:   authToken,
		HTTPClient:  &http.Client{Timeout: 30 * time.Second},
		Timeout:     30 * time.Second,
		ToolHandler: toolHandler,
	}
}

// jsonrpcRequest represents a JSON-RPC request.
type jsonrpcRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

// jsonrpcResponse represents a JSON-RPC response.
type jsonrpcResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *jsonrpcError   `json:"error,omitempty"`
}

type jsonrpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Connect establishes a connection to the MCP server.
func (c *Client) Connect(ctx context.Context) error {
	// Send initialize request
	resp, err := c.sendRequest(ctx, "initialize", map[string]interface{}{
		"protocolVersion": "2024-11-05",
		"clientInfo": map[string]string{
			"name":    "cone",
			"version": "0.1.0",
		},
		"capabilities": map[string]interface{}{
			"tools": map[string]bool{"supported": true},
		},
	})
	if err != nil {
		return fmt.Errorf("initialize failed: %w", err)
	}

	// Parse initialize result
	var initResult struct {
		ProtocolVersion string `json:"protocolVersion"`
		ServerInfo      struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"serverInfo"`
	}
	if err := json.Unmarshal(resp.Result, &initResult); err != nil {
		return fmt.Errorf("failed to parse initialize result: %w", err)
	}

	c.initialized = true
	return nil
}

// Analyze starts a connector analysis session.
// It handles the full interaction loop: calling connector_analyze,
// processing tool callbacks, and returning when complete.
func (c *Client) Analyze(ctx context.Context, connectorPath string, mode string) (*AnalysisResult, error) {
	if !c.initialized {
		if err := c.Connect(ctx); err != nil {
			return nil, err
		}
	}

	if mode == "" {
		mode = "interactive"
	}

	// Call connector_analyze tool
	resp, err := c.sendRequest(ctx, "tools/call", map[string]interface{}{
		"name": "connector_analyze",
		"arguments": map[string]interface{}{
			"connector_path": connectorPath,
			"mode":           mode,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("connector_analyze failed: %w", err)
	}

	// Process the response and any tool callbacks
	return c.processAnalysisResponse(ctx, resp)
}

// AnalysisResult contains the results of a connector analysis.
type AnalysisResult struct {
	Status       string                 `json:"status"`
	Message      string                 `json:"message"`
	IssuesFound  int                    `json:"issues_found"`
	FilesScanned int                    `json:"files_scanned"`
	Summary      map[string]interface{} `json:"summary,omitempty"`
}

// processAnalysisResponse handles the analysis response, including tool callback loops.
func (c *Client) processAnalysisResponse(ctx context.Context, resp *jsonrpcResponse) (*AnalysisResult, error) {
	for {
		var result struct {
			Status   string `json:"status"`
			Message  string `json:"message"`
			ToolCall *struct {
				Name      string                 `json:"name"`
				Arguments map[string]interface{} `json:"arguments"`
			} `json:"tool_call,omitempty"`
			Summary map[string]interface{} `json:"summary,omitempty"`
		}

		if err := json.Unmarshal(resp.Result, &result); err != nil {
			return nil, fmt.Errorf("failed to parse analysis response: %w", err)
		}

		switch result.Status {
		case "complete":
			// Analysis complete
			issuesFound := 0
			filesScanned := 0
			if result.Summary != nil {
				if v, ok := result.Summary["issues_found"].(float64); ok {
					issuesFound = int(v)
				}
				if v, ok := result.Summary["files_analyzed"].(float64); ok {
					filesScanned = int(v)
				}
			}
			return &AnalysisResult{
				Status:       "complete",
				Message:      result.Message,
				IssuesFound:  issuesFound,
				FilesScanned: filesScanned,
				Summary:      result.Summary,
			}, nil

		case statusToolCall:
			if result.ToolCall == nil {
				return nil, fmt.Errorf("%s status but no tool_call data", statusToolCall)
			}

			// Ensure we have a tool handler configured
			if c.ToolHandler == nil {
				return nil, fmt.Errorf("no ToolHandler configured for tool_call %s", result.ToolCall.Name)
			}

			// Execute the tool locally
			toolResult, err := c.ToolHandler.HandleToolCall(ctx, result.ToolCall.Name, result.ToolCall.Arguments)
			if err != nil {
				return nil, fmt.Errorf("tool %s failed: %w", result.ToolCall.Name, err)
			}

			// Send the result back to the server
			resp, err = c.sendRequest(ctx, "tool_result", map[string]interface{}{
				"tool":   result.ToolCall.Name,
				"result": toolResult,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to send tool result: %w", err)
			}

			// Continue the loop with the new response
			continue

		case "error":
			return &AnalysisResult{
				Status:  "error",
				Message: result.Message,
			}, nil

		default:
			return nil, fmt.Errorf("unexpected status: %s", result.Status)
		}
	}
}

// sendRequest sends a JSON-RPC request to the server.
func (c *Client) sendRequest(ctx context.Context, method string, params interface{}) (*jsonrpcResponse, error) {
	c.requestID++
	req := jsonrpcRequest{
		JSONRPC: "2.0",
		ID:      c.requestID,
		Method:  method,
		Params:  params,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.ServerURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.AuthToken != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.AuthToken)
	}

	httpResp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(httpResp.Body)
		return nil, fmt.Errorf("server returned %d: %s", httpResp.StatusCode, string(bodyBytes))
	}

	var resp jsonrpcResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("server error %d: %s", resp.Error.Code, resp.Error.Message)
	}

	return &resp, nil
}

// Close closes the client connection.
func (c *Client) Close() error {
	// HTTP client doesn't need explicit cleanup
	c.initialized = false
	return nil
}
