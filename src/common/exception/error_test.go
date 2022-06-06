package exception

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError(t *testing.T) {
	t.Run("it should create error correctly", func(t *testing.T) {
		message := "error message"
		code := 500
		err := Error{message: message, code: code}

		assert.Equal(t, message, err.Error())
		assert.Equal(t, code, err.Code())
	})
}
