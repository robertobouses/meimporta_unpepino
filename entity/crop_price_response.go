package entity

type CropPriceResponse struct {
	Crop   string  `json:"crop"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}
