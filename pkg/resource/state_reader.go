package resource

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

var closures = map[string]bool{
	"},": true,
	"],": true,
	"}":  true,
	"]":  true,
}
var resourceRegex = regexp.MustCompile(`\s+[a-zA-Z_0-9]+\s+= {`)
var fieldRegex = regexp.MustCompile(`\+\s+([a-zA-Z_]+)\s+=`)

type Stack struct {
	items []string
	// TODO: @anthony Make this into a struct, may help with Map Attributes
	mappings  map[string](map[string]map[string]FieldAttribute)
	resources map[string]TemplateData
	resource  TemplateData
	lock      bool
	Result    string
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	s.items = append(s.items, str)
}

// Pop removes and returns the top value from the stack. Returns empty string if stack is empty.
func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	index := len(s.items) - 1
	element := s.items[index]
	s.items = s.items[:index]
	return element
}

// Peek returns the top value of the stack without removing it. Returns empty string if stack is empty.
func (s *Stack) Peek() string {
	if len(s.items) == 0 {
		return ""
	}
	return s.items[len(s.items)-1]
}

// Empty returns true if the stack is empty, false otherwise.
func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack) Root() bool {
	return len(s.items) == 1
}

func (s *Stack) parseField(input string) error {
	if s.Root() {
		fieldMatch := fieldRegex.FindStringSubmatch(input)
		if len(fieldMatch) > 0 {
			fieldName := strings.TrimSpace(fieldMatch[1])
			fieldAtttribute := s.mappings[s.resource.GetType()]["root"][fieldName]
			if fieldAtttribute != READ_ONLY && fieldAtttribute != UNKNOWN {
				s.Result += strings.Repeat("\t", len(s.items)) + strings.TrimSpace(input)[2:] + "\n"
			}
		}
	}

	return nil
}
func (s *Stack) intializeResourceBlock(input string) error {
	// Extract the resource name
	matches := resourceRegex.FindStringSubmatch(input)
	if len(matches) > 0 {
		temp := strings.TrimSpace(matches[0])
		temp = strings.Split(temp, " ")[0]
		if r := s.resources[temp]; r != nil {
			s.resource = r
			s.Result += "resource \"" + r.GetType() + "\" " + "\"" + r.GetDatasourceId() + "\" " + "{\n"
		}
	} else {
		return errors.New("mapping error: unknown format, opening { without resource definition")
	}
	return nil
}

func (s *Stack) PushString(input string) error {
	// Parse the current line
	err := s.parseField(input)
	if err != nil {
		return err
	}

	var closure string
	if len(input) == 0 {
		return nil
	}
	if input[len(input)-1] == ',' && len(input) > 1 {
		closure = input[len(input)-2:]
	} else {
		closure = input[len(input)-1:]
	}

	// Check if input is a opening bracket
	if closure == "{" {
		if s.Peek() == "[" {
			s.Push("},")
		} else {
			if s.Empty() {
				err := s.intializeResourceBlock(input)
				if err != nil {
					return err
				}
			}
			s.Push("}")
		}
	}
	if closure == "[" {
		if s.Peek() == "[" {
			s.Push("],")
		} else {
			s.Push("]")
		}
	}

	// Check if input is a closure
	if closures[closure] {
		if s.Pop() != closure {
			return errors.New("mapping error: invalid closure")
		}
		if s.Empty() {
			s.Result += "}\n"
		}
	}

	return nil
}

func ParseHCLBlocks(outputDir string, mappings map[string](map[string]map[string]FieldAttribute), resources map[string]TemplateData) (string, error) {
	file, err := os.Open(outputDir + "/plan.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	stack := Stack{mappings: mappings, resources: resources}
	scanner := bufio.NewScanner(file)
	startRegex := regexp.MustCompile(`Changes to Outputs:`)
	start := false

	for scanner.Scan() {
		line := scanner.Text()

		if start {
			stack.PushString(line)
		}

		// Start parsing the output
		matches := startRegex.FindStringSubmatch(line)
		if len(matches) > 0 {
			start = true
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return stack.Result, nil
}
