// Package prompt provides simple interactive prompts for CLI applications.
// It uses basic fmt.Print + bufio.Scanner patterns (NOT Bubbletea).
// All prompts fail gracefully with an error when stdin is not a terminal.
package prompt

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"
)

// ErrNotInteractive is returned when prompts are called in non-interactive mode.
var ErrNotInteractive = errors.New("prompt: stdin is not an interactive terminal")

// ErrNoOptions is returned when Select is called with an empty options slice.
var ErrNoOptions = errors.New("prompt: no options provided")

// ErrCancelled is returned when the user cancels a prompt (e.g., Ctrl+C).
var ErrCancelled = errors.New("prompt: cancelled by user")

// IsInteractive returns true if stdin is an interactive terminal.
func IsInteractive() bool {
	return term.IsTerminal(int(os.Stdin.Fd()))
}

// requireInteractive returns an error if stdin is not interactive.
func requireInteractive() error {
	if !IsInteractive() {
		return ErrNotInteractive
	}
	return nil
}

// Confirm prompts the user with a yes/no question.
// Returns true for yes, false for no.
func Confirm(question string) (bool, error) {
	if err := requireInteractive(); err != nil {
		return false, err
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", question)
		input, err := reader.ReadString('\n')
		if err != nil {
			return false, fmt.Errorf("failed to read input: %w", err)
		}

		input = strings.TrimSpace(strings.ToLower(input))
		switch input {
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		default:
			fmt.Println("Please enter 'y' or 'n'.")
		}
	}
}

// ConfirmWithDefault prompts the user with a yes/no question with a default.
// Empty input returns the default value.
func ConfirmWithDefault(question string, defaultYes bool) (bool, error) {
	if err := requireInteractive(); err != nil {
		return false, err
	}

	reader := bufio.NewReader(os.Stdin)
	prompt := "[y/N]"
	if defaultYes {
		prompt = "[Y/n]"
	}

	for {
		fmt.Printf("%s %s: ", question, prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			return false, fmt.Errorf("failed to read input: %w", err)
		}

		input = strings.TrimSpace(strings.ToLower(input))
		switch input {
		case "":
			return defaultYes, nil
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		default:
			fmt.Println("Please enter 'y' or 'n'.")
		}
	}
}

// Input prompts the user for a single line of text input.
func Input(prompt string) (string, error) {
	if err := requireInteractive(); err != nil {
		return "", err
	}

	reader := bufio.NewReader(os.Stdin)
	_, _ = fmt.Fprint(os.Stdout, prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	return strings.TrimSpace(input), nil
}

// InputWithDefault prompts for text input with a default value.
func InputWithDefault(promptText, defaultValue string) (string, error) {
	if err := requireInteractive(); err != nil {
		return "", err
	}

	reader := bufio.NewReader(os.Stdin)
	if defaultValue != "" {
		fmt.Printf("%s [%s]: ", promptText, defaultValue)
	} else {
		fmt.Printf("%s: ", promptText)
	}

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue, nil
	}
	return input, nil
}

// Option represents a selectable option.
type Option struct {
	Label       string
	Description string
}

// Select prompts the user to select from a list of options.
// Returns the index of the selected option.
func Select(question string, options []Option) (int, error) {
	if err := requireInteractive(); err != nil {
		return -1, err
	}

	if len(options) == 0 {
		return -1, ErrNoOptions
	}

	reader := bufio.NewReader(os.Stdin)

	// Display the question and options
	fmt.Println(question)
	fmt.Println()
	for i, opt := range options {
		if opt.Description != "" {
			fmt.Printf("  %d) %s\n     %s\n", i+1, opt.Label, opt.Description)
		} else {
			fmt.Printf("  %d) %s\n", i+1, opt.Label)
		}
	}
	fmt.Println()

	for {
		fmt.Printf("Enter selection (1-%d): ", len(options))
		input, err := reader.ReadString('\n')
		if err != nil {
			return -1, fmt.Errorf("failed to read input: %w", err)
		}

		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(options) {
			fmt.Printf("Please enter a number between 1 and %d.\n", len(options))
			continue
		}

		return choice - 1, nil // Return 0-indexed
	}
}

// SelectString is a convenience wrapper that takes string options.
func SelectString(question string, options []string) (int, error) {
	opts := make([]Option, len(options))
	for i, s := range options {
		opts[i] = Option{Label: s}
	}
	return Select(question, opts)
}

// MultiSelect prompts the user to select multiple options.
// Returns the indices of selected options.
func MultiSelect(question string, options []Option) ([]int, error) {
	if err := requireInteractive(); err != nil {
		return nil, err
	}

	if len(options) == 0 {
		return nil, ErrNoOptions
	}

	reader := bufio.NewReader(os.Stdin)

	// Display the question and options
	fmt.Println(question)
	fmt.Println("(Enter comma-separated numbers, or 'all' for all, 'none' for none)")
	fmt.Println()
	for i, opt := range options {
		if opt.Description != "" {
			fmt.Printf("  %d) %s\n     %s\n", i+1, opt.Label, opt.Description)
		} else {
			fmt.Printf("  %d) %s\n", i+1, opt.Label)
		}
	}
	fmt.Println()

	for {
		fmt.Printf("Enter selections: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read input: %w", err)
		}

		input = strings.TrimSpace(strings.ToLower(input))

		if input == "all" {
			indices := make([]int, len(options))
			for i := range options {
				indices[i] = i
			}
			return indices, nil
		}

		if input == "none" || input == "" {
			return []int{}, nil
		}

		// Parse comma-separated numbers
		parts := strings.Split(input, ",")
		indices := make([]int, 0, len(parts))
		valid := true

		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			choice, err := strconv.Atoi(p)
			if err != nil || choice < 1 || choice > len(options) {
				fmt.Printf("Invalid selection: %s. Please enter numbers between 1 and %d.\n", p, len(options))
				valid = false
				break
			}
			indices = append(indices, choice-1) // 0-indexed
		}

		if valid {
			return indices, nil
		}
	}
}

// DisplayBox prints text in a simple box.
// Used for consent dialogs and other important messages.
func DisplayBox(title, content string) {
	width := 70

	// Top border
	fmt.Println("+" + strings.Repeat("-", width-2) + "+")

	// Title
	if title != "" {
		padding := (width - 2 - len(title)) / 2
		fmt.Printf("|%s%s%s|\n", strings.Repeat(" ", padding), title, strings.Repeat(" ", width-2-padding-len(title)))
		fmt.Println("+" + strings.Repeat("-", width-2) + "+")
	}

	// Content - word wrap
	lines := wrapText(content, width-4)
	for _, line := range lines {
		padding := width - 2 - len(line)
		fmt.Printf("| %s%s|\n", line, strings.Repeat(" ", padding-1))
	}

	// Bottom border
	fmt.Println("+" + strings.Repeat("-", width-2) + "+")
}

// wrapText wraps text to the given width.
func wrapText(text string, width int) []string {
	var lines []string
	paragraphs := strings.Split(text, "\n")

	for _, para := range paragraphs {
		if para == "" {
			lines = append(lines, "")
			continue
		}

		words := strings.Fields(para)
		if len(words) == 0 {
			lines = append(lines, "")
			continue
		}

		line := words[0]
		for _, word := range words[1:] {
			if len(line)+1+len(word) <= width {
				line += " " + word
			} else {
				lines = append(lines, line)
				line = word
			}
		}
		lines = append(lines, line)
	}

	return lines
}
