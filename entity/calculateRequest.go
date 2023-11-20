package entity

type CalculateRequest struct {
	Name            string `json:"name" form:"nombre" binding:"required"`
	MetrosCuadrados int    `json:"metros" form:"metros" binding:"required"`
}
