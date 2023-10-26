package repository

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino-pruebas/entity"
)

func (r *Repository) ExtractCultivos() ([]entity.Cultivo, error) {

	rows, err := r.db.Query(ExtractCultivosQuery)

	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL", err)
		return []entity.Cultivo{}, err
	}
	defer rows.Close()

	var cultivos []entity.Cultivo

	for rows.Next() {
		var cultivo entity.Cultivo
		if err := rows.Scan(
			&cultivo.IdCultivo,
			&cultivo.InformacionCultivo.Nombre,
		); err != nil {
			log.Printf("Error al escanear filas", err)
			return nil, err
		}
		cultivos = append(cultivos, cultivo)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al recorrer filas", err)
		return nil, err
	}

	return cultivos, nil
}
