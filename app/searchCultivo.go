package app

import "github.com/robertobouses/meimporta_unpepino/entity"

func (s *Service) SearchCultivo(request entity.SearchRequest) ([]entity.Cultivo, error) {

	result, err := s.repo.ExtractSearchCultivo(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
