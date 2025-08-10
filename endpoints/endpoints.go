package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"solarsystems.com/db"
	"strconv"
)

// Si los campos de las estructuras no estan capitlizados no podran ser vistos desde fuera
// Y por ende no podran ser mostados en la response de la API.
type Star struct {
	Name      string `json:"name"`
	SolarMass uint   `json:"solarMass"`
}

type Planet struct {
	Name      string `json:"name"`
	Mass      int    `json:"mass"`
	IsLibable bool   `json:"isLibable"`
	StarId    uint   `json:"starId"`
}

type UpdatePlanetInput struct {
	Name      string `json:"name"`
	Mass      int    `json:"mass"`
	IsLibable *bool   `json:"isLibable"`//*bool me permite manejar el caso en que la clave no sea enviada.
	//Entiendase *bool como "El valor de un puntero a un booleano".
}

func GetSolarSystem(ctx *gin.Context) {

	//Queda penidiente investigar como transformar esto a un uint : /
	starId, parseErr := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if parseErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": parseErr.Error()})
	}

	solarSystem, err := db.GetSolarSystem(starId)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusFound, solarSystem)

}

func AddStar(ctx *gin.Context) {

	var newSatar Star

	err := ctx.BindJSON(&newSatar)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newSatarId, createdError := db.AddStar(db.StarInput{Name: newSatar.Name, SolarMass: newSatar.SolarMass})

	if createdError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": createdError.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"starId": newSatarId})

}

func AddPlanetToStar(ctx *gin.Context) {

	var newPlanet Planet

	err := ctx.BindJSON(&newPlanet)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newPlanetId, createError := db.AddPlanetToStar(db.PlanetInput{
		Name:      newPlanet.Name,
		Mass:      newPlanet.Mass,
		IsLibable: newPlanet.IsLibable,
		StarID:    newPlanet.StarId,
	})

	if createError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": createError.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"planetId": newPlanetId})

}

// Como cada planeta tiene su propio id no es necesario espesificar el id de la estrella de la que se removera el planeta.
func RemovePlanetFromStar(ctx *gin.Context) {
	planetId, parseErr := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if parseErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": parseErr.Error()})
	}

	deleteError := db.RemovePlanetFromStar(planetId)

	if deleteError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": deleteError.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Planet id = %v delete succesfully", planetId)})

}

func DeleteSolarSystem(ctx *gin.Context) {

	starId, parseErr := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if parseErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": parseErr.Error()})
	}

	deleteError := db.DeleteSolarSystem(starId)

	if deleteError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": deleteError.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("The star with id = %v and all its planets have been deleted.", starId)})

}

func UpdateStar(ctx *gin.Context) {

	starId, parseErr := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if parseErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": parseErr.Error()})
	}

	var starBodyData Star

	err := ctx.BindJSON(&starBodyData)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updatedRows, updateError := db.UpdateStar(starId, db.StarInput{
		Name:      starBodyData.Name,
		SolarMass: starBodyData.SolarMass,
	})

	if updateError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": updateError.Error()})
		return
	}

	if updatedRows == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("there is no star whit id = %d", starId)})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Star whit id = %d updated successfully", starId)})

}

//Este endpoint no recisve starId para que no se pueda mover un planeta de una estrella a otra o no dejarlo como un planeta errante.
func UpdatePlanet(ctx *gin.Context) {

	planetId, parseErr := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if parseErr != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": parseErr.Error()})
	}

	var planetBodyData UpdatePlanetInput

	err := ctx.BindJSON(&planetBodyData)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//fmt.Println(*planetBodyData.IsLibable)

	updatedRows, updateError := db.UpdatePlanet(planetId, db.UpdatePlanetInput{
		Name:      planetBodyData.Name,
		Mass:      planetBodyData.Mass,
		IsLibable: planetBodyData.IsLibable,
	})

	if updateError != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": updateError.Error()})
		return
	}

	if updatedRows == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("there is no planet whit id = %d", planetId)})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("planet whit id = %d updated successfully", planetId)})

}
