package functools

func All4Slice[T any, E ~[]T](condition func(T) bool, entry E) bool {
	for _, item := range entry {
		if !condition(item) {
			return false
		}
	}
	return true
}

func All4String(condition func(string) bool, entry string) bool {
	for _, charCode := range entry {
		if !condition(string(charCode)) {
			return false
		}
	}
	return true
}

func All4Chan[T any, E ~chan T](condition func(T) bool, entry E) bool {
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
