package endpoints

type Star struct {
	Name      string `json:"name" binding:"required"`
	SolarMass uint   `json:"solarMass" binding:"required"`
}

type Planet struct {
	Name      string `json:"name" binding:"required"`
	Mass      int    `json:"mass" binding:"required"`
	IsLibable bool   `json:"isLibable" binding:"required"`
	StarId    uint   `json:"starId" binding:"required"`
}

type UpdatePlanetInput struct {
	Name      string `json:"name"`
	Mass      int    `json:"mass"`
	IsLibable *bool   `json:"isLibable"`
}
