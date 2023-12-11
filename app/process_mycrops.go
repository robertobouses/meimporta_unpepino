package app

import (
	"fmt"
	"log"

	"github.com/robertobouses/meimporta_unpepino/app/calculate"
	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
)

func (s *Service) ProcessMyCrops(name string) (mycrop.MyCropResult, error) {
	var result mycrop.MyCropResult

	crops, err := s.repo.ExtractCropsName(name)
	if err != nil {
		return result, err
	}
	if len(crops) == 0 {
		return result, fmt.Errorf("no crops found for name: %s", name)
	}
	firstCrop := crops[0]
	mycrops, err := s.repo.ExtractMyCropsName(name)
	if err != nil {
		return result, err
	}

	field, err := s.repo.ExtractMyFieldsId(mycrops.FieldId)

	result.Planting = mycrops.Planting
	plantingDate := result.Planting
	family := firstCrop.CropInformation.Family
	climate := firstCrop.CropRequirements.Climate[0]
	soil := firstCrop.CropRequirements.Soil
	production := firstCrop.CropFruit.Production
	result.Transplant = calculate.CalculateTransplant(plantingDate, family, climate)
	result.Harvest = calculate.CalculateHarvest(result.Transplant, family, climate)
	result.Water = calculate.CalculateWater(plantingDate, family, climate, soil)
	result.Production = calculate.CalculateProduction(climate, soil, production)
	pricefloat, err := calculate.CalculatePrice()
	if err != nil {
		log.Println(err)
	}
	result.Price = int(pricefloat)
	result.Amount = calculate.CalculateAmount(result.Price, field.SquareMeters)

	return result, nil
}
