package functools

import (
	"fmt"
	"strings"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	convey.Convey("slice", t, func() {
		arr := []int{1, 2, 3, 4, 5}
		f := func(a, b int) int {
			return a + b
		}
		total := 1 + 2 + 3 + 4 + 5
		assert.Equal(t, total, Reduce(f, arr)())
		convey.Convey("initial", func() {
			assert.Equal(t, total+6, Reduce(f, arr, 6)())
		})
		convey.Convey("no initial", func() {
			assert.Nil(t, Reduce(f, []int{}))
		})
	})
	convey.Convey("string", t, func() {
		seq := "golang"
		f := func(a, b string) string {
			return strings.ToUpper(a) + strings.ToUpper(b)
		}
		total := "GOLANG"
		assert.Equal(t, total, Reduce(f, seq)())
		convey.Convey("initial", func() {
			assert.Equal(t, "GOOGLE"+total, Reduce(f, seq, "GOOGLE")())
		})
		convey.Convey("no initial", func() {
			assert.Nil(t, Reduce(f, ""))
		})
		convey.Convey("invalid func", func() {
			f := func(a, b int) int {
				return a + b
			}
			assert.Nil(t, Reduce(f, seq))
		})
	})
	convey.Convey("chan", t, func() {
		ch := make(chan int, 10)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		f := func(a, b int) int {
			return a + b
		}
		total := 1 + 2 + 3 + 4 + 5
		assert.Equal(t, total, Reduce(f, ch)())
		convey.Convey("initial", func() {
			assert.Equal(t, total+6, Reduce(f, ch, 6)())
		})
	})
}

func ExampleReduce() {
	// [slice]
	{
		arr := []int{1, 2, 3, 4, 5}
		f := func(a, b int) int {
			return a + b
		}
		fmt.Println(Reduce(f, arr)())
		fmt.Println(Reduce(f, arr, 6)())
	} // [/]
	// [string]
	{
		seq := "golang"
		f := func(a, b string) string {
			return strings.ToUpper(a) + strings.ToUpper(b)
		}
		fmt.Println(Reduce(f, seq)())
		fmt.Println(Reduce(f, seq, "GOOGLE")())
	} // [/]
	// [chan]
	{
		ch := make(chan int, 10)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		f := func(a, b int) int {
			return a + b
		}
		fmt.Println(Reduce(f, ch)())
		fmt.Println(Reduce(f, ch, 6)())
	} // [/]
	// Output:
	// 15
	// 21
	// GOLANG
	// GOOGLEGOLANG
	// 15
	// 21
}
