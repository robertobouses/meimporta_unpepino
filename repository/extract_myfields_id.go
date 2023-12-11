package repository

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
)

func (r *Repository) ExtractMyFieldsId(id int) (mycrop.MyField, error) {

	log.Printf("Consulta SQL: %s", ExtractMyFieldsIdQuery)
	log.Printf("id del cultivo: %d", id)

	rows, err := r.db.Query(ExtractMyFieldsIdQuery, id)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		return mycrop.MyField{}, err
	}
	defer rows.Close()

	var myFieldResponse mycrop.MyField

	for rows.Next() {
		if err := rows.Scan(
			&myFieldResponse.IdMyField,
			&myFieldResponse.CatastralCode,
			&myFieldResponse.Name,
			&myFieldResponse.SquareMeters,
			&myFieldResponse.Soil,
			&myFieldResponse.City,
			&myFieldResponse.Province,
			&myFieldResponse.Climate,
		); err != nil {
			log.Printf("Error al escanear las filas: %v", err)
			return mycrop.MyField{}, err
		}
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al procesar las filas: %v", err)
		return mycrop.MyField{}, err
	}

	return myFieldResponse, nil
}
