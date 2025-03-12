package functools

func Filter4Slice[T any, E ~[]T](condition func(T) bool, entry E) E {
	output := make(E, 0)
	for _, item := range entry {
		if condition(item) {
			output = append(output, item)
		}
	}
	return output
}

func Filter4String(condition func(string) bool, entry string) string {
	output := make([]rune, 0)
	for _, charCode := range entry {
		if condition(string(charCode)) {
			output = append(output, charCode)
		}
	}
	return string(output)
}

func Filter4Chan[T any, E ~chan T](condition func(T) bool, entry E) E {
	output := make(E, len(entry))
	cache := make(chan T, len(entry))
	defer close(cache)
	for item := range entry {
		cache <- item
		if condition(item) {
			output <- item
		}
	}
	for item := range cache {
		entry <- item
	}
	return output
}
