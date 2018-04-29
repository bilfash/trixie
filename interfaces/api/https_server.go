package api

import (
	"net/http"

	"github.com/bilfash/trixie/interfaces/api/views/routers"
)

type HttpServer struct {
	router routers.Router
}

func NewHttpServer(router routers.Router) HttpServer {
	return HttpServer{router}
}

func (h *HttpServer) ListenAndServe() {
	if err := http.ListenAndServe(":8080", h.router.RouterHandler()); err != nil {
		panic(err)
	}
}
