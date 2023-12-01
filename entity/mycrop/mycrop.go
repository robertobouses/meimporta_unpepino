package mycrop

import (
	"time"
)

type Mycrop struct {
	IdMyCrop int       `json:"idmycrop"`
	Name     string    `json:"name"`
	Planting time.Time `json:"planting"`
}
