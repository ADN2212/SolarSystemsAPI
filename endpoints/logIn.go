package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"solarsystems.com/DB"
	"solarsystems.com/IO"
	"time"
)

var secret = (func() string {
	fmt.Println("Searching for the secret key")
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

	fmt.Println("SECRET found")

	return secretkey

})()

func LogIn(ctx *gin.Context) {

	var userCredentials IO.UserInput

	err := ctx.BindJSON(&userCredentials)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//Primero comprobamos que el user exista en la base de datos:
	user, userError := DB.GetUserByUserName(userCredentials.UserName)

	if userError != nil {
		ctx.IndentedJSON(http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("There is no user whit the username: '%s'", userCredentials.UserName)})
		return
	}

	//Luego comparamos la passqoed en texto plano con la password encriptada:
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredentials.Password))

	if passErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": passErr.Error()})
		return
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{"sub": user.Username,
			"exp": time.Now().Add(time.Minute * 5).Unix()})
	

	tokeStr, tokenErr := token.SignedString([]byte(secret))//why this methods argument has to be of type []byte

	if tokenErr != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : tokenErr.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"accesToken": tokeStr})
	//Una maneta alternativa (y recomendada de hacerlo) es a travez de las cookies:
	// ctx.SetSameSite(http.SameSiteLaxMode)	
	// ctx.SetCookie("Authorization", tokeStr, 36000 * 5, "", "", false, true,)
	// ctx.IndentedJSON(http.StatusOK, gin.H{})
}
