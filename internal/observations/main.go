package main

import (
	"fmt"
	faker "github.com/bxcodec/faker/v3"
	"github.com/go-gnss/data/protobuf"
	"github.com/golang/protobuf/jsonpb"
	"observations/domain"
)

func main() {
	svc := domain.NewObservationService(&InMemoryRepository{[]*protobuf.ObservationSet{}})
	var obs protobuf.ObservationSet
	faker.FakeData(&obs)
	svc.SubmitObservation(&obs)
	resp, _ := svc.GetAllObservations()
	marshaler := jsonpb.Marshaler{EnumsAsInts: false}
	for _, o := range resp {
		fmt.Println(marshaler.MarshalToString(o))
	}
}