package repository

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/lib/pq"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (repo *Repository) ExtractCropsCalendary(month, climate string) ([]entity.Crop, error) {
	var crops []entity.Crop

	rows, err := repo.db.Query(ExtractCropsCalendaryQuery, month, climate)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var crop entity.Crop
		var color sql.NullString
		var associations sql.NullString
		//var phArray string
		var climate sql.NullString

		err := rows.Scan(
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
			pq.Array(&crop.CropRequirements.Ph),
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
		)
		if err != nil {
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

		crops = append(crops, crop)
		log.Printf("Fila procesada: %+v", crop)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al procesar las filas: %v", err)
		return nil, err
	}
	log.Println("crops en repository", crops)
	return crops, nil
}
func parsePhArray(phArray string) ([]float64, error) {

	values := strings.Split(strings.Trim(phArray, "{}"), ",")
	var phValues []float64

	for _, value := range values {
		ph, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		phValues = append(phValues, ph)
	}

	return phValues, nil
}
