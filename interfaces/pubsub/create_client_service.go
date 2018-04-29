package pubsub

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bilfash/trixie/config"
	"github.com/bsm/sarama-cluster"
	"os"
	"os/signal"
)

type CreateClientService struct {
	configuration config.Configuration
}

func NewCreateClientService(configuration config.Configuration) CreateClientService {
	return CreateClientService{configuration}
}

func (ccs *CreateClientService) Start() error {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	topics := []string{
		ccs.configuration.KafkaProducerConfig.Topic.ClientCreateRequestedTopic,
	}
	address := []string{
		ccs.configuration.KafkaProducerConfig.Address,
	}
	consumer, err := cluster.NewConsumer(address, "logger", topics, config)

	if err != nil {
		return err
	}

	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	fmt.Println("consumer is now online.")

	go func() {
		for error := range consumer.Errors() {
			fmt.Println(error)
		}
	}()

	for {
		select {
		case msg, ok := <-consumer.Messages():
			fmt.Println("message consumed")
			if ok {
				eventData, err := ccs.parseRequest(msg.Value)
				if err != nil {
					fmt.Println("parse message failed")
				}
				fmt.Println("event data", eventData.ReqId, eventData.Client)
				fmt.Printf("event received : %s \n", msg.Topic)
				consumer.MarkOffset(msg, "complete")
			} else {
				fmt.Println("cluster not ok ", ok)
			}
		case <-signals:
			fmt.Println("consumer terminated")
			return nil
		}
	}
}

func (ccs *CreateClientService) parseRequest(rawData []byte) (*ClientMessage, error) {
	eventData := &ClientMessage{}
	err := json.Unmarshal(rawData, eventData)
	return eventData, err
}
