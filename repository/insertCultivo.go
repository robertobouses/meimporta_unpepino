package repository

import (
	"fmt"

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

	// _, err = tx.Stmt(r.insertCultivoStmt).Exec(cultivo.Siglas)
	// if err != nil {
	// 	return err
	// }

	_, err = tx.Stmt(r.insertInformacionCultivoStmt).Exec(
		cultivo.InformacionCultivo.Nombre,
		cultivo.InformacionCultivo.LitrosTierraMaceta,
		cultivoID,
	)
	if err != nil {
		return err
	}

	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
