package app

import (
	"github.com/robertobouses/meimporta_unpepino/entity"
	"github.com/robertobouses/meimporta_unpepino/repository"
)

type APP interface {
	CreateCrops(newCrop entity.Crop) error
	PrintCrops() ([]entity.Crop, error)
	DeleteCropsAll() error
	DeleteCropsId(id string) error
	ProcessCropsIssues(nameIntro string) ([]entity.ProblemResponse, error)
	ProcessCropsCalculate(newCalculate entity.CalculateRequest) (int, error)
	ProcessCropsSearch(request entity.SearchRequest) ([]entity.Crop, error)
}

type Service struct {
	repo repository.REPOSITORY
}

func NewAPP(repo repository.REPOSITORY) APP {
	return &Service{
		repo: repo,
	}
}
