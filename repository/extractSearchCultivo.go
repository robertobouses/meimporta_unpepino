package repository

import (
	"log"

	"github.com/lib/pq"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) ExtractSearchCrop(request entity.SearchRequest) ([]entity.Crop, error) {
	var crops []entity.Crop

	rows, err := r.db.Query(ExtractSearchCropsQuery,
		request.Name,
		pq.Array(request.Color),
		request.DensidadPlantacion,
		request.Water,
		request.Soil,
		request.Nutrition,
		request.Salinity,
		request.Cycle        ,
	)

	log.Printf("rows tras la Query en extractSearchCrop del repo", rows)
	log.Printf("Consulta SQL completa: %s", ExtractSearchCropsQuery)
	log.Printf("Valores de la consulta: name=%s, color=%v, water=%v, ...", request.Name, request.Color, request.Water)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		log.Print("escaneando query")
		var crop entity.Crop
		var colorBytes []byte
		var associations       CSV string
		var ph, clima float64

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

		log.Println("Crop en Repo:", crop)
		crops = append(crops, crop)
	}

	return crops, nil
}
