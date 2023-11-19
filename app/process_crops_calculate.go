package app

import (
	"fmt"
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (s *Service) ProcessCropsCalculate(newCalculate entity.CalculateRequest) (int, error) {
	crop, err := s.repo.ExtractCropsName(newCalculate.Name)
	if err != nil {
		return 0, err
	}

	log.Println("newCalculate en app", newCalculate)
	densidad := crop.DensidadPlantacion
	metros := newCalculate.MetrosCuadrados
	log.Println("densidad en app", densidad)
	log.Println("metros en app", metros)

	var factor int

	switch densidad {
	case "alta":
		factor = 10
	case "media":
		factor = 5
	case "baja":
		factor = 2
	default:
		return 0, fmt.Errorf("densidad de plantación desconocida: %s", densidad)
	}

	resultado := factor * metros
	log.Println("resultado en app", resultado)
	return resultado, nil
}
