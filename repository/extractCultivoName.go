package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) ExtractCultivoName(name string) (entity.CalculateResponse, error) {
	var response entity.CalculateResponse

	log.Printf("Consulta SQL: %s", ExtractCultivoNameQuery)
	log.Printf("Nombre del cultivo: %s", name)

	err := r.db.QueryRow(ExtractCultivoNameQuery, name).
		Scan(
			&response.Nombre,
			&response.DensidadPlantacion,
		)
	fmt.Println("densidad plantacion repo", response.DensidadPlantacion)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No se encontraron filas para el cultivo con nombre: %s", name)
			return entity.CalculateResponse{}, nil
		}
		log.Printf("Error al ejecutar la consulta SQL: %v", err)
		return entity.CalculateResponse{}, err
	}
	log.Println("cultivo en Repo", response)
	return response, nil
}
