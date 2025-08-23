package midlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"solarsystems.com/DB"
	"strings"
)

func CheckRol(ctx *gin.Context) {

	fmt.Println("In the second middleware")

	username, exists := ctx.Get("username")

	if !exists {
		ctx.IndentedJSON(http.StatusConflict, gin.H{"message": "username not found"})
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	usernameStr, ok := username.(string)

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

	fmt.Println(ctx.Request.RequestURI)

	requestURI := ctx.Request.RequestURI

	isStarsEnpoint := strings.Contains(requestURI, "stars")

	if isStarsEnpoint {
		if user.Rol != "god" {
			ctx.IndentedJSON(http.StatusConflict, gin.H{"message": "Only users whit rol 'god' can perform this action"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		} else {
			fmt.Println("Passing to a star endpoint ...")
			ctx.Next()
			return
		}
	}

	isPlanetsEnpoint := strings.Contains(requestURI, "planets")

	if isPlanetsEnpoint {
		if user.Rol != "human" {
			ctx.IndentedJSON(http.StatusConflict, gin.H{"message": "Only users whit rol 'human' can perform this action"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		} else {
			fmt.Println("Passing to a planet endpoint ...")
			ctx.Next()
			return
		}
	}

	//Is this correct ?
	panic("The execution flow should not reach this point.")

}
