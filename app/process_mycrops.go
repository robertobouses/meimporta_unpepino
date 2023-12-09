package app

import (
	"github.com/robertobouses/meimporta_unpepino/app/calculate"
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
	plantingDate := result.Planting
	family := crops.CropInformation.Family
	climate := crops.CropInformation.Climate
	soil := crops.CropInformation.Soil
	result.Transplant = calculate.CalculateTransplant(plantingDate, family, climate)
	result.Harvest = calculate.CalculateHarvest(result.Transplant, family, climate)
	result.Water = calculate.CalculateWater(plantingDate, family, climate, soil)
	result.Production = calculate.CalculateProduction()
	result.Price = CalculatePrice()
	result.Amount = CalculateAmount()

	return result, nil
}
