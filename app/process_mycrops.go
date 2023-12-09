package app

import (
	"time"

	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
)

func (s *Service) ProcessMyCrops(name string) (mycrop.MyCropResult, error) {
	var result mycrop.MyCropResult

	// Extract data from ExtractCropsName
	crops, err := s.repo.ExtractCropsName(name)
	if err != nil {
		return result, err
	}

	// Extract data from ExtractMyCrops
	mycrops, err := s.repo.ExtractMyCrops(name)
	if err != nil {
		return result, err
	}

	result.Planting = mycrops.Planting
	result.Transplant = time.Now()
	result.Harvest = time.Now()
	result.Water = []time.Time{time.Now(), time.Now()} // Example water data
	result.Production = 10
	result.Price = 5
	result.Amount = 100

	return result, nil
}
