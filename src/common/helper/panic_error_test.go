package helper

import (
	"github.com/edwintantawi/gomark/src/common/exception"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicError(t *testing.T) {
	t.Run("it should not throw panic when error not exist", func(t *testing.T) {
		assert.NotPanics(t, func() {
			PanicError(nil)
		})
	})

	t.Run("it should throw panic when error exist", func(t *testing.T) {
		err := exception.NewInvariantError("Bad request")

		assert.Panics(t, func() {
			PanicError(err)
		})
	})
}
