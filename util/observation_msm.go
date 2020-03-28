package util

// TODO: Consider what package this belongs in

import (
	"github.com/go-gnss/data"
	"github.com/go-gnss/rtcm/rtcm3"
)

func ParseSatelliteMask(satMask uint64) (prns []int) {
	for i, prn := 64, 1; i > 0; i-- {
		if (satMask >> uint64(i-1)) & 0x1 == 1 {
			prns = append(prns, prn)
		}
		prn++
	}
	return prns
}

func ParseSignalMask(sigMask uint32) (ids []int) {
	for i := 32; i > 0; i-- {
		if (sigMask >> uint32(i-1)) & 0x1 == 1 {
			ids = append(ids, i)
		}
	}
	return ids
}

func Utob(v uint64) bool {
	if v == 0 {
		return false
	}
	return true
}

func ParseCellMask(cellMask uint64, length int) (cells []bool) {
	for i := 0; i < length; i++ {
		cells = append([]bool{Utob((cellMask >> uint(i)) & 0x1)}, cells...)
	}
	return cells
}

var Constellations map[int]string = map[int]string{
	1077: "GPS",
	1087: "GLONASS",
	1097: "Galileo",
	1107: "SBAS",
	1117: "QZSS",
	1127: "BeiDou",
}

func ObservationMsm7(msg rtcm3.MessageMsm7) (obs data.Observation, err error) {
	obs = data.Observation{
		Constellation: Constellations[msg.Number()], // This is a bit dangerous
		ReferenceStationID: msg.ReferenceStationId,
		Epoch: msg.Epoch,
		ClockSteeringIndicator: msg.ClockSteeringIndicator,
		ExternalClockIndicator: msg.ExternalClockIndicator,
		SmoothingInterval: msg.SmoothingInterval,
		SatelliteData: []data.SatelliteData{},
	}

	satIDs := ParseSatelliteMask(msg.SatelliteMask)
	sigIDs := ParseSignalMask(msg.SignalMask)
	cellIDs := ParseCellMask(msg.CellMask, len(satIDs) * len(sigIDs))
	cellPos := 0
	sigPos := 0

	for i, satID := range satIDs {
		satData := data.SatelliteData{
			SatelliteID: satID,
			Extended: msg.SatelliteData.Extended[i],
			PhaseRangeRates: msg.SatelliteData.PhaseRangeRates[i],
			SignalData: []data.SignalData{},
		}
		for _, sigID := range sigIDs {
			if cellIDs[cellPos] {
				satData.SignalData = append(satData.SignalData, data.SignalData{
					SignalID: sigID,
					Pseudoranges: msg.SignalData.Pseudoranges[sigPos],
					PhaseRanges: msg.SignalData.PhaseRanges[sigPos],
					PhaseRangeLocks: msg.SignalData.PhaseRangeLocks[sigPos],
					HalfCycles: msg.SignalData.HalfCycles[sigPos],
					CNRs: msg.SignalData.Cnrs[sigPos],
					PhaseRangeRates: msg.SignalData.PhaseRangeRates[sigPos],
				})
				sigPos ++
			}
			cellPos ++
		}
		obs.SatelliteData = append(obs.SatelliteData, satData)
	}

	return obs, err
}
