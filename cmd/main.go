package main

import (
	"github.com/bilfash/trixie/config"
	"github.com/bilfash/trixie/interfaces/api"
	"github.com/bilfash/trixie/interfaces/api/views/routers"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	env = kingpin.Flag("env", "Environment").Short('e').Required().String()
)

func main() {
	kingpin.Parse()
	configuration := config.ConfigGenerator(*env)

	router := routers.NewRouter()

	httpServer := api.NewHttpServer(router, configuration.ApiConfig)
	httpServer.ListenAndServe()
}
