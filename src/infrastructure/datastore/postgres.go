package datastore

import (
	"database/sql"
	"github.com/edwintantawi/gomark/src/common/helper"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func NewPostgres() *sql.DB {
	dsn := os.Getenv("DSN_POSTGRES")
	log.Printf("[Presistent:Postgres] Open with DSN: %s\n", dsn)

	db, err := sql.Open("postgres", dsn)
	helper.PanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(15 * time.Minute)

	return db
}
