package app

import (
	"github.com/robertobouses/meimporta_unpepino-pruebas/entity"
	"github.com/robertobouses/meimporta_unpepino-pruebas/repository"
)

type APP interface {
	CreateCultivo(newCultivo entity.Cultivo) error
	PrintCultivos() ([]entity.Cultivo, error)
}

type Service struct {
	repo repository.REPOSITORY
}

func NewAPP(repo repository.REPOSITORY) APP {
	return &Service{
		repo: repo,
	}
}
