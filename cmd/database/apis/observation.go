package apis

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-gnss/data/cmd/database/daos"
	"github.com/google/uuid"
)

func GetObservation(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	if obs, err := daos.GetObservation(id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, obs)
	}
}

func GetObservations(c *gin.Context) {
	if obs, err := daos.GetObservations(); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, obs)
	}
}
