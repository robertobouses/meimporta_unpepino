package entity

type SearchRequest struct {
	Nombre             string   `json:"nombre"`
	Color              []string `json:"color"`
	DensidadPlantacion string   `json:"densidadplantacion"`
	Agua               string   `json:"agua"`
	Tierra             string   `json:"tierra"`
	Nutricion          string   `json:"nutricion"`
	Salinidad          string   `json:"salinidad"`
	Ciclo              string   `json:"ciclo"`
}
