package functools

import "strings"

func Map4Slice[T, U any, E ~[]T, R []U](handler FuncT2R[T, U], entry E) FuncNone2T[R] {
	return func() R {
		output := make(R, len(entry))
		for i, item := range entry {
			output[i] = handler(item)
		}
		return output
	}
}

func Map4String(handler FuncT2T[string], entry string) FuncNone2T[string] {
	return func() string {
		var builder strings.Builder
		builder.Grow(len(entry))
		for _, charCode := range entry {
			builder.WriteString(handler(string(charCode)))
		}
		return builder.String()
	}
}

func Map4Chan[T, U any, E ~chan T, R chan U](handler FuncT2R[T, U], entry E) FuncNone2T[R] {
	return func() R {
		output := make(R, cap(entry))
		for _, item := range extractChanElements(entry) {
			output <- handler(item)
		}
		return output
	}
}

// Map creates a new sequence composed of elements transformed from the input sequence.
func Map[R Sequence[U] | ~string, T, U any, E Sequence[T] | ~string](handler FuncT2R[T, U], entry E) FuncNone2T[R] {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return any(Map4Slice(handler, e)).(FuncNone2T[R])
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			goto END
		}
		if _, ok := any(*new(U)).(string); !ok {
			goto END
		}
		return any(Map4String(FuncT2T[string](any(handler).(FuncT2R[string, string])), e)).(FuncNone2T[R])
	case chan T:
		return any(Map4Chan(handler, e)).(FuncNone2T[R])
	}
END:
	return nil
}
