package app

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (s *Service) PrintCultivo() (entity.Cultivo, error) {

	cultivo, err := s.repo.ExtractCultivo()
	if err != nil {
		log.Println("Error al extraer cultivo", err)
		return entity.Cultivo{}, err
	}
	return cultivo, nil
}
