package app

import (
	"github.com/robertobouses/meimporta_unpepino/entity"
	"github.com/robertobouses/meimporta_unpepino/entity/mycrop"
	"github.com/robertobouses/meimporta_unpepino/repository"

	repoProvince "github.com/robertobouses/meimporta_unpepino/repository/province"
)

type APP interface {
	CreateCrops(newCrop entity.Crop) error
	PrintCrops() ([]entity.Crop, error)
	DeleteCropsAll() error
	DeleteCropsId(id string) error
	ProcessCropsIssues(nameIntro string) ([]entity.IssueResponse, error)
	ProcessCropsCalculate(newCalculate entity.CalculateRequest) (int, error)
	ProcessCropsSearch(request entity.SearchRequest) ([]entity.Crop, error)
	ProcessCropsCalendary(month, provinceName string) ([]entity.Crop, error)
	CreateFields(field mycrop.MyField) error
	CreateMyCrops(mycrop mycrop.MyCrop) error
	ProcessMyCrops(name string) (mycrop.MyCropResult, error)
}

type Service struct {
	repo         repository.REPOSITORY
	provinceRepo repoProvince.REPOSITORYProvince
}

func NewAPP(repo repository.REPOSITORY, provinceRepo repoProvince.REPOSITORYProvince) APP {
	return &Service{
		repo:         repo,
		provinceRepo: provinceRepo,
	}
}
