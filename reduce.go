package functools

func reduce4Slice[T any, E ~[]T](handler FuncMergeT[T], entry E, initial ...T) FuncNone2T[T] {
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

func reduce4String(handler FuncMergeT[string], entry string, initial ...string) FuncNone2T[string] {
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

func Reduce[T any, E ~[]T | ~string | ~chan T](handler FuncMergeT[T], entry E, initial ...T) FuncNone2T[T] {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return reduce4Slice(handler, e, initial...)
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			goto END
		}
		return any(reduce4String(any(handler).(FuncMergeT[string]), e, any(initial).([]string)...)).(FuncNone2T[T])
	case chan T:
		return reduce4Slice(handler, extractChanElements(e), initial...)
	}
END:
	return nil
}
