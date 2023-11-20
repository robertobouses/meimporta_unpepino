package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) ExtractCropsName(name string) (entity.CalculateResponse, error) {
	var response entity.CalculateResponse

	log.Printf("Consulta SQL: %s", ExtractCropsNameQuery)
	log.Printf("Name del crop: %s", name)

	err := r.db.QueryRow(ExtractCropsNameQuery, name).
		Scan(
			&response.Name,
			&response.PlantingDensity,
		)
	fmt.Println("densidad plantacion repo", response.PlantingDensity)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No se encontraron filas para el crop con name: %s", name)
			return entity.CalculateResponse{}, nil
		}
		log.Printf("Error al ejecutar la consulta SQL: %v", err)
		return entity.CalculateResponse{}, err
	}
	log.Println("crop en Repo", response)
	return response, nil
}
