package app

import (
	"log"
	"strings"

	"github.com/rainycape/unidecode"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (s *Service) ProcessCropsCalendary(month, provinceName string) ([]entity.Crop, error) {
	log.Println("datos en APP antes de repo mes:%v, ciudad:%v", month, provinceName)
	normalizedProvinceName := unidecode.Unidecode(strings.ToLower(provinceName))

	log.Printf("Nombre de provincia normalizado:%v\n", normalizedProvinceName)
	province, err := s.provinceRepo.ExtractProvincesName(normalizedProvinceName)
	log.Println("datos en APP extraidos de REPO Provincia:%v", province)

	if err != nil {
		return nil, err
	}

	log.Println("climate obtenido de la consulta al repositorio a la tabla province:", province.ClimateProvince)

	crops, err := s.repo.ExtractCropsCalendary(month, province.ClimateProvince)
	log.Println("datos en APP extraidos de REPO Crops:%v", crops)
	if err != nil {
		return nil, err
	}

	return crops, nil
}
