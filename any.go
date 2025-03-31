package functools

import "github.com/se-dev-pion/functools/types"

func Any4Slice[T any, E ~[]T](condition types.FuncT2Bool[T], entry E) bool {
	for _, item := range entry {
		if condition(item) {
			return true
		}
	}
	return false
}

func Any4String(condition types.FuncT2Bool[string], entry string) bool {
	for _, charCode := range entry {
		if condition(string(charCode)) {
			return true
		}
	}
	return false
}

func Any4Chan[T any, E ~chan T](condition types.FuncT2Bool[T], entry E) bool {
	return Any4Slice(condition, extractChanElements(entry))
}

// Any checks if any element in the given types.Sequence(slice/chan/string) satisfies the specified condition.
func Any[T any, E types.Sequence[T] | ~string](condition types.FuncT2Bool[T], entry E) bool {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return Any4Slice(condition, e)
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			goto END
		}
		return Any4String(any(condition).(types.FuncT2Bool[string]), e)
	case chan T:
		return Any4Chan(condition, e)
	}
END:
	return false
}
