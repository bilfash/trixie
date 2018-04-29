package kafka

import "github.com/Shopify/sarama"

type Producer interface {
	SendMessageToTopic(topic string, message []byte) error
}

type kafkaProducer struct {
	address []string
}

func getProducerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	return config
}

func newSyncProducer(address []string) (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(address, getProducerConfig())
	return producer, err
}

func NewProducer(address []string) Producer {
	return &kafkaProducer{
		address: address,
	}
}

func (kProducer *kafkaProducer) SendMessageToTopic(topic string, message []byte) error {
	kafkaMsg := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(message)}
	producer, err := newSyncProducer(kProducer.address)
	defer producer.Close()
	_, _, err = producer.SendMessage(kafkaMsg)
	return err
}
