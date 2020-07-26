package domain

import (
	"observations/domain/model"
)

type ObservationService interface {
	SubmitObservation(observation *model.Observation) error
	GetAllObservations() ([]*model.Observation, error) // Garbage example
}


type observationService struct {
	repo ObservationRepository
}

func NewObservationService(repo ObservationRepository) ObservationService {
	return &observationService{repo}
}

func (service observationService) SubmitObservation(observation *model.Observation) error {
	return service.repo.Store(observation)
}

func (service observationService) GetAllObservations() ([]*model.Observation, error) {
	return service.repo.GetAllObservations()
}