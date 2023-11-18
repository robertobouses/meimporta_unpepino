package app

import "github.com/robertobouses/meimporta_unpepino/entity"

func (s *Service) SearchCrop(request entity.SearchRequest) ([]entity.Crop, error) {

	result, err := s.repo.ExtractSearchCrop(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
