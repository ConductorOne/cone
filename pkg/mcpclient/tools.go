// Package mcpclient provides an MCP client for connecting to C1's connector analysis service.
package mcpclient

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/conductorone/cone/pkg/prompt"
)

// ToolHandler handles tool callbacks from the MCP server.
type ToolHandler struct {
	// ConnectorDir is the root directory of the connector being analyzed.
	// All file operations are restricted to this directory.
	ConnectorDir string

	// DryRun prevents actual file modifications when true.
	DryRun bool

	// Verbose enables detailed output.
	Verbose bool
}

// NewToolHandler creates a new tool handler for the given connector directory.
func NewToolHandler(connectorDir string) *ToolHandler {
	return &ToolHandler{
		ConnectorDir: connectorDir,
	}
}

// ToolResult is the result of executing a tool.
type ToolResult struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Error   string                 `json:"error,omitempty"`
}

// HandleToolCall dispatches a tool call to the appropriate handler.
func (h *ToolHandler) HandleToolCall(ctx context.Context, name string, args map[string]interface{}) (*ToolResult, error) {
	switch name {
	case "read_files":
		return h.handleReadFiles(ctx, args)
	case "ask_user":
		return h.handleAskUser(ctx, args)
	case "edit_file":
		return h.handleEditFile(ctx, args)
	case "write_file":
		return h.handleWriteFile(ctx, args)
	case "show_diff":
		return h.handleShowDiff(ctx, args)
	case "confirm":
		return h.handleConfirm(ctx, args)
	default:
		return &ToolResult{
			Success: false,
			Error:   fmt.Sprintf("unknown tool: %s", name),
		}, nil
	}
}

// handleReadFiles reads files matching the given glob patterns.
// Security: Only reads files within ConnectorDir.
func (h *ToolHandler) handleReadFiles(ctx context.Context, args map[string]interface{}) (*ToolResult, error) {
	patternsRaw, ok := args["patterns"]
	if !ok {
		return &ToolResult{Success: false, Error: "missing 'patterns' argument"}, nil
	}

	var patterns []string
	switch p := patternsRaw.(type) {
	case []string:
		patterns = p
	case []interface{}:
		for _, v := range p {
			if s, ok := v.(string); ok {
				patterns = append(patterns, s)
			}
		}
	default:
		return &ToolResult{Success: false, Error: "patterns must be an array of strings"}, nil
	}

	files := make([]map[string]string, 0)

	for _, pattern := range patterns {
		// Security: Resolve pattern relative to connector directory
		fullPattern := filepath.Join(h.ConnectorDir, pattern)

		matches, err := filepath.Glob(fullPattern)
		if err != nil {
			continue // Skip invalid patterns
		}

		for _, match := range matches {
			// Security: Verify file is within connector directory
			relPath, err := filepath.Rel(h.ConnectorDir, match)
			if err != nil || strings.HasPrefix(relPath, "..") {
				continue // Skip files outside connector dir
			}

			info, err := os.Stat(match)
			if err != nil || info.IsDir() {
				continue // Skip directories and inaccessible files
			}

			content, err := os.ReadFile(match)
			if err != nil {
				continue // Skip unreadable files
			}

			files = append(files, map[string]string{
				"path":    relPath,
				"content": string(content),
			})
		}
	}

	return &ToolResult{
		Success: true,
		Data: map[string]interface{}{
			"files": files,
		},
	}, nil
}

// handleAskUser prompts the user for input.
func (h *ToolHandler) handleAskUser(ctx context.Context, args map[string]interface{}) (*ToolResult, error) {
	question, _ := args["question"].(string)
	questionType, _ := args["type"].(string)

	if question == "" {
		return &ToolResult{Success: false, Error: "missing 'question' argument"}, nil
	}

	switch questionType {
	case "confirm":
		answer, err := prompt.Confirm(question)
		if err != nil {
			return &ToolResult{Success: false, Error: err.Error()}, nil
		}
		return &ToolResult{
			Success: true,
			Data:    map[string]interface{}{"answer": answer},
		}, nil

	case "text":
		answer, err := prompt.Input(question + ": ")
		if err != nil {
			return &ToolResult{Success: false, Error: err.Error()}, nil
		}
		return &ToolResult{
			Success: true,
			Data:    map[string]interface{}{"answer": answer},
		}, nil

	case "select":
		optionsRaw, _ := args["options"].([]interface{})
		var options []string
		for _, o := range optionsRaw {
			if s, ok := o.(string); ok {
				options = append(options, s)
			}
		}
		if len(options) == 0 {
			return &ToolResult{Success: false, Error: "select requires options"}, nil
		}

		idx, err := prompt.SelectString(question, options)
		if err != nil {
			return &ToolResult{Success: false, Error: err.Error()}, nil
		}
		return &ToolResult{
			Success: true,
			Data: map[string]interface{}{
				"answer": options[idx],
				"index":  idx,
			},
		}, nil

	default:
		// Default to text input
		answer, err := prompt.Input(question + ": ")
		if err != nil {
			return &ToolResult{Success: false, Error: err.Error()}, nil
		}
		return &ToolResult{
			Success: true,
			Data:    map[string]interface{}{"answer": answer},
		}, nil
	}
}

// handleEditFile applies an edit to a file.
// Security: Only edits files within ConnectorDir.
func (h *ToolHandler) handleEditFile(ctx context.Context, args map[string]interface{}) (*ToolResult, error) {
	path, _ := args["path"].(string)
	old, _ := args["old"].(string)
	new, _ := args["new"].(string)

	if path == "" || old == "" {
		return &ToolResult{Success: false, Error: "missing required arguments (path, old)"}, nil
	}

	// Security: Resolve path relative to connector directory
	fullPath := filepath.Join(h.ConnectorDir, path)
	relPath, err := filepath.Rel(h.ConnectorDir, fullPath)
	if err != nil || strings.HasPrefix(relPath, "..") {
		return &ToolResult{Success: false, Error: "path must be within connector directory"}, nil
	}

	// Read current content
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return &ToolResult{Success: false, Error: fmt.Sprintf("failed to read file: %v", err)}, nil
	}

	// Check if old string exists
	if !strings.Contains(string(content), old) {
		return &ToolResult{Success: false, Error: "old string not found in file"}, nil
	}

	// Show diff and ask for confirmation
	fmt.Printf("\n--- %s (before)\n", path)
	fmt.Printf("+++ %s (after)\n", path)
	fmt.Printf("@@ edit @@\n")
	fmt.Printf("-%s\n", strings.ReplaceAll(old, "\n", "\n-"))
	fmt.Printf("+%s\n", strings.ReplaceAll(new, "\n", "\n+"))
	fmt.Println()

	if h.DryRun {
		return &ToolResult{
			Success: true,
			Data:    map[string]interface{}{"applied": false, "reason": "dry run"},
		}, nil
	}

	accepted, err := prompt.Confirm("Apply this change?")
	if err != nil {
		return &ToolResult{Success: false, Error: err.Error()}, nil
	}

	if !accepted {
		return &ToolResult{
			Success: true,
			Data:    map[string]interface{}{"applied": false, "reason": "user declined"},
		}, nil
	}

	// Apply the edit
	newContent := strings.Replace(string(content), old, new, 1)
	if err := os.WriteFile(fullPath, []byte(newContent), 0644); err != nil {
		return &ToolResult{Success: false, Error: fmt.Sprintf("failed to write file: %v", err)}, nil
	}

	return &ToolResult{
		Success: true,
		Data:    map[string]interface{}{"applied": true},
	}, nil
}

// handleWriteFile writes a new file.
// Security: Only writes files within ConnectorDir.
func (h *ToolHandler) handleWriteFile(ctx context.Context, args map[string]interface{}) (*ToolResult, error) {
	path, _ := args["path"].(string)
	content, _ := args["content"].(string)

	if path == "" {
		return &ToolResult{Success: false, Error: "missing 'path' argument"}, nil
	}

	// Security: Resolve path relative to connector directory
	fullPath := filepath.Join(h.ConnectorDir, path)
	relPath, err := filepath.Rel(h.ConnectorDir, fullPath)
	if err != nil || strings.HasPrefix(relPath, "..") {
		return &ToolResult{Success: false, Error: "path must be within connector directory"}, nil
	}

	// Check if file exists
	if _, err := os.Stat(fullPath); err == nil {
		// File exists, ask for confirmation
		overwrite, err := prompt.Confirm(fmt.Sprintf("File %s exists. Overwrite?", path))
		if err != nil {
			return &ToolResult{Success: false, Error: err.Error()}, nil
		}
		if !overwrite {
			return &ToolResult{
				Success: true,
				Data:    map[string]interface{}{"written": false, "reason": "user declined overwrite"},
			}, nil
		}
	}

	if h.DryRun {
		return &ToolResult{
			Success: true,
			Data:    map[string]interface{}{"written": false, "reason": "dry run"},
		}, nil
	}

	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return &ToolResult{Success: false, Error: fmt.Sprintf("failed to create directory: %v", err)}, nil
	}

	// Write the file
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return &ToolResult{Success: false, Error: fmt.Sprintf("failed to write file: %v", err)}, nil
	}

	return &ToolResult{
		Success: true,
		Data:    map[string]interface{}{"written": true},
	}, nil
}

// handleShowDiff displays a diff for review.
func (h *ToolHandler) handleShowDiff(ctx context.Context, args map[string]interface{}) (*ToolResult, error) {
	path, _ := args["path"].(string)
	old, _ := args["old"].(string)
	new, _ := args["new"].(string)

	fmt.Printf("\n--- %s (before)\n", path)
	fmt.Printf("+++ %s (after)\n", path)
	fmt.Printf("@@ diff @@\n")

	// Simple line-by-line diff display
	oldLines := strings.Split(old, "\n")
	newLines := strings.Split(new, "\n")

	for _, line := range oldLines {
		fmt.Printf("-%s\n", line)
	}
	for _, line := range newLines {
		fmt.Printf("+%s\n", line)
	}
	fmt.Println()

	accepted, err := prompt.Confirm("Accept this change?")
	if err != nil {
		return &ToolResult{Success: false, Error: err.Error()}, nil
	}

	return &ToolResult{
		Success: true,
		Data:    map[string]interface{}{"accepted": accepted},
	}, nil
}

// handleConfirm asks for a simple yes/no confirmation.
func (h *ToolHandler) handleConfirm(ctx context.Context, args map[string]interface{}) (*ToolResult, error) {
	message, _ := args["message"].(string)
	if message == "" {
		message = "Continue?"
	}

	confirmed, err := prompt.Confirm(message)
	if err != nil {
		return &ToolResult{Success: false, Error: err.Error()}, nil
	}

	return &ToolResult{
		Success: true,
		Data:    map[string]interface{}{"confirmed": confirmed},
	}, nil
}
