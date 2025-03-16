package functools

func Reduce4Slice[T any, E ~[]T](handler func(T, T) T, entry E, initial ...T) FuncNone2T[T] {
	var result T
	switch {
	case len(initial) > 0:
		result = initial[0]
	case len(entry) > 0:
		result = entry[0]
		entry = entry[1:]
	default:
		return nil
	}
	return func() T {
		for _, item := range entry {
			result = handler(result, item)
		}
		return result
	}
}

func Reduce4String(handler func(string, string) string, entry string, initial ...string) FuncNone2T[string] {
	var result string
	switch {
	case len(initial) > 0:
		result = initial[0]
	case len(entry) > 0:
		for _, charCode := range entry {
			result = string(charCode)
			break
		}
	default:
		return nil
	}
	return func() string {
		for _, charCode := range entry {
			result = handler(result, string(charCode))
		}
		return result
	}
}

func Reduce4Chan[T any, E ~chan T](handler func(T, T) T, entry E, initial ...T) FuncNone2T[T] {
	if len(initial)|len(entry) == 0 {
		return nil
	}
	return func() (result T) {
		cache := make(chan T, len(entry))
		defer close(cache)
		switch {
		case len(initial) > 0:
			result = initial[0]
		case len(entry) > 0:
			result = <-entry
			cache <- result
		}
		for item := range entry {
			cache <- item
			result = handler(result, item)
		}
		for item := range cache {
			entry <- item
		}
		return
	}
}
