package scaffold

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	// Create temp directory
	tmpDir, err := os.MkdirTemp("", "scaffold-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	outputDir := filepath.Join(tmpDir, "baton-test-app")

	cfg := &Config{
		Name:        "test-app",
		ModulePath:  "github.com/example/baton-test-app",
		OutputDir:   outputDir,
		Description: "Test connector for testing",
	}

	if err := Generate(cfg); err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	// Verify expected files exist
	expectedFiles := []string{
		"go.mod",
		"main.go",
		"cmd/baton-test-app/config/config.go",
		"pkg/connector/connector.go",
		"pkg/connector/users.go",
		"pkg/connector/groups.go",
		"pkg/connector/roles.go",
		"pkg/client/client.go",
		".gitignore",
		"README.md",
		"Makefile",
	}

	for _, f := range expectedFiles {
		path := filepath.Join(outputDir, f)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("expected file %s to exist", f)
		}
	}

	// Verify go.mod contains module path
	goMod, err := os.ReadFile(filepath.Join(outputDir, "go.mod"))
	if err != nil {
		t.Fatalf("failed to read go.mod: %v", err)
	}
	if !strings.Contains(string(goMod), "github.com/example/baton-test-app") {
		t.Error("go.mod should contain module path")
	}

	// Verify main.go contains connector name
	mainGo, err := os.ReadFile(filepath.Join(outputDir, "main.go"))
	if err != nil {
		t.Fatalf("failed to read main.go: %v", err)
	}
	if !strings.Contains(string(mainGo), "baton-test-app") {
		t.Error("main.go should contain connector name")
	}
}

func TestGenerateDefaults(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "scaffold-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Change to temp dir so default output dir works
	oldWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(oldWd)

	cfg := &Config{
		Name: "my-service",
	}

	if err := Generate(cfg); err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	// Verify defaults were applied
	expectedDir := filepath.Join(tmpDir, "baton-my-service")
	if _, err := os.Stat(expectedDir); os.IsNotExist(err) {
		t.Error("expected default output directory baton-my-service")
	}

	// Verify default module path
	goMod, err := os.ReadFile(filepath.Join(expectedDir, "go.mod"))
	if err != nil {
		t.Fatalf("failed to read go.mod: %v", err)
	}
	if !strings.Contains(string(goMod), "github.com/conductorone/baton-my-service") {
		t.Error("go.mod should contain default module path")
	}
}

func TestGenerateMissingName(t *testing.T) {
	cfg := &Config{}

	err := Generate(cfg)
	if err == nil {
		t.Error("expected error for missing name")
	}
}

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"my-app", "MyApp"},
		{"test", "Test"},
		{"foo-bar-baz", "FooBarBaz"},
		{"", ""},
	}

	for _, tc := range tests {
		result := toPascalCase(tc.input)
		if result != tc.expected {
			t.Errorf("toPascalCase(%q) = %q, expected %q", tc.input, result, tc.expected)
		}
	}
}
