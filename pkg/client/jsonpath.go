package client

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var matchParentheses = regexp.MustCompile(`\[\d+\]`)

func getInsideParentheses(str string) (int, error) {
	matches := matchParentheses.FindStringSubmatch(str)
	if len(matches) == 0 {
		return -1, nil
	}
	if len(matches) > 1 {
		return -1, errors.New("jsonpath: invalid path, nested parentheses not supported")
	}

	match := matches[0][1 : len(matches[0])-1]
	if match == "*" {
		return -1, errors.New("jsonpath: invalid path, wildcard array indexing not supported")
	}
	index, err := strconv.Atoi(match)
	if err != nil {
		return -1, fmt.Errorf("jsonpath: invalid path, only array index operations are supported: %w", err)
	}
	return index, nil
}

// Implementation of JSONPath only intended to allow expansion of non nested search results using array indexing.
// If you need to support more complex JSONPath operations, please use a library.
// Example: matches "$.expanded[0]".
func GetJSONPathIndex(jsonpath *string) (int, error) {
	if jsonpath == nil {
		return -1, nil
	}
	// Only support dot notation for now
	path := strings.Split(*jsonpath, ".")
	if len(path) == 0 {
		return -1, nil
	}
	if len(path) > 1 {
		return -1, errors.New("jsonpath: invalid path, nested jsonpath operations are not supported")
	}

	if path[0] != "$" {
		return -1, errors.New("jsonpath: invalid path, no root element")
	}

	return getInsideParentheses(path[1])
}
