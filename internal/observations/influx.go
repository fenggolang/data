package main

// TODO: Where does this code belong? Where do we put domain interfaces vs implementations?

import (
	"github.com/go-gnss/data/protobuf"
)

type InfluxRepository struct {
	host   string
	org    string
	bucket string
	token  string
}

func NewInfluxRepository(host, org, bucket, token string) InfluxRepository {
	return InfluxRepository{
		host:   host,
		org:    org,
		bucket: bucket,
		token:  token,
	}
}

func (repo InfluxRepository) Store(observation protobuf.ObservationSet) error {
	return nil
}