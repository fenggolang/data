package daos

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/go-gnss/data/cmd/database/models"
)

func GetObservation(id uuid.UUID) (*models.Observation, error) {
	// This should be in config package
	db, err := gorm.Open("sqlite3", "../test.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	obs := models.Observation{SatelliteData: []models.SatelliteData{}}

	err = db.First(&obs, "id = ?", id).
		Preload("SignalData").
		Related(&obs.SatelliteData).
		Error

	return &obs, err
}

func GetObservations() (*[]models.Observation, error) {
	// This should be in config package
	db, err := gorm.Open("sqlite3", "../test.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	obs := []models.Observation{}
	err = db.Find(&obs).Error

	for i, _ := range obs {
		db.Model(obs[i]).Preload("SignalData").Related(&obs[i].SatelliteData)
	}

	return &obs, err
}

func PutObservation(obs *models.Observation) (id uuid.UUID, err error) {
	// This should be in config package
	db, err := gorm.Open("sqlite3", "../test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	obs.ID = uuid.New()
	db.Create(&obs)

	return obs.ID, nil
}
