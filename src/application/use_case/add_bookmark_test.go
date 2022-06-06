package use_case

import (
	"context"
	"database/sql"
	"github.com/edwintantawi/gomark/src/domain/bookmark"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockBookmarkRepository struct {
	mock.Mock
}

func (m *mockBookmarkRepository) Add(ctx context.Context, newBookmark bookmark.New) bookmark.ID {
	args := m.Called(ctx, newBookmark)
	return bookmark.ID(args[0].(string))
}

func (m *mockBookmarkRepository) WithTx(txHandle *sql.Tx) bookmark.Repository {
	args := m.Called(txHandle)
	return args[0].(bookmark.Repository)
}

func TestAddBookmark(t *testing.T) {
	t.Run("it should call repository WithTx and return it self", func(t *testing.T) {

		bookmarkRepository := new(mockBookmarkRepository)

		bookmarkRepository.On("WithTx", mock.Anything).Return(bookmarkRepository)

		useCase := NewAddBookmarkUseCase(bookmarkRepository)
		r := useCase.WithTx(&sql.Tx{})

		assert.Equal(t, r, useCase) // should return it self
		bookmarkRepository.AssertNumberOfCalls(t, "WithTx", 1)
	})

	t.Run("it should call repository and return bookmark id", func(t *testing.T) {
		ctx := context.Background()
		bookmarkRepository := new(mockBookmarkRepository)

		newBookmark := bookmark.New{
			Title:       "GoMark",
			Description: "an description",
			Url:         "https://example.com",
		}

		bookmarkId := "unique_bookmark_id"

		bookmarkRepository.On("Add", ctx, newBookmark).Return(bookmarkId)

		useCase := NewAddBookmarkUseCase(bookmarkRepository)
		result := useCase.Execute(ctx, newBookmark)

		assert.Equal(t, bookmark.ID(bookmarkId), result)
		bookmarkRepository.AssertNumberOfCalls(t, "Add", 1)
	})

}
