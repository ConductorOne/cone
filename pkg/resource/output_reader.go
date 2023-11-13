package resource

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

var resourceRegex = regexp.MustCompile(`\s+[a-zA-Z_0-9]+\s+= {`)
var fieldRegex = regexp.MustCompile(`\+\s+([a-zA-Z_]+)\s+=`)

type levelAttribute struct {
	Locked    bool
	FieldName string
}

type levelAttributes struct {
	levelAttributes []levelAttribute
	levelString     string
}

func (l *levelAttributes) GetLevel() string {
	if l.levelString == "" {
		return "root"
	} else {
		return l.levelString
	}
}

func (l *levelAttributes) AddLevel(lock bool, fieldName string) {
	if l.IsLocked() {
		// TODO: @anthony add error handling
		l.levelAttributes = append(l.levelAttributes, levelAttribute{true, fieldName})
	} else {
		l.levelAttributes = append(l.levelAttributes, levelAttribute{lock, fieldName})
	}

	if fieldName != "" {
		if l.levelString == "" {
			l.levelString = fieldName
		} else {
			l.levelString += "." + fieldName
		}
	}
}

// Peek and determine if the current level is locked.
func (l *levelAttributes) IsLocked() bool {
	if len(l.levelAttributes) == 0 {
		return false
	}
	element := l.levelAttributes[len(l.levelAttributes)-1]
	return element.Locked
}

func (l *levelAttributes) RemoveLevel() {
	if len(l.levelAttributes) == 0 {
		return
	}

	index := len(l.levelAttributes) - 1
	element := l.levelAttributes[index]
	l.levelAttributes = l.levelAttributes[:index]

	if element.FieldName != "" {
		l.levelString = l.levelString[:len(l.levelString)-len(element.FieldName)]
		if len(l.levelString) != 0 && l.levelString[len(l.levelString)-1:] == "." {
			l.levelString = l.levelString[:len(l.levelString)-1]
		}
	}
}

func (l *levelAttributes) GetNextLevel(fieldName string) string {
	if l.levelString == "" {
		return fieldName
	} else {
		return l.levelString + "." + fieldName
	}
}

type Stack struct {
	items           []string
	levelAttributes levelAttributes
	mappings        map[string](map[string]map[string]FieldAttribute)
	resources       map[string]TemplateData
	resource        TemplateData
	Result          string
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
		} else {
			return errors.New("mapping error: unknown resource import")
		}
	} else {
		return nil
	}
	ok := s.Push("{", false, "")
	if !ok {
		return errors.New("mapping error: unable to push onto stack")
	}
	return nil
}

func (s *Stack) parseField(input string, fieldName string) error {
	// Check if the input is nested
	isNested := s.mappings[s.resource.GetType()][s.levelAttributes.GetNextLevel(fieldName)] != nil

	// Only Include the field if it is not read-only
	// TODO: @anthony unknown fields are allowed for now, fix this later
	fieldAtttribute := s.mappings[s.resource.GetType()][s.levelAttributes.GetLevel()][fieldName]
	if fieldAtttribute != ReadOnly {
		s.Result += strings.Repeat("\t", len(s.items)) + strings.TrimSpace(input)[2:] + "\n"
	}

	if isNested {
		s.Push(input, fieldAtttribute != ReadOnly, fieldName)
	} else {
		s.Push(input, false, "")
	}
	return nil
}

func (s *Stack) checkLine(input string) error {
	var closure string
	if len(input) == 0 {
		return nil
	}
	if input[len(input)-1:] == "," && len(input) > 1 {
		closure = input[len(input)-2:]
	} else {
		closure = input[len(input)-1:]
	}

	switch closure {
	case "{":
		s.Push(input, false, "")
		s.Result += strings.Repeat("\t", len(s.items)) + closure + "\n"
		return nil
	case "[":
		s.Push(input, false, "")
		s.Result += strings.Repeat("\t", len(s.items)) + closure + "\n"
		return nil
	case "}":
		x := s.Pop()
		if x == "{" {
			s.Result += strings.Repeat("\t", len(s.items)) + closure + "\n"
		} else {
			return errors.New("mapping error: unknown format, closing } without opening {")
		}
		return nil
	case "]":
		x := s.Pop()
		if x == "[" {
			s.Result += strings.Repeat("\t", len(s.items)) + closure + "\n"
		} else {
			return errors.New("mapping error: unknown format, closing ] without opening [")
		}
		return nil
	case "],":
		x := s.Pop()
		if x == "[" && s.Peek() == "[" {
			s.Result += strings.Repeat("\t", len(s.items)) + closure + "\n"
		} else {
			return errors.New("mapping error: unknown format, closing ] without opening [")
		}
		return nil
	case "},":
		x := s.Pop()
		if x == "{" && s.Peek() == "[" {
			s.Result += strings.Repeat("\t", len(s.items)) + closure + "\n"
		} else {
			return errors.New("mapping error: unknown format, closing ] without opening [")
		}
		return nil
	}

	str := strings.TrimSpace(input)[2:]
	s.Result += strings.Repeat("\t", len(s.items)) + str + "\n"
	return nil
}

func (s *Stack) CheckLine(input string) error {
	// Check if the input is the start of a new resource block
	if s.Empty() {
		return s.intializeResourceBlock(input)
	}

	// Check if the input is a field
	fieldName := ""
	if fieldMatch := fieldRegex.FindStringSubmatch(input); len(fieldMatch) > 0 {
		fieldName = strings.TrimSpace(fieldMatch[1])
	}

	if fieldName != "" {
		return s.parseField(input, fieldName)
	}

	// All other cases
	return s.checkLine(input)
}

// Push a new value onto the stack.
func (s *Stack) Push(str string, lock bool, fieldName string) bool {
	x := str[len(str)-1:]
	if x != "{" && x != "[" {
		return false
	}
	s.items = append(s.items, x)
	s.levelAttributes.AddLevel(lock, fieldName)
	return true
}

// Pop removes and returns the top value from the stack. Returns empty string if stack is empty.
func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	index := len(s.items) - 1
	element := s.items[index]
	s.items = s.items[:index]
	s.levelAttributes.RemoveLevel()
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

func (s *Stack) Level() int {
	return len(s.items)
}

func ParseHCLBlocks(outputPath string, mappings map[string](map[string]map[string]FieldAttribute), resources map[string]TemplateData) (string, error) {
	file, err := os.Open(outputPath)
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
			err := stack.CheckLine(line)
			if err != nil {
				return "", err
			}
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
