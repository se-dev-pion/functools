package functools

type FuncT2R[T, R any] func(T) R

type FuncT2T[T any] FuncT2R[T, T]

type FuncT2Ts[T any] FuncT2R[T, []T]

type FuncNone2T[T any] func() T

type FuncTs2R[T, R any] func(...T) R

type FuncT2Bool[T any] FuncT2R[T, bool]

type FuncMergeT[T any] func(T, T) T
