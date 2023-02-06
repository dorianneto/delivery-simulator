package kafka

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	Message chan *ckafka.Message
}

func NewKafkcaConsumer(message chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{Message: message}
}

func (k *KafkaConsumer) Consume() {
	consumer, err := ckafka.NewConsumer(&ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	})

	if err != nil {
		log.Fatalln("error consuming kafka message:" + err.Error())
		return
	}

	topics := []string{os.Getenv("KafkaReadTopic")}
	consumer.SubscribeTopics(topics, nil)

	log.Println("kafka consumer has been started!")

	for {
		data, err := consumer.ReadMessage(-1)
		if err == nil {
			k.Message <- data
		}
	}
}
