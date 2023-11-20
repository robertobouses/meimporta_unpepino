package http

import (
	"fmt"
	nethttp "net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func printType(v interface{}) {
	fmt.Printf("Type: %T\n", v)
}

func (h *Http) GetCropsCalculate(ctx *gin.Context) {
	nombre := ctx.Query("nombre")
	metrosStr := ctx.Query("metros")

	printType(metrosStr)

	metros, err := strconv.Atoi(metrosStr)
	printType(metros)

	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error al convertir metros a int": err.Error()})
		return
	}

	newCalculate := entity.CalculateRequest{
		Name:            nombre,
		MetrosCuadrados: metros,
	}
	result, err := h.service.ProcessCropsCalculate(newCalculate)

	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	ctx.JSON(nethttp.StatusOK, gin.H{"mensaje": "Crop insertado correctamente", "result": result})
}
