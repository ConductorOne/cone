package resource

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type FieldAttribute int

const (
	UNKNOWN FieldAttribute = iota
	READ_ONLY
	REQUIRED
	OPTIONAL
)

func ParseFieldAttributes(filename string) (map[string]FieldAttribute, error) {
	file, err := os.Open("pkg/resource/md/" + filename + ".md")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fieldMappings := make(map[string]FieldAttribute)
	scanner := bufio.NewScanner(file)
	currentAttribute := UNKNOWN
	pattern := "- `([a-zA-Z_]+)`"
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "### Required" {
			currentAttribute = REQUIRED
			continue
		} else if line == "### Optional" {
			currentAttribute = OPTIONAL
			continue
		} else if line == "### Read-Only" {
			currentAttribute = READ_ONLY
			continue
		}
		// Extract the field name and map it to the current attribute
		if currentAttribute != UNKNOWN {
			matches := r.FindStringSubmatch(line)
			if len(matches) > 0 {
				match := matches[0][3 : len(matches[0])-1]
				fieldMappings[match] = currentAttribute
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return fieldMappings, nil
}
