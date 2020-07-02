package data

import "time"

// Flat structure version - more suitable for time series DB
type Observationx struct {
	ReferenceStationID string
	Epoch              time.Time
	// This should be normalized to a type - spec says 0-4 is not applied, applied,
	// unknown, and reseverd
	ClockSteeringIndicator uint8
	// This should be normalized to a type - spec says 0-4 is internal, external
	// (locked), external (not locked), and unknown
	ExternalClockIndicator uint8
	// This could probably be normalized to SmoothingType - spec says true means
	// divergence-free smoothing and false means any other smoothing type
	SmoothingTypeIndicator bool
	// Could be normalized to seconds (or null for no smoothing)
	SmoothingInterval uint8
	Constellation     string
	SatelliteID       int
	// This is specific for each constellation
	Extended  uint8
	Frequency string
	Signal    string
	Pseudorange float64
	PhaseRange  int32
	// Could be some time range type
	PhaseRangeLock uint16
	HalfCycle      bool
	SNR            float64
	PhaseRangeRate float64
}
