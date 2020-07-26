package main

import (
	"observations/domain/model"
)

type InMemoryRepository struct {
	obs []*model.Observation
}

func (r *InMemoryRepository) Store(observation *model.Observation) error {
	r.obs = append(r.obs, observation)
	return nil
}

func (r *InMemoryRepository) GetAllObservations() ([]*model.Observation, error) {
	return r.obs, nil
}