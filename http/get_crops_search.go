package http

import (
	"fmt"
	"log"
	nethttp "net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (h *Http) GetCropsSearch(ctx *gin.Context) {

	request := entity.SearchRequest{
		Name:            ctx.Query("name"),
		Color:           strings.Split(ctx.Query("color"), ","),
		PlantingDensity: ctx.Query("densidad"),
		Water:           ctx.Query("water"),
		Soil:            ctx.Query("soil"),
		Nutrition:       ctx.Query("nutrition"),
		Salinity:        ctx.Query("salinity"),
		Cycle:           ctx.Query("cycle"),
	}

	fmt.Println("Valor de Name:", request.Name)
	fmt.Println("Valor de Water:", request.Water)
	results, err := h.service.ProcessCropsSearch(request)
	fmt.Printf("Resultados obtenidos: %+v\n", results)

	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(results) == 0 {
		ctx.JSON(nethttp.StatusOK, "No se encontraron resultados para la búsqueda.")
		return
	}
	ctx.JSON(nethttp.StatusOK, results)
}
