package functools

import (
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

func TestPack(t *testing.T) {
	convey.Convey("pack", t, func() {
		assert.Equal(t, []int{1, 2, 3}, Pack(1, 2, 3))
		assert.Equal(t, []string{"a", "b", "c"}, Pack("a", "b", "c"))
	})
}

func TestLazy(t *testing.T) {
	convey.Convey("lazy", t, func() {
		f := func(x int) int {
			return x * 2
		}
		assert.Equal(t, f(2), Lazy(f, 2)())
	})
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

func TestPipe(t *testing.T) {
	convey.Convey("pipe", t, func() {

	})
}

func TestBatch(t *testing.T) {
	convey.Convey("batch", t, func() {

	})
}

func TestCached(t *testing.T) {
	convey.Convey("cache", t, func() {

	})
}

func TestCopy(t *testing.T) {
	convey.Convey("copy", t, func() {

	})
}
