package main

import (
	"fmt"

	"github.com/go-gnss/data/cmd/database/util"
	"github.com/go-gnss/data/cmd/database/models"
	"github.com/go-gnss/ntrip"
	"github.com/go-gnss/rtcm/rtcm3"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := gorm.Open("sqlite3", "../test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.Observation{})
	db.AutoMigrate(&models.SatelliteData{})
	db.AutoMigrate(&models.SignalData{})

	streamId := "LAUT01AUS0"

	client, err := ntrip.NewClient("https://stream.geops.team/" + streamId)
	resp, err := client.Connect()
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode, err)
	}

	scanner := rtcm3.NewScanner(resp.Body)
	for frame, err := scanner.NextFrame(); err == nil; frame, err = scanner.NextFrame() {
		switch frame.MessageNumber() {
		case 1077, 1087, 1097, 1107, 1117, 1127:
			obs, _ := util.ObservationMsm7(rtcm3.DeserializeMessageMsm7(frame.Payload), streamId)
			obs.ID = uuid.New()
			db.Create(&obs)
		}
	}
	panic(err)
}
