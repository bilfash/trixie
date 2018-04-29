package main

import (
	"github.com/bilfash/trixie/config"
	"github.com/bilfash/trixie/infrastructures/kafka"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	env = kingpin.Flag("env", "Environment").Short('e').Required().String()
)

func main() {
	kingpin.Parse()
	configuration := config.ConfigGenerator(*env)

	kafka.StartConsumer(*configuration)
}
