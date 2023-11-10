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

func ParseFieldAttributes(filename string) (map[string]map[string]FieldAttribute, error) {
	file, err := os.Open("pkg/resource/md/" + filename + ".md")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	mapLevel := "root"
	fieldMappings := make(map[string]map[string]FieldAttribute)
	fieldMappings[mapLevel] = make(map[string]FieldAttribute)
	scanner := bufio.NewScanner(file)
	currentAttribute := UNKNOWN
	nestedPattern, err := regexp.Compile("### Nested Schema for `([^`]*)`")
	if err != nil {
		return nil, err
	}
	fieldPattern, err := regexp.Compile("- `([a-zA-Z_]+)`")
	if err != nil {
		return nil, err
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if matches := nestedPattern.FindStringSubmatch(line); len(matches) > 0 {
			mapLevel = matches[1]
			fieldMappings[mapLevel] = make(map[string]FieldAttribute)
		}

		if line == "### Required" || line == "Required:" {
			currentAttribute = REQUIRED
			continue
		} else if line == "### Optional" || line == "Optional:" {
			currentAttribute = OPTIONAL
			continue
		} else if line == "### Read-Only" || line == "Read-Only:" {
			currentAttribute = READ_ONLY
			continue
		}
		// Extract the field name and map it to the current attribute
		if currentAttribute != UNKNOWN {
			matches := fieldPattern.FindStringSubmatch(line)
			if len(matches) > 0 {
				match := matches[0][3 : len(matches[0])-1]
				fieldMappings[mapLevel][match] = currentAttribute
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return fieldMappings, nil
}
