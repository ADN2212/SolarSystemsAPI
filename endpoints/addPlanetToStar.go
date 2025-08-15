package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"solarsystems.com/DB"
	"solarsystems.com/IO"
)

func AddPlanetToStar(ctx *gin.Context) {

	var newPlanet IO.PlanetInput

	err := ctx.BindJSON(&newPlanet)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newPlanetId, createError := DB.AddPlanetToStar(IO.PlanetInput{
		Name:      newPlanet.Name,
		Mass:      newPlanet.Mass,
		IsLibable: newPlanet.IsLibable,
		StarId:    newPlanet.StarId,
	})

	if createError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": createError.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"planetId": newPlanetId})

}
