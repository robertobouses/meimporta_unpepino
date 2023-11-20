package repository

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) ExtractCropsSearch(request entity.SearchRequest) ([]entity.Crop, error) {
	var crops []entity.Crop

	log.Printf("Consulta SQL completa: %s", ExtractCropsSearchQuery)
	log.Printf("Antes de la ejecución de la consulta: %v", request)
	log.Printf("Valores de la consulta: name=%s, color=%v, water=%v, ...", request.Name, request.Color, request.Water)

	rows, err := r.db.Query(ExtractCropsSearchQuery,
		request.Name,
		pq.Array(request.Color),
		request.PlantingDensity,
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
		var crop entity.Crop
		var color, associations, climate sql.NullString
		var ph sql.NullFloat64

		if err := rows.Scan(
			&crop.IdCrop,
			&crop.Abbreviation,
			&crop.CropInformation.Name,
			&color, &crop.CropInformation.Family,
			&crop.CropInformation.PlantingDensity,
			&crop.CropInformation.LitersPottingSoil,
			&associations,
			&crop.CropRequirements.Water,
			&crop.CropRequirements.Soil,
			&crop.CropRequirements.Nutrition,
			&crop.CropRequirements.Salinity,
			&ph,
			&climate,
			&crop.CropRequirements.Depth,
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
			log.Printf("Error al escanear las filas: %v", err)
			return nil, err
		}

		if color.Valid && color.String != "" {
			crop.CropInformation.Color = append(crop.CropInformation.Color, color.String)
		}
		if associations.Valid && associations.String != "" {
			crop.CropInformation.Associations = append(crop.CropInformation.Associations, associations.String)
		}
		if climate.Valid && climate.String != "" {
			crop.CropRequirements.Climate = append(crop.CropRequirements.Climate, climate.String)
		}
		if ph.Valid {
			phValues := ph.Float64
			crop.CropRequirements.Ph = append(crop.CropRequirements.Ph, phValues)
		}

		crops = append(crops, crop)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al procesar las filas: %v", err)
		return nil, err
	}

	return crops, nil
}
