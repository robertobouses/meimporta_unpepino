package app

import (
	"log"
	"strconv"
	"time"

	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
)

func (s *Service) ProcessMyCrops(name string) (mycropresponse mycrop.MyCropResponse) {

	crops, err := s.repo.ExtractCrops()
	if err != nil {
		return nil, err
	}

	threshold := 2
	var result mycrop.CropResponse

	for _, crop := range crops {
		distance := LevenshteinDistance(name, crop.CropInformation.Name)
		if distance <= threshold {
			result := mycrop.CropResponse{
				Name: crop.CropInformation.Name,
				Family: crop.CropInformation.Family,
				Water: crop.CropRequirements.Water.water,
				Soil: crop.CropRequirements.Soil,
				Nutrition: crop.CropRequirements.Nutrition,
				Climate: crop.CropRequirements.Climate,
				Planting: crop.CropDates.Planting,
				Transplant: crop.CropDates.Transplant,
				Harvest: crop.CropDates.Harvest,
				Cycle: crop.CropDates.Cycle,
				Production: crop.CropFruit.Production,	
			}

			

	mycrop, err := s.repo.ExtractMyCrops(nameInt)
	if err != nil {
		log.Print("Error al eliminar el crop:", err)
		return err
	}

	return nil
}

type MyCropResponse struct {
	Planting   time.Time
	Transplant time.Time
	Harvest    time.Time
	Water      []time.Time
	Production int
	Price      int
	Amount     int
}

type MyCrop struct {
	IdMyCrop int       `json:"namemycrop"`
	Name     string    `json:"name"`
	Planting time.Time `json:"planting"`
	FieldId  int       `json:"fieldname"`
}
