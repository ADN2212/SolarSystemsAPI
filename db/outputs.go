package db

// Hay alguna forma de no tener que repetir tanto la cracion de estas strucs con casi los mismos campos ??
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

