package id_gen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUuidGenerator(t *testing.T) {
	t.Run("it should return an not empty string", func(t *testing.T) {
		idGen := NewUuid()
		r := idGen.Generate()

		assert.NotEmpty(t, r)
	})

	t.Run("it should return string of unique id", func(t *testing.T) {
		idGen := NewUuid()
		r1 := idGen.Generate()
		r2 := idGen.Generate()

		assert.NotEqual(t, r1, r2)
	})
}
