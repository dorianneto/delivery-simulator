package kafka

import (
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	Handler *ckafka.Producer
}

func NewProducer() *Producer {
	p, err := ckafka.NewProducer(&ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
	})

	if err != nil {
		log.Println(err.Error())
	}

	return &Producer{
		Handler: p,
	}
}

func (p *Producer) Publish(message string, topic string) error {
	err := p.Handler.Produce(&ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		return err
	}

	return nil
}
