package data

import "time"

// Flat structure version - potentially more suitable for use with InfluxDB
type Observationx struct {
	Epoch time.Time

	// Tag Keys
	ReferenceStationID string
	Constellation      string
	SatelliteID        int
	// TODO: Consider combining Frequency and Signal as they are in RINEX
	Band     string
	Frequency string

	// Field Keys
	Pseudorange float64
	PhaseRange  int32
	// Could be some time range type
	PhaseRangeLock uint16
	HalfCycle      bool
	SNR            float64
	PhaseRangeRate float64

	// TODO: These should be field keys
	//SmoothingInterval  uint8
	// TODO: This should be normalized to a type - spec says 0-4 is not applied, applied,
	// unknown, and reserved
	//ClockSteeringIndicator uint8
	// TODO: This should be normalized to a type - spec says 0-4 is internal, external
	// (locked), external (not locked), and unknown
	//ExternalClockIndicator uint8
	// TODO: This could probably be normalized to SmoothingType - spec says true means
	// divergence-free smoothing and false means any other smoothing type
	//SmoothingTypeIndicator bool
	// Could be normalized to seconds (or null for no smoothing)
	// TODO: This is specific for each constellation, what should we do?
	//Extended uint8
}
