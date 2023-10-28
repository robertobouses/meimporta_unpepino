package repository

import (
	"fmt"
	"log"

	"github.com/lib/pq"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (r *Repository) InsertCultivo(cultivo entity.Cultivo) error {
	fmt.Println("InsertCultivo - IdCultivo:", cultivo.IdCultivo)
	fmt.Printf("Valor que se pasa a la sentencia SQL: %d\n", cultivo.IdCultivo)

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Println("Error al iniciar la transacción:", err)
		return err
	}

	var cultivoID int
	err = tx.Stmt(r.insertCultivoStmt).QueryRow(cultivo.Siglas).Scan(&cultivoID)

	if err != nil {
		return err
	}

	log.Println("el cultivoid tras hacer el returning", cultivoID)

	if err != nil {
		fmt.Println("Error al escanear cultivoID:", err)
		return err
	}

	_, err = tx.Stmt(r.insertInformacionCultivoStmt).Exec(
		cultivo.InformacionCultivo.Nombre,
		pq.Array(cultivo.InformacionCultivo.Color),
		cultivo.InformacionCultivo.Familia,
		cultivo.InformacionCultivo.DensidadPlantacion,
		cultivo.InformacionCultivo.LitrosTierraMaceta,
		pq.Array(cultivo.InformacionCultivo.Asociaciones),
		cultivoID,
	)
	if err != nil {

		return err
	}

	_, err = tx.Stmt(r.insertRequisitosCultivoStmt).Exec(
		cultivo.RequisitosCultivo.Agua,
		cultivo.RequisitosCultivo.Tierra,
		cultivo.RequisitosCultivo.Nutricion,
		cultivo.RequisitosCultivo.Salinidad,
		pq.Array(cultivo.RequisitosCultivo.Ph),
		pq.Array(cultivo.RequisitosCultivo.Clima),
		cultivo.RequisitosCultivo.Profundidad,
		cultivoID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Stmt(r.insertFechasCultivoStmt).Exec(
		cultivo.FechasCultivo.Siembra,
		cultivo.FechasCultivo.Transplante,
		cultivo.FechasCultivo.Cosecha,
		cultivo.FechasCultivo.Ciclo,
		cultivoID,
	)
	if err != nil {
		return err
	}
	_, err = tx.Stmt(r.insertFrutoCultivoStmt).Exec(
		cultivo.FrutoCultivo.Produccion,
		cultivo.FrutoCultivo.Nutrientes,
		cultivoID,
	)
	if err != nil {
		return err
	}
	_, err = tx.Stmt(r.insertSemillaCultivoStmt).Exec(
		cultivo.SemillaCultivo.Semilla,
		cultivo.SemillaCultivo.SemillasGramo,
		cultivo.SemillaCultivo.VidaSemilla,
		cultivoID,
	)
	if err != nil {
		return err
	}
	_, err = tx.Stmt(r.insertProblemasCultivoStmt).Exec(
		cultivo.ProblemasCultivo.Plagas,
		cultivo.ProblemasCultivo.Dificultades,
		cultivo.ProblemasCultivo.Cuidados,
		cultivo.ProblemasCultivo.Miscelanea,
		cultivoID,
	)
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
