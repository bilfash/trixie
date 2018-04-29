package fasthttp

import (
	"github.com/bilfash/trixie/config"
	"github.com/valyala/fasthttp"
)

type HttpServer struct {
	router Router
	config config.Configuration
}

func NewHttpServer(router Router, config config.Configuration) HttpServer {
	return HttpServer{router, config}
}

func (h *HttpServer) ListenAndServe() {
	panic(fasthttp.ListenAndServe(h.config.ApiConfig.Port, h.router.RouterHandler().HandleRequest))
}
