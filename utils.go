package functools

type FuncT2T[T any] func(T) T

type FuncT2R[T, R any] func(T) R

type FuncNone2T[T any] func() T

type FuncTs2R[T, R any] func(...T) R

type FuncT2Ts[T any] func(T) []T

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

func Stream[T any](f ...FuncT2T[T]) FuncT2T[T] {
	return func(param T) T {
		output := param
		for _, fn := range f {
			output = fn(output)
		}
		return output
	}
}

func Batch[T any](f ...FuncT2T[T]) FuncT2Ts[T] {
	return func(param T) []T {
		output := make([]T, len(f))
		for i, fn := range f {
			output[i] = fn(param)
		}
		return output
	}
}
