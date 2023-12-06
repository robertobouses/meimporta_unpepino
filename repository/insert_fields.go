package repository

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
)

func (r *Repository) InsertFields(field mycrop.MyField, climate string) error {
	log.Println("Antes de ejecutar la consulta preparada")
	_, err := r.insertFieldsStmt.Exec(
		field.CatastralCode,
		field.Name,
		field.SquareMeters,
		field.Soil,
		field.City,
		field.Province,
		climate,
	)

	if err != nil {
		log.Print("Error en InsertFields repo", err)
		return err
	}
	log.Println("Después de ejecutar la consulta preparada")
	return nil
}
