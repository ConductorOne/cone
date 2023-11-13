package resource

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type FieldAttribute int

const (
	Unknown FieldAttribute = iota
	ReadOnly
	Required
	Optional
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
	currentAttribute := Unknown
	nestedPattern := regexp.MustCompile("### Nested Schema for `([^`]*)`")
	fieldPattern := regexp.MustCompile("- `([a-zA-Z_]+)`")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if matches := nestedPattern.FindStringSubmatch(line); len(matches) > 0 {
			mapLevel = matches[1]
			fieldMappings[mapLevel] = make(map[string]FieldAttribute)
		}

		switch line {
		case "### Required":
			fallthrough
		case "Required:":
			currentAttribute = Required
			continue
		case "### Optional":
			fallthrough
		case "Optional:":
			currentAttribute = Optional
			continue
		case "### Read-Only":
			fallthrough
		case "Read-Only:":
			currentAttribute = ReadOnly
			continue
		}
		// Extract the field name and map it to the current attribute
		if currentAttribute != Unknown {
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
