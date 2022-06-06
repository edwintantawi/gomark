package middleware

import "net/http"

type MWFunc func(next http.Handler) http.HandlerFunc
