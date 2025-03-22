package functools

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestDecorate(t *testing.T) {
	convey.Convey("decorate", t, func() {
		wrapper := func(f func(int, int) int) func(int, int) int {
			return func(x, y int) int {
				return 2 * f(x, y)
			}
		}
		f := func(a, b int) int {
			return a + b
		}
		assert.Equal(t, 6, Decorate(wrapper, f)(1, 2))
	})
}

func ExampleDecorate() {
	wrapper := func(f func(int, int) int) func(int, int) int {
		return func(x, y int) int {
			return 2 * f(x, y)
		}
	}
	f := func(a, b int) int {
		return a + b
	}
	fmt.Println(Decorate(wrapper, f)(1, 2))
	// Output:
	// 6
}

func TestPack(t *testing.T) {
	convey.Convey("pack", t, func() {
		assert.Equal(t, []int{1, 2, 3}, Pack(1, 2, 3))
		assert.Equal(t, []string{"a", "b", "c"}, Pack("a", "b", "c"))
	})
}

func ExamplePack() {
	fmt.Println(Pack(1, 2, 3))
	fmt.Println(Pack("a", "b", "c"))
	// Output:
	// [1 2 3]
	// [a b c]
}

func TestLazy(t *testing.T) {
	convey.Convey("lazy", t, func() {
		f := func(x int) int {
			return x * 2
		}
		assert.Equal(t, f(2), Lazy(f, 2)())
	})
}

func ExampleLazy() {
	f := func(x int) int {
		return x * 2
	}
	fmt.Println(Lazy(f, 2)())
	// Output:
	// 4
}

func TestPartial(t *testing.T) {
	convey.Convey("partial", t, func() {
		f := func(l ...int) int {
			s := 0
			for _, x := range l {
				s += x
			}
			return s
		}
		g := Partial(f, 1, 2, 3)
		assert.Equal(t, 1+2+3+4, g(4))
	})
}

func ExamplePartial() {
	f := func(l ...int) int {
		s := 0
		for _, x := range l {
			s += x
		}
		return s
	}
	g := Partial(f, 1, 2, 3)
	fmt.Println(g(4))
	// Output:
	// 10
}

func TestFlow(t *testing.T) {
	convey.Convey("flow", t, func() {
		f1 := func(a int) int {
			return a + 1
		}
		f2 := func(a int) int {
			return a * 2
		}
		f3 := func(a int) int {
			return a * a
		}
		assert.Equal(t, 16, Flow(f1, f2, f3)(1))
	})
}

func ExampleFlow() {
	f1 := func(a int) int {
		return a + 1
	}
	f2 := func(a int) int {
		return a * 2
	}
	f3 := func(a int) int {
		return a * a
	}
	fmt.Println(Flow(f1, f2, f3)(1))
	// Output:
	// 16
}

func TestBatch(t *testing.T) {
	convey.Convey("batch", t, func() {
		f1 := func(a int) int {
			return a + 1
		}
		f2 := func(a int) int {
			return a * 2
		}
		f3 := func(a int) int {
			return a * a
		}
		assert.Equal(t, []int{2, 2, 1}, Batch(f1, f2, f3)(1))
	})
}

func ExampleBatch() {
	f1 := func(a int) int {
		return a + 1
	}
	f2 := func(a int) int {
		return a * 2
	}
	f3 := func(a int) int {
		return a * a
	}
	fmt.Println(Batch(f1, f2, f3)(1))
	// Output:
	// [2 2 1]
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func TestCached(t *testing.T) {
	convey.Convey("cache", t, func() {
		f := Cached(fibonacci)
		for i := 0; i <= 10; i++ {
			assert.Equal(t, fibonacci(10), f(10))
		}
	})
}

func ExampleCached() {
	f := Cached(fibonacci)
	fmt.Println(f(10))
	// Output:
	// 55
}

func BenchmarkCached(b *testing.B) {
	b.Run("raw", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fibonacci(10)
		}
	})
	f := Cached(fibonacci)
	b.Run("cached", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f(10)
		}
	})
}

func TestCopy(t *testing.T) {
	convey.Convey("copy", t, func() {
		convey.Convey("slice", func() {
			arr := []int{1, 2, 3, 4, 5}
			assert.Equal(t, arr, Copy[int](arr))
		})
		convey.Convey("chan", func() {
			ch := make(chan int, 10)
			for i := 1; i <= 5; i++ {
				ch <- i
			}
			chCopy := Copy[int](ch)
			assert.Equal(t, []int{1, 2, 3, 4, 5}, extractChanElements(chCopy))
		})
	})
}

func ExampleCopy() {
	// [slice]
	{
		arr := []int{1, 2, 3, 4, 5}
		fmt.Println(Copy[int](arr))
	} // [/]
	// [chan]
	{
		ch := make(chan int, 10)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		chCopy := Copy[int](ch)
		fmt.Println(extractChanElements(chCopy))
	} // [/]
	// Output:
	// [1 2 3 4 5]
	// [1 2 3 4 5]
}
