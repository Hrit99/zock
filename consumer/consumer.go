package consumer

import (
	// Change With Your Package Name
	kafkaConfig "github.com/Hrit99/zock.git/config"
	"github.com/IBM/sarama"

	"log"
)

func Consume(topic string) {
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer([]string{kafkaConfig.CONST_HOST}, config)
	if err != nil {
		log.Fatal("NewConsumer err: ", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		log.Printf("product_id passed value: %s\n", string(message.Value))

	}
}
