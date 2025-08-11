package main

import (
	"github.com/gin-gonic/gin"
	"solarsystems.com/endpoints"
)

func main() {
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

	router.Run("localhost:8080")
}
