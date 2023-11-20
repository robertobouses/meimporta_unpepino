package repository

import (
	"fmt"
	"log"

	"github.com/lib/pq"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) InsertCrops(crop entity.Crop) error {
	fmt.Println("InsertCrop - IdCrop:", crop.IdCrop)
	fmt.Printf("Valor que se pasa a la sentencia SQL: %d\n", crop.IdCrop)
	fmt.Printf("Cultivo recibido: %+v\n", crop)

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Println("Error al iniciar la transacción:", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			log.Println("Panic recuperado durante la transacción:", p)
		}
	}()
	var cropID int
	err = tx.Stmt(r.insertCropsStmt).QueryRow(crop.Abbreviation).Scan(&cropID)
	if err != nil {
		log.Println("Error al ejecutar la consulta SQL y escanear cropID:", err)
		return tx.Rollback()
	}

	log.Println("el cropid tras hacer el returning", cropID)

	if err != nil {
		fmt.Println("Error al escanear cropID:", err)
		return err
	}

	_, err = tx.Stmt(r.insertCropsInformationStmt).Exec(
		crop.CropInformation.Name,
		pq.Array(crop.CropInformation.Color),
		crop.CropInformation.Family,
		crop.CropInformation.PlantingDensity,
		crop.CropInformation.LitersPottingSoil,
		pq.Array(crop.CropInformation.Associations),
		cropID,
	)
	if err != nil {
		log.Println("Error al ejecutar la sentencia SQL para insertar CropsInformation:", err)
		return tx.Rollback()
	}

	_, err = tx.Stmt(r.insertCropsRequirementsStmt).Exec(
		crop.CropRequirements.Water,
		crop.CropRequirements.Soil,
		crop.CropRequirements.Nutrition,
		crop.CropRequirements.Salinity,
		pq.Array(crop.CropRequirements.Ph),
		pq.Array(crop.CropRequirements.Climate),
		crop.CropRequirements.Depth,
		cropID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Stmt(r.insertCropsDatesStmt).Exec(
		crop.CropDates.Planting,
		crop.CropDates.Transplant,
		crop.CropDates.Harvest,
		crop.CropDates.Cycle,
		cropID,
	)
	if err != nil {
		return err
	}
	_, err = tx.Stmt(r.insertCropsFruitStmt).Exec(
		crop.CropFruit.Production,
		crop.CropFruit.Nutrients,
		cropID,
	)
	if err != nil {
		return err
	}
	_, err = tx.Stmt(r.insertCropsSeedStmt).Exec(
		crop.CropSeed.Seed,
		crop.CropSeed.SeedsPerGram,
		crop.CropSeed.SeedLifespan,
		cropID,
	)
	if err != nil {
		return err
	}
	_, err = tx.Stmt(r.insertCropsIssuesStmt).Exec(
		crop.CropIssues.Pests,
		crop.CropIssues.Difficulties,
		crop.CropIssues.Care,
		crop.CropIssues.Miscellaneous,
		cropID,
	)
	if err != nil {
		return tx.Rollback()
	}
	if err := tx.Commit(); err != nil {
		log.Println("Error al hacer commit de la transacción:", err)
		return err
	}

	return nil
}
