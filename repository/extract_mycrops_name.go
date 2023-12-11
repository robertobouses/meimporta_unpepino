package repository

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
)

func (r *Repository) ExtractMyCropsName(name string) (mycrop.MyCrop, error) {

	log.Printf("Consulta SQL: %s", ExtractMyCropsNameQuery)
	log.Printf("name del cultivo: %d", name)

	rows, err := r.db.Query(ExtractMyCropsNameQuery, name)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		return mycrop.MyCrop{}, err
	}
	defer rows.Close()

	var myCropResponse mycrop.MyCrop

	for rows.Next() {
		if err := rows.Scan(
			&myCropResponse.IdMyCrop,
			&myCropResponse.Name,
			&myCropResponse.Planting,
			&myCropResponse.FieldId,
		); err != nil {
			log.Printf("Error al escanear las filas: %v", err)
			return mycrop.MyCrop{}, err
		}
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al procesar las filas: %v", err)
		return mycrop.MyCrop{}, err
	}

	return myCropResponse, nil
}
