package main

import (
	"log"

	"github.com/dorianneto/delivery-simulator/simulator/application/usecase"
	"github.com/dorianneto/delivery-simulator/simulator/model"
)

func main() {
	route := model.NewModel()
	route.ID = "1"
	route.ClientID = "1"

	route.LoadDestination()

	parser := usecase.NewParser()

	log.Println(parser.ToString(route))
}
