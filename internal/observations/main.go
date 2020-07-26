package main

import (
	"encoding/json"
	"fmt"
	faker "github.com/bxcodec/faker/v3"
	"observations/domain"
	"observations/domain/model"
)

func main() {
	svc := domain.NewObservationService(&InMemoryRepository{[]*model.Observation{}})
	var obs model.Observation
	faker.FakeData(&obs)
	svc.SubmitObservation(&obs)
	resp, _ := svc.GetAllObservations()
	for _, o := range resp {
		j, _ := json.Marshal(o)
		fmt.Println(string(j))
	}
}