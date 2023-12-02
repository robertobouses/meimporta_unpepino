package mycrop

import "time"

type MyCropResponse struct {
	IdMyCrop     int
	Name         string
	Planting     time.Time
	SquareMeters int
	Soil         string
	Climate      string
}
