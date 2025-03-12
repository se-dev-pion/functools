package functools

func Reduce4Slice[T any, E ~[]T](handler func(T, T) T, entry E, initial ...T) (result T) {
	switch {
	case len(initial) > 0:
		result = initial[0]
	case len(entry) > 0:
		result = entry[0]
		entry = entry[1:]
	default:
		panic("No initial value")
	}
	for _, item := range entry {
		result = handler(result, item)
	}
	return
}

func Reduce4String(handler func(string, string) string, entry string, initial ...string) (result string) {
	switch {
	case len(initial) > 0:
		result = initial[0]
	case len(entry) > 0:
		for _, charCode := range entry {
			result = string(charCode)
			break
		}
	default:
		panic("No initial value")
	}
	for _, charCode := range entry {
		result = handler(result, string(charCode))
	}
	return
}

func Reduce4Chan[T any, E ~chan T](handler func(T, T) T, entry E, initial ...T) (result T) {
	var cache chan T
	switch {
	case len(initial) > 0:
		cache = make(chan T, len(entry))
		result = initial[0]
	case len(entry) > 0:
		cache = make(chan T, len(entry))
		result = <-entry
		cache <- result
	default:
		panic("No initial value")
	}
	defer close(cache)
	for item := range entry {
		cache <- item
		result = handler(result, item)
	}
	for item := range cache {
		entry <- item
	}
	return
}
