package functools

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	convey.Convey("slice", t, func() {
		arr := []int{1, 2, 3, 4, 5}
		f := func(x int) bool {
			return x%2 == 0
		}
		assert.True(t, Any(f, arr))
		assert.False(t, Any(f, []int{}))
	})
	convey.Convey("string", t, func() {
		seq := "golang"
		f := func(x string) bool {
			return x >= "g"
		}
		assert.True(t, Any(f, seq))
		assert.False(t, Any(f, ""))
		convey.Convey("invalid func", func() {
			f := func(x int) bool {
				return x > 0
			}
			assert.False(t, Any(f, seq))
			assert.False(t, Any(f, ""))
		})
	})
	convey.Convey("chan", t, func() {
		ch := make(chan int, 10)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		f := func(x int) bool {
			return x%2 == 0
		}
		assert.True(t, Any(f, ch))
		assert.False(t, Any(f, make(chan int, 1)))
		convey.Convey("close chan", func() {
			close(ch)
			assert.Panics(t, func() {
				Any(f, ch)
			})
		})
	})
}

func ExampleAny() {
	// [slice]
	{
		arr := []int{1, 2, 3, 4, 5}
		f := func(x int) bool {
			return x%2 == 0
		}
		fmt.Println(Any(f, arr))
		fmt.Println(Any(f, []int{}))
	} // [/]
	// [string]
	{
		seq := "golang"
		f := func(x string) bool {
			return x >= "g"
		}
		fmt.Println(Any(f, seq))
		fmt.Println(Any(f, ""))
	} // [/]
	// [chan]
	{
		ch := make(chan int, 10)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		f := func(x int) bool {
			return x%2 == 0
		}
		fmt.Println(Any(f, ch))
		fmt.Println(Any(f, make(chan int, 1)))
	} // [/]
	// Output:
	// true
	// false
	// true
	// false
	// true
	// false
}

func BenchmarkAny(b *testing.B) {
	// [slice]
	{
		arr := []int{1, 2, 3, 4, 5}
		f := func(x int) bool {
			return x%2 == 0
		}
		b.Run("slice", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Any(f, arr)
			}
		})
		b.Run("slice raw", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Any4Slice(f, arr)
			}
		})
	} // [/]
	// [string]
	{
		seq := "golang"
		f := func(x string) bool {
			return x >= "g"
		}
		b.Run("string", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Any(f, seq)
			}
		})
		b.Run("string raw", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Any4String(f, seq)
			}
		})
	} // [/]
	// [chan]
	{
		ch := make(chan int, 10)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		f := func(x int) bool {
			return x%2 == 0
		}
		b.Run("chan", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Any(f, ch)
			}
		})
		b.Run("chan raw", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Any4Chan(f, ch)
			}
		})
	} // [/]
}
