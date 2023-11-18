package app

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (s *Service) PrintCrops() ([]entity.Crop, error) {

	crops, err := s.repo.ExtractCrops()
	if err != nil {
		log.Println("Error al extraer crop", err)
		return []entity.Crop{}, err
	}
	return crops, nil
}
