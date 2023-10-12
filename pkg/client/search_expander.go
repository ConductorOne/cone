package client

import (
	"regexp"
	"strconv"
)

type Expanded interface {
	GetAtType() *string
	GetAdditionalProperties() map[string]interface{}
}

type Expandable[E Expanded] interface {
	GetPaths() []*string
}

type ExpandableReponse[E Expanded, T Expandable[E]] struct {
	// 1D array of expanded objects
	Expanded []E
	// List of objects with references to expanded objects
	List []T
}

func (o *ExpandableReponse[E, T]) GetExpanded() []map[string]*E {
	rv := make([]map[string]*E, len(o.List))
	for _, v := range o.List {
		temp := make(map[string]*E)
		for _, path := range v.GetPaths() {
			expanded := o.GetJSONPath(path)
			if expanded != nil {
				temp[*path] = expanded
			}
		}
		rv = append(rv, temp)
	}
	return rv
}

func (o *ExpandableReponse[E, T]) GetJSONPath(jsonpath *string) *E {
	if jsonpath == nil || *jsonpath == "" {
		return nil
	}
	// TODO: support more complex paths?
	// Matches "$.expanded[0]"
	re := regexp.MustCompile(`\[\d+\]`)

	matches := re.FindStringSubmatch(*jsonpath)
	if len(matches) > 0 {
		// Remove the brackets to leave just the number
		indexStr, err := strconv.Atoi(matches[0][1 : len(matches[0])-1])
		if err != nil {
			return nil
		}
		return &o.Expanded[indexStr]
	}
	return nil
}
