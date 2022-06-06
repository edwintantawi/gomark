package router

import (
	"github.com/edwintantawi/gomark/src/interface/http/middleware"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func fakeMiddleware() middleware.MWFunc {
	return func(next http.Handler) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
		}
	}
}

func TestNewAppRouter(t *testing.T) {
	t.Run("it should add middleware correctly", func(t *testing.T) {
		r := NewAppRouter()

		assert.Equal(t, 0, len(r.middlewares))

		mw1 := fakeMiddleware()
		mw2 := fakeMiddleware()

		r.Use(mw1)
		r.Use(mw2)

		assert.Equal(t, 2, len(r.middlewares))
	})
}
