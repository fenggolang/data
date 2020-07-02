package util

// TODO: Consider what package this belongs in

import (
	"math"
	"time"

	"github.com/go-gnss/data"
	"github.com/go-gnss/rtcm/rtcm3"
)

func ParseSatelliteMask(satMask uint64) (prns []int) {
	for i, prn := 64, 1; i > 0; i-- {
		if (satMask>>uint64(i-1))&0x1 == 1 {
			prns = append(prns, prn)
		}
		prn++
	}
	return prns
}

func ParseSignalMask(sigMask uint32) (ids []int) {
	for i, id := 32, 1; i > 0; i-- {
		if (sigMask>>uint32(i-1))&0x1 == 1 {
			ids = append(ids, id)
		}
		id++
	}
	return ids
}

type freq_sig struct {
	frequency string
	signal    string
}

var (
	signals map[string]map[int]freq_sig = map[string]map[int]freq_sig{
		"GPS": {
			2:  freq_sig{"L1", "C/A"},
			3:  freq_sig{"L1", "P"},
			4:  freq_sig{"L1", "Z"},
			8:  freq_sig{"L2", "C/A"},
			9:  freq_sig{"L2", "P"},
			10: freq_sig{"L2", "Z"},
			15: freq_sig{"L2", "L2C(M)"},
			16: freq_sig{"L2", "L2C(L)"},
			17: freq_sig{"L2", "L2C(M+L)"},
			22: freq_sig{"L5", "I"},
			23: freq_sig{"L5", "Q"},
			24: freq_sig{"L5", "I+Q"},
			30: freq_sig{"L1", "L1C-D"},
			31: freq_sig{"L1", "L1C-P"},
			32: freq_sig{"L1", "L1C-(D+P)"},
		},
		"GLONASS": {
			2: freq_sig{"G1", "C/A"},
			3: freq_sig{"G1", "P"},
			8: freq_sig{"G2", "C/A"},
			9: freq_sig{"G2", "P"},
		},
		"Galileo": {
			2:  freq_sig{"E1", "C"},
			3:  freq_sig{"E1", "A"},
			4:  freq_sig{"E1", "B"},
			5:  freq_sig{"E1", "B+C"},
			6:  freq_sig{"E1", "A+B+C"},
			8:  freq_sig{"E6", "C"},
			9:  freq_sig{"E6", "A"},
			10: freq_sig{"E6", "B"},
			11: freq_sig{"E6", "B+C"},
			12: freq_sig{"E6", "A+B+C"},
			14: freq_sig{"E5B", "I"},
			15: freq_sig{"E5B", "Q"},
			16: freq_sig{"E5B", "I+Q"},
			18: freq_sig{"E5(A+B)", "I"},
			19: freq_sig{"E5(A+B)", "Q"},
			20: freq_sig{"E5(A+B)", "I+Q"},
			22: freq_sig{"E5A", "I"},
			23: freq_sig{"E5A", "Q"},
			24: freq_sig{"E5A", "I+Q"},
		},
		"QZSS": {
			2:  freq_sig{"L1", "C/A"},
			9:  freq_sig{"LEX", "S"},
			10: freq_sig{"LEX", "L"},
			11: freq_sig{"LEX", "S+L"},
			15: freq_sig{"L2", "L2C(M)"},
			16: freq_sig{"L2", "L2C(L)"},
			17: freq_sig{"L2", "L2C(M+L)"},
			22: freq_sig{"L5", "I"},
			23: freq_sig{"L5", "Q"},
			24: freq_sig{"L5", "I+Q"},
			30: freq_sig{"L1", "L1C(D)"},
			31: freq_sig{"L1", "L1C(P)"},
			32: freq_sig{"L1", "L1C(D+P)"},
		},
		"BeiDou": {
			2:  freq_sig{"B1", "I"},
			3:  freq_sig{"B1", "Q"},
			4:  freq_sig{"B1", "I+Q"},
			8:  freq_sig{"B3", "I"},
			9:  freq_sig{"B3", "Q"},
			10: freq_sig{"B3", "I+Q"},
			14: freq_sig{"B2", "I"},
			15: freq_sig{"B2", "Q"},
			16: freq_sig{"B2", "I+Q"},
		},
	}
)

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

func MessageNumberToConstellation(messageNumber int) string {
	switch messageNumber {
	case 1071, 1072, 1073, 1074, 1075, 1076, 1077:
		return "GPS"
	case 1081, 1082, 1083, 1084, 1085, 1086, 1087:
		return "GLONASS"
	case 1091, 1092, 1093, 1094, 1095, 1096, 1097:
		return "Galileo"
	case 1111, 1112, 1113, 1114, 1115, 1116, 1117:
		return "QZSS"
	case 1121, 1122, 1123, 1124, 1125, 1126, 1127:
		return "BeiDou"
	}
	return "" // TODO
}

// Converting rough ranges into pseduorange in KMs
//   e.g. ((84 + (220 * math.Pow(2, -10)) + (121087 * math.Pow(2, -29))) * 299792.458) / 1000
// SatelliteData[x].RoughRangeMilliseconds +
// SatelliteData[x].RoughRange * math.Pow(2, -10) +
// SatelliteData[x].SignalData[y].Pseudorange * math.Pow(2, -29) *
// SpeedOfLightPerMillisecond / 1000
func Pseudorange(roughRangeMillis uint8, roughRange uint16, fineRange int32) float64 {
	// TODO: This is irreversible - find out whether the satellite data rough ranges are truly different observations
	// to the signal data fine ranges
	return ((float64(roughRangeMillis) + (float64(roughRange) * math.Pow(2, -10)) + (float64(fineRange) * math.Pow(2, -29))) * 299792.458) / 1000
}

func ObservationMsm7(msg rtcm3.MessageMsm7) (obs data.Observation, err error) {
	var t time.Time
	switch msg.Number() {
	case 1077:
		t = rtcm3.Message1077{msg}.Time()
	case 1087:
		t = rtcm3.Message1087{msg}.Time()
	case 1097:
		t = rtcm3.Message1097{msg}.Time()
	case 1117:
		t = rtcm3.Message1117{msg}.Time()
	case 1127:
		t = rtcm3.Message1127{msg}.Time()
	}
	obs = data.Observation{
		Constellation: MessageNumberToConstellation(msg.Number()),
		// TODO: ReferenceStationID, leaving up to client to add
		Epoch:                  t,
		ClockSteeringIndicator: msg.ClockSteeringIndicator,
		ExternalClockIndicator: msg.ExternalClockIndicator,
		SmoothingInterval:      msg.SmoothingInterval,
		SatelliteData:          []data.SatelliteData{},
	}

	satIDs := ParseSatelliteMask(msg.SatelliteMask)
	sigIDs := ParseSignalMask(msg.SignalMask)
	cellIDs := ParseCellMask(msg.CellMask, len(satIDs)*len(sigIDs))
	cellPos := 0
	sigPos := 0

	for i, satID := range satIDs {
		satData := data.SatelliteData{
			SatelliteID: satID,
			Extended:    msg.SatelliteData.Extended[i],
			SignalData:  []data.SignalData{},
		}
		for _, sigID := range sigIDs {
			if cellIDs[cellPos] {
				satData.SignalData = append(satData.SignalData, data.SignalData{
					Frequency: signals[obs.Constellation][sigID].frequency,
					Signal:    signals[obs.Constellation][sigID].signal,
					Pseudorange: Pseudorange(
						msg.SatelliteData.RangeMilliseconds[i],
						msg.SatelliteData.Ranges[i],
						msg.SignalData.Pseudoranges[sigPos]),
					PhaseRange:     msg.SignalData.PhaseRanges[sigPos],
					PhaseRangeLock: msg.SignalData.PhaseRangeLocks[sigPos],
					HalfCycle:      msg.SignalData.HalfCycles[sigPos],
					SNR:            float64(msg.SignalData.Cnrs[sigPos]) * math.Pow(2, -4),
					PhaseRangeRate: float64(msg.SatelliteData.PhaseRangeRates[i]) + (float64(msg.SignalData.PhaseRangeRates[sigPos]) * math.Pow(2, -31)),
				})
				sigPos++
			}
			cellPos++
		}
		obs.SatelliteData = append(obs.SatelliteData, satData)
	}

	return obs, err
}
