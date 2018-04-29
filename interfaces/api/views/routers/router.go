package routers

import (
	"github.com/bilfash/trixie/interfaces/api/controllers"
	"github.com/qiangxue/fasthttp-routing"
)

type Router struct {
	router           *routing.Router
	clientController controllers.ClientController
}

func NewRouter() Router {
	return Router{
		routing.New(),
		controllers.NewClientController(),
	}
}

func (r *Router) RouterHandler() *routing.Router {
	r.initiateRouter()
	return r.router
}

func (r *Router) initiateRouter() {
	r.router.Post("/client", r.clientController.ClientApiPostHandler)
}
