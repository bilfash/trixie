package fasthttp

import (
	"github.com/bilfash/trixie/config"
	"github.com/bilfash/trixie/interfaces/api/handlers"
	"github.com/qiangxue/fasthttp-routing"
)

type Router struct {
	router        *routing.Router
	clientHandler handlers.ClientHandler
}

func NewRouter(kafkaProducer handlers.IClientHandler, config config.Configuration) Router {
	return Router{
		routing.New(),
		handlers.NewClientHandler(config, kafkaProducer),
	}
}

func (r *Router) RouterHandler() *routing.Router {
	r.initiateRouter()
	return r.router
}

func (r *Router) initiateRouter() {
	r.router.Post("/client", r.clientHandler.ClientApiPostHandler)
}
