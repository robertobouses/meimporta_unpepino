package app

import "github.com/robertobouses/meimporta_unpepino/entity"

func (s *Service) ProcessCropsSearch(request entity.SearchRequest) ([]entity.Crop, error) {

	result, err := s.repo.ExtractCropsSearch(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
