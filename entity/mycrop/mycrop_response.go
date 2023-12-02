package mycrop

import "time"

type MyCropResponse struct {
	Planting   time.Time
	Transplant time.Time
	Harvest    time.Time
	Water      []time.Time
	Production int
	Price      int
	Amount     int
}

//date of care
