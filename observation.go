package data

// TODO: Consider what package this belongs in

// Observation is a normalized structure which should be able to represent any
// format of GNSS Observable
type Observation struct {
	// Could put this into SatelliteData and have each constellation nested under
	// the same "Observation" which is unique for <Epoch + ReferenceStationId>
	Constellation string
	// This does not seem to be correctly implemented anywhere at the moment -
	// could use the station name instead (otherwise have the ID link to a table
	// of ID + station name)
	// Still can't assume that the ReferenceStationId from RTCM is correct
	ReferenceStationID uint16
	// TODO: normalize constellation epochs as timestamp
	Epoch uint32
	// This can be normalized to a type - spec says 0-4 is not applied, applied,
	// unknown, and reseverd
	ClockSteeringIndicator uint8
	// This can be normalized to a type - spec says 0-4 is internal, external
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

type SatelliteData struct {
	// PRN
	SatelliteID int
	// Merge RoughRangeMilliseconds and and RoughRange into a float
	RoughRangeMilliseconds uint8
	// This is specific for each constellation
	Extended uint8
	// Consider merging rough pseudo and phase ranges (sat data) and fine ranges (sig data)
	// into a float value in signal data
	RoughRange     uint16
	PhaseRangeRate int16
	SignalData     []SignalData
}

type SignalData struct {
	// Use some Signal type which gives frequency band and channel
	SignalID       int
	// See comment on SatelliteData.RoughRange
	Pseudorange    int32
	PhaseRange     int32
	// Could be some time range type
	PhaseRangeLock uint16
	HalfCycle      bool
	SNR            uint16
	PhaseRangeRate int16
}
