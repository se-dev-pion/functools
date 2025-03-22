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
		for _, item := range extractChanElements(entry) {
			if condition(item) {
				output <- item
			}
		}
		return output
	}
}

// Filter creates a new sequence consisting of elements from the input sequence that meet the conditions.
func Filter[T any, E Sequence[T] | ~string](condition FuncT2Bool[T], entry E) FuncNone2T[E] {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return any(filter4Slice(condition, e)).(FuncNone2T[E])
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			goto END
		}
		return any(filter4String(any(condition).(FuncT2Bool[string]), e)).(FuncNone2T[E])
	case chan T:
		return any(filter4Chan(condition, e)).(FuncNone2T[E])
	}
END:
	return nil
}
