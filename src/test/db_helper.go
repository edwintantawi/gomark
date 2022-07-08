package test

import (
	"database/sql"
	"fmt"
	"github.com/edwintantawi/gomark/src/common/helper"
)

func CleanTable(db *sql.DB) {
	_, err := db.Query("TRUNCATE TABLE bookmarks")
	helper.PanicError(err)
}

func AddManyBookmarks(db *sql.DB) {
	_, err := db.Exec(
		"INSERT INTO bookmarks VALUES " +
			fmt.Sprintf("('id001', 'Bookmark 1', 'Bookmark description 1', 'https://bookmark-1'),") +
			fmt.Sprintf("('id002', 'Bookmark 2', 'Bookmark description 2', 'https://bookmark-2'),") +
			fmt.Sprintf("('id003', 'Bookmark 3', 'Bookmark description 3', 'https://bookmark-3')"))

	helper.PanicError(err)
}
