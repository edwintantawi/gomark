package middleware

import (
	"context"
	"database/sql"
	"github.com/edwintantawi/gomark/src/common/helper"
	"net/http"
)

func DBTransactionMiddleware(db *sql.DB) MWFunc {
	return func(next http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tx, err := db.Begin()
			helper.PanicError(err)
			defer helper.HandleDeferTX(tx)

			txCtx := context.WithValue(r.Context(), "dbTx", tx)

			next.ServeHTTP(w, r.WithContext(txCtx))
		}
	}
}
