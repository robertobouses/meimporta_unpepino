package http

import (
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/app"
)

type HTTP interface {
	PostCultivo(ctx *gin.Context)
	GetCultivo(ctx *gin.Context)
}

type Http struct {
	service app.APP
}

func NewHTTP(service app.APP) HTTP {
	return &Http{
		service: service,
	}
}
