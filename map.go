package functools

func Map4Slice[T, U any, E ~[]T, R ~[]U](handler func(T) U, entry E) FuncNone2T[R] {
	return func() R {
		output := make(R, len(entry))
		for i, item := range entry {
			output[i] = handler(item)
		}
		return output
	}
}

func Map4String(handler func(string) string, entry string) FuncNone2T[string] {
	return func() string {
		output := make([]rune, 0)
		for _, charCode := range entry {
			output = append(output, []rune(handler(string(charCode)))...)
		}
		return string(output)
	}
}

func Map4Chan[T, U any, E ~chan T, R ~chan U](handler func(T) U, entry E) FuncNone2T[R] {
	return func() R {
		output := make(R, len(entry))
		cache := make(chan T, len(entry))
		defer close(cache)
		for item := range entry {
			cache <- item
			output <- handler(item)
		}
		for item := range cache {
			entry <- item
		}
		return output
	}
}
