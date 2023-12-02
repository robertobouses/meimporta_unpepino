package repository

import (
	"database/sql"
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) ExtractMyCrops(name string) (entity.Crop, error) {
	var crops []entity.Crop

	log.Printf("Consulta SQL: %s", ExtractCropsNameQuery)
	log.Printf("name del cultivo: %d", name)

	rows, err := r.db.Query(ExtractCropsNameQuery, name)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var crop entity.Crop
		var color sql.NullString
		var associations sql.NullString
		var ph string
		var climate sql.NullString

		if err := rows.Scan(
			&crop.IdCrop,
			&crop.Abbreviation,
			&crop.CropInformation.Name,
			&color,
			&crop.CropInformation.Family,
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

		if color.Valid {
			crop.CropInformation.Color = append(crop.CropInformation.Color, color.String)
		}
		if associations.Valid {
			crop.CropInformation.Associations = append(crop.CropInformation.Associations, associations.String)
		}
		if climate.Valid {
			crop.CropRequirements.Climate = append(crop.CropRequirements.Climate, climate.String)
		}
		if ph != "" {
			phValues, err := ConvertStringToFloatSlice(ph)
			if err != nil {
				log.Printf("Error al convertir el valor de ph: %v", err)
				return nil, err
			}
			crop.CropRequirements.Ph = append(crop.CropRequirements.Ph, phValues...)
		}

		crops = append(crops, crop)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al procesar las filas: %v", err)
		return nil, err
	}

	return crops, nil
}
