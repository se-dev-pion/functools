package functools

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	convey.Convey("slice", t, func() {
		arr := []int{1, 2, 3, 4, 5}
		f := func(x int) bool {
			return x%2 == 0
		}
		assert.Equal(t, []int{2, 4}, Filter(f, arr)())
	})
	convey.Convey("string", t, func() {
		seq := "golang"
		f := func(x string) bool {
			return x >= "g"
		}
		assert.Equal(t, "golng", Filter(f, seq)())
		convey.Convey("invalid func", func() {
			f := func(x int) bool {
				return x > 0
			}
			assert.Nil(t, Filter(f, seq))
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
		filtered := Filter(f, ch)()
		assert.Equal(t, []int{2, 4}, extractChanElements(filtered))
		convey.Convey("close chan", func() {
			close(ch)
			assert.Panics(t, func() {
				Any(f, ch)
			})
		})
	})
}

func ExampleFilter() {
	// [slice]
	{
		arr := []int{1, 2, 3, 4, 5}
		f := func(x int) bool {
			return x%2 == 0
		}
		fmt.Println(Filter(f, arr)())
	} // [/]
	// [string]
	{
		seq := "golang"
		f := func(x string) bool {
			return x >= "g"
		}
		fmt.Println(Filter(f, seq)())
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
		filtered := Filter(f, ch)()
		fmt.Println(extractChanElements(filtered))
	} // [/]
	// Output:
	// [2 4]
	// golng
	// [2 4]
}

func BenchmarkFilter(b *testing.B) {
	// [slice]
	{
		arr := []int{1, 2, 3, 4, 5}
		f := func(x int) bool {
			return x%2 == 0
		}
		b.Run("slice", func(b *testing.B) {
			ff := Filter(f, arr)
			for i := 0; i < b.N; i++ {
				ff()
			}
		})
		b.Run("slice raw", func(b *testing.B) {
			ff := Filter4Slice(f, arr)
			for i := 0; i < b.N; i++ {
				ff()
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
			ff := Filter(f, seq)
			for i := 0; i < b.N; i++ {
				ff()
			}
		})
		b.Run("string raw", func(b *testing.B) {
			ff := Filter4String(f, seq)
			for i := 0; i < b.N; i++ {
				ff()
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
			ff := Filter(f, ch)
			for i := 0; i < b.N; i++ {
				ff()
			}
		})
		b.Run("chan raw", func(b *testing.B) {
			ff := Filter4Chan(f, ch)
			for i := 0; i < b.N; i++ {
				ff()
			}
		})
	} // [/]
}
