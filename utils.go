package functools

import "sync"

// Decorate wraps something with a function, mainly used for enhancing a function with another function.
func Decorate[T any](wrapper FuncT2T[T], f T) T {
	return wrapper(f)
}

// Pack creates a slice for input elements.
func Pack[T any](params ...T) []T {
	return params
}

// Lazy creates a function with no params that returns the result of input function called with input params.
func Lazy[T, R any](f FuncT2R[T, R], param T) FuncNone2T[R] {
	return func() R {
		return f(param)
	}
}

// Partial creates a function that fixes input params with the input function.
func Partial[T, R any](f FuncTs2R[T, R], params ...T) FuncTs2R[T, R] {
	return func(others ...T) R {
		return f(append(params, others...)...)
	}
}

// Flow creates a function connecting the handling process of input functions.
func Flow[T any](f ...FuncT2T[T]) FuncT2T[T] {
	return func(param T) T {
		output := param
		for _, fn := range f {
			output = fn(output)
		}
		return output
	}
}

// Batch creates a function merging the handling process of input functions.
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

// Cached creates a function with cache that works in the same way as the input function.
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

func extractChanElements[T any](ch chan T) []T {
	output := make([]T, len(ch))
	defer func() {
		for _, item := range output {
			ch <- item
		}
	}()
	i := 0
	for {
		select {
		case item, ok := <-ch:
			if !ok {
				goto END
			}
			output[i] = item
			i++
		default:
			goto END
		}
	}
END:
	return output
}

// Copy creates a shallow copy of slice/chan
func Copy[T any, E Sequence[T]](entry E) (copy E) {
	v := any(entry)
	switch e := v.(type) {
	case []T:
		copy = any(Pack(e...)).(E)
	case chan T:
		output := make(chan T, cap(e))
		for _, item := range extractChanElements(e) {
			output <- item
		}
		copy = any(output).(E)
	}
	return
}
