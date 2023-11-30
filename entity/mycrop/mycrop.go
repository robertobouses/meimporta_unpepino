package mycrop

import (
	"time"
)

type mycrop struct {
	IdMyCrop int       `json:"idmycrop"`
	Name     string    `json:"name"`
	Planting time.Time `json:"planting"`
}
