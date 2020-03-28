package observations

import (
	"github.com/go-gnss/data"
	"github.com/go-gnss/rtcm/rtcm3"
	"github.com/mishudark/eventhus"
)

var (
	ObservationsAggregateType string = "Observations"
)

// Observations Aggregate provides source of data and where to find the data(?)
// Not sure if a stream of data from a source really aggregates into anything
type ObservationsAggregate struct {
	eventhus.BaseAggregate
	DataSource string // TODO: Some sort of unique ID
	Data string // TODO: No idea what this should be
}

type ObservationReceivedEvent struct {
	Observation data.Observation
}

// TODO: consider moving commands into individual files, or even sub packages
type SubmitObservationCommand struct {
	eventhus.BaseCommand
	Observation data.Observation
}

// Multiple Commands can produce the same Event
// For example this could generate a ObservationReceived event
type SubmitRtcmObservationCommand struct {
	eventhus.BaseCommand
	RtcmObservation rtcm3.Observable
}

// Not sure it makes sense to have a SubmitRtcmObservation command result in
// ObservationReceived event because you lose the original rtcm3.Observable
// type - though could have a RtcmObservationReceived event which a process
// subscribes to and produces SubmitRtcmObservation commands? (but then why
// would you submit observations instead of having the same projection handle
// the data)
// If there's a separate event, then this code would be in model/rtcm

// Similarly you could have the following, which a process subscribes to and
// generates multiple SubmitObservation commands
type SubmitRinexFileCommand struct {
	eventhus.BaseCommand
	RinexFile string //RinexFile
}
