package http

import (
	"fmt"
	nethttp "net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (h *Http) GetCropsCalculate(ctx *gin.Context) {
	name := ctx.Query("name")
	metersStr := ctx.Query("meters")

	printType(metersStr)

	meters, err := strconv.Atoi(metersStr)
	printType(meters)

	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error al convertir meters a int": err.Error()})
		return
	}

	newCalculate := entity.CalculateRequest{
		Name:         name,
		SquareMeters: meters,
	}
	result, err := h.service.ProcessCropsCalculate(newCalculate)

	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	ctx.JSON(nethttp.StatusOK, gin.H{"mensaje": "Crop insertado correctamente", "result": result})
}

func printType(v interface{}) {
	fmt.Printf("Type: %T\n", v)
}
