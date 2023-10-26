package app

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (s *Service) CreateCultivo(newCultivo entity.Cultivo) error {
	log.Println("CreateCultivo - IdCultivo:", newCultivo.IdCultivo)
	// existingCultivo, err := s.repo.ExtractCultivo(newCultivo.IdCultivo)
	// if err != nil {

	// 	return err
	// }

	// if existingCultivo.Id != "" {

	// 	return fmt.Errorf("El registro con ID %s ya existe", newCultivo.IdCultivo)
	// }
	err := s.repo.InsertCultivo(newCultivo)
	if err != nil {
		log.Println("Error en InsertCultivo:", err)
		return err
	}
	return nil
}
