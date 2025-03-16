package functools

func filter4Slice[T any, E ~[]T](condition FuncT2Bool[T], entry E) FuncNone2T[E] {
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

func filter4String(condition FuncT2Bool[string], entry string) FuncNone2T[string] {
	return func() string {
		output := make([]rune, 0)
		for _, charCode := range entry {
			if condition(string(charCode)) {
				output = append(output, charCode)
			}
		}
		return string(output)
	}
}

func filter4Chan[T any, E ~chan T](condition FuncT2Bool[T], entry E) FuncNone2T[E] {
	return func() E {
		output := make(E, cap(entry))
		cache := make([]T, len(entry))
		i := 0
		for {
			select {
			case item, ok := <-entry:
				if !ok {
					goto END
				}
				cache[i] = item
				i++
				if condition(item) {
					output <- item
				}
			default:
				goto END
			}
		}
	END:
		for _, item := range cache {
			entry <- item
		}
		return output
	}
}

func Filter[T any, E ~[]T | ~string | ~chan T](condition FuncT2Bool[T], entry E) FuncNone2T[E] {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return any(filter4Slice(condition, e)).(FuncNone2T[E])
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			return nil
		}
		return any(filter4String(any(condition).(FuncT2Bool[string]), e)).(FuncNone2T[E])
	case chan T:
		return any(filter4Chan(condition, e)).(FuncNone2T[E])
	default:
		return nil
	}
}
