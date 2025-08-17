package midlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strings"
	"time"
)

// Estiy repitiendo esta misma logica dentro de LogIn en el endpoints pakage.
var secret = (func() string {
	//fmt.Println("Searching for the secret key")
	envErr := godotenv.Load(".env")

	if envErr != nil {
		panic(".env file not found")
	}

	secretkey := os.Getenv("SECRET")

	//Ramper la aplicacion si el secret no es hallado
	//Se pdoria crear una en vez de hacer esto pero no se que tan seguro seria.
	if len(secretkey) == 0 {
		panic("SECRET not found")
	}

	//fmt.Println("SECRET found")

	return secretkey

})()

func RequireAuth(ctx *gin.Context) {
	fmt.Println("In the auth middleware ..")
	authHeader := ctx.GetHeader("Authorization")

	if len(authHeader) == 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Authorization header is required"})
		ctx.AbortWithStatus(http.StatusBadRequest) //Sin esto el middleware no detendria la llamada al endpoit de turno.
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")

	if len(authHeaderParts) != 2 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Authorization header is malformed"})
		ctx.AbortWithStatus(http.StatusBadRequest) //Sin esto el middleware no detendria la llamada al endpoit de turno.
		return
	}

	//fmt.Println(authHeaderParts)

	tokenStr := authHeaderParts[1]

	token, tokenErr := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})

	if tokenErr != nil {
		//fmt.Println(tokenErr.Error())
		ctx.IndentedJSON(http.StatusConflict, gin.H{"message": tokenErr.Error()})
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)//Invrstigar esta sintaxis

	if ok && token.Valid {
		//fmt.Println(claims)
		//Comprovar que el token no haya expirado:
		//Al parecer el hecho de que el token este expirado se verifica en la funcion anterior.
		//Entonces esta parte no sera necesaria.
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "This token has expired."})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next() //Esta funcion permite que se pase a la siguiente funcion en la ruta.
		return
	} else {
		//ctx.IndentedJSON(http.StatusUnauthorized, gin.H{})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
