package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/go-gnss/data/util"
	"github.com/go-gnss/ntrip"
	"github.com/go-gnss/rtcm/rtcm3"
)

func main() {
	mount := flag.String("m", "http://auscors.ga.gov.au:2101/ALIC00AUS0", "")
	username := flag.String("u", "", "")
	password := flag.String("p", "", "")
	flag.Parse()

	client, err := ntrip.NewClient(*mount)
	client.SetBasicAuth(*username, *password)
	resp, err := client.Connect()
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode, err)
	}

	scanner := rtcm3.NewScanner(resp.Body)
	for msg, err := scanner.NextMessage(); err == nil; msg, err = scanner.NextMessage() {
		if obs, ok := msg.(rtcm3.Message1077); ok {
			r, _ := util.ObservationMsm7(obs.MessageMsm7)
			j, _ := json.Marshal(r)
			fmt.Printf("%+v\n", string(j))
		} else if obs, ok := msg.(rtcm3.Message1087); ok {
			r, _ := util.ObservationMsm7(obs.MessageMsm7)
			j, _ := json.Marshal(r)
			fmt.Printf("%+v\n", string(j))
		} else if obs, ok := msg.(rtcm3.Message1097); ok {
			r, _ := util.ObservationMsm7(obs.MessageMsm7)
			j, _ := json.Marshal(r)
			fmt.Printf("%+v\n", string(j))
		} else if obs, ok := msg.(rtcm3.Message1117); ok {
			r, _ := util.ObservationMsm7(obs.MessageMsm7)
			j, _ := json.Marshal(r)
			fmt.Printf("%+v\n", string(j))
		} else if obs, ok := msg.(rtcm3.Message1127); ok {
			r, _ := util.ObservationMsm7(obs.MessageMsm7)
			j, _ := json.Marshal(r)
			fmt.Printf("%+v\n", string(j))
		}
	}
}
