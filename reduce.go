package functools

func Reduce4Slice[T any, E ~[]T](handler FuncMergeT[T], entry E, initial ...T) FuncNone2T[T] {
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

func Reduce4String(handler FuncMergeT[string], entry string, initial ...string) FuncNone2T[string] {
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

func Reduce4Chan[T any, E ~chan T](handler FuncMergeT[T], entry E, initial ...T) FuncNone2T[T] {
	return Reduce4Slice(handler, extractChanElements(entry), initial...)
}

// Reduce calculates the progressive processing of elements in the input sequence using the specified function to obtain the result
func Reduce[T any, E Sequence[T] | ~string](handler FuncMergeT[T], entry E, initial ...T) FuncNone2T[T] {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return Reduce4Slice(handler, e, initial...)
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			goto END
		}
		return any(Reduce4String(any(handler).(FuncMergeT[string]), e, any(initial).([]string)...)).(FuncNone2T[T])
	case chan T:
		return Reduce4Chan(handler, e, initial...)
	}
END:
	return nil
}
