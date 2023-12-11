package entity

type Crop struct {
	IdCrop           int              `json:"idcrop"`
	Abbreviation     string           `json:"abbreviation"`
	CropInformation  CropInformation  `json:"cropinformation"`
	CropRequirements CropRequirements `json:"croprequirements"`
	CropDates        CropDates        `json:"cropdates"`
	CropFruit        CropFruit        `json:"cropfruit"`
	CropSeed         CropSeed         `json:"cropseed"`
	CropIssues       CropIssues       `json:"cropissues"`
}

//uint   gorm:"primary_key

type CropInformation struct {
	Name              string   `json:"name"`
	Color             []string `json:"color"`
	Family            string   `json:"family"`
	PlantingDensity   string   `json:"plantingdensity"`
	LitersPottingSoil int      `json:"literspottingsoil"`
	Associations      []string `json:"associations"`
}
type CropRequirements struct {
	Water     string    `json:"water"`
	Soil      string    `json:"soil"`
	Nutrition string    `json:"nutrition"`
	Salinity  string    `json:"salinity"`
	Ph        []float64 `json:"ph"`
	Climate   []string  `json:"climate"`
	Depth     string    `json:"depth"`
}
type CropDates struct {
	Planting   string `json:"planting"`
	Transplant string `json:"trasplante"`
	Harvest    string `json:"harvest"`
	Cycle      string `json:"cycle"`
}
type CropFruit struct {
	Production int    `json:"production"`
	Nutrients  string `json:"nutrients"`
}

type CropSeed struct {
	Seed         string `json:"seed"`
	SeedsPerGram string `json:"seedspergram"`
	SeedLifespan string `json:"seedlifespan"`
}

type CropIssues struct {
	Pests         string `json:"pests"`
	Difficulties  string `json:"difficulties"`
	Care          string `json:"care"`
	Miscellaneous string `json:"miscellaneous"`
}
