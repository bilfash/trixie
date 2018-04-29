package api

import (
	"github.com/bilfash/trixie/interfaces/api/views/routers"
	"github.com/valyala/fasthttp"
)

type HttpServer struct {
	router routers.Router
}

func NewHttpServer(router routers.Router) HttpServer {
	return HttpServer{router}
}

func (h *HttpServer) ListenAndServe() {
	panic(fasthttp.ListenAndServe(":8080", h.router.RouterHandler().HandleRequest))
}
