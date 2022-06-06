package v1

import (
	"context"
	"github.com/edwintantawi/gomark/src/application/use_case"
	"github.com/edwintantawi/gomark/src/common/helper"
	"github.com/edwintantawi/gomark/src/domain/response"
	"github.com/edwintantawi/gomark/src/infrastructure/datastore"
	"github.com/edwintantawi/gomark/src/infrastructure/id_gen"
	"github.com/edwintantawi/gomark/src/infrastructure/repository"
	"github.com/edwintantawi/gomark/src/test"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostBookmarkHandler(t *testing.T) {
	db := datastore.NewPostgres()
	idGen := id_gen.NewUuid()
	repo := repository.NewBookmarkRepository(db, idGen)
	useCase := use_case.NewAddBookmarkUseCase(repo)

	t.Run("it should response 201 and persisted bookmark", func(t *testing.T) {
		defer test.CleanTable(db)

		handler := NewPostBookmarkHandler(useCase)

		raw := "{\"title\": \"Golang\",\"description\": \"Amazing docs\",\"url\": \"https://example.com\"}"
		body := ioutil.NopCloser(strings.NewReader(raw))

		tx, _ := db.Begin()
		defer helper.HandleDeferTX(tx)

		req := httptest.NewRequest(http.MethodPost, "/bookmarks", body)
		txCtx := context.WithValue(req.Context(), "dbTx", tx)

		w := httptest.NewRecorder()

		handler.Handle(w, req.WithContext(txCtx), httprouter.Params{})

		res := w.Result()
		defer func(Body io.ReadCloser) {
			closeErr := Body.Close()
			helper.PanicError(closeErr)
		}(res.Body)

		var target response.Body
		helper.BodyParser(res.Body, &target)

		assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, http.StatusText(http.StatusCreated), target.Status)
		assert.NotNil(t, target.Result)
	})
}
