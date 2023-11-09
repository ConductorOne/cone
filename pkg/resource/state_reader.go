package resource

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ParseHCLBlocks(outputDir string, filename string, mappings map[string]map[string]FieldAttribute) (string, error) {
	file, err := os.Open(outputDir + "/plan.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pattern := `\s+#\s+([a-zA-Z_.0-9]+) will be imported`
	rStart := regexp.MustCompile(pattern)
	resourceFmt := regexp.MustCompile(`\s*#\s*`)
	rAttribute := regexp.MustCompile(`\s+[a-zA-Z_]+\s+= \"`)

	res := ""
	start := false
	resource := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "}" {
			res = res + line + "\n"
			start = false
			resource = ""
			continue
		}
		if start && resource != "" {
			matches := rAttribute.FindStringSubmatch(line)
			if len(matches) > 0 {
				attribute := strings.TrimSpace(matches[0])
				attribute = strings.Split(attribute, " ")[0]
				if mappings[resource][attribute] == READ_ONLY {
					continue
				}
			}
			res = res + line + "\n"
		}
		matches := rStart.FindStringSubmatch(line)
		if len(matches) > 0 {
			start = true
			unformattedResource := strings.Split(matches[0], ".")[0]
			resource = resourceFmt.ReplaceAllString(unformattedResource, "")
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return res, nil
}
