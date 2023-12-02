package mycrop

type CropResponse struct {
	Name       string   `json:"name"`
	Family     string   `json:"family"`
	Water      string   `json:"water"`
	Soil       string   `json:"soil"`
	Nutrition  string   `json:"nutrition"`
	Climate    []string `json:"climate"`
	Planting   string   `json:"planting"`
	Transplant string   `json:"trasplante"`
	Harvest    string   `json:"harvest"`
	Cycle      string   `json:"cycle"`
	Production string   `json:"production"`
}
