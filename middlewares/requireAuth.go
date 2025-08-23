package midlewares

import (
	"net/http"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"solarsystems.com/DB"
)

// Estoy repitiendo esta misma logica dentro de LogIn en el endpoints pakage.
var secret = (func() string {
	
	envErr := godotenv.Load(".env")

	if envErr != nil {
		panic(".env file not found")
	}

	secretkey := os.Getenv("SECRET")

	if len(secretkey) == 0 {
		panic("SECRET not found")
	}

	return secretkey

})()

func RequireAuth(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if len(authHeader) == 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Authorization header is required"})
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")

	if len(authHeaderParts) != 2 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Authorization header is malformed"})
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tokenStr := authHeaderParts[1]

	if DB.TokenIsBlackListed(tokenStr) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, tokenErr := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})

	if tokenErr != nil {
		ctx.IndentedJSON(http.StatusConflict, gin.H{"message": tokenErr.Error()})
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}
	
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		ctx.Set("username", claims["sub"])
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
