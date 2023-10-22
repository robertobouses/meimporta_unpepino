package app

import (
	"github.com/robertobouses/meimporta_unpepino/entity"
	"github.com/robertobouses/meimporta_unpepino/repository"
)

type APP interface {
	CreateCultivo(newCultivo entity.Cultivo) error
	PrintCultivo() (entity.Cultivo, error)
}

type Service struct {
	repo repository.REPOSITORY
}

func NewAPP(repo repository.REPOSITORY) APP {
	return &Service{
		repo: repo,
	}
}
