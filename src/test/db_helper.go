package test

import (
	"database/sql"
	"github.com/edwintantawi/gomark/src/common/helper"
)

func CleanTable(db *sql.DB) {
	_, err := db.Query("TRUNCATE TABLE bookmarks")
	helper.PanicError(err)
}
