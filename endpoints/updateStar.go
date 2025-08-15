package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"solarsystems.com/db"
	"strconv"
	"solarsystems.com/IO"
)

func UpdateStar(ctx *gin.Context) {

	starId, parseErr := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if parseErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": parseErr.Error()})
	}

	var starBodyData IO.StarInput

	err := ctx.BindJSON(&starBodyData)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updatedRows, updateError := db.UpdateStar(starId, IO.StarInput{
		Name:      starBodyData.Name,
		SolarMass: starBodyData.SolarMass,
	})

	if updateError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": updateError.Error()})
		return
	}

	if updatedRows == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("there is no star whit id = %d", starId)})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Star whit id = %d updated successfully", starId)})

}
