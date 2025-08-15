package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"solarsystems.com/DB"
	"strconv"
)

func DeleteSolarSystem(ctx *gin.Context) {

	starId, parseErr := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if parseErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": parseErr.Error()})
	}

	deleteError := DB.DeleteSolarSystem(starId)

	if deleteError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": deleteError.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("The star with id = %v and all its planets have been deleted.", starId)})

}