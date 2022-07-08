package use_case

import (
	"context"
	"database/sql"
	"github.com/edwintantawi/gomark/src/domain/bookmark"
)

type addBookmarkUseCase struct {
	bookmarkRepository bookmark.Repository
}

func NewAddBookmarkUseCase(bookmarkRepository bookmark.Repository) *addBookmarkUseCase {
	return &addBookmarkUseCase{
		bookmarkRepository: bookmarkRepository,
	}
}

func (u *addBookmarkUseCase) WithTx(txHandle *sql.Tx) bookmark.AddBookmarkUseCase {
	u.bookmarkRepository.WithTx(txHandle)
	return u
}

func (u *addBookmarkUseCase) Execute(ctx context.Context, payload bookmark.New) bookmark.ID {
	return u.bookmarkRepository.Add(ctx, payload)
}
