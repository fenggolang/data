package main

import (
	"github.com/go-gnss/data/protobuf"
)

type InMemoryRepository struct {
	obs []*protobuf.ObservationSet
}

func (r *InMemoryRepository) Store(observation *protobuf.ObservationSet) error {
	r.obs = append(r.obs, observation)
	return nil
}

func (r *InMemoryRepository) GetAllObservations() ([]*protobuf.ObservationSet, error) {
	return r.obs, nil
}