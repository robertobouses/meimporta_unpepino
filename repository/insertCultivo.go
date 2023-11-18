package repository

import (
	"fmt"
	"log"

	"github.com/lib/pq"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) InsertCrop(crop entity.Crop) error {
	fmt.Println("InsertCrop - IdCrop:", crop.IdCrop)
	fmt.Printf("Valor que se pasa a la sentencia SQL: %d\n", crop.IdCrop)

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Println("Error al iniciar la transacción:", err)
		return err
	}

	var cropID int
	err = tx.Stmt(r.insertCropStmt).QueryRow(crop.Abbreviation).Scan(&cropID)

	if err != nil {
		return err
	}

	log.Println("el cropid tras hacer el returning", cropID)

	if err != nil {
		fmt.Println("Error al escanear cropID:", err)
		return err
	}

	_, err = tx.Stmt(r.insertCropInformation Stmt).Exec(
		crop.CropInformation .Name,
		pq.Array(crop.CropInformation .Color),
		crop.CropInformation .Family,
		crop.CropInformation .DensidadPlantacion,
		crop.CropInformation .LitersPottingSoil,
		pq.Array(crop.CropInformation .Associations),
		cropID,
	)
	if err != nil {

		return err
	}

	_, err = tx.Stmt(r.insertCropRequirements Stmt).Exec(
		crop.CropRequirements .Water,
		crop.CropRequirements .Soil,
		crop.CropRequirements .Nutrition,
		crop.CropRequirements .Salinity,
		pq.Array(crop.CropRequirements .Ph),
		pq.Array(crop.CropRequirements .Clima),
		crop.CropRequirements .Profundidad,
		cropID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Stmt(r.insertCropDatesStmt).Exec(
		crop.CropDates.Planting,
		crop.CropDates.Transplant,
		crop.CropDates.Harvest,
		crop.CropDates.Cycle,
		cropID,
	)
	if err != nil {
		return err
	}
	_, err = tx.Stmt(r.insertCropFruitStmt).Exec(
		crop.CropFruit.Production,
		crop.CropFruit.Nutrients  ,
		cropID,
	)
	if err != nil {
		return err
	}
	_, err = tx.Stmt(r.insertCropSeedStmt).Exec(
		crop.CropSeed.Seed,
		crop.CropSeed.SeedsPerGram,
		crop.CropSeed.SeedLifespan,
		cropID,
	)
	if err != nil {
		return err
	}
	_, err = tx.Stmt(r.insertCropIssuesStmt).Exec(
		crop.CropIssues.Pests,
		crop.CropIssues.Difficulties ,
		crop.CropIssues.Care,
		crop.CropIssues.Miscellaneous,
		cropID,
	)
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
