package mocks

import (
	"context"
	"database/sql"
	"github.com/edwintantawi/gomark/src/domain/bookmark"
	"github.com/stretchr/testify/mock"
)

type MockBookmarkRepository struct {
	mock.Mock
}

func (m *MockBookmarkRepository) WithTx(txHandle *sql.Tx) bookmark.Repository {
	args := m.Called(txHandle)
	return args[0].(bookmark.Repository)
}

func (m *MockBookmarkRepository) Add(ctx context.Context, newBookmark bookmark.New) bookmark.ID {
	args := m.Called(ctx, newBookmark)
	return bookmark.ID(args[0].(string))
}

func (m *MockBookmarkRepository) GetAll(ctx context.Context) []bookmark.Added {
	args := m.Called(ctx)
	return args[0].([]bookmark.Added)
}
