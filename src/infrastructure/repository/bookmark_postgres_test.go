package repository

import (
	"context"
	"database/sql"
	"github.com/edwintantawi/gomark/src/domain/bookmark"
	"github.com/edwintantawi/gomark/src/infrastructure/datastore"
	"github.com/edwintantawi/gomark/src/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func afterAll(db *sql.DB) {
	test.CleanTable(db)
}

type mockIdGen struct {
	mock.Mock
}

func (m *mockIdGen) Generate() string {
	args := m.Called()
	return args[0].(string)
}

func TestBookmarkRepositoryAdd(t *testing.T) {
	db := datastore.NewPostgres()

	t.Run("it should change DB to Tx and return it self", func(t *testing.T) {
		dummyId := "1234567890"
		idGen := new(mockIdGen)
		idGen.On("Generate").Return(dummyId)
		bookmarkRepo := NewBookmarkRepository(db, idGen)
		txHandle := &sql.Tx{}

		r := bookmarkRepo.WithTx(txHandle)

		assert.Equal(t, r, bookmarkRepo)
		assert.Equal(t, bookmarkRepo.DB, txHandle)
	})

	t.Run("it should save bookmark to db and return correct bookmark id", func(t *testing.T) {
		defer afterAll(db)

		dummyId := "1234567890"
		idGen := new(mockIdGen)
		idGen.On("Generate").Return(dummyId)
		bookmarkRepo := NewBookmarkRepository(db, idGen)

		newBookmark := bookmark.New{
			Title:       "Go documentation",
			Description: "documentation web pages for GO",
			Url:         "https://go.dev",
		}

		ctx := context.Background()
		r := bookmarkRepo.Add(ctx, newBookmark)

		assert.Equal(t, bookmark.ID(dummyId), r)
	})

}
