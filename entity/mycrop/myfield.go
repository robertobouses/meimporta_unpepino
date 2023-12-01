package mycrop

type MyField struct {
	IdMyField     int    `json:"idmyfield"`
	CatastralCode int    `json:"catastralcode"`
	Name          string `json:"name"`
	SquareMeters  int    `json:"squaremeters"`
	Soil          string `json:"soil"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Climate       string `json:"climate"`
}
