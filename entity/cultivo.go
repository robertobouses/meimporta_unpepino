package entity

type Cultivo struct {
	IdCultivo          int                `json:"idcultivo"`
	Siglas             string             `json:"siglas"`
	InformacionCultivo InformacionCultivo `json: "informacioncultivo"`
	RequisitosCultivo  RequisitosCultivo  `json: "requisitoscultivo`
	FechasCultivo      FechasCultivo      `json: "fechascultivo"`
	FrutoCultivo       FrutoCultivo       `json: "frutocultivo"`
	SemillaCultivo     SemillaCultivo     `json: "semillacultivo"`
	ProblemasCultivo   ProblemasCultivo   `json: "problemascultivo"`
}

//uint   gorm:"primary_key

type InformacionCultivo struct {
	Nombre             string   `json:"nombre"`
	Color              []string `json:"color"`
	Familia            string   `json:"familia"`
	DensidadPlantacion string   `json:"densidadplantacion"`
	LitrosTierraMaceta int      `json:"litrostierramaceta"`
	Asociaciones       []string `json:"asociaciones"`
}
type RequisitosCultivo struct {
	Agua        string    `json:"agua"`
	Tierra      string    `json:"tierra"`
	Nutricion   string    `json:"nutricion"`
	Salinidad   string    `json:"salinidad"`
	Ph          []float64 `json:"ph"`
	Clima       []string  `json:"clima"`
	Profundidad string    `json:"profuncidad"`
}
type FechasCultivo struct {
	Siembra     string `json:"siembra"`
	Transplante string `json:"trasplante"`
	Cosecha     string `json:"cosecha"`
	Ciclo       string `json:"ciclo"`
}
type FrutoCultivo struct {
	Produccion string `json:"produccion"`
	Nutrientes string `json:"nutrientes"`
}

type SemillaCultivo struct {
	Semilla       string `json:"semilla"`
	SemillasGramo string `json:"semillasgramo"`
	VidaSemilla   string `json:"vidasemilla"`
}

type ProblemasCultivo struct {
	Plagas       string `json:"plagas"`
	Dificultades string `json:"dificultades"`
	Cuidados     string `json:"cuidados"`
	Miscelanea   string `json:"miscelanea"`
}
