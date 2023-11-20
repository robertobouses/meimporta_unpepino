package app

import (
	"log"

	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (s *Service) ProcessCropsSearch(request entity.SearchRequest) ([]entity.Crop, error) {
	log.Println("request en APP", request)
	log.Println("name en APP", request.Name)
	log.Println("water en APP", request.Water)
	result, err := s.repo.ExtractCropsSearch(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
