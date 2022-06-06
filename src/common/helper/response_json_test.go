package helper

import (
	"github.com/edwintantawi/gomark/src/domain/response"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func wrapFakeHandler(statusCode int, body response.Body) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ResponseJson(w, statusCode, body)
	}
}

func TestResponseJson(t *testing.T) {
	t.Run("it should write response correctly", func(t *testing.T) {
		code := http.StatusOK
		msg := "response message ok"
		body := response.Body{
			Status:  http.StatusText(code),
			Message: msg,
			Result:  nil,
		}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler := wrapFakeHandler(code, body)

		handler(w, req, httprouter.Params{})

		res := w.Result()
		defer func(Body io.ReadCloser) {
			closeErr := Body.Close()
			PanicError(closeErr)
		}(res.Body)

		var target response.Body
		BodyParser(res.Body, &target)

		assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
		assert.Equal(t, code, res.StatusCode)
		assert.Equal(t, http.StatusText(code), target.Status)
		assert.Equal(t, msg, target.Message)
		assert.Nil(t, target.Result)
	})
}
