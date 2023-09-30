package consumer

import (
	"log"
	"testing"

	kafkaConfig "github.com/Hrit99/zock.git/config"
	"github.com/IBM/sarama/mocks"
)

func Test_Consume(t *testing.T) {
	config := mocks.NewTestConfig()

	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	mockConsumer := mocks.NewConsumer(t, config)
	defer mockConsumer.Close()

	mockpc := mockConsumer.ExpectConsumePartition(kafkaConfig.CONST_TOPIC, int32(1), int64(-2))

	defer mockpc.Close()

	for message := range mockpc.Messages() {
		log.Printf("[test] product_id passed value: %s\n", string(message.Value))
	}
}
