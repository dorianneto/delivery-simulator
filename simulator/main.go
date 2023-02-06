package main

import (
	"log"

	"github.com/dorianneto/delivery-simulator/simulator/model"
)

func main() {
	route := model.NewModel()
	route.ID = "1"

	route.LoadDestination()

	log.Println(route.Positions)
}
