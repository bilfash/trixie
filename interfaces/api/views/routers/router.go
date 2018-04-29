package routers

import (
	"github.com/bilfash/trixie/interfaces/api/controllers"
	"github.com/husobee/vestigo"
)

type Router struct {
	router           *vestigo.Router
	clientController controllers.ClientController
}

func NewRouter() Router {
	return Router{
		vestigo.NewRouter(),
		controllers.NewClientController(),
	}
}

func (r *Router) RouterHandler() *vestigo.Router {
	r.initiateRouter()
	return r.router
}

func (r *Router) initiateRouter() {
	r.router.Post("/client", r.clientController.ClientApiPostHandler)
}
