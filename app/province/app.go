package app

import (
	provinceEntity "github.com/robertobouses/meimporta_unpepino/entity/province"
	provinceRepo "github.com/robertobouses/meimporta_unpepino/repository/province"
)

type APPProvince interface {
	CreateProvinces(newProvince provinceEntity.Province) error
}

type ServiceProvince struct {
	repo provinceRepo.REPOSITORYProvince
}

func NewAPP(repo provinceRepo.REPOSITORYProvince) APPProvince {
	return &ServiceProvince{
		repo: repo,
	}
}
