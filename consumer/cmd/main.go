package main

import (
	// Change With Your Package Name
	kafkaConfig "github.com/Hrit99/zock.git/config"
	"github.com/Hrit99/zock.git/consumer"
)

func main() {
	topic := kafkaConfig.CONST_TOPIC
	consumer.Consume(topic)
}
