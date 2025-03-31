package functools

import (
	"strings"

	"github.com/se-dev-pion/functools/types"
)

func Filter4Slice[T any, E ~[]T](condition types.FuncT2Bool[T], entry E) types.FuncNone2T[E] {
	return func() E {
		output := make(E, 0)
		for _, item := range entry {
			if condition(item) {
				output = append(output, item)
			}
		}
		return output
	}
}

func Filter4String(condition types.FuncT2Bool[string], entry string) types.FuncNone2T[string] {
	return func() string {
		var builder strings.Builder
		builder.Grow(len(entry))
		for _, charCode := range entry {
			if condition(string(charCode)) {
				builder.WriteRune(charCode)
			}
		}
		return builder.String()
	}
}

func Filter4Chan[T any, E ~chan T](condition types.FuncT2Bool[T], entry E) types.FuncNone2T[E] {
	return func() E {
		output := make(E, cap(entry))
		for _, item := range extractChanElements(entry) {
			if condition(item) {
				output <- item
			}
		}
		return output
	}
}

// Filter creates a new types.Sequence consisting of elements from the input types.Sequence that meet the conditions.
func Filter[T any, E types.Sequence[T] | ~string](condition types.FuncT2Bool[T], entry E) types.FuncNone2T[E] {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return any(Filter4Slice(condition, e)).(types.FuncNone2T[E])
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			goto END
		}
		return any(Filter4String(any(condition).(types.FuncT2Bool[string]), e)).(types.FuncNone2T[E])
	case chan T:
		return any(Filter4Chan(condition, e)).(types.FuncNone2T[E])
	}
END:
	return nil
}
