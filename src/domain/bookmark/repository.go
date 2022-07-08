package bookmark

import (
	"context"
	"database/sql"
)

type Repository interface {
	WithTx(txHandle *sql.Tx) Repository
	Add(ctx context.Context, newBookmark New) ID
	GetAll(ctx context.Context) []Added
}
