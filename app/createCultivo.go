package app

import (
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (s *Service) CreateCultivo(newCultivo entity.Cultivo) error {

	// existingCultivo, err := s.repo.ExtractCultivo(newCultivo.IdCultivo)
	// if err != nil {

	// 	return err
	// }

	// if existingCultivo.Id != "" {

	// 	return fmt.Errorf("El registro con ID %s ya existe", newCultivo.IdCultivo)
	// }

	return s.repo.InsertCultivo(newCultivo)
}
