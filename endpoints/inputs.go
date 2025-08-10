package endpoints

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