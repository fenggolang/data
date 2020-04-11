package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-gnss/data/cmd/database/apis"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	v1.GET("/observations/:id", apis.GetObservation)
	v1.GET("/observations", apis.GetObservations)
	// /sources/:id/observations

	r.Run(fmt.Sprintf(":%v", 8000))
}
