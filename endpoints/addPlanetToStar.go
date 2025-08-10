package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"solarsystems.com/db"
)

func AddPlanetToStar(ctx *gin.Context) {

	var newPlanet Planet

	err := ctx.BindJSON(&newPlanet)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newPlanetId, createError := db.AddPlanetToStar(db.PlanetInput{
		Name:      newPlanet.Name,
		Mass:      newPlanet.Mass,
		IsLibable: newPlanet.IsLibable,
		StarID:    newPlanet.StarId,
	})

	if createError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": createError.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"planetId": newPlanetId})

}
