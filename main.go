package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"solarsystems.com/endpoints"
)

func main() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		panic(".env file not found")
	} else {
		fmt.Println(".env file loaded succesfully")
	}
	
	router := gin.Default()
	//Star Endpoints:
	router.POST("stars", endpoints.AddStar)
	router.GET("stars/:id", endpoints.GetSolarSystem)
	router.DELETE("stars/:id", endpoints.DeleteSolarSystem)
	router.PATCH("stars/:id", endpoints.UpdateStar)

	//Planet Endpoints:
	router.POST("planets", endpoints.AddPlanetToStar)
	router.DELETE("planets/:id", endpoints.RemovePlanetFromStar)
	router.PATCH("planets/:id", endpoints.UpdatePlanet)

	//user endpoints:
	router.POST("users", endpoints.SingUp)


	port := os.Getenv("PORT")
	
	if len(port) == 0 {
		fmt.Println("PORT not found, setting to 8080")
		port = "8080"
	}
	
	router.Run(fmt.Sprintf("localhost:%s", port))
}
