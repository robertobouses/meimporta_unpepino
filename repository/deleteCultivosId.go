package repository

import "log"

func (r *Repository) DeleteCultivosId(id int) error {
	log.Print("el valor del id en la función de la capa http es", id)
	_, err := r.db.Exec(DeleteCultivosIdQuery, id)
	if err != nil {
		return err
	}
	return nil
}
