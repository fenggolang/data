package main

import (
	"context"
	"flag"
	"github.com/go-gnss/data"
	"github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	"strconv"

	"github.com/go-gnss/data/util"
	"github.com/go-gnss/ntrip"
	"github.com/go-gnss/rtcm/rtcm3"
)

const (
	token string = "vV3SnGjqeQbKn6t8X53jmgWvCR7__ZieTfiVGLRaa1sMby46C1yKW6ED8UicPkFVdvoDMP6B6wvi8Tt0ua8Mbw=="
)

func main() {
	mount := flag.String("m", "ALIC00AUS0", "")
	caster := flag.String("c", "http://auscors.ga.gov.au:2101/", "")
	username := flag.String("u", "", "")
	password := flag.String("p", "", "")
	flag.Parse()

	ntripClient, err := ntrip.NewClient(*caster + *mount)
	ntripClient.SetBasicAuth(*username, *password)
	resp, err := ntripClient.Connect()
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic(resp.StatusCode)
	}

	// create new client with default option for server url authenticate by token
	client := influxdb2.NewClient("http://localhost:9999", token)
	defer client.Close()
	// user blocking write client for writes to desired bucket
	writeApi := client.WriteApiBlocking("Garbage Town", "observations")

	scanner := rtcm3.NewScanner(resp.Body)
	for msg, err := scanner.NextMessage(); err == nil; msg, err = scanner.NextMessage() {
		if obs, ok := msg.(rtcm3.Message1077); ok {
			o, _ := util.ObservationMsm7(obs.MessageMsm7)
			WriteToInflux(o, writeApi, mount)
		} else if obs, ok := msg.(rtcm3.Message1087); ok {
			o, _ := util.ObservationMsm7(obs.MessageMsm7)
			WriteToInflux(o, writeApi, mount)
		} else if obs, ok := msg.(rtcm3.Message1097); ok {
			o, _ := util.ObservationMsm7(obs.MessageMsm7)
			WriteToInflux(o, writeApi, mount)
		} else if obs, ok := msg.(rtcm3.Message1117); ok {
			o, _ := util.ObservationMsm7(obs.MessageMsm7)
			WriteToInflux(o, writeApi, mount)
		} else if obs, ok := msg.(rtcm3.Message1127); ok {
			o, _ := util.ObservationMsm7(obs.MessageMsm7)
			WriteToInflux(o, writeApi, mount)
		}
	}

	panic(err)
}

func WriteToInflux(o data.Observation, writeApi api.WriteApiBlocking, mount *string)  {
	for _, sat := range o.SatelliteData {
		for _, sig := range sat.SignalData {
			// create point using full params constructor
			p := influxdb2.NewPointWithMeasurement("obs").
				AddTag("ReferenceStationID", *mount).
				AddTag("Constellation", o.Constellation).
				AddTag("SatelliteID", strconv.Itoa(sat.SatelliteID)).
				AddTag("Band", sig.Band).
				AddTag("Frequency", sig.Frequency).
				AddField("Pseudorange", sig.Pseudorange).
				AddField("PhaseRange", sig.PhaseRange).
				AddField("PhaseRangeLock", sig.PhaseRangeLock).
				AddField("PhaseRangeRate", sig.PhaseRangeRate).
				AddField("HalfCycle", sig.HalfCycle).
				AddField("SNR", sig.SNR).
				SetTime(o.Epoch)

			// write point immediately
			writeApi.WritePoint(context.Background(), p)
		}
	}
}
