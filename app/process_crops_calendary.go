package app

import "github.com/robertobouses/meimporta_unpepino/entity"

func (s *Service) ProcessCropsCalendary(month, provinceName string) ([]entity.Crop, error) {
	province, err := s.provinceRepo.ExtractProvincesName(provinceName)
	if err != nil {
		return nil, err
	}

	crops, err := s.repo.ExtractCropsCalendary(month, province.NameProvince)
	if err != nil {
		return nil, err
	}

	return crops, nil
}
