package main

import (
	"github.com/bilfash/trixie/config"
	"github.com/bilfash/trixie/infrastructures/fasthttp"
	"github.com/bilfash/trixie/infrastructures/kafka"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	env = kingpin.Flag("env", "Environment").Short('e').Required().String()
)

func main() {
	kingpin.Parse()
	configuration := config.ConfigGenerator(*env)

	producer := kafka.NewProducer([]string{configuration.KafkaProducerConfig.Address})
	router := fasthttp.NewRouter(producer, *configuration)
	httpServer := fasthttp.NewHttpServer(router, *configuration)
	httpServer.ListenAndServe()
}
