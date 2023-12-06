package mycrop

import "time"

type MyCropResult struct {
	Planting   time.Time
	Transplant time.Time
	Harvest    time.Time
	Water      []time.Time
	Production int
	Price      int
	Amount     int
}

//date of care?
//son los datos que obtendrá el usuario como resultado a su consulta
