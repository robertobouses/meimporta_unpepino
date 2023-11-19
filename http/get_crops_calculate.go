package http

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (h *Http) GetCropsCalculate(ctx *gin.Context) {
	var newCalculate entity.CalculateRequest
	log.Println("Cuerpo de la solicitud:", ctx.Request.Body)
	err := ctx.BindJSON(&newCalculate)
	log.Printf("newCalculate: %+v", newCalculate)

	log.Println("newCalculate en http", newCalculate)

	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error al hacer BindJSON": err.Error()})
		return
	}

	result, err := h.service.ProcessCropsCalculate(newCalculate)
	log.Println("resultado en http", result)
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	ctx.JSON(nethttp.StatusOK, gin.H{"mensaje": "Crop insertado correctamente", "result": result})
}
