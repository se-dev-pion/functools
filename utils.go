package functools

import "sync"

func Decorate[T any](wrapper FuncT2T[T], f T) T {
	return wrapper(f)
}

func Pack[T any](params ...T) []T {
	return params
}

func Lazy[T, R any](f FuncT2R[T, R], param T) FuncNone2T[R] {
	return func() R {
		return f(param)
	}
}

func Partial[T, R any](f FuncTs2R[T, R], params ...T) FuncTs2R[T, R] {
	return func(others ...T) R {
		return f(append(params, others...)...)
	}
}

func Flow[T any](f ...FuncT2T[T]) FuncT2T[T] {
	return func(param T) T {
		output := param
		for _, fn := range f {
			output = fn(output)
		}
		return output
	}
}

func Batch[T any](parallel bool, f ...FuncT2T[T]) FuncT2Ts[T] {
	return func(param T) []T {
		output := make([]T, len(f))
		if !parallel {
			for i, fn := range f {
				output[i] = fn(param)
			}
		} else {
			var wg sync.WaitGroup
			wg.Add(len(f))
			for i, fn := range f {
				go func(idx int, fun FuncT2T[T]) {
					output[idx] = fun(param)
					wg.Done()
				}(i, fn)
			}
			wg.Wait()
		}
		return output
	}
}

func Cached[T comparable, R any](f FuncT2R[T, R]) FuncT2R[T, R] {
	cache := make(map[T]R)
	return func(param T) R {
		if v, ok := cache[param]; ok {
			return v
		}
		cache[param] = f(param)
		return cache[param]
	}
}

func Copy[T any, E ~[]T | ~chan T](entry E) E {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		return any(Pack(e...)).(E)
	case chan T:
		cache := make(chan T, len(e))
		defer close(cache)
		output := make(chan T, len(e))
		for item := range e {
			cache <- item
			output <- item
		}
		for item := range cache {
			e <- item
		}
		return any(output).(E)
	default:
		return nil
	}
}
