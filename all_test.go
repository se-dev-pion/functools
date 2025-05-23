package functools

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	convey.Convey("slice", t, func() {
		arr := []int{1, 2, 3, 4, 5}
		f := func(x int) bool {
			return x%2 == 0
		}
		assert.False(t, All(f, arr))
		assert.True(t, All(f, []int{}))
	})
	convey.Convey("string", t, func() {
		seq := "golang"
		f := func(x string) bool {
			return x >= "g"
		}
		assert.False(t, All(f, seq))
		assert.True(t, All(f, ""))
		convey.Convey("invalid func", func() {
			f := func(x int) bool {
				return x > 0
			}
			assert.False(t, All(f, seq))
			assert.False(t, All(f, ""))
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
		assert.False(t, All(f, ch))
		assert.True(t, All(f, make(chan int, 1)))
		convey.Convey("close chan", func() {
			close(ch)
			assert.Panics(t, func() {
				All(f, ch)
			})
		})
	})
}

func ExampleAll() {
	// [slice]
	{
		arr := []int{1, 2, 3, 4, 5}
		f := func(x int) bool {
			return x%2 == 0
		}
		fmt.Println(All(f, arr))
		fmt.Println(All(f, []int{}))
	} // [/]
	// [string]
	{
		seq := "golang"
		f := func(x string) bool {
			return x >= "g"
		}
		fmt.Println(All(f, seq))
		fmt.Println(All(f, ""))
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
		fmt.Println(All(f, ch))
		fmt.Println(All(f, make(chan int, 1)))
	} // [/]
	// Output:
	// false
	// true
	// false
	// true
	// false
	// true
}

func BenchmarkAll(b *testing.B) {
	// [slice]
	{
		arr := []int{1, 2, 3, 4, 5}
		f1 := func(x int) bool {
			return x%2 == 0
		}
		b.Run("slice", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				All(f1, arr)
			}
		})
		b.Run("slice raw", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				All4Slice(f1, arr)
			}
		})
	} // [/]
	// [string]
	{
		seq := "golang"
		f2 := func(x string) bool {
			return x >= "g"
		}
		b.Run("string", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				All(f2, seq)
			}
		})
		b.Run("string raw", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				All4String(f2, seq)
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
				All(f, ch)
			}
		})
		b.Run("chan raw", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				All4Chan(f, ch)
			}
		})
	} // [/]
}
