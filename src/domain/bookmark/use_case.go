package bookmark

import (
	"context"
	"database/sql"
)

type AddBookmarkUseCase interface {
	WithTx(txHandle *sql.Tx) AddBookmarkUseCase
	Execute(ctx context.Context, payload New) ID
}

type GetBookmarksUseCase interface {
	WithTx(txHandle *sql.Tx) GetBookmarksUseCase
	Execute(ctx context.Context) []Added
}
