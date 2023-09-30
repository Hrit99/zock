package producer

import (
	// Change With Your Package Name
	kafkaConfig "github.com/Hrit99/zock.git/config"

	"log"
	"strconv"

	"github.com/IBM/sarama"
)

func Produce(topic string, product_id int) {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewAsyncProducer([]string{kafkaConfig.CONST_HOST}, config)
	if err != nil {
		log.Fatal("failed to initialize NewSyncProducer, err:", err)
		return
	}
	defer producer.Close()
	msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(strconv.Itoa(product_id))}
	producer.Input() <- msg
	log.Printf("product_id passed to queue: %s\n", strconv.Itoa(product_id))

}
