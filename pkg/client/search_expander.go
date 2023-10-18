package client

import (
	"encoding/json"
	"errors"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

const (
	atTypeApp               = "type.googleapis.com/c1.api.app.v1.App"
	atTypeAppResource       = "type.googleapis.com/c1.api.app.v1.AppResource"
	atTypeAppResourceType   = "type.googleapis.com/c1.api.app.v1.AppResourceType"
	ExpandedApp             = "App"
	ExpandedAppResource     = "AppResource"
	ExpandedAppResourceType = "AppResourceType"
)

type PathDetails struct {
	Name string
	Path *string
}

type Expandable interface {
	GetPaths() []PathDetails
	SetPath(pathname string, value int)
}

// Populate the expanded map with references to the related objects.
func PopulateExpandedMap(expandMap map[string]int, expanded []any) map[string]*any {
	rv := make(map[string]*any)
	for k, v := range expandMap {
		rv[k] = &expanded[v]
	}
	return rv
}

type ExpandableReponse[T Expandable] struct {
	List []T
}

func (o ExpandableReponse[T]) PopulateExpandedIndexes() error {
	for _, v := range o.List {
		for _, path := range v.GetPaths() {
			expanded, err := GetJSONPathIndex(path.Path)
			if err != nil {
				return err
			}
			if expanded != -1 {
				v.SetPath(path.Name, expanded)
			}
		}
	}
	return nil
}

type ExpandedReponse interface {
	GetExpanded() map[string]*any
}

func GetExpanded[T any](e ExpandedReponse, key string) *T {
	var rv *T
	if x, ok := e.GetExpanded()[key]; ok {
		if x == nil {
			return nil
		}
		if x, ok := (*x).(*T); ok {
			rv = x
		}
	}
	return rv
}

type AnyType interface {
	MarshalJSON() ([]byte, error)
}

func As[T AnyType, V any](input T) (*V, error) {
	d, err := input.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var rv V
	err = json.Unmarshal(d, &rv)
	if err != nil {
		return nil, err
	}

	return &rv, nil
}

func UnmarshalAnyType[T AnyType, PT interface {
	*T
	GetAtType() *string
}](input PT) (any, error) {
	inputType := input.GetAtType()
	if inputType == nil {
		return nil, errors.New("input type is nil")
	}

	switch *inputType {
	case atTypeApp:
		return As[T, shared.App](*input)
	case atTypeAppResource:
		return As[T, shared.AppResource](*input)
	case atTypeAppResourceType:
		return As[T, shared.AppResourceType](*input)
	default:
		return nil, errors.New("unknown type")
	}
}
