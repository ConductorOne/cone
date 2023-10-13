package client

import (
	"testing"
)

func TestGetJSONPathIndexNil(t *testing.T) {
	want := -1
	got, err := GetJSONPathIndex(nil)
	if err != nil {
		t.Errorf("GetJSONPathIndex(nil) returned error: %v", err)
	}
	if got != want {
		t.Errorf("GetJSONPathIndex(nil) = %v, want %v", got, want)
	}
}

func TestGetJsonPathIndexEmpty(t *testing.T) {
	want := -1
	str := ""
	got, err := GetJSONPathIndex(&str)
	if err != nil {
		t.Errorf("GetJSONPathIndex(\"\") returned error: %v", err)
	}
	if got != want {
		t.Errorf("GetJSONPathIndex(\"\") = %v, want %v", got, want)
	}
}

func TestGetJsonPathIndexBadRoot(t *testing.T) {
	str := "badroot"
	_, err := GetJSONPathIndex(&str)
	if err == nil {
		t.Errorf("GetJSONPathIndex(\"badroot\") returned no error")
	}
}

func TestGetJsonPathIndexNestedPath(t *testing.T) {
	str := "$.nested.path"
	_, err := GetJSONPathIndex(&str)
	if err == nil {
		t.Errorf("GetJSONPathIndex(\"$.nested.path\") returned no error")
	}
}

func TestGetInsideParenthesesNoParentheses(t *testing.T) {
	str := "no parentheses"
	_, err := getInsideParentheses(str)
	if err == nil {
		t.Errorf("GetJSONPathIndex(\"no parentheses\") returned no error")
	}
}

func TestGetInsideParenthesesNestedParentheses(t *testing.T) {
	str := "[[0]]"
	_, err := getInsideParentheses(str)
	if err == nil {
		t.Errorf("GetJSONPathIndex(\"[[0]]\") returned no error")
	}
}

func TestGetInsideParenthesesWildcard(t *testing.T) {
	str := "a[*]"
	_, err := getInsideParentheses(str)
	if err == nil {
		t.Errorf("GetJSONPathIndex(\"[*]\") returned no error")
	}
}

func TestGetInsideParenthesesInvalidIndex(t *testing.T) {
	str := "aa[invalid]"
	_, err := getInsideParentheses(str)
	if err == nil {
		t.Errorf("GetJSONPathIndex(\"[invalid]\") returned no error")
	}
}

func TestGetInsideParenthesesNegativeIndex(t *testing.T) {
	str := "asd[-1]"
	_, err := getInsideParentheses(str)
	if err == nil {
		t.Errorf("GetJSONPathIndex(\"asd[-1]\") returned no error")
	}
}
