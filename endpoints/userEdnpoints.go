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
	"strings"
	"time"
)

var secret = (func() string {

	envErr := godotenv.Load(".env")

	if envErr != nil {
		panic(".env file not found")
	}

	secretkey := os.Getenv("SECRET")

	if len(secretkey) == 0 {
		panic("SECRET not found")
	}

	fmt.Println("SECRET found")

	return secretkey

})()

func LogIn(ctx *gin.Context) {

	var userCredentials IO.UserLoginInput

	err := ctx.BindJSON(&userCredentials)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, userError := DB.GetUserByUserName(userCredentials.UserName)

	if userError != nil {
		ctx.IndentedJSON(http.StatusBadRequest,
			gin.H{"message": fmt.Sprintf("There is no user whit the username: '%s'", userCredentials.UserName)})
		return
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredentials.Password))

	if passErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": passErr.Error()})
		return
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{"sub": user.Username,
			"exp": time.Now().Add(time.Minute * 5).Unix()})

	tokeStr, tokenErr := token.SignedString([]byte(secret))

	if tokenErr != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": tokenErr.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"accesToken": tokeStr})
}

func SingUp(ctx *gin.Context) {

	var newUser IO.UserSinginInput

	err := ctx.BindJSON(&newUser)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if len(newUser.UserName) < 5 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "username must have at least 5 characters"})
		return
	}

	if len(newUser.Password) < 6 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "password must have at least 6 characters"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newUser.Password = string(hash)
	userId, createError := DB.AddUser(newUser)

	if createError != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": createError.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"userId": userId})
}

func Logout(ctx *gin.Context) {
	tokenStr := strings.Split(ctx.GetHeader("Authorization"), " ")[1]

	err := DB.AddTokenToBlackList(tokenStr)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusFound, gin.H{"message": "You are logged out"})

}
