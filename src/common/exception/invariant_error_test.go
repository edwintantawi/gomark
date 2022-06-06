package exception

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestInvariant(t *testing.T) {
	message := "Field is Required"
	err := NewInvariantError(message)

	assert.Equal(t, message, err.Error())
	assert.IsType(t, Error{}, err)
	assert.Equal(t, http.StatusBadRequest, err.Code())
}
