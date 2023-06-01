package main

type CacheKey interface {
	CacheKey() string
}

type ItemCache[T CacheKey] struct {
	items map[string]T
}

func (i *ItemCache[T]) Run(cacheKey string, run func() (T, error)) (T, error) {
	if i.items == nil {
		i.items = make(map[string]T)
	}

	item, ok := i.items[cacheKey]
	if ok {
		return item, nil
	}

	t, err := run()
	if err != nil {
		var t T
		return t, err
	}
	i.items[t.CacheKey()] = t
	return t, nil
}
