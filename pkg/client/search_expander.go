package client

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
