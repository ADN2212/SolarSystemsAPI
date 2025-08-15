package DB

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"solarsystems.com/IO"
)

func AddStar(star IO.StarInput) (uint, error) {
	
	newStar := Star{Name: star.Name, SolarMas: star.SolarMass}
	createError := gorm.G[Star](db).Create(dbContext, &newStar)

	if createError != nil {
		return 0, createError
	}

	return newStar.ID, nil
}

func AddPlanetToStar(planet IO.PlanetInput) (uint, error) {

	_, starError := gorm.G[Star](db).Where("id = ?", planet.StarId).First(dbContext)

	if starError != nil {
		return 0, fmt.Errorf("there is no star whit id = %d", planet.StarId)
	}

	newPlanet := Planet{
		Name:      planet.Name,
		Mass:      planet.Mass,
		IsLibable: planet.IsLibable,
		StarID:    planet.StarId,
	}

	createError := gorm.G[Planet](db).Create(dbContext, &newPlanet)

	if createError != nil {
		return 0, createError
	}

	return newPlanet.ID, nil

}

func GetSolarSystem(starId uint64) (IO.SolarSystemOutput, error) {
	star, starError := gorm.G[Star](db).Where("id = ?", starId).First(dbContext)

	var solarSystem IO.SolarSystemOutput

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

	star, starError := gorm.G[Star](db).Where("id = ?", starId).First(dbContext)

	if starError != nil {
		return fmt.Errorf("there is no star whit id = %d", starId)
	}

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

func UpdateStar(starId uint64, starBodyData IO.StarInput) (int, error) {

	updatedStar := Star{
		Name:     starBodyData.Name,
		SolarMas: starBodyData.SolarMass,
	}

	return gorm.G[Star](db).Where("id = ?", starId).Updates(dbContext, updatedStar)
}

func UpdatePlanet(planetId uint64, planetBodyData IO.UpdatePlanetInput) (int, error) {

	updatedPlanet := Planet{
		Name: planetBodyData.Name,
		Mass: planetBodyData.Mass,
	}

	if planetBodyData.IsLibable == nil {
		return gorm.G[Planet](db).Where("id = ?", planetId).Omit("is_libable", "star_id").Updates(dbContext, updatedPlanet)
	}

	updatedPlanet.IsLibable = *planetBodyData.IsLibable
	return gorm.G[Planet](db).Where("id = ?", planetId).Select("name", "mass", "is_libable").Updates(dbContext, updatedPlanet)

}
