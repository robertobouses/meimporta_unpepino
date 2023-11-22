package app

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (s *Service) ProcessCropsCalendary(month, provinceName string) ([]entity.Crop, error) {
	province, err := s.provinceRepo.ExtractProvincesName(provinceName)
	if err != nil {
		return nil, err
	}
	log.Println("climate obtenido de la consulta al repositorio a la tabla province:", province.ClimateProvince)
	crops, err := s.repo.ExtractCropsCalendary(month, province.ClimateProvince)
	if err != nil {
		return nil, err
	}

	return crops, nil
}
