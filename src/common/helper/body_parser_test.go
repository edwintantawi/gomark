package helper

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestBodyParser(t *testing.T) {
	t.Run("it should panic when request body is invalid", func(t *testing.T) {
		raw := "an body"
		body := ioutil.NopCloser(strings.NewReader(raw))

		var target string

		assert.PanicsWithError(t, "Invalid request body", func() {
			BodyParser(body, &target)
		})

		assert.Equal(t, "", target)
	})

	t.Run("it should parse body correctly", func(t *testing.T) {
		raw := "{\"ping\":\"pong\"}"
		body := ioutil.NopCloser(strings.NewReader(raw))

		type targetStruct struct {
			Ping string `json:"ping"`
		}

		actual := targetStruct{Ping: "pong"}
		var target targetStruct

		BodyParser(body, &target)

		assert.Equal(t, actual, target)
	})
}
