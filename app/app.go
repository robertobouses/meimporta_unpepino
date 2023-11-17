package app

import (
	"github.com/robertobouses/meimporta_unpepino/entity"
	"github.com/robertobouses/meimporta_unpepino/repository"
)

type APP interface {
	CreateCultivo(newCultivo entity.Cultivo) error
	PrintCultivos() ([]entity.Cultivo, error)
	DeleteCultivosAll() error
	DeleteCultivosId(id string) error
	SearchProblemsCultivo(nameIntro string) ([]entity.ProblemResponse, error)
	CreateCalculateCultivo(newCalculate entity.CalculateRequest) (int, error)
	SearchCultivo(request entity.SearchRequest) ([]entity.Cultivo, error)
}

type Service struct {
	repo repository.REPOSITORY
}

func NewAPP(repo repository.REPOSITORY) APP {
	return &Service{
		repo: repo,
	}
}
