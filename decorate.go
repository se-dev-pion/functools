package functools

func Decorate[T any](wrapper func(T) T, f T) T {
	return wrapper(f)
}
