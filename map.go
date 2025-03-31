package functools

import (
	"strings"

	"github.com/se-dev-pion/functools/types"
)

func Map4Slice[T, U any, E ~[]T, R []U](handler types.FuncT2R[T, U], entry E) types.FuncNone2T[R] {
	return func() R {
		output := make(R, len(entry))
		for i, item := range entry {
			output[i] = handler(item)
		}
		return output
	}
}

func Map4String(handler types.FuncT2T[string], entry string) types.FuncNone2T[string] {
	return func() string {
		var builder strings.Builder
		builder.Grow(len(entry))
		for _, charCode := range entry {
			builder.WriteString(handler(string(charCode)))
		}
		return builder.String()
	}
}

func Map4Chan[T, U any, E ~chan T, R chan U](handler types.FuncT2R[T, U], entry E) types.FuncNone2T[R] {
	return func() R {
		output := make(R, cap(entry))
		for _, item := range extractChanElements(entry) {
			output <- handler(item)
		}
		return output
	}
}

// Map creates a new types.Sequence composed of elements transformed from the input types.Sequence.
func Map[R types.Sequence[U] | ~string, T, U any, E types.Sequence[T] | ~string](handler types.FuncT2R[T, U], entry E) types.FuncNone2T[R] {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return any(Map4Slice(handler, e)).(types.FuncNone2T[R])
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			goto END
		}
		if _, ok := any(*new(U)).(string); !ok {
			goto END
		}
		return any(Map4String(types.FuncT2T[string](any(handler).(types.FuncT2R[string, string])), e)).(types.FuncNone2T[R])
	case chan T:
		return any(Map4Chan(handler, e)).(types.FuncNone2T[R])
	}
END:
	return nil
}
