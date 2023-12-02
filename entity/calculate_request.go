package entity

type CalculateRequest struct {
	Name         string `json:"name" form:"name" binding:"required"`
	SquareMeters int    `json:"meters" form:"meters" binding:"required"`
}
