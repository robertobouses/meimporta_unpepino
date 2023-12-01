package app

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
)

func (s *Service) CreateFields(field mycrop.MyField) error {
	log.Println("Antes de llamar a InsertFields en el servicio")
	err := s.repo.InsertFields(field)
	if err != nil {
		log.Println("Error en InsertFields del servicio:", err)
		return err
	}
	log.Println("Después de llamar a InsertFields en el servicio")
	return nil
}
