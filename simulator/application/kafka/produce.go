package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/dorianneto/delivery-simulator/simulator/application/usecase"
	ikafka "github.com/dorianneto/delivery-simulator/simulator/infra/kafka"
	"github.com/dorianneto/delivery-simulator/simulator/model"
)

type Produce struct {
	Producer *ikafka.Producer
	Parser   *usecase.Parser
}

func NewProduce(producer *ikafka.Producer, parser *usecase.Parser) *Produce {
	return &Produce{
		Producer: producer,
		Parser:   parser,
	}
}

func (p *Produce) Produce(message *ckafka.Message) {
	route := model.NewRoute()

	json.Unmarshal(message.Value, &route)
	route.LoadDestination()
	positions, err := p.Parser.ToString(route)

	if err != nil {
		log.Println(err.Error())
	}

	for _, position := range positions {
		err := p.Producer.Publish(position, os.Getenv("KafkaProduceTopic"))
		time.Sleep(time.Millisecond * 500)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
