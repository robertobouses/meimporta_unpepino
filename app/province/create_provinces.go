package app

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity/province"
)

func (s *ServiceProvince) CreateProvinces(newProvince province.Province) error {

	err := s.repo.InsertProvinces(newProvince)
	if err != nil {
		log.Println("Error en InsertProvince:", err)
		return err
	}
	return nil
}
