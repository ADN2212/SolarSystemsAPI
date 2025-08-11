package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"solarsystems.com/db"
	"strconv"
)

func UpdatePlanet(ctx *gin.Context) {

	planetId, parseErr := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if parseErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": parseErr.Error()})
	}

	var planetBodyData UpdatePlanetInput

	err := ctx.BindJSON(&planetBodyData)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updatedRows, updateError := db.UpdatePlanet(planetId, db.UpdatePlanetInput{
		Name:      planetBodyData.Name,
		Mass:      planetBodyData.Mass,
		IsLibable: planetBodyData.IsLibable,
	})

	if updateError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": updateError.Error()})
		return
	}

	if updatedRows == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("there is no planet whit id = %d", planetId)})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("planet whit id = %d updated successfully", planetId)})

}
