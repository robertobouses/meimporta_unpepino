package app

import "github.com/robertobouses/meimporta_unpepino/entity"

func (s *Service) ProcessCropsCalendary(month, province string) (entity.Crop, error) {

	crops, err := s.repo.ExtractCropsCalendary(month, province)
	if err != nil {
		return entity.Crop{}, err
	}

	return crops, nil
}
