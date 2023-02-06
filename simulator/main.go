package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/dorianneto/delivery-simulator/simulator/infra/kafka"
	"github.com/dorianneto/delivery-simulator/simulator/model"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file")
	}
}

func main() {
	route := model.NewModel()
	route.ID = "1"
	route.ClientID = "1"

	route.LoadDestination()

	// parser := usecase.NewParser()
	// data, err := parser.ToString(route)

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	message := make(chan *ckafka.Message)
	consumer := kafka.NewKafkcaConsumer(message)

	go consumer.Consume()

	for m := range message {
		fmt.Println(string(m.Value))
	}
}
