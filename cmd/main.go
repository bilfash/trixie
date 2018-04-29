package main

import (
	"github.com/bilfash/trixie/interfaces/api"
	"github.com/bilfash/trixie/interfaces/api/views/routers"
)

func main() {
	router := routers.NewRouter()

	httpServer := api.NewHttpServer(router)
	httpServer.ListenAndServe()
}
