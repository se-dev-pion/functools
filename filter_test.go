package functools

import (
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
		nums := make([]int, 0)
		for {
			select {
			case v, ok := <-filtered:
				if !ok {
					goto END
				}
				nums = append(nums, v)
			default:
				goto END
			}
		}
	END:
		assert.Equal(t, []int{2, 4}, nums)
	})
}
