package domain

import (
	"observations/domain/model"
)

type ObservationRepository interface {
	Store(observation *model.Observation) error
	GetAllObservations() ([]*model.Observation, error) // Garbage example
}