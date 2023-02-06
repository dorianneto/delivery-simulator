package main

import (
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	akafka "github.com/dorianneto/delivery-simulator/simulator/application/kafka"
	"github.com/dorianneto/delivery-simulator/simulator/application/usecase"
	ikafka "github.com/dorianneto/delivery-simulator/simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file")
	}
}

// Input:
// {"clientId":"1","routeId":"1"}
// {"clientId":"2","routeId":"2"}
// {"clientId":"3","routeId":"3"}

func main() {
	message := make(chan *ckafka.Message)

	parser := usecase.NewParser()
	consumer := ikafka.NewConsumer(message)
	producer := ikafka.NewProducer()

	go consumer.Consume()

	produce := akafka.NewProduce(producer, parser)

	for m := range message {
		go produce.Produce(m)
	}
}
