package app

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (s *Service) PrintCultivos() ([]entity.Cultivo, error) {

	cultivos, err := s.repo.ExtractCultivos()
	if err != nil {
		log.Println("Error al extraer cultivo", err)
		return []entity.Cultivo{}, err
	}
	return cultivos, nil
}
