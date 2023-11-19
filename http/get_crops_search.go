package http

import (
	nethttp "net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (h *Http) GetCropsSearch(ctx *gin.Context) {

	request := entity.SearchRequest{
		Name:               ctx.Query("name"),
		Color:              strings.Split(ctx.Query("color"), ","),
		DensidadPlantacion: ctx.Query("densidad"),
		Water:              ctx.Query("water"),
		Soil:               ctx.Query("soil"),
		Nutrition:          ctx.Query("nutrition"),
		Salinity:           ctx.Query("salinity"),
		Cycle:              ctx.Query("cycle"),
	}

	results, err := h.service.ProcessCropsSearch(request)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(results) == 0 {
		ctx.JSON(nethttp.StatusOK, "No se encontraron resultados para la búsqueda.")
		return
	}
	ctx.JSON(nethttp.StatusOK, results)
}
