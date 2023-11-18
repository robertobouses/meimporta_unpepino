package repository

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) ExtractCrops() ([]entity.Crop, error) {
	rows, err := r.db.Query(ExtractCropsQuery)
	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL: %v", err)
		return []entity.Crop{}, err
	}
	defer rows.Close()

	var crops []entity.Crop

	for rows.Next() {
		var crop entity.Crop
		var ph string
		var requirements entity.CropRequirements 
		var clima string
		var colorBytes []byte
		var colorSlice []string
		var associations       CSV string

		if err := rows.Scan(
			&crop.IdCrop,
			&crop.CropInformation .Name,
			&colorBytes,
			&crop.CropInformation .Family,
			&crop.CropInformation .DensidadPlantacion,
			&crop.CropInformation .LitersPottingSoil  ,
			&associations       CSV,
			&crop.CropRequirements .Water,
			&crop.CropRequirements .Soil,
			&crop.CropRequirements .Nutrition,
			&crop.CropRequirements .Salinity,
			&ph,
			&clima,
			&crop.CropRequirements .Profundidad,
			&crop.CropDates.Planting     ,
			&crop.CropDates.Transplant   ,
			&crop.CropDates.Harvest      ,
			&crop.CropDates.Cycle        ,
			&crop.CropFruit.Production,
			&crop.CropFruit.Nutrients  ,
			&crop.CropSeed.Seed,
			&crop.CropSeed.SeedsPerGram  ,
			&crop.CropSeed.SeedLifespan,
			&crop.CropIssues.Pests       ,
			&crop.CropIssues.Difficulties ,
			&crop.CropIssues.Care,
			&crop.CropIssues.Miscellaneous,
		); err != nil {
			log.Printf("Error al escanear filas: %v", err)
			return nil, err
		}

		log.Printf("Name del crop: %s", crop.CropInformation .Name) // Agrega esta línea

		ph = strings.Replace(ph, "{", "", -1)
		ph = strings.Replace(ph, "}", "", -1)
		phValues := strings.Split(ph, ",")
		phFloats := make([]float64, len(phValues))
		for i, phStr := range phValues {
			phFloat, err := strconv.ParseFloat(phStr, 64)
			if err != nil {
				log.Printf("Error al convertir 'ph' en float64: %v", err)
				return nil, err
			}
			phFloats[i] = phFloat
		}

		requirements.Ph = phFloats

		requirements.Clima = []string{clima}

		err = json.Unmarshal(colorBytes, &colorSlice)
		if err != nil {

		}

		crop.CropInformation .Color = colorSlice
		crop.CropInformation .Associations        = strings.Split(associations       CSV, ",")

		crop.CropRequirements  = requirements

		crops = append(crops, crop)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al recorrer filas: %v", err)
		return nil, err
	}

	return crops, nil
}
