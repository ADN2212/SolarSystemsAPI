package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"solarsystems.com/db"
)

func AddStar(ctx *gin.Context) {

	var newSatar Star

	err := ctx.BindJSON(&newSatar)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newSatarId, createdError := db.AddStar(db.StarInput{Name: newSatar.Name, SolarMass: newSatar.SolarMass})

	if createdError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": createdError.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"starId": newSatarId})

}