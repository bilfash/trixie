package kafka

import "github.com/Shopify/sarama"

type TopicProducer interface {
	SendMessage(message []byte) error
}

type kafkaProducer struct {
	address []string
	topic   string
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

func NewTopicProducer(topic string, address []string) TopicProducer {
	return &kafkaProducer{
		address: address,
		topic:   topic,
	}
}

func (kProducer *kafkaProducer) SendMessage(message []byte) error {
	kafkaMsg := &sarama.ProducerMessage{Topic: kProducer.topic, Value: sarama.StringEncoder(message)}
	producer, err := newSyncProducer(kProducer.address)
	defer producer.Close()
	_, _, err = producer.SendMessage(kafkaMsg)
	return err
}
