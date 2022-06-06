package repository

import (
	"context"
	"database/sql"
	"github.com/edwintantawi/gomark/src/common/helper"
	"github.com/edwintantawi/gomark/src/domain/bookmark"
	"github.com/edwintantawi/gomark/src/infrastructure/datastore"
	"github.com/edwintantawi/gomark/src/infrastructure/id_gen"
)

type bookmarkRepository struct {
	DB    datastore.DBConn
	idGen id_gen.IdGenerator
}

func NewBookmarkRepository(DB *sql.DB, idGen id_gen.IdGenerator) *bookmarkRepository {
	return &bookmarkRepository{DB: DB, idGen: idGen}
}

func (r *bookmarkRepository) WithTx(txHandle *sql.Tx) bookmark.Repository {
	r.DB = txHandle
	return r
}

func (r *bookmarkRepository) Add(ctx context.Context, newBookmark bookmark.New) bookmark.ID {
	var bookmarkId bookmark.ID
	id := r.idGen.Generate()

	SQL := "INSERT INTO bookmarks (id, title, description, url) VALUES ($1, $2, $3, $4) RETURNING id"
	rows := r.DB.QueryRowContext(ctx, SQL, id, newBookmark.Title, newBookmark.Description, newBookmark.Url)

	err := rows.Scan(&bookmarkId)
	helper.PanicError(err)

	return bookmarkId
}
