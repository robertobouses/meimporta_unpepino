package http

import (
	"fmt"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (h *Http) PostProblemsCrop(ctx *gin.Context) {
	var request entity.NameCropRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error ShouldBindJSON PostProblemsCrop": err.Error()})
		return
	}

	results, err := h.service.SearchProblemsCrop(request.Name)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(results) == 0 {
		ctx.JSON(nethttp.StatusOK, "No se encontraron resultados para la búsqueda.")
		return
	}

	response := "Resultados:\n"
	for _, result := range results {
		response += fmt.Sprintf("Crop: %s\nPests       : %s\nDifficulties : %s\nCare: %s\nMiscellaneous: %s\n\n",
			result.Name, result.Pests, result.Difficulties, result.Care, result.Miscellaneous)
	}

	ctx.JSON(nethttp.StatusOK, response)
}
