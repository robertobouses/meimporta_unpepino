package entity

type ProblemResponse struct {
	Nombre       string `json:"nombre"`
	Plagas       string `json:"plagas"`
	Dificultades string `json:"dificultades"`
	Cuidados     string `json:"cuidados"`
	Miscelanea   string `json:"miscelanea"`
}
