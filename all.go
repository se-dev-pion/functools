package functools

func All4Slice[T any, E ~[]T](condition FuncT2Bool[T], entry E) bool {
	for _, item := range entry {
		if !condition(item) {
			return false
		}
	}
	return true
}

func All4String(condition FuncT2Bool[string], entry string) bool {
	for _, charCode := range entry {
		if !condition(string(charCode)) {
			return false
		}
	}
	return true
}

func All4Chan[T any, E ~chan T](condition FuncT2Bool[T], entry E) bool {
	return All4Slice(condition, extractChanElements(entry))
}

// All checks whether all elements in the given sequence(slice/chan/string) satisfy the specified condition.
func All[T any, E Sequence[T] | ~string](condition FuncT2Bool[T], entry E) bool {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return All4Slice(condition, e)
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			goto END
		}
		return All4String(any(condition).(FuncT2Bool[string]), e)
	case chan T:
		return All4Chan(condition, e)
	}
END:
	return false
}
