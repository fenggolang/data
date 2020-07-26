package domain

import (
	"github.com/go-gnss/data/protobuf"
)

type ObservationRepository interface {
	Store(observation *protobuf.ObservationSet) error
	GetAllObservations() ([]*protobuf.ObservationSet, error) // Garbage example
}