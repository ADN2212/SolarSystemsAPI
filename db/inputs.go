package db

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

type UpdatePlanetInput struct {
	Name      string
	Mass      int
	IsLibable *bool
}
