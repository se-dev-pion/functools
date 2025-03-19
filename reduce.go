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

func reduce4Chan[T any, E ~chan T](handler FuncMergeT[T], entry E, initial ...T) FuncNone2T[T] {
	if len(initial)|len(entry) == 0 {
		return nil
	}
	return func() (result T) {
		cache := make([]T, len(entry))
		i := 0
		switch {
		case len(initial) > 0:
			result = initial[0]
		case len(entry) > 0:
			result = <-entry
			cache[i] = result
			i++
		}
		for {
			select {
			case item, ok := <-entry:
				if !ok {
					goto END
				}
				cache[i] = item
				i++
				result = handler(result, item)
			default:
				goto END
			}
		}
	END:
		for _, item := range cache {
			entry <- item
		}
		return
	}
}

func Reduce[T any, E ~[]T | ~string | ~chan T](handler FuncMergeT[T], entry E, initial ...T) FuncNone2T[T] {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return reduce4Slice(handler, e, initial...)
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			return nil
		}
		return any(reduce4String(any(handler).(FuncMergeT[string]), e, any(initial).([]string)...)).(FuncNone2T[T])
	case chan T:
		return reduce4Chan(handler, e, initial...)
	default:
		return nil
	}
}
