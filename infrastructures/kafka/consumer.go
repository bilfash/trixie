package kafka

import (
	"fmt"
	"github.com/bilfash/trixie/config"
	"github.com/bilfash/trixie/interfaces/pubsub"
)

func StartConsumer(configuration config.Configuration) {
	createClientService := pubsub.NewCreateClientService(configuration)
	createClientService.Start()
	fmt.Println("OKEEEEEEE")
}
