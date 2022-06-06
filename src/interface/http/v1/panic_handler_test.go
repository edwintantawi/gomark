package v1

import (
	"errors"
	"github.com/edwintantawi/gomark/src/common/exception"
	"github.com/edwintantawi/gomark/src/common/helper"
	"github.com/edwintantawi/gomark/src/domain/response"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func wrapPanicHandler(error error) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		h := NewPanicHandler()
		h.Handle(w, r, error)
	}
}

func TestPanicHandler(t *testing.T) {
	t.Run("it should response correctly if error is exception error", func(t *testing.T) {
		errMsg := "An exception invariant error"
		err := exception.NewInvariantError(errMsg)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		w := httptest.NewRecorder()

		handler := wrapPanicHandler(err)

		handler(w, req, httprouter.Params{})

		res := w.Result()
		defer func(Body io.ReadCloser) {
			closeErr := Body.Close()
			helper.PanicError(closeErr)
		}(res.Body)

		var target response.Body
		helper.BodyParser(res.Body, &target)

		assert.Equal(t, res.Header.Get("Content-Type"), "application/json")
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, http.StatusText(http.StatusBadRequest), target.Status)
		assert.Equal(t, errMsg, target.Message)
		assert.Nil(t, target.Result)
	})

	t.Run("it should response correctly if error is unexpected error", func(t *testing.T) {
		errMsg := "an other unexpected error"
		err := errors.New(errMsg)

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		w := httptest.NewRecorder()

		handler := wrapPanicHandler(err)

		handler(w, req, httprouter.Params{})

		res := w.Result()
		defer func(Body io.ReadCloser) {
			closeErr := Body.Close()
			helper.PanicError(closeErr)
		}(res.Body)

		var target response.Body
		helper.BodyParser(res.Body, &target)

		assert.Equal(t, res.Header.Get("Content-Type"), "application/json")
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Equal(t, http.StatusText(http.StatusInternalServerError), target.Status)
		assert.Equal(t, errMsg, target.Message)
		assert.Nil(t, target.Result)
	})
}
