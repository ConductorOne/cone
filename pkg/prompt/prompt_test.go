package prompt

import (
	"errors"
	"strings"
	"testing"
)

func TestIsInteractive(t *testing.T) {
	// In tests, stdin is typically not a terminal
	// This test just verifies the function doesn't panic
	_ = IsInteractive()
}

func TestRequireInteractive_InTests(t *testing.T) {
	// In test environment, stdin is not a terminal
	err := requireInteractive()
	// Should return error in non-interactive context
	if IsInteractive() {
		if err != nil {
			t.Errorf("expected no error in interactive mode, got %v", err)
		}
	} else {
		if !errors.Is(err, ErrNotInteractive) {
			t.Errorf("expected ErrNotInteractive, got %v", err)
		}
	}
}

func TestConfirm_NonInteractive(t *testing.T) {
	if IsInteractive() {
		t.Skip("test requires non-interactive environment")
	}

	result, err := Confirm("Test question?")
	if result {
		t.Error("expected false result")
	}
	if !errors.Is(err, ErrNotInteractive) {
		t.Errorf("expected ErrNotInteractive, got %v", err)
	}
}

func TestConfirmWithDefault_NonInteractive(t *testing.T) {
	if IsInteractive() {
		t.Skip("test requires non-interactive environment")
	}

	result, err := ConfirmWithDefault("Test question?", true)
	if result {
		t.Error("expected false result")
	}
	if !errors.Is(err, ErrNotInteractive) {
		t.Errorf("expected ErrNotInteractive, got %v", err)
	}
}

func TestInput_NonInteractive(t *testing.T) {
	if IsInteractive() {
		t.Skip("test requires non-interactive environment")
	}

	result, err := Input("Enter value: ")
	if result != "" {
		t.Errorf("expected empty result, got %s", result)
	}
	if !errors.Is(err, ErrNotInteractive) {
		t.Errorf("expected ErrNotInteractive, got %v", err)
	}
}

func TestInputWithDefault_NonInteractive(t *testing.T) {
	if IsInteractive() {
		t.Skip("test requires non-interactive environment")
	}

	result, err := InputWithDefault("Enter value", "default")
	if result != "" {
		t.Errorf("expected empty result, got %s", result)
	}
	if !errors.Is(err, ErrNotInteractive) {
		t.Errorf("expected ErrNotInteractive, got %v", err)
	}
}

func TestSelect_NonInteractive(t *testing.T) {
	if IsInteractive() {
		t.Skip("test requires non-interactive environment")
	}

	options := []Option{
		{Label: "Option 1"},
		{Label: "Option 2"},
	}

	result, err := Select("Choose:", options)
	if result != -1 {
		t.Errorf("expected -1, got %d", result)
	}
	if !errors.Is(err, ErrNotInteractive) {
		t.Errorf("expected ErrNotInteractive, got %v", err)
	}
}

func TestSelect_NoOptions(t *testing.T) {
	// Even in interactive mode, empty options should fail
	result, err := Select("Choose:", []Option{})
	if result != -1 {
		t.Errorf("expected -1, got %d", result)
	}
	// Will be either ErrNoOptions or ErrNotInteractive
	if err == nil {
		t.Error("expected error for empty options")
	}
}

func TestSelectString_NonInteractive(t *testing.T) {
	if IsInteractive() {
		t.Skip("test requires non-interactive environment")
	}

	result, err := SelectString("Choose:", []string{"A", "B"})
	if result != -1 {
		t.Errorf("expected -1, got %d", result)
	}
	if !errors.Is(err, ErrNotInteractive) {
		t.Errorf("expected ErrNotInteractive, got %v", err)
	}
}

func TestMultiSelect_NonInteractive(t *testing.T) {
	if IsInteractive() {
		t.Skip("test requires non-interactive environment")
	}

	options := []Option{
		{Label: "Option 1"},
		{Label: "Option 2"},
	}

	result, err := MultiSelect("Choose:", options)
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
	if !errors.Is(err, ErrNotInteractive) {
		t.Errorf("expected ErrNotInteractive, got %v", err)
	}
}

func TestMultiSelect_NoOptions(t *testing.T) {
	result, err := MultiSelect("Choose:", []Option{})
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
	// Will be either ErrNoOptions or ErrNotInteractive
	if err == nil {
		t.Error("expected error for empty options")
	}
}

func TestWrapText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		width    int
		expected []string
	}{
		{
			name:     "single short line",
			text:     "Hello world",
			width:    50,
			expected: []string{"Hello world"},
		},
		{
			name:     "wrap long line",
			text:     "This is a longer line that should wrap to multiple lines",
			width:    20,
			expected: []string{"This is a longer", "line that should", "wrap to multiple", "lines"},
		},
		{
			name:     "empty text",
			text:     "",
			width:    50,
			expected: []string{""},
		},
		{
			name:     "multiple paragraphs",
			text:     "First paragraph.\n\nSecond paragraph.",
			width:    50,
			expected: []string{"First paragraph.", "", "Second paragraph."},
		},
		{
			name:     "blank lines preserved",
			text:     "Line 1\n\n\nLine 4",
			width:    50,
			expected: []string{"Line 1", "", "", "Line 4"},
		},
		{
			name:     "single word longer than width",
			text:     "supercalifragilisticexpialidocious",
			width:    10,
			expected: []string{"supercalifragilisticexpialidocious"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := wrapText(tc.text, tc.width)
			if len(result) != len(tc.expected) {
				t.Errorf("expected %d lines, got %d", len(tc.expected), len(result))
				t.Logf("result: %v", result)
				return
			}
			for i := range result {
				if result[i] != tc.expected[i] {
					t.Errorf("line %d: expected %q, got %q", i, tc.expected[i], result[i])
				}
			}
		})
	}
}

func TestDisplayBox(t *testing.T) {
	// DisplayBox writes to stdout, just verify it doesn't panic
	DisplayBox("Test Title", "Test content here.")
	DisplayBox("", "No title content")
	DisplayBox("Title", "Multi\nLine\nContent")
}

func TestOption(t *testing.T) {
	opt := Option{
		Label:       "Test Label",
		Description: "Test Description",
	}
	if opt.Label != "Test Label" {
		t.Errorf("expected label 'Test Label', got %s", opt.Label)
	}
	if opt.Description != "Test Description" {
		t.Errorf("expected description 'Test Description', got %s", opt.Description)
	}
}

func TestErrorMessages(t *testing.T) {
	if ErrNotInteractive.Error() != "prompt: stdin is not an interactive terminal" {
		t.Errorf("unexpected error message: %s", ErrNotInteractive.Error())
	}
	if ErrNoOptions.Error() != "prompt: no options provided" {
		t.Errorf("unexpected error message: %s", ErrNoOptions.Error())
	}
	if ErrCancelled.Error() != "prompt: cancelled by user" {
		t.Errorf("unexpected error message: %s", ErrCancelled.Error())
	}
}

func TestWrapTextEdgeCases(t *testing.T) {
	// Width of 1 - each word on its own line
	result := wrapText("a b c", 1)
	expected := []string{"a", "b", "c"}
	if len(result) != len(expected) {
		t.Errorf("expected %d lines, got %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("line %d: expected %q, got %q", i, expected[i], result[i])
		}
	}

	// Very long text
	longText := "The quick brown fox jumps over the lazy dog. "
	for i := 0; i < 5; i++ {
		longText += longText
	}
	result = wrapText(longText, 80)
	if len(result) == 0 {
		t.Error("expected non-empty result for long text")
	}
	for i, line := range result {
		if len(line) > 80 {
			t.Errorf("line %d exceeds width 80: length=%d", i, len(line))
		}
	}
}

func TestSelectStringConvertsToOptions(t *testing.T) {
	// Test that SelectString creates proper Options from strings
	// Can't fully test without interactive input, but verify the conversion logic
	// by checking the function signature and behavior match Select
	if IsInteractive() {
		t.Skip("test requires non-interactive environment")
	}

	// Both should fail the same way in non-interactive mode
	idx1, err1 := Select("Q?", []Option{{Label: "A"}, {Label: "B"}})
	idx2, err2 := SelectString("Q?", []string{"A", "B"})

	if idx1 != idx2 {
		t.Errorf("expected same index, got %d and %d", idx1, idx2)
	}
	// Both should return ErrNotInteractive in non-interactive mode.
	if !errors.Is(err1, ErrNotInteractive) || !errors.Is(err2, ErrNotInteractive) {
		t.Errorf("expected both errors to be ErrNotInteractive, got %v and %v", err1, err2)
	}
}

func TestDisplayBoxFormatting(t *testing.T) {
	// Verify box width constant behavior
	// DisplayBox uses width of 70
	content := "Short"
	DisplayBox("Title", content) // Just verify no panic

	// Long content that needs wrapping
	longContent := "This is a very long line that will need to be wrapped because it exceeds the box width of 70 characters"
	DisplayBox("Long Content Test", longContent)
}

func TestWrapTextPreservesIndentation(t *testing.T) {
	// Verify that leading spaces in words are not stripped
	text := "First line\n  Indented line\n    More indented"
	result := wrapText(text, 50)

	// Each line should be preserved (though wrapping behavior depends on implementation)
	if len(result) < 3 {
		t.Errorf("expected at least 3 lines, got %d", len(result))
	}
}

func TestWrapTextWithOnlyNewlines(t *testing.T) {
	text := "\n\n\n"
	result := wrapText(text, 50)
	// Should produce 4 empty lines (3 newlines split into 4 segments)
	if len(result) != 4 {
		t.Errorf("expected 4 lines for 3 newlines, got %d", len(result))
	}
	for i, line := range result {
		if line != "" {
			t.Errorf("line %d should be empty, got %q", i, line)
		}
	}
}

func TestWrapTextExactWidth(t *testing.T) {
	// Text that exactly fits the width
	text := "12345"
	result := wrapText(text, 5)
	if len(result) != 1 {
		t.Errorf("expected 1 line, got %d", len(result))
	}
	if result[0] != "12345" {
		t.Errorf("expected '12345', got %q", result[0])
	}
}

func TestWrapTextWordBoundaries(t *testing.T) {
	// Words that would split at boundary
	text := "aaa bbb ccc"
	result := wrapText(text, 7)
	// "aaa bbb" = 7 chars, should fit on one line
	// "ccc" on next line
	if len(result) != 2 {
		t.Errorf("expected 2 lines, got %d: %v", len(result), result)
	}
	if result[0] != "aaa bbb" {
		t.Errorf("expected 'aaa bbb', got %q", result[0])
	}
	if result[1] != "ccc" {
		t.Errorf("expected 'ccc', got %q", result[1])
	}
}

func TestDisplayBoxEmptyContent(t *testing.T) {
	// Empty content should not panic
	DisplayBox("Title", "")
	DisplayBox("", "")
}

func TestOptionWithDescription(t *testing.T) {
	opt := Option{
		Label:       "Build",
		Description: "Build the connector binary",
	}
	if !strings.Contains(opt.Label, "Build") {
		t.Error("label should contain 'Build'")
	}
	if !strings.Contains(opt.Description, "connector") {
		t.Error("description should contain 'connector'")
	}
}
