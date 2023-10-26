package entity

type Cultivo struct {
	IdCultivo int    `json:"idcultivo"`
	Siglas    string `json:"siglas"`
	InformacionCultivo
	RequisitosCultivo
	FechasCultivo
	FrutoCultivo
	SemillaCultivo
	ProblemasCultivo
}

//uint   gorm:"primary_key

type InformacionCultivo struct {
	Nombre             string `json:"nombre"`
	LitrosTierraMaceta int    `json:"litrostierramaceta"`
}
