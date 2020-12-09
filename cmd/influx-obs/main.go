package main

import (
	"context"
	"flag"
	"strconv"

	"github.com/go-gnss/data"
	"github.com/go-gnss/data/util"
	"github.com/go-gnss/ntrip"
	"github.com/go-gnss/rtcm/rtcm3"
	"github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
)

func main() {
	caster := flag.String("c", "http://auscors.ga.gov.au:2101/", "NTRIP Caster")
	mount := flag.String("m", "ALIC00AUS0", "NTRIP Mountpoint")
	username := flag.String("u", "", "NTRIP Username")
	password := flag.String("p", "", "NTRIP Password")

	host := flag.String("e", "", "InfluxDB Endpoint")
	org := flag.String("o", "", "InfluxDB Organization")
	bucket := flag.String("b", "", "InfluxDB Bucket")
	token := flag.String("t", "", "InfluxDB Token")
	flag.Parse()

	writeApi := influxdb2.NewClient(*host, *token).WriteApiBlocking(*org, *bucket)

	client, err := ntrip.NewClient(*caster + *mount)
	client.SetBasicAuth(*username, *password)
	resp, err := client.Connect()
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic(resp.StatusCode)
	}

	scanner := rtcm3.NewScanner(resp.Body)
	for msg, err := scanner.NextMessage(); err == nil; msg, err = scanner.NextMessage() {
		if obs, ok := msg.(rtcm3.Message1077); ok {
			r, _ := util.ObservationMsm7(obs.MessageMsm7)
			WriteToInflux(r, writeApi, *mount)
		} else if obs, ok := msg.(rtcm3.Message1087); ok {
			r, _ := util.ObservationMsm7(obs.MessageMsm7)
			WriteToInflux(r, writeApi, *mount)
		} else if obs, ok := msg.(rtcm3.Message1097); ok {
			r, _ := util.ObservationMsm7(obs.MessageMsm7)
			WriteToInflux(r, writeApi, *mount)
		} else if obs, ok := msg.(rtcm3.Message1117); ok {
			r, _ := util.ObservationMsm7(obs.MessageMsm7)
			WriteToInflux(r, writeApi, *mount)
		} else if obs, ok := msg.(rtcm3.Message1127); ok {
			r, _ := util.ObservationMsm7(obs.MessageMsm7)
			WriteToInflux(r, writeApi, *mount)
		}
	}
}

func WriteToInflux(obs data.Observation, writeApi api.WriteApiBlocking, mount string) {
	for _, sat := range obs.SatelliteData {
		for _, sig := range sat.SignalData {
			p := influxdb2.NewPointWithMeasurement("obs").
				AddTag("ReferenceStationID", mount).
				AddTag("Constellation", obs.Constellation).
				AddTag("SatelliteID", strconv.Itoa(sat.SatelliteID)).
				AddTag("Band", sig.Band).
				AddTag("Frequency", sig.Frequency).
				AddField("Pseudorange", sig.Pseudorange).
				AddField("PhaseRange", sig.PhaseRange).
				AddField("PhaseRangeLock", sig.PhaseRangeLock).
				AddField("PhaseRangeRate", sig.PhaseRangeRate).
				AddField("HalfCycle", sig.HalfCycle).
				AddField("SNR", sig.SNR).
				SetTime(obs.Epoch)

			writeApi.WritePoint(context.Background(), p)
		}
	}
}
