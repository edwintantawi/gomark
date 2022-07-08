package use_case

import (
	"context"
	"database/sql"
	"github.com/edwintantawi/gomark/src/domain/bookmark"
	"github.com/edwintantawi/gomark/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestGetAllBookmarks(t *testing.T) {
	t.Run("it should call repository WithTx and return it self", func(t *testing.T) {
		bookmarkRepository := new(mocks.MockBookmarkRepository)
		bookmarkRepository.On("WithTx", mock.Anything).Return(bookmarkRepository)

		useCase := NewGetBookmarksUseCase(bookmarkRepository)
		r := useCase.WithTx(&sql.Tx{})

		assert.Equal(t, r, useCase) // should return it self
		bookmarkRepository.AssertNumberOfCalls(t, "WithTx", 1)
	})

	t.Run("it should call repository and return bookmarks", func(t *testing.T) {
		expectedBookmarks := []bookmark.Added{
			{
				ID:          "001",
				Title:       "Bookmark 01",
				Description: "Desc for bookmark 01",
				Url:         "https://bookmark-01",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				ID:          "001",
				Title:       "Bookmark 01",
				Description: "Desc for bookmark 01",
				Url:         "https://bookmark-01",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}

		ctx := context.Background()

		bookmarkRepository := new(mocks.MockBookmarkRepository)
		bookmarkRepository.On("GetAll", ctx).Return(expectedBookmarks)

		useCase := NewGetBookmarksUseCase(bookmarkRepository)

		result := useCase.Execute(ctx)

		assert.Equal(t, len(expectedBookmarks), len(result))
		assert.ObjectsAreEqualValues(expectedBookmarks, result)
	})
}
