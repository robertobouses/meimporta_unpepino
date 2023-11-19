package app

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (s *Service) CreateCrops(newCrop entity.Crop) error {
	log.Println("CreateCrop - IdCrop:", newCrop.IdCrop)
	// existingCrop, err := s.repo.ExtractCrop(newCrop.IdCrop)
	// if err != nil {

	// 	return err
	// }

	// if existingCrop.Id != "" {

	// 	return fmt.Errorf("El registro con ID %s ya existe", newCrop.IdCrop)
	// }
	err := s.repo.InsertCrops(newCrop)
	if err != nil {
		log.Println("Error en InsertCrop:", err)
		return err
	}
	return nil
}
