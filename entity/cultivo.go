package entity

type Cultivo struct {
	IdCultivo          int                `json:"idcultivo"`
	Siglas             string             `json:"siglas"`
	InformacionCultivo InformacionCultivo `json:"informacioncultivo"`
}

type InformacionCultivo struct {
	Nombre             string `json:"nombre"`
	LitrosTierraMaceta int    `json:"litrostierramaceta"`
}
