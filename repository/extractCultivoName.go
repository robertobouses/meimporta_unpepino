package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) ExtractCropName(name string) (entity.CalculateResponse, error) {
	var response entity.CalculateResponse

	log.Printf("Consulta SQL: %s", ExtractCropNameQuery)
	log.Printf("Name del crop: %s", name)

	err := r.db.QueryRow(ExtractCropNameQuery, name).
		Scan(
			&response.Name,
			&response.DensidadPlantacion,
		)
	fmt.Println("densidad plantacion repo", response.DensidadPlantacion)
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
