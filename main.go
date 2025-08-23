package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"solarsystems.com/endpoints"
	"solarsystems.com/middlewares"
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
	router.POST("stars", midlewares.RequireAuth, midlewares.CheckRol, endpoints.AddStar)
	router.GET("stars/:id", midlewares.RequireAuth, midlewares.CheckRol, endpoints.GetSolarSystem)
	router.DELETE("stars/:id", midlewares.RequireAuth, midlewares.CheckRol, endpoints.DeleteSolarSystem)
	router.PATCH("stars/:id", midlewares.RequireAuth, midlewares.CheckRol, endpoints.UpdateStar)

	//Planet Endpoints:
	router.POST("planets", midlewares.RequireAuth, midlewares.CheckRol,endpoints.AddPlanetToStar)
	router.DELETE("planets/:id", midlewares.RequireAuth,midlewares.CheckRol ,endpoints.RemovePlanetFromStar)
	router.PATCH("planets/:id", midlewares.RequireAuth, midlewares.CheckRol,endpoints.UpdatePlanet)

	//user endpoints:
	router.POST("users", endpoints.SingUp)
	router.GET("users", endpoints.LogIn)
	router.DELETE("users", midlewares.RequireAuth, endpoints.Logout)

	port := os.Getenv("PORT")
	
	if len(port) == 0 {
		fmt.Println("PORT not found, setting to 8080")
		port = "8080"
	}
	
	router.Run(fmt.Sprintf("localhost:%s", port))
}
