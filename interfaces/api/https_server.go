package api

import (
	"github.com/bilfash/trixie/config"
	"github.com/bilfash/trixie/interfaces/api/views/routers"
	"github.com/valyala/fasthttp"
)

type HttpServer struct {
	router routers.Router
	config config.APIConfig
}

func NewHttpServer(router routers.Router, config config.APIConfig) HttpServer {
	return HttpServer{router, config}
}

func (h *HttpServer) ListenAndServe() {
	panic(fasthttp.ListenAndServe(h.config.Port, h.router.RouterHandler().HandleRequest))
}
