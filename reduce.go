package functools

import "github.com/se-dev-pion/functools/types"

func Reduce4Slice[T any, E ~[]T](handler types.FuncMergeT[T], entry E, initial ...T) types.FuncNone2T[T] {
	var result T
	switch {
	case len(initial) > 0:
		result = initial[0]
	case len(entry) > 0:
		result = entry[0]
		entry = entry[1:]
	default:
		return nil
	}
	return func() T {
		for _, item := range entry {
			result = handler(result, item)
		}
		return result
	}
}

func Reduce4String(handler types.FuncMergeT[string], entry string, initial ...string) types.FuncNone2T[string] {
	var result string
	switch {
	case len(initial) > 0:
		result = initial[0]
	case len(entry) > 0:
		for _, charCode := range entry {
			result = string(charCode)
			break
		}
		entry = entry[1:]
	default:
		return nil
	}
	return func() string {
		for _, charCode := range entry {
			result = handler(result, string(charCode))
		}
		return result
	}
}

func Reduce4Chan[T any, E ~chan T](handler types.FuncMergeT[T], entry E, initial ...T) types.FuncNone2T[T] {
	return Reduce4Slice(handler, extractChanElements(entry), initial...)
}

// Reduce calculates the progressive processing of elements in the input types.Sequence using the specified function to obtain the result
func Reduce[T any, E types.Sequence[T] | ~string](handler types.FuncMergeT[T], entry E, initial ...T) types.FuncNone2T[T] {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return Reduce4Slice(handler, e, initial...)
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			goto END
		}
		return any(Reduce4String(any(handler).(types.FuncMergeT[string]), e, any(initial).([]string)...)).(types.FuncNone2T[T])
	case chan T:
		return Reduce4Chan(handler, e, initial...)
	}
END:
	return nil
}
