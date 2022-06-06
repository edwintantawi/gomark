package router

import (
	"github.com/edwintantawi/gomark/src/interface/http/middleware"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AppRouter struct {
	httprouter.Router
	middlewares []middleware.MWFunc
}

func NewAppRouter() *AppRouter {
	r := httprouter.New()
	var mw []middleware.MWFunc
	return &AppRouter{*r, mw}
}

func (ar *AppRouter) Use(middleware middleware.MWFunc) {
	ar.middlewares = append(ar.middlewares, middleware)
}

func (ar *AppRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isEmptyMiddleware := len(ar.middlewares) == 0

	if isEmptyMiddleware {
		ar.Router.ServeHTTP(w, r)
		return
	}

	// loop and chain middleware
	var curr http.Handler
	for _, mw := range ar.middlewares {
		if curr != nil {
			curr = mw(curr)
		} else {
			curr = mw(&ar.Router)
		}
	}

	ar.injectMiddleware(curr)(w, r)
}

func (ar *AppRouter) injectMiddleware(mw http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mw.ServeHTTP(w, r)
	}
}
