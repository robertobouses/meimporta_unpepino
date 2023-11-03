package app

import (
	"log"
	"strconv"
)

func (s *Service) DeleteCultivosId(id string) error {
	idInt, err := strconv.Atoi(id)
	//var ErrCultivoNoExiste = errors.New("El cultivo no existe")

	log.Print("el valor del id en la función de la capa app es", idInt)
	if err != nil {
		log.Print("Error al convertir id a int:", err)
		return err
	}

	exists, err := s.repo.CheckExistCultivoId(idInt)
	if err != nil {
		return err
	}

	if !exists {
		return ErrCultivoNoExiste
	}

	err = s.repo.DeleteCultivosId(idInt)
	if err != nil {
		log.Print("Error al eliminar el cultivo:", err)
		return err
	}

	return nil
}
