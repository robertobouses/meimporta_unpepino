package http

import (
	"fmt"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (h *Http) PostProblemsCultivo(ctx *gin.Context) {
	var request entity.NombreCultivoRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error ShouldBindJSON PostProblemsCultivo": err.Error()})
		return
	}

	results, err := h.service.SearchProblemsCultivo(request.Name)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, err)
		return
	}

	if len(results) == 0 {
		ctx.JSON(nethttp.StatusOK, "No se encontraron resultados para la búsqueda.")
		return
	}

	response := "Resultados:\n"
	for _, result := range results {
		response += fmt.Sprintf("Cultivo: %s\nPlagas: %s\nDificultades: %s\nCuidados: %s\nMiscelanea: %s\n\n",
			result.Nombre, result.Plagas, result.Dificultades, result.Cuidados, result.Miscelanea)
	}

	ctx.JSON(nethttp.StatusOK, response)
}
