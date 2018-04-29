package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bilfash/trixie/config"
	"github.com/bilfash/trixie/interfaces/pubsub"
	"github.com/bsm/sarama-cluster"
	"os"
	"os/signal"
)

func StartConsumer(configuration config.Configuration) {
	createClientService := pubsub.NewCreateClientService()

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	topics := []string{
		configuration.KafkaProducerConfig.Topic.ClientCreateRequestedTopic,
	}
	address := []string{
		configuration.KafkaProducerConfig.Address,
	}
	consumer, err := cluster.NewConsumer(address, "logger", topics, config)

	if err != nil {
		return
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
				eventData, err := parseRequest(msg.Value)
				if err != nil {
					fmt.Println("parse message failed")
				}
				createClientService.Execute(*eventData)
				fmt.Printf("event received : %s \n", msg.Topic)
				consumer.MarkOffset(msg, "complete")
			} else {
				fmt.Println("cluster not ok ", ok)
			}
		case <-signals:
			fmt.Println("consumer terminated")
			return
		}
	}
}

func parseRequest(rawData []byte) (*pubsub.ClientMessage, error) {
	eventData := &pubsub.ClientMessage{}
	err := json.Unmarshal(rawData, eventData)
	return eventData, err
}
