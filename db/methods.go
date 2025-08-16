package DB

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"solarsystems.com/IO"
)

func AddStar(starInput IO.StarInput) (uint, error) {
	
	newStar := star{Name: starInput.Name, SolarMas: starInput.SolarMass}
	createError := gorm.G[star](db).Create(dbContext, &newStar)

	if createError != nil {
		return 0, createError
	}

	return newStar.ID, nil
}

func AddPlanetToStar(planetInput IO.PlanetInput) (uint, error) {

	_, starError := gorm.G[star](db).Where("id = ?", planetInput.StarId).First(dbContext)

	if starError != nil {
		return 0, fmt.Errorf("there is no star whit id = %d", planetInput.StarId)
	}

	newPlanet := planet{
		Name:      planetInput.Name,
		Mass:      planetInput.Mass,
		IsLibable: planetInput.IsLibable,
		StarID:    planetInput.StarId,
	}

	createError := gorm.G[planet](db).Create(dbContext, &newPlanet)

	if createError != nil {
		return 0, createError
	}

	return newPlanet.ID, nil

}

func GetSolarSystem(starId uint64) (IO.SolarSystemOutput, error) {
	starResult, starError := gorm.G[star](db).Where("id = ?", starId).First(dbContext)

	var solarSystem IO.SolarSystemOutput

	if starError != nil {
		return solarSystem, fmt.Errorf("there is no star whit id = %d", starId)
	}

	solarSystem.StarId = starResult.ID
	solarSystem.StarName = starResult.Name
	solarSystem.StarSolarMass = starResult.SolarMas

	planets, planetsError := gorm.G[planet](db).Where("star_id = ?", starId).Find(dbContext)

	if planetsError != nil {
		return solarSystem, errors.New("an error happened while searchinf the planets")
	}

	var planetsSlice []IO.PlanetOutput

	for i := range planets {
		planetsSlice = append(planetsSlice,
			IO.PlanetOutput{
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

	planet, planetError := gorm.G[planet](db).Where("id = ?", planetId).First(dbContext)

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

	star, starError := gorm.G[star](db).Where("id = ?", starId).First(dbContext)

	if starError != nil {
		return fmt.Errorf("there is no star whit id = %d", starId)
	}

	deleteStarError := db.Unscoped().Delete(&star).Error

	if deleteStarError != nil {
		return deleteStarError
	}

	deletePlanetsError := db.Unscoped().Where("star_id", starId).Delete(&planet{}).Error

	if deletePlanetsError != nil {
		return deletePlanetsError
	}

	return nil

}

func UpdateStar(starId uint64, starBodyData IO.StarInput) (int, error) {

	updatedStar := star{
		Name:     starBodyData.Name,
		SolarMas: starBodyData.SolarMass,
	}

	return gorm.G[star](db).Where("id = ?", starId).Updates(dbContext, updatedStar)
}

func UpdatePlanet(planetId uint64, planetBodyData IO.UpdatePlanetInput) (int, error) {

	updatedPlanet := planet{
		Name: planetBodyData.Name,
		Mass: planetBodyData.Mass,
	}

	if planetBodyData.IsLibable == nil {
		return gorm.G[planet](db).Where("id = ?", planetId).Omit("is_libable", "star_id").Updates(dbContext, updatedPlanet)
	}

	updatedPlanet.IsLibable = *planetBodyData.IsLibable
	return gorm.G[planet](db).Where("id = ?", planetId).Select("name", "mass", "is_libable").Updates(dbContext, updatedPlanet)

}

func AddUser(userInput IO.UserInput) (uint, error) {

	newUser := user{
		Username: userInput.UserName,
		Password: userInput.Password,
	}

	createError := gorm.G[user](db).Create(dbContext, &newUser)

	if createError != nil {
		return 0, createError
	}

	return  newUser.ID, nil

}
