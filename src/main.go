package main

import (
	"github.com/edwintantawi/gomark/src/infrastructure/http/router"
	"github.com/edwintantawi/gomark/src/infrastructure/http/server"
)

func main() {
	appRouter := router.NewAppRouter()
	svr := server.NewServer(*appRouter)
	svr.Listen("5000")
}
