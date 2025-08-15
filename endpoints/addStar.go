package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"solarsystems.com/DB"
	"solarsystems.com/IO"
)

func AddStar(ctx *gin.Context) {

	var newSatar IO.StarInput

	err := ctx.BindJSON(&newSatar)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newSatarId, createdError := DB.AddStar(IO.StarInput{Name: newSatar.Name, SolarMass: newSatar.SolarMass})

	if createdError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": createdError.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"starId": newSatarId})

}
