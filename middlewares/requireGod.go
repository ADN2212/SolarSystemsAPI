package midlewares

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"solarsystems.com/DB"
)

func RequireGod(ctx *gin.Context) {

	fmt.Println("In the second middleware")

	username, exists := ctx.Get("username")

	if !exists {
		ctx.IndentedJSON(http.StatusConflict, gin.H{"message": "username not found"})
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	usernameStr, ok := username.(string)//Esta es una forma de parcear !!

	if !ok {
		ctx.IndentedJSON(http.StatusConflict, gin.H{"message": "could not parce username to string"})
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user, err := DB.GetUserByUserName(usernameStr)

	if err != nil {
		ctx.IndentedJSON(http.StatusConflict, gin.H{"message": "user not found in data base"})
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if user.Rol != "god" {
		ctx.IndentedJSON(http.StatusConflict, gin.H{"message": "Only users whit rol 'god' can perform this action"})
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	fmt.Println("Passing to the endpoint ...")
	ctx.Next()

}
