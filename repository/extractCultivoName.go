package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) ExtractCultivoName(name string) (entity.Cultivo, error) {
	var cultivo entity.Cultivo

	log.Printf("Consulta SQL: %s", ExtractCultivoNameQuery)
	log.Printf("Nombre del cultivo: %s", name)

	err := r.db.QueryRow(ExtractCultivoNameQuery, name).
		Scan(
			&cultivo.InformacionCultivo.Nombre,
			&cultivo.InformacionCultivo.DensidadPlantacion,
		)
	fmt.Println("densidad plantacion repo", cultivo.InformacionCultivo.DensidadPlantacion)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No se encontraron filas para el cultivo con nombre: %s", name)
			return entity.Cultivo{}, nil
		}
		log.Printf("Error al ejecutar la consulta SQL: %v", err)
		return entity.Cultivo{}, err
	}
	log.Println("cultivo en Repo", cultivo)
	return cultivo, nil
}
