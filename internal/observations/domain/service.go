package domain

import (
	"github.com/go-gnss/data/protobuf"
)

type ObservationService interface {
	SubmitObservation(observation *protobuf.ObservationSet) error
	GetAllObservations() ([]*protobuf.ObservationSet, error) // Garbage example
}


type observationService struct {
	repo ObservationRepository
}

func NewObservationService(repo ObservationRepository) ObservationService {
	return &observationService{repo}
}

func (service observationService) SubmitObservation(observation *protobuf.ObservationSet) error {
	return service.repo.Store(observation)
}

func (service observationService) GetAllObservations() ([]*protobuf.ObservationSet, error) {
	return service.repo.GetAllObservations()
}