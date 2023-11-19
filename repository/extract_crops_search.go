package repository

import (
	"log"
	"strings"

	"github.com/lib/pq"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) ExtractCropsSearch(request entity.SearchRequest) ([]entity.Crop, error) {
	var crops []entity.Crop

	rows, err := r.db.Query(ExtractCropsSearchQuery,
		request.Name,
		pq.Array(request.Color),
		request.DensidadPlantacion,
		request.Water,
		request.Soil,
		request.Nutrition,
		request.Salinity,
		request.Cycle,
	)

	log.Printf("rows tras la Query en extractSearchCrop del repo: %v", rows)
	log.Printf("Consulta SQL completa: %s", ExtractCropsSearchQuery)
	log.Printf("Valores de la consulta: name=%s, color=%v, water=%v, ...", request.Name, request.Color, request.Water)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		log.Print("escaneando query")
		var crop entity.Crop
		var associationsCSV string
		var ph, clima float64
		var color string

		if err := rows.Scan(
			&crop.IdCrop,
			&crop.CropInformation.Name,
			&color,
			&crop.CropInformation.Family,
			&crop.CropInformation.DensidadPlantacion,
			&crop.CropInformation.LitersPottingSoil,
			&associationsCSV,
			&crop.CropRequirements.Water,
			&crop.CropRequirements.Soil,
			&crop.CropRequirements.Nutrition,
			&crop.CropRequirements.Salinity,
			&ph,
			&clima,
			&crop.CropRequirements.Profundidad,
			&crop.CropDates.Planting,
			&crop.CropDates.Transplant,
			&crop.CropDates.Harvest,
			&crop.CropDates.Cycle,
			&crop.CropFruit.Production,
			&crop.CropFruit.Nutrients,
			&crop.CropSeed.Seed,
			&crop.CropSeed.SeedsPerGram,
			&crop.CropSeed.SeedLifespan,
			&crop.CropIssues.Pests,
			&crop.CropIssues.Difficulties,
			&crop.CropIssues.Care,
			&crop.CropIssues.Miscellaneous,
		); err != nil {
			log.Printf("Error al escanear filas: %v", err)
			return nil, err
		}

		log.Println("Crop en Repo:", crop)

		crop.CropInformation.Associations = strings.Split(associationsCSV, ",")

		crops = append(crops, crop)
	}

	return crops, nil
}
