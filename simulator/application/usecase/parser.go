package usecase

import (
	"encoding/json"

	"github.com/dorianneto/delivery-simulator/simulator/model"
)

type partialRoute struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ToString(route *model.Route) ([]string, error) {
	var temp partialRoute
	var result []string

	total := len(route.Positions)

	for k, v := range route.Positions {
		temp.ID = route.ID
		temp.ClientID = route.ClientID
		temp.Position = []float64{v.Lat, v.Long}
		temp.Finished = false

		if k == total-1 {
			temp.Finished = true
		}

		json, err := json.Marshal(temp)

		if err != nil {
			return nil, err
		}

		result = append(result, string(json))
	}

	return result, nil
}
