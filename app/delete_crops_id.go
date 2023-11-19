package app

import (
	"log"
	"strconv"
)

func (s *Service) DeleteCropsId(id string) error {
	idInt, err := strconv.Atoi(id)
	//var ErrCropNoExiste = errors.New("El crop no existe")

	log.Print("el valor del id en la función de la capa app es", idInt)
	if err != nil {
		log.Print("Error al convertir id a int:", err)
		return err
	}

	exists, err := s.repo.CheckExistCropsId(idInt)
	if err != nil {
		return err
	}

	if !exists {
		return ErrCropNoExiste
	}

	err = s.repo.DeleteCropsId(idInt)
	if err != nil {
		log.Print("Error al eliminar el crop:", err)
		return err
	}

	return nil
}
