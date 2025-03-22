package functools

import (
	"strings"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	convey.Convey("slice", t, func() {
		arr := []int{1, 2, 3, 4, 5}
		f := func(x int) float64 {
			return float64(x*2) / 10
		}
		assert.Equal(t,
			[]float64{0.2, 0.4, 0.6, 0.8, 1},
			Map[[]float64](f, arr)(),
		)
	})
	convey.Convey("string", t, func() {
		seq := "golang"
		assert.Equal(t, "GOLANG", Map[string](strings.ToUpper, seq)())
		convey.Convey("invalid func", func() {
			f1 := func(x int) bool {
				return x > 0
			}
			assert.Nil(t, Map[string](f1, seq))
			f2 := func(x string) int {
				return len(x)
			}
			assert.Nil(t, Map[string](f2, seq))
		})
	})
	convey.Convey("chan", t, func() {
		ch := make(chan int, 10)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		f := func(x int) float64 {
			return float64(x*2) / 10
		}
		mapped := Map[chan float64](f, ch)()
		assert.Equal(t, []float64{0.2, 0.4, 0.6, 0.8, 1}, extractChanElements(mapped))
		convey.Convey("close chan", func() {
			close(ch)
			assert.Panics(t, func() {
				Map[chan float64](f, ch)()
			})
		})
	})
}
