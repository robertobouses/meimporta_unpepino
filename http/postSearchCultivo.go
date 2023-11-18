package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (h *Http) PostSearchCrop(ctx *gin.Context) {
	var request entity.SearchRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error ShouldBindJSON PostProblemsCrop": err.Error()})
		return
	}

	results, err := h.service.SearchCrop(request)
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
