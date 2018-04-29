package main

import (
	"github.com/bilfash/trixie/config"
	"github.com/bilfash/trixie/infrastructures/kafka"
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

	producer := kafka.NewTopicProducer(configuration.KafkaProducerConfig.Topic.ClientCreateRequestedTopic,
		[]string{configuration.KafkaProducerConfig.Address})
	router := routers.NewRouter(producer)

	httpServer := api.NewHttpServer(router, configuration.ApiConfig)
	httpServer.ListenAndServe()
}
