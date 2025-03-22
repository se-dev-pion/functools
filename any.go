package functools

func any4Slice[T any, E ~[]T](condition FuncT2Bool[T], entry E) bool {
	for _, item := range entry {
		if condition(item) {
			return true
		}
	}
	return false
}

func any4String(condition FuncT2Bool[string], entry string) bool {
	for _, charCode := range entry {
		if condition(string(charCode)) {
			return true
		}
	}
	return false
}

// Any checks if any element in the given sequence(slice/chan/string) satisfies the specified condition.
func Any[T any, E Sequence[T] | ~string](condition FuncT2Bool[T], entry E) bool {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return any4Slice(condition, e)
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			goto END
		}
		return any4String(any(condition).(FuncT2Bool[string]), e)
	case chan T:
		return any4Slice(condition, extractChanElements(e))
	}
END:
	return false
}
