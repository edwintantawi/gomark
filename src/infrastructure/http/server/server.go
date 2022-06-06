package server

import (
	"fmt"
	"github.com/edwintantawi/gomark/src/application/use_case"
	"github.com/edwintantawi/gomark/src/common/helper"
	"github.com/edwintantawi/gomark/src/infrastructure/datastore"
	"github.com/edwintantawi/gomark/src/infrastructure/http/router"
	"github.com/edwintantawi/gomark/src/infrastructure/id_gen"
	"github.com/edwintantawi/gomark/src/infrastructure/repository"
	"github.com/edwintantawi/gomark/src/interface/http/middleware"
	v1 "github.com/edwintantawi/gomark/src/interface/http/v1"
	"net/http"
)

type server struct {
	router router.AppRouter
}

func NewServer(router router.AppRouter) *server {
	return &server{router}
}

func (s *server) Listen(port string) {
	p := fmt.Sprintf(":%s", port)

	db := datastore.NewPostgres()
	idGen := id_gen.NewUuid()

	// deps
	bookmarkRepository := repository.NewBookmarkRepository(db, idGen)
	addBookmarkUseCase := use_case.NewAddBookmarkUseCase(bookmarkRepository)

	// handler
	panicHandler := v1.NewPanicHandler()
	postBookmarkHandler := v1.NewPostBookmarkHandler(addBookmarkUseCase)

	// middleware
	requestLoggerMiddleware := middleware.RequestLoggerMiddleware()
	dBTransactionMiddleware := middleware.DBTransactionMiddleware(db)

	s.router.Use(requestLoggerMiddleware)
	s.router.Use(dBTransactionMiddleware)

	// routes
	s.router.POST("/bookmarks", postBookmarkHandler.Handle)

	// exception handler
	s.router.PanicHandler = panicHandler.Handle

	// run server
	err := http.ListenAndServe(p, &s.router)
	helper.PanicError(err)
}
