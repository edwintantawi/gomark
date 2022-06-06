package bookmark

import (
	"context"
	"database/sql"
)

type UseCase interface {
	WithTx(txHandle *sql.Tx) UseCase
	Execute(ctx context.Context, payload New) ID
}
