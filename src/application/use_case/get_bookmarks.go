package use_case

import (
	"context"
	"database/sql"
	"github.com/edwintantawi/gomark/src/domain/bookmark"
)

type getBookmarksUseCase struct {
	bookmarkRepository bookmark.Repository
}

func NewGetBookmarksUseCase(bookmarkRepository bookmark.Repository) *getBookmarksUseCase {
	return &getBookmarksUseCase{
		bookmarkRepository: bookmarkRepository,
	}
}

func (u *getBookmarksUseCase) WithTx(txHandle *sql.Tx) bookmark.GetBookmarksUseCase {
	u.bookmarkRepository.WithTx(txHandle)
	return u
}

func (u *getBookmarksUseCase) Execute(ctx context.Context) []bookmark.Added {
	return u.bookmarkRepository.GetAll(ctx)
}
