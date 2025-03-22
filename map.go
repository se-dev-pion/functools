package functools

func map4Slice[T, U any, E ~[]T, R ~[]U](handler FuncT2R[T, U], entry E) FuncNone2T[R] {
	return func() R {
		output := make(R, len(entry))
		for i, item := range entry {
			output[i] = handler(item)
		}
		return output
	}
}

func map4String(handler FuncT2T[string], entry string) FuncNone2T[string] {
	return func() string {
		output := make([]rune, 0)
		for _, charCode := range entry {
			output = append(output, []rune(handler(string(charCode)))...)
		}
		return string(output)
	}
}

func map4Chan[T, U any, E ~chan T, R ~chan U](handler FuncT2R[T, U], entry E) FuncNone2T[R] {
	return func() R {
		output := make(R, cap(entry))
		for _, item := range extractChanElements(entry) {
			output <- handler(item)
		}
		return output
	}
}

func Map[R ~[]U | ~string | ~chan U, T, U any, E ~[]T | ~string | ~chan T](handler FuncT2R[T, U], entry E) FuncNone2T[R] {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return any(map4Slice[T, U, []T, []U](handler, e)).(FuncNone2T[R])
	case string:
		if _, ok := any(*new(T)).(string); !ok {
			goto END
		}
		if _, ok := any(*new(U)).(string); !ok {
			goto END
		}
		return any(map4String(FuncT2T[string](any(handler).(FuncT2R[string, string])), e)).(FuncNone2T[R])
	case chan T:
		return any(map4Chan[T, U, chan T, chan U](handler, e)).(FuncNone2T[R])
	}
END:
	return nil
}
