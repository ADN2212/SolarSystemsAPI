package endpoints

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
	IsLibable *bool   `json:"isLibable"`
}
