package http

import (
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/app"
)

type HTTP interface {
	PostCultivo(ctx *gin.Context)
	GetCultivos(ctx *gin.Context)
	DeleteCultivosAll(ctx *gin.Context)
	DeleteCultivosId(ctx *gin.Context, id string)
	PostProblemsCultivo(ctx *gin.Context)
	PostCalculateCultivo(ctx *gin.Context)
}

type Http struct {
	service app.APP
}

func NewHTTP(service app.APP) HTTP {
	return &Http{
		service: service,
	}
}
