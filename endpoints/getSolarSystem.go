package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"solarsystems.com/db"
	"strconv"
)

func GetSolarSystem(ctx *gin.Context) {
	
	starId, parseErr := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if parseErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": parseErr.Error()})
	}

	solarSystem, err := db.GetSolarSystem(starId)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusFound, solarSystem)

}
