package client

import (
	"testing"
)

func TestGetJSONPathIndex(t *testing.T) {
	tests := []struct {
		name    string
		input   *string
		want    int
		wantErr bool
	}{
		{name: "Nil", input: nil, want: -1, wantErr: false},
		{name: "Empty", input: newString(""), want: -1, wantErr: false},
		{name: "BadRoot", input: newString("badroot"), want: -1, wantErr: true},
		{name: "NestedPath", input: newString("$.nested.path"), want: -1, wantErr: true},
		{name: "WorkingPath", input: newString("$.expanded[0]"), want: 0, wantErr: false},
	}

	for _, tc := range tests {
		got, err := GetJSONPathIndex(tc.input)
		if (err != nil) != tc.wantErr {
			t.Errorf("%s: expected error: %v, got: %v", tc.name, tc.wantErr, err)
		}
		if got != tc.want {
			t.Errorf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestGetInsideParentheses(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{name: "NoParentheses", input: "no parentheses", want: -1, wantErr: true},
		{name: "NestedParentheses", input: "[[0]]", want: -1, wantErr: true},
		{name: "NumerousParentheses", input: "[1][1]", want: -1, wantErr: true},
		{name: "Wildcard", input: "a[*]", want: -1, wantErr: true},
		{name: "InvalidIndex", input: "aa[invalid]", want: -1, wantErr: true},
		{name: "NegativeIndex", input: "asd[-1]", want: -1, wantErr: true},
		{name: "ProperPath", input: "expanded[1]", want: 1, wantErr: false},
	}

	for _, tc := range tests {
		got, err := getIndexInsideParentheses(tc.input)
		if (err != nil) != tc.wantErr {
			t.Errorf("%s: expected error: %v, got: %v", tc.name, tc.wantErr, err)
		}
		if got != tc.want { // Adjust this based on the type of 'want'
			t.Errorf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

// Helper function to return a pointer to a string.
func newString(s string) *string {
	return &s
}
