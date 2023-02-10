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
	ClientID  string     `json:"clientId"`
	Positions []Position `json:"positions"`
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) LoadDestination() error {
	if r.ID == "" {
		return errors.New("couldn't find the destination for this route")
	}

	file, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		lat, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			log.Println("couldn't convert lat")
			return nil
		}

		long, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			log.Println("couldn't convert long")
			return nil
		}

		r.Positions = append(r.Positions, Position{Lat: lat, Long: long})
	}

	return nil
}
