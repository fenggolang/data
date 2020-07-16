package data

import "time"

// Observation is a normalized structure which should be able to represent any
// format of GNSS Observable
type Observation struct {
	// Could put this into SatelliteData and have each constellation nested under
	// the same "Observation" which is unique for <Epoch + ReferenceStationId>
	Constellation      string
	ReferenceStationID string
	Epoch              time.Time
	// This should be normalized to a type - spec says 0-4 is not applied, applied,
	// unknown, and reserved
	ClockSteeringIndicator uint8
	// This should be normalized to a type - spec says 0-4 is internal, external
	// (locked), external (not locked), and unknown
	ExternalClockIndicator uint8
	// This could probably be normalized to SmoothingType - spec says true means
	// divergence-free smoothing and false means any other smoothing type
	SmoothingTypeIndicator bool
	// Could be normalized to seconds (or null for no smoothing)
	SmoothingInterval uint8
	SatelliteData     []SatelliteData
}

// TODO: Need to add some precision information which in RTCM is inferred by the MSM type
// TODO: Find out whether the MSM SatelliteData rough ranges and the SignalData fine ranges are separate observations -
// because merging them into a single value (as we have done) is irreversible.
type SatelliteData struct {
	// PRN
	SatelliteID int
	// This is specific for each constellation
	Extended   uint8
	SignalData []SignalData
}

type SignalData struct {
	Band        string
	Frequency   string
	Pseudorange float64 // m
	PhaseRange  float64 // ??
	// TODO: This requires a lookup table - Could be some time range type
	PhaseRangeLock uint16
	HalfCycle      bool
	SNR            float64 // dB-Hz
	PhaseRangeRate float64 // m/s
}
