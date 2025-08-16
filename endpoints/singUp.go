package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"solarsystems.com/DB"
	"solarsystems.com/IO"
)

func SingUp(ctx *gin.Context) {

	var newUser IO.UserInput

	err := ctx.BindJSON(&newUser)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//Ojo: esta logica deberia estar en un validador o ser un constraint en la definicion del modelo de la tabla.
	if len(newUser.UserName) < 5 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "username must have at least 5 characters"})
		return
	}
	
	if len(newUser.Password) < 6 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "password must have at least 6 characters"})
		return
	}
	
	//Que hace exactamente esta funcion ???
	//retorna un hash
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	newUser.Password = string(hash)
	userId, createError := DB.AddUser(newUser)

	if createError != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message" : createError.Error()})
		return
	} 

	ctx.IndentedJSON(http.StatusCreated, gin.H{"userId": userId})
}
