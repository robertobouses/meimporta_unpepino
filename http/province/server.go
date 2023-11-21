package province

import (
	"github.com/gin-gonic/gin"
	province "github.com/robertobouses/meimporta_unpepino/app/province"
)

type HTTPProvince interface {
	PostProvinces(ctx *gin.Context)
}

type Http struct {
	serviceProvince province.APPProvince
}

func NewHTTP(serviceProvince province.APPProvince) HTTPProvince {
	return &Http{
		serviceProvince: serviceProvince,
	}
}
