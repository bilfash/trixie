package config

import (
	"github.com/spf13/viper"
	"log"
)

func ConfigGenerator(env string) *Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.UnmarshalKey(env, &configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &configuration
}
