package producer

import (
	"math/rand"
	"strconv"
	"testing"

	kafkaConfig "github.com/Hrit99/zock.git/config"
	"github.com/IBM/sarama"
	"github.com/IBM/sarama/mocks"
)

func Test_Produce(t *testing.T) {
	config := mocks.NewTestConfig()

	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	mockproducer := mocks.NewAsyncProducer(t, config)
	defer mockproducer.Close()

	mockproducer.ExpectInputAndSucceed()

	msg := &sarama.ProducerMessage{Topic: kafkaConfig.CONST_TOPIC, Key: nil, Value: sarama.StringEncoder(strconv.Itoa(rand.Intn(500)))}
	mockproducer.Input() <- msg
}
