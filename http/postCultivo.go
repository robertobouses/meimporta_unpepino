package http

import (
	"fmt"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino-pruebas/entity"
)

func (h *Http) PostCultivo(ctx *gin.Context) {
	var newCultivo entity.Cultivo
	err := ctx.BindJSON(&newCultivo)
	fmt.Print("cultivo en PostCultivo", newCultivo)
	fmt.Print("Id cultivo en PostCultivo", newCultivo.IdCultivo)

	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error al hacer BindJSON": err.Error()})
		return
	}

	fmt.Printf("Valores de newCultivo: %+v\n", newCultivo)

	err = h.service.CreateCultivo(newCultivo)
	fmt.Print("cultivo en PostCultivo2", newCultivo)
	fmt.Print("Id cultivo en PostCultivo2", newCultivo.IdCultivo)
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	//cultivoID := newCultivo.IdCultivo
	ctx.JSON(nethttp.StatusOK, gin.H{"idcultivo": newCultivo.IdCultivo})
	ctx.JSON(nethttp.StatusOK, newCultivo)
	fmt.Printf("Cultivo insertado con ID: %d\n", newCultivo.IdCultivo)
	ctx.JSON(nethttp.StatusOK, gin.H{"mensaje": "Cultivo insertado correctamente"})
}
