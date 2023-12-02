package app

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
)

func (s *Service) CreateMyCrops(mycrop mycrop.MyCrop) error {
	log.Println("Antes de llamar a InsertFields en el servicio")
	err := s.repo.InsertMyCrops(mycrop)
	if err != nil {
		log.Println("Error en InsertCrops del servicio:", err)
		return err
	}
	log.Println("Después de llamar a InsertCrops en el servicio")
	return nil
}
