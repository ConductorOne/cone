package client

import (
	"regexp"
	"strconv"
)

type PathDetails struct {
	Name string
	Path *string
}

type Expandable interface {
	GetPaths() []PathDetails
	SetPath(pathname string, value int)
}

func PopulateExpandedMap[T any](expandMap map[string]int, expanded []T) map[string]*T {
	rv := make(map[string]*T)
	for k, v := range expandMap {
		rv[k] = &expanded[v]
	}
	return rv
}

type ExpandableReponse[T Expandable] struct {
	List []T
}

func (o ExpandableReponse[T]) PopulateExpandedIndexes() {
	for _, v := range o.List {
		for _, path := range v.GetPaths() {
			expanded := o.GetJSONPath(path.Path)
			if expanded != -1 {
				v.SetPath(path.Name, expanded)
			}
		}
	}
}

func (o ExpandableReponse[T]) GetJSONPath(jsonpath *string) int {
	if jsonpath == nil || *jsonpath == "" {
		return -1
	}
	// TODO: support more complex paths?
	// Matches "$.expanded[0]"
	re := regexp.MustCompile(`\[\d+\]`)

	matches := re.FindStringSubmatch(*jsonpath)
	if len(matches) > 0 {
		// Remove the brackets to leave just the number
		indexStr, err := strconv.Atoi(matches[0][1 : len(matches[0])-1])
		if err != nil {
			return -1
		}

		return indexStr
	}
	return -1
}
