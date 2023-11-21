package entity

type SearchRequest struct {
	Name            string   `json:"name"`
	Color           []string `json:"color"`
	PlantingDensity string   `json:"plantingdensity"`
	Water           string   `json:"water"`
	Soil            string   `json:"soil"`
	Nutrition       string   `json:"nutrition"`
	Salinity        string   `json:"salinity"`
	Cycle           string   `json:"cycle"`
}
