package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"solarsystems.com/DB"
	"strconv"
)

func RemovePlanetFromStar(ctx *gin.Context) {
	planetId, parseErr := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if parseErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": parseErr.Error()})
	}

	deleteError := DB.RemovePlanetFromStar(planetId)

	if deleteError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": deleteError.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Planet id = %v deleted succesfully", planetId)})

}


