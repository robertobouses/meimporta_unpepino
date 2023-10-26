package app

import (
	"log"
	"strconv"
)

func (s *Service) DeleteCultivosId(id string) error {
	idInt, err := strconv.Atoi(id)
	log.Print("el valor del id en la función de la capa app es", idInt)
	if err != nil {
		log.Print("Error al convertir id a int:", err)
		return err
	}

	err = s.repo.DeleteCultivosId(idInt)
	if err != nil {
		log.Print("Error:", err)
		return err
	}

	return nil
}
