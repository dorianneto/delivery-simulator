package model

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string     `json:"routeId"`
	ClientId  string     `json:"clientId"`
	Positions []Position `json:"positions"`
}

func NewModel() *Route {
	return &Route{}
}

func (route *Route) LoadDestination() error {
	if route.ID == "" {
		return errors.New("couldn't find the destination for this route")
	}

	file, err := os.Open("destinations/" + route.ID + ".txt")
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		lat, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			log.Println("couldn't convert lat")
			return nil
		}

		long, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			log.Println("couldn't convert long")
			return nil
		}

		route.Positions = append(route.Positions, Position{Lat: lat, Long: long})
	}

	return nil
}
