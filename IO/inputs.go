package IO

type StarInput struct {
	Name      string `json:"name" binding:"required"`
	SolarMass uint   `json:"solarMass" binding:"required"`
}

type PlanetInput struct {
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

type UserInput struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
	Rol string `json:"rol" binding:"required"`
}
