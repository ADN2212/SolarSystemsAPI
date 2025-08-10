package db

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Star struct {
	gorm.Model
	ID        uint
	Name      string
	SolarMas uint
}

type Planet struct {
	gorm.Model
	ID        uint
	Name      string
	Mass      int
	IsLibable bool
	StarID    uint `gorm:"constraint:OnDelete:CASCADE;"` //Is not working, o tal vez solo funciona desde el ORM?
}

// Estos cuentan como DTOs?:
type StarInput struct {
	Name      string
	SolarMass uint
}

type PlanetInput struct {
	Name      string
	Mass      int
	IsLibable bool
	StarID    uint
}

// esto deberia ser una variable de entorno.
const dsn string = "host=localhost user=postgres password=123456 dbname=SolarSystemsDB2 port=5432 sslmode=disable TimeZone=Asia/Shanghai"

var dbContext = context.Background()
var db *gorm.DB = (func() *gorm.DB {
	fmt.Println("Initializing database from IFE")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Star{})
	db.AutoMigrate(&Planet{})

	return db
})()

func AddStar(star StarInput) (uint, error) {

	newStar := Star{Name: star.Name, SolarMas: star.SolarMass}
	createError := gorm.G[Star](db).Create(dbContext, &newStar)

	if createError != nil {
		return 0, createError
	}

	return newStar.ID, nil

}

func AddPlanetToStar(planet PlanetInput) (uint, error) {

	_, starError := gorm.G[Star](db).Where("id = ?", planet.StarID).First(dbContext)

	if starError != nil {
		return 0, fmt.Errorf("there is no star whit id = %d", planet.StarID)
	}

	newPlanet := Planet{
		Name:      planet.Name,
		Mass:      planet.Mass,
		IsLibable: planet.IsLibable,
		StarID:    planet.StarID,
	}

	createError := gorm.G[Planet](db).Create(dbContext, &newPlanet)

	if createError != nil {
		return 0, createError
	}

	//Si la query es exitosa el ID se agrega al newPlanet a travez del puntero.
	return newPlanet.ID, nil

}

// Hay alguna forma de no tener que repetir tanto la cracion de estas strucs con casi los mismos campos??
type PlanetOutput struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Mass      int    `json:"mass"`
	IsLibable bool   `json:"isLibable"`
}

type SolarSystemOutput struct {
	StarId        uint           `json:"id"`
	StarName      string         `json:"name"`
	StarSolarMass uint           `json:"solarMass"`
	Planets       []PlanetOutput `json:"planets"`
}

func GetSolarSystem(starId uint64) (SolarSystemOutput, error) {

	star, starError := gorm.G[Star](db).Where("id = ?", starId).First(dbContext)

	var solarSystem SolarSystemOutput

	if starError != nil {
		return solarSystem, fmt.Errorf("there is no star whit id = %d", starId)
	}

	solarSystem.StarId = star.ID
	solarSystem.StarName = star.Name
	solarSystem.StarSolarMass = star.SolarMas

	planets, planetsError := gorm.G[Planet](db).Where("star_id = ?", starId).Find(dbContext)

	if planetsError != nil {
		return solarSystem, errors.New("an error happened while searchinf the planets")
	}

	var planetsSlice []PlanetOutput

	for i := range planets {
		planetsSlice = append(planetsSlice,
			PlanetOutput{
				Id:        planets[i].ID,
				Name:      planets[i].Name,
				Mass:      planets[i].Mass,
				IsLibable: planets[i].IsLibable,
			})
	}

	solarSystem.Planets = planetsSlice

	return solarSystem, nil

}

func RemovePlanetFromStar(planetId uint64) error {

	planet, planetError := gorm.G[Planet](db).Where("id = ?", planetId).First(dbContext)

	if planetError != nil {
		return fmt.Errorf("there is no planet whit id = %d", planetId)
	}

	deletePlanetError := db.Unscoped().Delete(&planet).Error

	if deletePlanetError != nil {
		return deletePlanetError
	}

	return nil

}

func DeleteSolarSystem(starId uint64) error {

	//Esta es la Generics API
	star, starError := gorm.G[Star](db).Where("id = ?", starId).First(dbContext)

	if starError != nil {
		return fmt.Errorf("there is no star whit id = %d", starId)
	}

	//Esta es la Tradicional API (Deberia usar una de las dos, no las dos ligadas dude : |)
	deleteStarError := db.Unscoped().Delete(&star).Error

	if deleteStarError != nil {
		return deleteStarError
	}

	deletePlanetsError := db.Unscoped().Where("star_id", starId).Delete(&Planet{}).Error

	if deletePlanetsError != nil {
		return deletePlanetsError
	}

	return nil

}


func UpdateStar(starId uint64, starBodyData StarInput) (int, error) {

	updatedStar := Star{
		Name:      starBodyData.Name,
		SolarMas: starBodyData.SolarMass,
	}

	//Updates actuliza solo los nom zeros values.
	return gorm.G[Star](db).Where("id = ?", starId).Updates(dbContext, updatedStar)
}

type UpdatePlanetInput struct {
	Name      string
	Mass      int
	IsLibable *bool
}

func UpdatePlanet(planetId uint64, planetBodyData UpdatePlanetInput) (int, error) {

	updatedPlanet := Planet{
		Name:      planetBodyData.Name,
		Mass:      planetBodyData.Mass,
	}

	if planetBodyData.IsLibable == nil {
		fmt.Println("Omitiendo IsLibable: ")
		//Omit hace que sean omitidas las claves espesificadas en sus argumentos
		//Si no se pone star id se setearia a 0 lo cual cambiaria el planeta a una estrella enexistente.
		return gorm.G[Planet](db).Where("id = ?", planetId).Omit("is_libable", "star_id").Updates(dbContext, updatedPlanet)	
	}
	
	fmt.Println("Tomando en consideracion IsLibable")
	updatedPlanet.IsLibable = *planetBodyData.IsLibable
	//En este caso los zero values como false si son actualizados.
	return gorm.G[Planet](db).Where("id = ?", planetId).Select("naem", "mass", "is_libable").Updates(dbContext, updatedPlanet)

}
