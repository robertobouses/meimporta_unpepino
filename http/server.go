package http

import (
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/app"
)

type HTTP interface {
	PostCrops(ctx *gin.Context)
	GetCrops(ctx *gin.Context)
	DeleteCropsAll(ctx *gin.Context)
	DeleteCropsId(ctx *gin.Context, id string)
	GetCropsIssues(ctx *gin.Context)
	GetCropsCalculate(ctx *gin.Context)
	GetCropsSearch(ctx *gin.Context)
	GetCropsCalendary(ctx *gin.Context)
	PostMyCropsFields(ctx *gin.Context)
	PostMyCropsDefine(ctx *gin.Context)
}

type Http struct {
	service app.APP
}

func NewHTTP(service app.APP) HTTP {
	return &Http{
		service: service,
	}
}
