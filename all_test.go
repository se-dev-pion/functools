package functools

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	convey.Convey("slice", t, func() {
		var err error
		assert.Nil(t, err)
	})
	convey.Convey("string", t, func() {
		var err error
		assert.Nil(t, err)
	})
	convey.Convey("chan", t, func() {
		var err error
		assert.Nil(t, err)
	})
}
