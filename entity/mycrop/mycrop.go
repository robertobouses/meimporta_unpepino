package mycrop

import (
	"time"
)

type MyCrop struct {
	IdMyCrop int       `json:"idmycrop"`
	Name     string    `json:"name"`
	Planting time.Time `json:"planting"`
	FieldId  int       `json:"fieldid"`
}
