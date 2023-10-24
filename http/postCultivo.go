package http

import (
	"fmt"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/meimporta_unpepino/entity"
)

func (h *Http) PostCultivo(ctx *gin.Context) {
	var newCultivo entity.Cultivo
	err := ctx.BindJSON(&newCultivo)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error al hacer BindJSON": err.Error()})
		return
	}

	fmt.Printf("Valores de newCultivo: %+v\n", newCultivo)

	err = h.service.CreateCultivo(newCultivo)
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	ctx.JSON(nethttp.StatusOK, newCultivo)
}
