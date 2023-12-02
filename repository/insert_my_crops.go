package repository

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
)

func (r *Repository) InsertMyCrops(mycrop mycrop.MyCrop) error {
	log.Println("Antes de ejecutar la consulta preparada")
	_, err := r.insertMyCropsStmt.Exec(
		mycrop.Name,
		mycrop.Planting,
		mycrop.FieldId,
	)

	if err != nil {
		log.Print("Error en InsertMyCrops repo", err)
		return err
	}
	log.Println("Después de ejecutar la consulta preparada")
	return nil
}
