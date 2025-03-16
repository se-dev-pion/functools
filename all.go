package functools

func all4Slice[T any, E ~[]T](condition FuncT2Bool[T], entry E) bool {
	for _, item := range entry {
		if !condition(item) {
			return false
		}
	}
	return true
}

func all4String(condition FuncT2Bool[string], entry string) bool {
	for _, charCode := range entry {
		if !condition(string(charCode)) {
			return false
		}
	}
	return true
}

func all4Chan[T any, E ~chan T](condition FuncT2Bool[T], entry E) bool {
	success := true
	cache := make(chan T, len(entry))
	defer close(cache)
	for item := range entry {
		cache <- item
		if success && !condition(item) {
			success = false
		}
	}
	for item := range cache {
		entry <- item
	}
	return success
}

func All[T any, E ~[]T | ~string | ~chan T](condition FuncT2Bool[T], entry E) bool {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return all4Slice(condition, e)
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			return false
		}
		return all4String(any(condition).(FuncT2Bool[string]), e)
	case chan T:
		return all4Chan(condition, e)
	default:
		return false
	}
}
