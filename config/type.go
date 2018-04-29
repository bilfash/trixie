package config

type Configuration struct {
	ApiConfig           APIConfig
	KafkaProducerConfig KafkaProducerConfig
}

type APIConfig struct {
	Port string
}

type KafkaProducerConfig struct {
	Address string
	Topic   Topic
}

type Topic struct {
	ClientCreateRequestedTopic string
}
